package tree
import (
	"testing"
	"fmt"
)

func TestInitTree(t *testing.T) {
	e := Entry{43}
	myTree := InitTree(e)
	if &myTree == nil {
		t.Error("Tree instance have to be initialized")
	}
}

func createTree() Tree{
	rootEntry := Entry{2}
	leftBranchEntry := Entry{1}
	rightBranchEntry := Entry{3}

	tree := InitTree(rootEntry)
	tree.Insert(leftBranchEntry)
	tree.Insert(rightBranchEntry)
	return tree
}

func TestTree_Insert(t *testing.T) {
	tree := createTree()
	if tree.root.left.info.Key != 1{
		t.Error("Error: adding in left branch is incorrect ")
	}
	if tree.root.right.info.Key != 3{
		t.Error("Error: adding in right branch is incorrect ")
	}
	fmt.Println(tree.root)
	fmt.Println(tree.root.left)
	fmt.Println(tree.root.right)

}

func TestTree_Search(t *testing.T) {
	tree := createTree()
	searchedEntry :=  Entry{1}
	notExistEntry :=  Entry{4}
	if(!tree.Search(searchedEntry)){
		t.Error("Search works incorectly")
	}
	if(tree.Search(notExistEntry)){
		t.Error("Search works incorectly")
	}
}

func BenchmarkTree_Insert(b *testing.B) {
	for i:= 0; i < b.N; i++{
		createTree()
	}
}

func BenchmarkTree_Search(b *testing.B) {
	tree := createTree()
	searchedEntry :=  Entry{3}
	notExistEntry :=  Entry{3}
	for i:= 0; i < b.N; i++{
		tree.Search(searchedEntry)
		tree.Search(notExistEntry)
	}
}

