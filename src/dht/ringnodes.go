package main

import (
	"fmt"			// standard input, output etc
	"errors"		// error handling
	"math"			
	"sync"			// sync muting, allowing concurrency
	"sync/atomic" 	// atomic operations
	rbt "github.com/arriqaaq/rbt"
	xxhash "github.com/cespare/xxhash"
)


var (
	ERR_EMPTY_RING = errors.New("Empty ring")
	ERR_KEY_NOT_FOUND = errors.New("Key was not found in the ring")
	ERR_HEAVY_LOAD = errors.New("Servers under load, retry after a while")
	ERR_UNLOADED_HOST = errors.New("Can not find jobs on the said cluster")
	ERR_UNLOADED_SERVER = errors.New("Can not find jobs on the said node")
)

type hashFunction interface {
	hash(string)	int64 					// the hash function returns int32 key to let us search
											// TODO : implement using perfect hash tables, re-hash on the go in each node independently								
}

type sha256HashFunction struct {
}

func (x sha256HashFunction) hash(data string) int64 {
	xhash := xxhash.New()
	xhash.Write([]byte(data))
	r := xhash.Sum64()
	xhash.Reset()
	return int64(r)
}

func newHashFunction() hashFunction {
	return sha256HashFunction{}
}



// End of code of hash function
// Starting code for Node from here 

type node struct {
	name string 							// name of the sake of clarity, and printing and a make-do for node identity (should be hashed value of string, but trying to make simple thing right now
	active bool 							// activity flag
	load int64 								// load count
}

func (n *node) Load() int64 {
	return n.load
}

func newNode(name string) *node {
	return &node {
		name : name , 
		active : false,
		load : 0,
	}
}


// End of code for Node

// Start of code for configuration 

type Config struct {
	VirtualNodes int 						// virtual nodes count
	LoadFactor float64 						// load factor of key-value store
}

// End of code for configuration

// Start of code for ring

type Ring struct {
	store *rbt.Tree 						// red black tree for nodes
	nodeMap map[string]*node 				// mapping from string (input) to node that has to be accessed 
	hashfn hashFunction 					// hashfn which implements interface hashFunction
	VirtualNodes int 						// count of virtual nodes
	LoadFactor float64 						// stored for convenience 
	TotalLoad int64  						// total load on the system

	mu sync.RWMutex 						// semaphore for maintaining operations atomic
}


// code for adding a node in the ring
func (r *Ring) Add(node string){
	r.mu.Lock()
	defer r.mu.Unlock()						// given a fair scheduler, this will execute before any other statement
	if _, ok := r.nodeMap[node] ;  !ok {
		r.nodeMap[node] = newNode(node); 	// adding original, true node to the ring
		hashKey := r.hashfn.hash(node);			// adding mapping for the ring from the RBT
		r.store.Insert(hashKey, node);		
		for i:=0 ; i < r.VirtualNodes ; i++ { 			// add multiple copies of the key in order to get uniforly hashed copies
			vNodeKey := fmt.Sprintf("%s-%d", node, i)
			r.nodeMap[vNodeKey] = newNode(vNodeKey)		// create a node present on every virtualNode corresponding to a given node
			hashKey := r.hashfn.hash(vNodeKey)
			r.store.Insert(hashKey, node)				// store name of the node using hashKey
		}
	}

}

func (r *Ring) Remove(key string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.nodeMap[key] ; !ok {
		return
	}
	hashKey := r.hashfn.hash(key)
	r.store.Delete(hashKey)
	delete(r.nodeMap, key)
	for i := 0 ; i < r.VirtualNodes ; i++ {
		vNodeKey := fmt.Sprintf("%s-%d", key, i)
		hashKey := r.hashfn.hash(vNodeKey)
		r.store.Delete(hashKey)
		delete(r.nodeMap, key)
	}
}

func (r *Ring) checkLoad(key string) bool {
	if r.TotalLoad < 0 {
		r.TotalLoad = 0
	}
	var avgLoadPerNode float64
	avgLoadPerNode = float64(int(r.TotalLoad + 1) / len(r.nodeMap))
	if avgLoadPerNode == 0 {
		avgLoadPerNode = 1
	}
	avgLoadPerNode = math.Ceil(avgLoadPerNode * r.LoadFactor)
	vnode, ok := r.nodeMap[key]
	if !ok {
		fmt.Printf("given host(%s) not in loadsMap\n", vnode.name)
		return false
	}
	if float64(vnode.load) > avgLoadPerNode {
		return false
	}
	return true
}

func (r *Ring) Get (key string) (string, error){
	r.mu.RLock()											// block reading, blocks for WUnlock on other threads
	defer r.mu.RUnlock()
	if r.store.Size() == 0 {
		return "", ERR_EMPTY_RING
	}
	hashKey := r.hashfn.hash(key)
	q := r.store.Nearest(hashKey)
	var count int
	count = 0
	for {
		if count >= r.store.Size() {
			fmt.Println("Looked through everything, could not find free node")
			return "", ERR_HEAVY_LOAD
		}
		if hashKey > q.GetKey()  {							// find next bucket in ring
			g := rbt.FindSuccessor(q)
			if g != nil {
				q = g
			} else {
				q = r.store.Root()
			}
		}
		if r.checkLoad(q.GetValue()) {
			break; 											// found a node where we can work
		}
		count ++;
		h := rbt.FindSuccessor(q)							// go to the beginning if you can not find it
		if h != nil {
			q = h
		} else {
			q = r.store.Minimum()
		}
	}
	atomic.AddInt64(&r.nodeMap[q.GetValue()].load, 1)
	atomic.AddInt64(&r.TotalLoad, 1)
	return q.GetValue(), nil
}

func (r *Ring) Done (key string) error {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if _, ok := r.nodeMap[key] ; !ok {
		return ERR_KEY_NOT_FOUND
	}
	if r.nodeMap[key].load <= 0 {
		return ERR_UNLOADED_SERVER
	}
	if r.TotalLoad <= 0 {
		return ERR_UNLOADED_HOST
	}
	atomic.AddInt64(&r.nodeMap[key].load, -1)
	atomic.AddInt64(&r.TotalLoad, -1)
	return nil
}

func NewRing() *Ring {
	return &Ring {
		store : rbt.NewTree(),
		nodeMap : make(map[string]*node),
		hashfn : newHashFunction(),
	}
}


func NewRingWithConfig(nodes []string, cfg Config) *Ring {
	r := &Ring{
		store : rbt.NewTree(),
		nodeMap : make(map[string]*node),
		hashfn : newHashFunction(),
		VirtualNodes : cfg.VirtualNodes,
		LoadFactor : cfg.LoadFactor,
	}
	if r.LoadFactor <= 0 {
		r.LoadFactor = 1
	}

	for _, node := range nodes {
		r.Add(node);
	}
	return r
}

// End of code for ring

// Start of tests of individual functionality


func testHashFunction(){ 					// test for hash function
	data := "HelloWorld"
	hashFunction := newHashFunction()
	fmt.Print(hashFunction.hash(data))
	fmt.Println()
}

func testRingInitComplete() {
	nodeList := []string{"node1", "node2", "node3"}
	var cfg Config;
	cfg.VirtualNodes = 1
	cfg.LoadFactor = 4
	ring := NewRingWithConfig(nodeList, cfg)
	ring.Remove("node1")
	fmt.Println(ring.Get("node3"))
	fmt.Println(ring.Done("node3"))
}

func main() {
	testHashFunction()
	testRingInitComplete()
	fmt.Println("Main : test. Compile works. Run works. Exiting cleanly")
}
