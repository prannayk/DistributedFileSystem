package main

import (
	"github.com/prannayk/DistributedFileSystem/src/node"
	"google.golang.org/grpc"
	"sync"
	"net"
	"golang.org/x/net/context"
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

func (node *Node) GetData(ctx context.Context, key *ChordStructure.GetRequest) (*ChordStructure.GetResponse, error) {
	val := []byte("HelloDhish")
	return &ChordStructure.GetResponse{Value : val}, nil
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
	// node.fingerTable = Ne
	return node, nil
}

func NewNode(id []byte, Name string) *Node {
	n := &Node {}
	n.Id = id
	n.Name = Name
	n.ShutDownChannel = make(chan struct{})
	n.ClientConnections = make(map[string]*ClientConn)

	return n
}
