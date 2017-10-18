package queue

type (
	Queue struct {
		first  *node
		last   *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new queue
func Create() *Queue {
	return &Queue{nil, nil, 0}
}

// Len is length of the queue
func (queue *Queue) Len() int {
	return queue.length
}

// First : View the first item of the queue
func (queue *Queue) First() interface{} {
	if queue.length == 0 {
		return nil
	}
	return queue.first.value
}

// Last : View the last item of the queue
func (queue *Queue) Last() interface{} {
	if queue.length == 0 {
		return nil
	}
	return queue.last.value
}

// Pop the first item of the queue and return it
func (queue *Queue) Pop() interface{} {
	if queue.length == 0 {
		return nil
	}
	i := queue.first
	if queue.length == 1 {
		queue.first = nil
		queue.last = nil
	} else {
		queue.first = i.prev
	}
	queue.length--
	return i.value
}

// Push the new item to end of the queue
func (queue *Queue) Push(value interface{}) {
	i := &node{value, nil}
	if queue.length == 0 {
		queue.first = i
	} else if queue.length > 0 {
		queue.last.prev = i
	}
	queue.last = i
	queue.length++
}
