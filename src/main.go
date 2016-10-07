package main

import (
	"math/rand"
	"fmt"
	"tree"
	//"time"
)

var N = 1000
var ITER = 10000000


func main() {
	rootEntry := createEntry()
	myTree := tree.InitTree(rootEntry)
	for i:= 0; i<ITER; i++{
		fmt.Println("-----------------------------------")
		newEntry := createEntry()
		myTree.Insert(newEntry)
		fmt.Printf("New entry inserted: %d\n", newEntry.Key)

		searchedEntry := createEntry()
		fmt.Printf("Entry with value: %d found ?: %t\n",searchedEntry.Key, myTree.Search(searchedEntry))
		//time.Sleep(1 * time.Second)
	}
}

func createEntry() tree.Entry{
	return tree.Entry{Key:rand.Intn(N)}
}