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


func (node *Node) FindSuccessor(nextHash []byte) (*Node, error) {
	// TODO(prannayk) : unimplemented method
	return node, nil;
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


func (node *Node) stabilize() {
	fmt.Println("Unimplemented method called");
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

func NewNode(id []byte, Name string) *Node {
	n := &Node {}
	n.Id = id
	n.Name = Name
	n.ShutDownChannel = make(chan struct{})
	n.ClientConnections = make(map[string]*ClientConn)

	return n
}
