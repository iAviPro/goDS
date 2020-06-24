package main

// Program to implement thread safe stacks using slice and linked list.

import (
	"errors"
	"fmt"
	"sync"
)

// #### IMPLEMENTATION OF STACK USING SLICE ####

type stack struct {
	mux  sync.Mutex
	item []interface{}
}

// NewStack : to create new stack
func NewStack() *stack {
	return &stack{sync.Mutex{}, make([]interface{}, 0)}
}

func (s *stack) Push(val interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.item = append(s.item, val)
}

func (s *stack) Pop() (interface{}, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	// check for the length of the array
	l := len(s.item)
	// if array is empty
	if l == 0 {
		return 0, errors.New("Given Stack Is Empty")
	}
	// pick out the last item
	res := s.item[l-1]
	// reduce the size of the array by last item
	s.item = s.item[:l-1]
	// return value
	return res, nil
}

func (s *stack) Size() int {
	l := len(s.item)
	return l
}

// ##### IMPLEMENTATION OF STACK USING LINKED LIST #####

// Item defines the structure of single node of linkedlist
type Item struct {
	val  interface{}
	next *Item
}

type llStack struct {
	mux  sync.Mutex
	head *Item
	size int
}

func (s *llStack) ListPush(data interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()
	item := new(Item)
	item.val = data
	tmp := s.head
	item.next = tmp
	s.head = item
	s.size++
}

func (s *llStack) ListPop() (interface{}, error) {
	// null check
	if s.head == nil {
		return 0, errors.New("Empty Stack")
	}

	s.mux.Lock()
	defer s.mux.Unlock()
	tmp := s.head.val
	s.head = s.head.next
	s.size--
	return tmp, nil
}

func NewListStack() *llStack {
	stk := new(llStack)
	stk.mux = sync.Mutex{}
	return stk

}

func (s *llStack) Size() int {
	return s.size
}

// StacksDriver : Driver code to run stacks
func StacksDriver() {
	// Driver code for slice based Stack implementation
	st := NewStack()
	st.Push(0)
	st.Push(1)
	st.Push(4)
	fmt.Printf("Size of Stack: %d \n", st.Size())
	v, er := st.Pop()
	if er == nil {
		fmt.Printf("Last Item of Stack: %d \n", v)
	} else {
		fmt.Printf("Error: %s \n", er)
	}
	fmt.Println()

	//Driver code for list based Stack implementation
	llStk := NewListStack()
	llStk.ListPush(10)
	llStk.ListPush(11)
	llStk.ListPush(12)
	fmt.Printf("Size of List Stack: %d \n", llStk.Size())
	val, e := llStk.ListPop()
	if e == nil {
		fmt.Printf("Last Item of List Stack: %d \n", val)
	} else {
		fmt.Printf("Error: %s \n", e)
	}

}
