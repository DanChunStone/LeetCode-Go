package main

import "fmt"

func main()  {
	fmt.Println(isValid(""))
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	array := []byte(s)
	stack := charStack(make([]byte,0,len(s)))

	for i := 0; i < len(array); i++ {
		if array[i] - stack.top() <= 2 && array[i] - stack.top() != 0{
			stack.pop()
		} else {
			stack.push(array[i])
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}

type charStack []byte

func(s *charStack) push(b byte) {
	*s = append(*s,b)
}

func(s *charStack) pop() byte {
	if len(*s) == 0 {
		return 0
	}

	b := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return b
}

func(s *charStack) top() byte {
	if len(*s) == 0 {
		return 0
	}

	return (*s)[len(*s)-1]
}