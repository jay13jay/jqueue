package jqueue

// Create new ring queue.
// accepts cap and resize as params
// cap is an int, the size of the queue
// resize is a flag, if true the queue will
// resize when enqueing while full.
// if false, the queue will overwrite elements
// in a circular fashion.
func NewRingQueue(cap int, resize bool) *ringQueue {
	return &ringQueue{
		data: make([]interface{}, cap),
		start: 0,
		end: 0,
		size: 0,
		isFull: false,
		resize: resize,
	}
}

// resize the queue takes 1 argument, m
// which is the multiple to increase the size by
func (q *ringQueue) ResizeQueue(m float64) {
    newCapacity := cap(q.data) * m
    newData := make([]interface{}, newCapacity)

    if q.start < q.end {
        copy(newData, q.data[q.start:q.end])
    } else {
        n := copy(newData, q.data[q.start:])
        copy(newData[n:], q.data[:q.end])
    }

    q.start = 0
    q.end = len(q.data)
    q.data = newData
    q.isFull = false
		q.resize = true
}