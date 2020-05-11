package main

import (
	"./student"
	"github.com/01-edu/public/go/lib"
)

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	i := 0
	for i < 30 {
		a := lib.RandInt()
		b := lib.RandInt()
		aCopy := a
		bCopy := b
		student.Swap(&a, &b)
		if a != bCopy {
			lib.Fatalf("Swap(%d, %d), a == %d instead of %d", aCopy, bCopy, a, bCopy)
		}
		if b != aCopy {
			lib.Fatalf("Swap(%d, %d), b == %d instead of %d", aCopy, bCopy, b, aCopy)
		}
		i++
	}
}