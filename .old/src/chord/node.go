package main

import (
	"github.com/prannayk/DistributedFileSystem/src/node"
	"google.golang.org/grpc"
	"errors"
	"sync"
	"net"
	"time"
	"fmt"
)
	
type Node struct {
	ChordStructure.Node
	Options					NodeOptions 
	GRPCServer 				*grpc.Server
	FTMutex 				sync.RWMutex
	FingerTable 			fingerTable
	Pred 					*Node 
	Successor 				*Node
	predMtx 				sync.RWMutex
	succMtx 				sync.RWMutex
	ShutDownChannel 		chan struct{}
	DataStore 				map[string][]byte
	DSMtx 					sync.RWMutex
	ClientConnections 		map[string]*ClientConn
	CCMtx 					sync.RWMutex
}

type NodeOptions struct {
	ID 				[]byte
	Name 			string
	ServerOptions 	[]grpc.ServerOption
	DialOptions 	[]grpc.DialOption
}

// Following 3 functions have been implemented from the paper
// They follow the exact methods, except the involvement of RPC
func (node *Node) FindSuccessor(nextHash []byte) (*Node, error) {
	pred , err := node.FindPredecessor(nextHash)
	if err != nil {
		return nil, err
	}
	if pred.Name == "" {
		return node, nil
	}
	succ, err := node.GetSuccessorRPC(pred)
	if err != nil {
		return nil, err
	}
	if succ.Name == "" {
		return node, nil
	}
	return succ, nil;
}

func (node *Node) FindPredecessor(nextHash []byte) (*Node, error) {
	pred := node
	for {
		succ, _ := pred.GetSuccessorRPC(pred)
		if succ == nil || succ.Name == "" {
			return pred, nil
		}
		if between(nextHash, pred.Id, succ.Id) {
			break
		}
		pred,err := pred.ClosestPrecedingFingerRPC(nextHash)
		if err != nil {
			return nil, err
		}
		if pred.Name == "" {
			return node, nil
		}
	}
	return pred, nil
}

func (node *Node) ClosestPrecedingFinger(nextHash []byte) *Node {
	node.FTMutex.Lock()
	defer node.FTMutex.Unlock()
	for i := range node.FingerTable {
		if node.FingerTable[i].RemoteNode == nil {
			continue
		}
		if between(node.FingerTable[i].RemoteNode.Id, node.Id, nextHash) {
			return node.FingerTable[i].RemoteNode
		}
	}
	return node
}

type NodeOption func(o *NodeOptions)

func SetID(id []byte) NodeOption {
	return func (o *NodeOptions){
		o.ID = id
	}
}

func SetAddress (name string) NodeOption {
	return func(o *NodeOptions) {
		o.Name = name
	}
}

func SetServerOptions (opts ...grpc.ServerOption) NodeOption {
	return func(o *NodeOptions) {
		o.ServerOptions = opts
	}
}

func SetDialOptions (opts ...grpc.DialOption) NodeOption {
	return func(o *NodeOptions){
		o.DialOptions = opts
	}
}

func NewNodeFromParent(parent * Node, opts ... NodeOption) (* Node , error) {
	node := &Node{
		ShutDownChannel : make(chan struct{}),
		ClientConnections : make(map[string]*ClientConn),
	}

	for _, opt := range opts {
		opt(&node.Options)
	}

	id := node.Options.ID
	node.Id = id

	name := node.Options.Name
	node.Name = name

	lis, err := net.Listen("tcp", name)
	if err != nil {
		return nil, err
	}
	node.GRPCServer = grpc.NewServer(node.Options.ServerOptions...)
	// chord.RegisterChordServer(node.GRPCServer, node) 		// TODO : implement chord
	// ChordStructure.RegisterChordDHTServer(node.GRPCServer, parent) 		// TODO : implement API

	if id != nil {
		if len(id) != config.IDLength {
			return nil, ERR_BAD_ID_LEN
		}
	} else {
		id, err := HashKey(lis.Addr().String())
		if err != nil {
			return nil, err
		} 
		node.Id = id
	}
	node.Name = lis.Addr().String()
	node.DataStore = make(map[string][]byte)
	node.FingerTable = newFingerTable(node)
	go node.GRPCServer.Serve(lis) 							// make the node available on the said TCP server
	var joinNode *Node
	if parent != nil {
		remoteNode, err := node.FindSuccessorRPC(parent, node.Id) // unimplemented right now
		if err != nil {
			return nil, err
		}
		if EqualID(remoteNode.Id, node.Id){
			return nil, errors.New("Node with the set ID already exists. Better luck next time!")
		}
		joinNode = parent
		
	} else {
		joinNode = node
	}

	if err := node.Join(joinNode) ; err != nil {
		return nil, err
	}
	// timer to handle the timer periodically
	go func() {
		ticker := time.NewTicker(config.StabilizeInterval)
		for {
			select {
				case <-ticker.C :
					node.stabilize()
				case <-node.ShutDownChannel :
					ticker.Stop()
					return
			}
		}
	} ()
	// timer to fix finger table periodically
	go func() {
		next := 0
		ticker := time.NewTicker(config.FixNextFingerInterval)
		for {
			select {
			case <-ticker.C :
				next = node.FixNextFinger(next)
			case <-node.ShutDownChannel :
				ticker.Stop()
				return
			}
		}
	}()
	<-time.After(config.StabilizeInterval)				// put to sleep for StabilizeInterval, return after stabilization is complete for the first time
	return node, nil
}


