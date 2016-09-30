package tree

type Entry struct {
	Key int
	//value string
}

type node struct {
	left  *node
	right *node
	info  Entry
}

type Tree struct {
	root *node
}

func InitTree(rootEntry Entry) Tree {
	root := node{nil, nil, rootEntry}
	return Tree{&root}
}
func (t Tree) Insert(e Entry){
	t.add(t.root, &node{nil, nil, e})
}

func (t Tree) Search(e Entry) bool {
	if t.find(t.root, e) != nil {
		return true
	}
	return false
}

func (t Tree) find(n *node, e Entry) *node {
	if n == nil {
		return nil
	}
	info := n.info.Key
	if info > e.Key {
		return t.find(n.left, e)
	} else if info < e.Key {
		return t.find(n.right, e)
	} else {
		return n
	}

}

func (t Tree ) add(currentNode *node, insertedNode *node)  {
	if(currentNode == nil){
		t.root = insertedNode
		return
	}
	currentKey := currentNode.info.Key
	insertedKey := insertedNode.info.Key
	if currentKey < insertedKey {
		if(currentNode.right == nil){
			currentNode.right = insertedNode
		}else{
			t.add(currentNode.right, insertedNode)
		}
	}else{
		if(currentNode.left == nil){
			currentNode.left = insertedNode
		}else{
			t.add(currentNode.left, insertedNode)
		}
	}
}


