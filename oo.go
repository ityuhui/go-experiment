package main

import "fmt"

type A struct {
	m_a int
}

func (a *A) print() {
	fmt.Println("---a print---")
	fmt.Println(a.m_a)
}

type B struct {
	A
	m_b string
}

func (b *B) print() {
	fmt.Println("---b print---")
	println(b.m_a)
	println(b.m_b)
}

type I interface {
	print()
}

func testoo() {
	a := new(A)
	a.m_a = 1

	b := new(B)
	b.m_a = 2
	b.m_b = "b"

	a.print()
	b.print()

}
