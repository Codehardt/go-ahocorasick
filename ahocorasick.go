package ahocorasick

import (
	"sort"
)

// AhoCorasick is an interface that returns all matching strings in a text. Use New() to initialize a new AhoCorasick interface.
type AhoCorasick interface {
	// Match returns all indices of strings that were found in the passed text
	Match(text string) []int
}

type ahoCorasick struct {
	root *node
}

// Match is the interface implementation of AhoCorasick's Match function
func (a *ahoCorasick) Match(text string) (res []int) {
	resMap := make(map[int]struct{})
	a.root.find(text, resMap)
	res = make([]int, len(resMap))
	var i int
	for elem := range resMap {
		res[i] = elem
		i++
	}
	sort.Ints(res)
	return
}

// New builds a new AhoCorasick interface.
func New(allStrings []string) AhoCorasick {
	ac := &ahoCorasick{root: new(node)}
	ac.root.children = make(map[byte]*node)
	ac.root.failLink = ac.root
	for i, s := range allStrings {
		ac.root.add(i, s)
	}
	setFailLinks(ac.root, ac.root, true)
	setOutputLinks(ac.root)
	return ac
}

type node struct {
	children      map[byte]*node
	output        *int
	failLink      *node
	outputLink    *node
	outputLinkSet bool
}

func (n *node) add(i int, s string) {
	child := n.children[s[0]]
	if child == nil {
		child = new(node)
		child.children = make(map[byte]*node)
		n.children[s[0]] = child
	}
	s = s[1:]
	if s == "" {
		child.output = &i
	} else {
		child.add(i, s)
	}
}

func setFailLinks(root *node, n *node, isRoot bool) {
	for b, child := range n.children {
		if isRoot {
			child.failLink = root
		} else {
			child.failLink = getFailLink(root, n, b)
		}
		setFailLinks(root, child, false)
	}
}

func getFailLink(root *node, parent *node, b byte) *node {
	if failChild, ok := parent.failLink.children[b]; ok {
		return failChild
	}
	return root
}

func setOutputLinks(n *node) {
	for _, child := range n.children {
		setOutputLink(child)
		setOutputLinks(child)
	}
}

func setOutputLink(n *node) *node {
	if !n.outputLinkSet {
		n.outputLinkSet = true
		if n.failLink.output != nil {
			n.outputLink = n.failLink
		} else {
			n.outputLink = setOutputLink(n.failLink)
		}
	}
	return n.outputLink
}

func (n *node) find(s string, res map[int]struct{}) {
	if n.output != nil {
		res[*n.output] = struct{}{}
	}
	n.followOutputLink(res)
	if s == "" {
		return
	}
	if child, ok := n.children[s[0]]; ok {
		child.find(s[1:], res)
	} else {
		n.failLink.find(s, res)
	}
	return
}

func (n *node) followOutputLink(res map[int]struct{}) {
	if n.outputLink != nil {
		res[*n.outputLink.output] = struct{}{}
		n.outputLink.followOutputLink(res)
	}
}
