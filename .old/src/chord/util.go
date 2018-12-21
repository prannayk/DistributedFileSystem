package main

import (
	"fmt"
	"math/big"
	xxhash "github.com/cespare/xxhash"
)

func HashKey (key string) ([]byte,error) {
	h := xxhash.New()
	if _, err := h.Write([]byte(key)) ; err != nil {
		return nil, err
	}
	r := h.Sum(nil)
	h.Reset()
	return r[:config.IDLength], nil
}

func PadID(NodeID []byte) []byte {
	n := config.IDLength - len(NodeID)
	if n < 0 {
		n = 0
	}
	_NodeID := make([]byte, n)
	NodeID = append(_NodeID, NodeID...) 	// handles smaller than n case
	return NodeID[:config.IDLength] // handles large case
}

func IDToString(id []byte) string {
	keyInt := big.Int{}
	keyInt.SetBytes(id)
	return keyInt.String()
}

func TestID() {
	InitConfigWrapper()
	t := []byte("one")
	fmt.Println(len(PadID(t)), PadID(t), IDToString(PadID(t)))
	t = []byte("qwertyuiop[asdfghjklzxcvbnmqwertyuioasdfghjklzxcvvbnbmyuawteiujkhalkjsdfhalksdjhfalksdjfh")
	fmt.Println(len(PadID(t)), PadID(t), IDToString(PadID(t)))
}

func EqualID(first, second []byte) bool {
	if len(first) != len(second) {	
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func between(middle, first, second []byte) bool {
	middleInt := (&big.Int{}).SetBytes(middle)
	firstInt := (&big.Int{}).SetBytes(middle)
	secondInt := (&big.Int{}).SetBytes(middle)
	switch firstInt.Cmp(secondInt) {
		case -1 : 
			return (middleInt.Cmp(firstInt) > 0) && (middleInt.Cmp(secondInt) < 0)
		case 1 : 
			return (middleInt.Cmp(firstInt) < 0) && (middleInt.Cmp(secondInt) > 0)
		case 0 :
			return middleInt.Cmp(firstInt) != 0
	}
	return false
}
