package datastructures

type Stack []interface{}

func(s *Stack) Push(e interface{}){
	*s = append(*s, e)
}

func (s*Stack) IsEmpty() bool{
	return len(*s) == 0
}

func (s*Stack)Len()int{
	return len(*s)
}

func (s*Stack)Peek()interface{}{
	return (*s)[len(*s)-1]
}

func (s*Stack)Top()interface{}{
	if s.IsEmpty(){return nil}

	return (*s)[len(*s)-1]
}

func (s*Stack)Pop()interface{}{
	if s.IsEmpty(){return nil}

	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return x
}
