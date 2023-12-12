// avl/avl_tree.go
package avltree

import "fmt"

// Node is a structure that represents a node in the AVL tree.
// Each node contains a key (Key), pointers to its left (Left) and right (Right) child nodes,
// and the height (Height) of the node in the tree.
type Node struct {
	Key         int
	Left, Right *Node
	Height      int
}

// A função height calcula e retorna a altura de um nó específico na árvore AVL.
func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// A função max retorna o maior entre dois inteiros.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// A função rotateLeft realiza uma rotação para a esquerda em um nó específico da árvore AVL.
func rotateLeft(node *Node) *Node {
	right := node.Right
	left := right.Left

	// Realiza a rotação
	right.Left = node
	node.Right = left

	// Atualiza alturas
	node.Height = max(height(node.Left), height(node.Right)) + 1
	right.Height = max(height(right.Left), height(right.Right)) + 1

	return right
}

// A função rotateRight realiza uma rotação para a direita em um nó específico da árvore AVL.
func rotateRight(node *Node) *Node {
	left := node.Left
	right := left.Right

	// Realiza a rotação
	left.Right = node
	node.Left = right

	// Atualiza alturas
	node.Height = max(height(node.Left), height(node.Right)) + 1
	left.Height = max(height(left.Left), height(left.Right)) + 1

	return left
}

// A função getBalance calcula e retorna o fator de equilíbrio de um nó específico na árvore AVL.
func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

// A função insert insere um novo nó com uma chave específica na árvore AVL e retorna a nova raiz da árvore.
func insert(root *Node, key int) *Node {
	// 1: Insira o nó como uma BST
	if root == nil {
		return &Node{Key: key, Height: 1}
	}

	if key < root.Key {
		root.Left = insert(root.Left, key)
	} else if key > root.Key {
		root.Right = insert(root.Right, key)
	} else { // Chaves iguais não são permitidas
		return root
	}

	// Passo 2: Atualizar a altura do nó atual
	root.Height = 1 + max(height(root.Left), height(root.Right))

	// Passo 3: Obter o fator de balanceamento deste nó
	balance := getBalance(root)

	// Casos de desequilíbrio e rotações
	// Caso da esquerda-esquerda
	if balance > 1 && key < root.Left.Key {
		return rotateRight(root)
	}
	// Caso da direita-direita
	if balance < -1 && key > root.Right.Key {
		return rotateLeft(root)
	}
	// Caso da esquerda-direita
	if balance > 1 && key > root.Left.Key {
		root.Left = rotateLeft(root.Left)
		return rotateRight(root)
	}
	// Caso da direita-esquerda
	if balance < -1 && key < root.Right.Key {
		root.Right = rotateRight(root.Right)
		return rotateLeft(root)
	}

	return root
}

// A função inorderTraversal percorre a árvore AVL em ordem (esquerda, raiz, direita) e imprime as chaves dos nós.
func inorderTraversal(root *Node) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Printf("%d ", root.Key)
		inorderTraversal(root.Right)
	}
}
