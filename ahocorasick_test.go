package ahocorasick

import "fmt"

func ExampleMatch() {
	allStrings := []string{"a", "ab", "bc", "bd", "bbc", "ab", "aab", "bcd"}
	ac := New(allStrings)
	for _, s := range append(allStrings, "cd") {
		fmt.Printf("----- %s -----\n", s)
		for _, m := range ac.Match(s) {
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
	// bbc
	// ----- ab -----
	// a
	// ab
	// ----- aab -----
	// a
	// aab
	// ----- bcd -----
	// bc
	// bcd
	// ----- cd -----
}
