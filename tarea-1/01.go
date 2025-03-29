package main

import "fmt"

// ----------------- Stack -----------------
type Stack struct {
	items []int
}

func (s *Stack) push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) pop() int {
	// remove and return last item
	if s.isEmpty() {
		return -1
	}

	var last int = s.top()
	s.items = s.items[:len(s.items)-1]
	return last
}

func (s *Stack) top() int {
	// return last index item
	return s.items[len(s.items)-1]
}

func (s *Stack) isEmpty() bool {
	// returns true if length == 0 else false
	return len(s.items) == 0
}

// ----------------- Queue -----------------
type Queue struct {
	items []int
}

func (q *Queue) enqueue(data int) {
	// Insert at back
	q.items = append(q.items, data)
}

func (q *Queue) dequeue() int {
	// Remove from front
	if q.isEmpty() {
		return -1
	}

	var front int = q.head()
	q.items = q.items[1:len(q.items)]
	return front
}

func (q *Queue) head() int {
	// Front index of array
	if q.isEmpty() {
		return -1
	}

	return q.items[0]
}

func (q *Queue) tail() int {
	// Last index of array
	if q.isEmpty() {
		return -1
	}

	return q.items[len(q.items)-1]
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

// ----------------- Dictionary -----------------

func main() {
	// Stack Test Cases
	fmt.Println("--------- STACK ---------")
	s := Stack{}

	s.push(1)
	s.push(2)
	s.push(3)

	fmt.Println(s.top())
	fmt.Println(s.pop())
	fmt.Println(s.top())

	fmt.Println(s.isEmpty())

	fmt.Println(s.pop())
	fmt.Println(s.pop())

	fmt.Println(s.isEmpty())

	// Queue Test Cases
	fmt.Println("--------- QUEUE ---------")
	q := Queue{}

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)

	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())

	fmt.Println(q.tail())

	q.enqueue(4)

	fmt.Println(q.tail())

	// Dictionary Test Cases
	fmt.Println("--------- DICTIONARY ---------")

	d := make(map[string]int)

	d["Diego"] = 22
	d["Roberto"] = 21
	d["Jose"] = 20

	fmt.Println(d["Roberto"])
	fmt.Println(len(d))
	delete(d, "Roberto")
	fmt.Println(len(d))
}
