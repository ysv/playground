package main

import "fmt"

func main() {
	var nilBar *Bar
	var barWithNilFoo = &Bar{foo: nil}
	var bar = &Bar{foo: &Foo{value: 1}}

	fmt.Printf("nilBar.GetFoo().GetValue(): %v\n", nilBar.GetFoo().GetValue())
	fmt.Printf("barWithNilFoo.GetFoo().GetValue(): %v\n", barWithNilFoo.GetFoo().GetValue())
	fmt.Printf("bar.GetFoo().GetValue(): %v\n", bar.GetFoo().GetValue())
}

type Foo struct {
	value int
}

type Bar struct {
	foo *Foo
}

func (f *Foo) GetValue() int {
	if f == nil {
		return 0
	}

	return f.value
}

func (b *Bar) GetFoo() *Foo {
	if b == nil {
		return nil
	}

	return b.foo
}
