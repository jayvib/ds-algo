package avl

// KeyValue represents the key-value interface
type KeyValue interface {
	LessThan(KeyValue) bool
	EqualTo(KeyValue) bool
}

type TreeNode struct {
	KeyValue KeyValue
	BalanceValue int
	LinkedNodes [2]*TreeNode
}

func opposite(nodeValue int) int { // either 0 or 1
	return 1 - nodeValue
}

func singleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {
	saveNode := rootNode.LinkedNodes[opposite(nodeValue)] // the the opposite node of the linked nodeValue
	//       B
	//      / \    -R Height: 0  -L Height: 2
	//     A
	//    /
	//   C

	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	//      B
	//     /



	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}
