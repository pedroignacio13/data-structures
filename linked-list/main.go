package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedList struct {
	length int
	head   *node
}

//O(N) + O(1) time complexity
func (l *linkedList) InserirHead(value int) {
	newNode := node{data: value}
	if l.length == 0 {
		l.head = &newNode
		l.length++
		return
	}
	aux := l.head
	for aux.next != nil {
		aux = aux.next
	}
	if aux.next == nil {
		aux.next = &newNode
		l.length++
		return
	}
}

//O(1) time complexity
func (l *linkedList) TiraHead() {
	node := l.head
	prox := node.next
	for prox.next != nil {
		node = node.next
		prox = prox.next
	}
	node.next = nil
}

//O(1) time complexity
func (l *linkedList) MostraHead() {
	fmt.Println("Node de posição 0 valor:", l.head.data)
}

//O(1) time complexity
func (l *linkedList) InsereTail(val int) {
	nodeToInsert := node{data: val}
	nodeToInsert.next = l.head
	l.head = &nodeToInsert
}

//O(N) time complexity
func (l *linkedList) MostraTail() {
	pos := 0
	node := l.head
	for node.next != nil {
		node = node.next
		pos++
	}
	fmt.Println("Node de posição", pos, "valor:", node.data)
}

//O(1) time complexity
func (l *linkedList) TiraTail() {
	l.head = l.head.next
}

//O(N) time complexity
func (l *linkedList) MostraNaPosicao(pos int) {
	if pos < 0 || pos > l.length {
		fmt.Println("Posição Inválida!")
		return
	}
	cont := 0
	node := l.head
	for cont < pos {
		node = node.next
		cont++
	}
	fmt.Println("Node de valor:", node.data)
}

//O(N) + O(1) time complexity
func (l *linkedList) DeletaNaPosicao(pos int) {
	if pos < 0 || pos > l.length {
		fmt.Println("Posição Inválida!")
		return
	}
	node := l.head
	cont := 0
	for cont < pos-1 {
		node = node.next
		cont++
	}
	node.next = node.next.next
}

//O(N) + 1 time complexity
func (l *linkedList) InsereNaPosicao(val int, pos int) {
	if pos < 0 || pos > l.length {
		fmt.Println("Posição Inválida!")
		return
	}
	cont := 0
	nodeAux := l.head
	for cont < pos-1 {
		nodeAux = nodeAux.next
		cont++
	}
	nodeToInsert := node{data: val}
	nodeToInsert.next = nodeAux.next
	nodeAux.next = &nodeToInsert
}

//O(N) time complexity
func MostraLista(l *linkedList) {
	node := l.head
	for node.next != nil {
		fmt.Print(node.data, " -> ")
		node = node.next
	}
	fmt.Print(node.data)
}

// O(N) time complexity
func (l *linkedList) ProcuraValor(val int) bool {
	node := l.head
	for node.next != nil {
		if node.data == val {
			return true
		}
		node = node.next
	}
	if node.data == val {
		return true
	}
	return false
}

func main() {
	minhaLista := linkedList{}
	minhaLista.InserirHead(0)
	minhaLista.InserirHead(1)
	minhaLista.InserirHead(2)
	minhaLista.InserirHead(3)
	minhaLista.InserirHead(4)
	minhaLista.InserirHead(5)
	minhaLista.InserirHead(6)
	minhaLista.InserirHead(7)
	minhaLista.InserirHead(8)
	minhaLista.InserirHead(9)
	// MostraLista(&minhaLista)
	// minhaLista.MostraTail()
	// minhaLista.MostraHead()
	// minhaLista.MostraNaPosicao(0)
	// minhaLista.TiraHead()
	// minhaLista.TiraTail()
	// minhaLista.InsereTail(100)
	// minhaLista.DeletaNaPosicao(3)
	// fmt.Println(minhaLista.ProcuraValor(77))
	minhaLista.InsereNaPosicao(100, 3)
	MostraLista(&minhaLista)
}
