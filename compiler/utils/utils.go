package utils

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
