package datastructures

type Queue []interface{}

func (q *Queue)IsEmpty()bool{
	return len(*q) == 0
}

func (q *Queue) Push(e interface{}) {
	*q = append(*q, e)
}

func (q *Queue)Pop()interface{}{
	if q.IsEmpty(){return nil}

	x := (*q)[0]
	*q = (*q)[1:]

	return x
}
