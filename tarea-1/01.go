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
// Uses already implemented library

func main() {
	// Stack Test Cases
	fmt.Println("--------- STACK ---------")
	s := Stack{}

	s.push(1)
	fmt.Println("TOP: Should be 1 is:", s.top())

	s.push(2)
	fmt.Println("TOP: Should be 2 is:", s.top())

	s.push(3)
	fmt.Println("TOP: Should be 3 is:", s.top())

	fmt.Println("POP: Should be 3 is:", s.pop())
	fmt.Println("TOP: Should be 2 is:", s.top())

	fmt.Println("IsEmpty: Should be false is:", s.isEmpty())

	fmt.Println("POP: Should be 2 is:", s.pop())
	fmt.Println("POP: Should be 1 is:", s.pop())

	fmt.Println("IsEmpty: Should be true is: ", s.isEmpty())

	// Queue Test Cases
	fmt.Println("--------- QUEUE ---------")
	q := Queue{}

	q.enqueue(1)
	fmt.Println("HEAD: Should be 1 is:", q.head())
	fmt.Println("TAIL: Should be 1 is:", q.tail())

	q.enqueue(2)
	fmt.Println("HEAD: Should be 1 is:", q.head())
	fmt.Println("TAIL: Should be 2 is:", q.tail())

	q.enqueue(3)
	fmt.Println("HEAD: Should be 1 is:", q.head())
	fmt.Println("TAIL: Should be 3 is:", q.tail())

	fmt.Println("DEQUEUE: Should be 1 is:", q.dequeue())
	fmt.Println("HEAD: Should be 2 is:", q.head())

	fmt.Println("DEQUEUE: Should be 2 is:", q.dequeue())
	fmt.Println("HEAD: Should be 3 is:", q.head())
	fmt.Println("TAIL: Should be 3 is:", q.tail())

	q.enqueue(4)

	fmt.Println("TAIL: Should be 4 is:", q.tail())

	fmt.Println("IsEmpty: Should be false is:", q.isEmpty())

	q.dequeue()
	q.dequeue()

	fmt.Println("Length: Should be 0 is:", len(q.items))

	fmt.Println("IsEmpty: Should be true is:", q.isEmpty())

	// Dictionary Test Cases
	fmt.Println("--------- DICTIONARY ---------")

	d := make(map[string]int)

	d["Diego"] = 22
	d["Roberto"] = 21

	fmt.Println("Should be 21 is:", d["Roberto"])
	fmt.Println("Should be 2 is:", len(d))

	d["Roberto"] = 32
	fmt.Println("Should be 32 is:", d["Roberto"])

	delete(d, "Roberto")
	fmt.Println("Should be 1 is:", len(d))

	d["Jose"] = 20

	fmt.Println("Should be 20 is:", d["Jose"])
	fmt.Println("Should be 2 is:", len(d))

	fmt.Println("Should be 22 is:", d["Diego"])
}
