package main

//arbol completo cuando todos tienen el maximo o 0 nodos hijo

import (
	"fmt"
)

//arbol binario prueba

var RootNode TreeNode

type TreeNode struct {
	valor int

	//Derecha son numeros mayores, Izquierda numeros menores
	izquierda, derecha *TreeNode
}

func main() {

	InitArbol()
	fmt.Println(1)
	GetOneElement(&RootNode, 8)
	GetOneElement(&RootNode, 9)

	// fmt.Println(r)
	// p := &r
	// fmt.Println(p)
	// o := *p
	// fmt.Println(o)

}

func InsertInArbol(tree *TreeNode, nodo *TreeNode) {
	//si el nodo actual que se le fue pasado al coso es nulo lo inicializa
	if tree == nil {
		tree = nodo
		return
	}
	//Menores a la izquierda, Mayores a la derecha
	if nodo.valor < tree.valor {
		//Si es nil inmediatamente lo inicializo, si no se hace asi se haria una referencia a un posible nil
		// por ende error de memoria, go no permite referencias a nil
		if tree.izquierda == nil {
			tree.izquierda = nodo
		} else {
			//si esta ocupado usa recursividad para continuar recorriendo el arbol hasta encontrar donde insertar el nodo
			InsertInArbol(tree.izquierda, nodo)
		}

		//en otro caso a la derecha(el otro caso es que sea igual o mayor)
	} else {

		//Si es nil inmediatamente lo inicializo, si no se hace asi se haria una referencia a un posible nil
		// por ende error de memoria, go no permite referencias a nil
		if tree.derecha == nil {
			tree.derecha = nodo
		} else {
			//si esta ocupado usa recursividad para continuar recorriendo el arbol hasta encontrar donde insertar el nodo
			InsertInArbol(tree.derecha, nodo)
		}
	}
}

func GetOneElement(tree *TreeNode, number int) {
	if tree.valor == number {
		fmt.Println(tree)
	} else if tree.valor < tree.valor {
		GetOneElement(tree.izquierda, number)
	} else if tree.valor >= tree.valor {
		GetOneElement(tree.derecha, number)
	}
}

func DeleteNode(TreeNode *TreeNode, valor int) {
	//deberia tener 3 variables, el nodo pasado, el actual y masomenos el futuro
	//lo que deberia hacer es, cuando este en el elemento que desee borrar, guarde su nodo anterior, y el futuro, y hacer que el anterior nodo tenga como puntero ya no el nodo
	//que se desea borrar sino el hijo de el nodo a borrar, pero hay un problema son 2

	if TreeNode.valor == valor {
		fmt.Println(TreeNode)
	} else if TreeNode.valor < TreeNode.valor {
		DeleteNode(TreeNode.izquierda, valor)
	} else if TreeNode.valor >= TreeNode.valor {
		DeleteNode(TreeNode.derecha, valor)
	}

}

func FindLowestValueInTree(TreeNode *TreeNode) *TreeNode {

	// Si ya no tiene valores a la izquierda es el valor minimo y lo retorna

	if TreeNode.izquierda == nil {
		//Retorna el noto de menor
		return TreeNode
	}

	//Programacion resursiva para encontrar el valor minimo

	if TreeNode.izquierda != nil {
		FindLowestValueInTree(TreeNode.izquierda)
	}

	//Caso de que no devuelva nada

	return nil
}

func IsNodoHoja(Node TreeNode) bool {
	if Node.izquierda == nil && Node.derecha == nil {
		return true
	}

	return false
}

func InitArbol() {
	RootNode = TreeNode{7, nil, nil}
	valores := []int{1, 3, 5, 8, 9, 12, 0, 32, 6, 23, 7, 9}
	for i := range valores {
		node := &TreeNode{valores[i], nil, nil}
		InsertInArbol(&RootNode, node)
	}
}

//no borrar
