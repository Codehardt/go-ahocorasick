package ahocorasick

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

func ExampleMatch() {
	allStrings := []string{"a", "ab", "bc", "bd", "bbc", "ab", "aab", "bcd"}
	ac := New(allStrings)
	//ac.(*ahoCorasick).root.print('S', 0)
	for _, s := range append(allStrings) {
		fmt.Printf("----- %s -----\n", s)
		matches := ac.Match(s)
		sort.Ints(matches)
		for _, m := range matches {
			fmt.Println(allStrings[m])
		}
	}
	// Output:
	// ----- a -----
	// a
	// ----- ab -----
	// a
	// ab
	// ----- bc -----
	// bc
	// ----- bd -----
	// bd
	// ----- bbc -----
	// bc
	// bbc
	// ----- ab -----
	// a
	// ab
	// ----- aab -----
	// a
	// ab
	// aab
	// ----- bcd -----
	// bc
	// bcd
}

func TestFailLink(t *testing.T) {
	ac := New([]string{"abce", "bcd"})
	if matches := ac.Match("abcd"); !reflect.DeepEqual(matches, []int{1}) {
		t.Fatalf("wrong abcd match: %+v", matches)
	}
}

func TestOutputLink(t *testing.T) {
	ac := New([]string{"abcd", "bcd", "c", "abd"})
	matches := ac.Match("abcd")
	sort.Ints(matches)
	if !reflect.DeepEqual(matches, []int{0, 1, 2}) {
		t.Fatalf("wrong abcd match: %+v", matches)
	}
}

func (n *node) print(b byte, indent int) {
	indentStr := strings.Repeat("    ", indent)
	val := "-"
	if n.output != nil {
		val = strconv.Itoa(*n.output)
	}
	fmt.Printf("%s%s @ %x (%s) [fail: %x; output: %x]\n", indentStr, string(b), unsafe.Pointer(n), val, unsafe.Pointer(n.failLink), unsafe.Pointer(n.outputLink))
	for b, child := range n.children {
		child.print(b, indent+1)
	}
}
