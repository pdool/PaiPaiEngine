package DataStruct

// BinarySearchTree ...
type BinarySearchTree struct {
	root *treeNode
}

type treeNode struct {
	key   interface{}
	left  *treeNode
	right *treeNode
}

// Insert ...
func (b *BinarySearchTree) Insert(key interface{}) {
	newNode := treeNode{}
	newNode.key = key

	if b.root == nil {
		b.root = &newNode
	} else {
		insertNode(b.root, &newNode)
	}
}

func insertNode(node, newNode *treeNode) {
	newkey, _ := newNode.key.(int)
	key, _ := node.key.(int)
	if newkey < key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// Search ...
func (b *BinarySearchTree) Search(key interface{}) bool {
	return searchNode(b.root, key)
}

func searchNode(node *treeNode, key interface{}) bool {
	if node == nil {
		return false
	}
	nk, _ := node.key.(int)
	k, _ := key.(int)
	if k < nk {
		return searchNode(node.left, k)
	}
	if k > nk {
		return searchNode(node.right, k)
	}
	if k == nk {
		return true
	}
	return false
}

type callbackFunc func(key interface{})

// InOrderTraverse ...
func (b *BinarySearchTree) InOrderTraverse(callback callbackFunc) {
	inOrderTraverse(b.root, callback)
}

func inOrderTraverse(node *treeNode, callback callbackFunc) {
	if node != nil {
		inOrderTraverse(node.left, callback)
		callback(node.key)
		inOrderTraverse(node.right, callback)
	}
}

// PreOrderTraverse ...
func (b *BinarySearchTree) PreOrderTraverse(callback callbackFunc) {
	preOrderTraverse(b.root, callback)
}

func preOrderTraverse(node *treeNode, callback callbackFunc) {
	if node != nil {
		callback(node.key)
		preOrderTraverse(node.left, callback)
		preOrderTraverse(node.right, callback)
	}
}

// PostOrderTraverse ...
func (b *BinarySearchTree) PostOrderTraverse(callback callbackFunc) {
	postOrderTraverse(b.root, callback)
}

func postOrderTraverse(node *treeNode, callback callbackFunc) {
	if node != nil {
		postOrderTraverse(node.left, callback)
		postOrderTraverse(node.right, callback)
		callback(node.key)
	}
}

// Min ...
func (b *BinarySearchTree) Min() interface{} {
	current := b.root
	parrent := current
	for {
		if current == nil {
			return parrent.key
		}
		parrent = current
		current = current.left
	}
}

// Max ...
func (b *BinarySearchTree) Max() interface{} {
	current := b.root
	parrent := current
	for {
		if current == nil {
			return parrent.key
		}
		parrent = current
		current = current.right
	}
}

// Remove ...
func (b *BinarySearchTree) Remove() interface{} {
	return nil
}
