package main

import (
	"golang.org/x/net/context"
	"errors"
	"github.com/prannayk/DistributedFileSystem/src/node"
)

var (
	ERR_NO_SUCCESSOR = errors.New("Could not find successor in RPC")
)

type ClientConn struct {
}

func (node *Node) GetData(ctx context.Context, key *ChordStructure.GetRequest) (*ChordStructure.GetResponse, error) {
	val := []byte("HelloDhish")
	return &ChordStructure.GetResponse{Value : val}, nil
}

func (node *Node) FindSuccessorRPC(parent *Node, NodeID []byte) (*Node, error){
	return nil, ERR_NO_SUCCESSOR
}

func (node *Node) GetPredecessorRPC(succ *Node) (*Node, error){
	return succ, nil
}

func (node *Node) TransferKeysRPC(succ, prevPred *Node, NodeID []byte) error {
	return errors.New("Unimplemented function TransferKeysRPC")
}
