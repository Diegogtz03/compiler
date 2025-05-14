package utils

import "compiler/types"

// ----------------- Stack -----------------
type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() int {
	// remove and return last item
	if s.IsEmpty() {
		return -1
	}

	var last int = s.Top()
	s.items = s.items[:len(s.items)-1]
	return last
}

func (s *Stack) Top() int {
	// return last index item
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	// returns true if length == 0 else false
	return len(s.items) == 0
}

// ----------------- Type Stack -----------------
type TypeStack struct {
	items []types.Operator
}

func (s *TypeStack) Push(data types.Operator) {
	s.items = append(s.items, data)
}

func (s *TypeStack) Pop() types.Operator {
	// remove and return last item
	if s.IsEmpty() {
		return types.ErrorOperator
	}

	var last types.Operator = s.Top()
	s.items = s.items[:len(s.items)-1]
	return last
}

func (s *TypeStack) Top() types.Operator {
	// return last index item
	return s.items[len(s.items)-1]
}

func (s *TypeStack) IsEmpty() bool {
	// returns true if length == 0 else false
	return len(s.items) == 0
}

// ----------------- Queue -----------------
type Queue struct {
	items []string
}

func (q *Queue) Enqueue(data string) {
	// Insert at back
	q.items = append(q.items, data)
}

func (q *Queue) Dequeue() string {
	// Remove from front
	if q.IsEmpty() {
		return ""
	}

	var front string = q.Head()
	q.items = q.items[1:len(q.items)]
	return front
}

func (q *Queue) Head() string {
	// Front index of array
	if q.IsEmpty() {
		return ""
	}

	return q.items[0]
}

func (q *Queue) Tail() string {
	// Last index of array
	if q.IsEmpty() {
		return ""
	}

	return q.items[len(q.items)-1]
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
