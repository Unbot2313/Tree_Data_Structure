package utilities

import (
	"github.com/Unbot2313/models"
)

func PointerleftOrRight(NodoObjetivo *models.TreeNode, NodoPadre *models.TreeNode) int {

	//Devuelve 1 si es a la derecha, devuelve -1 si es a la izquierda
	// no se me ocurre otra forma de representar donde esta

	if NodoObjetivo.Valor < NodoPadre.Valor {

		return -1

	}

	return 1

}

func IsNodoHoja(Node *models.TreeNode) bool {
	if Node.Izquierda == nil && Node.Derecha == nil {
		return true
	}

	return false
}

func FindLowestValueInTree(TreeNode *models.TreeNode) *models.TreeNode {

	// Si ya no tiene Valores a la Izquierda es el Valor minimo y lo retorna

	if TreeNode.Izquierda == nil {
		//Retorna el noto de menor
		return TreeNode
	}

	//Programacion resursiva para encontrar el Valor minimo

	if TreeNode.Izquierda != nil {
		FindLowestValueInTree(TreeNode.Izquierda)
	}

	//Caso de que no devuelva nada

	return nil
}
