package main

import (
	"math/rand"
	//"fmt"
	"tree"
	//"time"
	//"time"
	//"lesson2"
	//"fmt"
	//"time"
	//"lesson2"
	"lesson3"
	"time"
)

var N = 10000000
var ITER = 10000000


func main() {
	//rootEntry := createEntry()
	//myTree := tree.InitTree(rootEntry)
	//for i:= 0; i<ITER; i++{
	//	fmt.Println("-----------------------------------")
	//	newEntry := createEntry()
	//	myTree.Insert(newEntry)
	//	fmt.Printf("New entry inserted: %d\n", newEntry.Key)
	//
	//	searchedEntry := createEntry()
	//	fmt.Printf("Entry with value: %d found ?: %t\n",searchedEntry.Key, myTree.Search(searchedEntry))
	//	//time.Sleep(1 * time.Second)
	//}
	lesson3.Lesson3()
	time.Sleep(60 * time.Second)
	lesson3.PrintA()
}

func createEntry() tree.Entry{
	return tree.Entry{Key:rand.Intn(N)}
}