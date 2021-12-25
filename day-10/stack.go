package main

type Stack []string

func (s *Stack) Push(value string) {
	*s = append(*s, value)
}

func (s *Stack) Pop() string {
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}