func (node * Node) Join(parent * Node) error {
	succ, err := node.FindSuccessorRPC(parent, node.Id)
	if err != nil {
		return err
	}
	node.succMtx.Lock()
	node.Successor = succ
	node.succMtx.Unlock()
	return node.ObtainNewKeys()
}

func (node *Node) ObtainNewKeys() error {
	node.succMtx.RLock()
	succ := node.Successor
	node.succMtx.RUnlock()

	prevPredecessor, err := node.GetPredecessorRPC(succ)
	if err != nil {
		return err
	}
	return node.TransferKeysRPC(succ, prevPredecessor, node.Id)
}

func (node *Node) stabilize(){
	node.succMtx.RLock()
	_succ := node.Successor
	node.succMtx.RUnlock()
	if _succ == nil {
		return
	}
	
	succ, err := node.GetPredecessorRPC(_succ)
	if succ == nil || err != nil {

		return 
	}
	if succ.Id != nil && between(succ.Id, node.Id, _succ.Id) {
		node.succMtx.Lock()
		node.Successor = succ
		node.succMtx.Unlock()
	}

	_ = node.NotifyRPC(_succ, node)				// at this step we inform the old successor that it has beenr reset
}	

func (node *Node) notify(remoteNode *Node){
	node.predMtx.Lock()
	defer node.predMtx.Unlock()
	pred := node.Pred
	if pred != nil && !between(remoteNode.Id, node.Pred.Id, node.Id) {
		return
	}
	var PrevID []byte
	if pred != nil {
		PrevID = pred.Id
	}
	node.Pred = remoteNode
	if between(node.Pred.Id, PrevID, node.Id){
		_ = node.TransferKeys(PrevID, node.Pred)
	}
}

func (node *Node) ShutDown() {
	close(node.ShutDownChannel)
	node.GRPCServer.Stop()
	node.succMtx.RLock()
	node.predMtx.RLock()
	if node.Name != node.Successor.Name && node.Pred != nil {
		_ = node.TransferKeys(node.Pred.Id, node.Successor)
		_ = node.SetSuccessorRPC(node.Pred, node.Successor)
		_ = node.SetPredecessorRPC(node.Successor, node.Pred)
	}
	node.predMtx.RUnlock()
	node.succMtx.RUnlock()

	node.CCMtx.Lock()
	for _, conn := range node.ClientConnections {
		_ = conn.conn.Close()
	}
	defer node.CCMtx.Unlock()
}

func (node *Node) String() string {
	var succ []byte
	var pred []byte

	node.succMtx.RLock()
	if node.Successor != nil {
		succ = node.Successor.Id
	}	
	node.succMtx.RUnlock()

	node.predMtx.RLock()
	if node.Pred != nil {
		pred = node.Pred.Id
	}
	node.predMtx.RUnlock()

	return fmt.Sprintf("Node-%v : Address : %s {succ :%v, pred : %v}",
						IDToString(node.Id), node.Name,
						IDToString(succ), IDToString(pred),
						)
}


func NewNode(id []byte, Name string) *Node {
	n := &Node {}
	n.Id = id
	n.Name = Name
	n.ShutDownChannel = make(chan struct{})
	n.ClientConnections = make(map[string]*ClientConn)

	return n
}
