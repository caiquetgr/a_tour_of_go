package main

type Foo struct {
	x int
	S string
}

func (f Foo) Hello() string {
	return "hello"
}

func (f Foo) goodbye() string {
	return "goodbye"
}

type Bar = Foo
