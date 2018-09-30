package main

import "fmt"

import (
	"fmt"
	"github.com/prannayk/DistributedFileSystem/src/chord/node"
)

type fingerTable []*fingerEntry
type Node  ChordStructure.Node
type fingerEntry struct {
	ID []byte 				// ID hash of the relevant entry
	RemoteNode *Node
}

func newFingerTable(node *Node) fingerTable {
	ft := make([]*Node, config)
	return ft
}

func main() {
	fmt.Println("vim-go")
}
