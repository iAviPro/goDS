package main

import (
	"errors"
	"fmt"
	"sync"
)

// ###### IMPLEMENTATION USING ARRAYS ########

type queue struct {
	mux   sync.Mutex
	items []interface{}
}

// NewQueue creates a new Queue
func NewQueue() *queue {
	return &queue{mux: sync.Mutex{}, items: make([]interface{}, 0)}
}

// QPush pushes item into the queue
func (q *queue) QPush(val interface{}) {
	q.mux.Lock()
	defer q.mux.Unlock()
	q.items = append(q.items, val)
}

func (q *queue) QPop() (interface{}, error) {
	q.mux.Lock()
	defer q.mux.Unlock()

	if len(q.items) == 0 {
		return 0, errors.New("Empty Queue")
	}

	temp := q.items[0]
	q.items = q.items[1:]
	return temp, nil
}

func (q *queue) QSize() int {
	return len(q.items)
}

// QueuesDriver : Driver code to run Queues
func QueuesDriver() {
	q := NewQueue()
	q.QPush(1)
	q.QPush(2)
	q.QPush(3)
	q.QPush(4)
	fmt.Printf("Size of Queue: %d \n", q.QSize())
	v, er := q.QPop()
	if er == nil {
		fmt.Printf("First Item of Queue: %d \n", v)
	} else {
		fmt.Printf("Error: %s \n", er)
	}
	fmt.Println(q.items...)
}
