package main

const SIZE = 10

func ReferencePassMain() {
	s := [SIZE]int{}

	for i := 0; i < SIZE; i += 1 {
		s[i] = i
	}
	ReferencePassFunction(&s)
}

func ReferencePassFunction(s *[SIZE]int) {
	for range *s {
	}
}

func ValuePassMain() {
	s := [SIZE]int{}

	for i := 0; i < SIZE; i += 1 {
		s[i] = i
	}

	ValuePassFunction(s)
}

func ValuePassFunction(s [SIZE]int) {
	for range s {
	}
}
