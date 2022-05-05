package memoryhttp

import "testing"

type foo struct {
	Name string
}

func TestWatch(t *testing.T) {
	var1 := []foo{{"a"}, {"b"}, {"c"}}
	var2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Watch("var1", var1)
	Watch("var2", var2)
}
