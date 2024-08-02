package controllers

import (
	"fmt"

	"github.com/Unbot2313/models"
	"github.com/Unbot2313/utilities"
)

func InitArbol(RootNode *models.TreeNode) {

	fmt.Println(&RootNode)

	RootNode = &models.TreeNode{7, nil, nil}

	fmt.Println(&RootNode)

	Valores := []int{1, 3, 5, 8, 9, 12, 0, 32, 6, 23, 7, 9}
	for i := range Valores {
		node := &models.TreeNode{Valores[i], nil, nil}
		InsertInArbol(RootNode, node)
	}
}

func InsertInArbol(RootNode *models.TreeNode, NewNodo *models.TreeNode) {
	//si el nodo actual que se le fue pasado al coso es nulo lo inicializa

	//Menores a la Izquierda, Mayores a la Derecha
	if NewNodo.Valor < RootNode.Valor {
		//Si es nil inmediatamente lo inicializo, si no se hace asi se haria una referencia a un posible nil
		// por ende error de memoria, go no permite referencias a nil
		if RootNode.Izquierda == nil {
			RootNode.Izquierda = NewNodo
		} else {
			//si esta ocupado usa recursividad para continuar recorriendo el arbol hasta encontrar donde insertar el nodo
			InsertInArbol(RootNode.Izquierda, NewNodo)
		}

		//en otro caso a la Derecha(el otro caso es que sea igual o mayor)
	} else {

		//Si es nil inmediatamente lo inicializo, si no se hace asi se haria una referencia a un posible nil
		// por ende error de memoria, go no permite referencias a nil
		if RootNode.Derecha == nil {
			RootNode.Derecha = NewNodo
		} else {
			//si esta ocupado usa recursividad para continuar recorriendo el arbol hasta encontrar donde insertar el nodo
			InsertInArbol(RootNode.Derecha, NewNodo)
		}
	}
}

func GetOneElement(tree *models.TreeNode, Valor int) {
	if tree == nil {
		fmt.Println("nil pointer referencessssssss")
		panic(tree)
	}
	if tree.Valor == Valor {
		fmt.Println(tree)
	} else if Valor < tree.Valor {
		GetOneElement(tree.Izquierda, Valor)
	} else if Valor >= tree.Valor {
		GetOneElement(tree.Derecha, Valor)
	}
}

func DeleteNode(TreeNode *models.TreeNode, Valor int) {
	//lo que deberia hacer es, cuando este en el elemento que desee borrar, guarde su nodo anterior, y el futuro, y hacer que el anterior nodo tenga como puntero ya no el nodo
	//que se desea borrar sino el hijo de el nodo a borrar, pero hay un problema son 2

	NodeActual := TreeNode

	if NodeActual.Valor == Valor {

		//Si no tiene padre seguro es la raiz

		AdminDeleteNodo(nil, NodeActual)

	} else if Valor < NodeActual.Valor {

		NodoHijo := NodeActual.Izquierda

		if NodoHijo.Valor == Valor {

			//si el nodo hijo es lo que estoy buscando administro el borrado

			AdminDeleteNodo(NodeActual, NodoHijo)

			return
		}

		//recursividad

		DeleteNode(NodeActual.Izquierda, Valor)

	} else if Valor >= NodeActual.Valor {

		NodoHijo := NodeActual.Derecha

		if NodoHijo.Valor == Valor {

			//si el nodo hijo es lo que estoy buscando administro el borrado

			AdminDeleteNodo(NodeActual, NodoHijo)

			return
		}

		//recursividad

		DeleteNode(NodeActual.Derecha, Valor)

	}

}

func DeleteAndHasAnNode(DadNode *models.TreeNode, NodoABorrar *models.TreeNode) {

	//Si el nodo derecho es nulo entonces el izquiedo debe estar siendo usado
	//pues ya se validaron todos el resto de casos

	if NodoABorrar.Izquierda != nil && NodoABorrar.Derecha == nil {

		if NodoABorrar.Valor < DadNode.Valor {

			DadNode.Izquierda = NodoABorrar.Izquierda

			return

		}

		DadNode.Derecha = NodoABorrar.Izquierda

		return

	} else if NodoABorrar.Derecha != nil && NodoABorrar.Izquierda == nil {

		if NodoABorrar.Valor < DadNode.Valor {

			DadNode.Izquierda = NodoABorrar.Derecha

			return

		}

		DadNode.Derecha = NodoABorrar.Derecha

		return

	}

	//caso en que tenga punteros a ambos

	//ver quien es sucesor
	LowestNode := utilities.FindLowestValueInTree(NodoABorrar.Derecha)

	if NodoABorrar.Valor < DadNode.Valor {

		DadNode.Izquierda = LowestNode

		return
	}

	DadNode.Derecha = LowestNode

}

func AdminDeleteNodo(DadNode *models.TreeNode, NodoHijo *models.TreeNode) {

	if utilities.IsNodoHoja(NodoHijo) {

		if NodoHijo.Valor < DadNode.Valor {
			DadNode.Izquierda = nil
			return
		}

		DadNode.Derecha = nil

	} else {

		DeleteAndHasAnNode(DadNode, NodoHijo)

	}

}
