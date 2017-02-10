package main

// http://www.golangbootcamp.com/book/tricks_and_tips#cid59

import (
	"log"
	"reflect"
)

// UniqStr return a copy of the passed slice with only unique string results
func UniqStr(col []string) []string {
	// map of empty structs
	m := map[string]struct{}{}

	for _, v := range col {

		// create a map with the keys being the values we want to be unique
		// the associated value does not matter much so it could be anything
		// m := map[string]bool{}
		// we chose an empty structure because it will be fast as a boolean
		// but does not allocate as much memory
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}

	list := make([]string, len(m))

	i := 0
	for v := range m {
		list[i] = v
		i++
	}

	return list
}

func main() {
	data := []struct{ in, out []string }{
		{[]string{}, []string{}},
		{[]string{"", "", ""}, []string{""}},
		{[]string{"a", "a"}, []string{"a"}},
		{[]string{"a", "b", "a"}, []string{"a", "b"}},
		{[]string{"a", "b", "a", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "b", "a", "b"}, []string{"a", "b"}},
		{[]string{"a", "a", "b", "b", "a", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "c", "a", "b", "c"}, []string{"a", "b", "c"}},
	}

	for _, exp := range data {
		res := UniqStr(exp.in)
		if !reflect.DeepEqual(res, exp.out) {
			log.Fatalf("%q didn't match %q\n", res, exp.out)
		}
	}
}
