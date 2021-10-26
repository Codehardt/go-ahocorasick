package ahocorasick

// AhoCorasick is an interface that returns all matching strings in a text. Use New() to initialize a new AhoCorasick interface.
type AhoCorasick interface {
	// Match returns all indices of strings that were found in the passed text
	Match(text string) []int
}

// Match is the interface implementation of AhoCorasick's Match function
func (n *node) Match(text string) (res []int) {
	// find all matches and write them in a map
	resMap := make(map[int]struct{})
	n.find(text, resMap)
	// convert the matches to a slice
	res = make([]int, len(resMap))
	var i int
	for elem := range resMap {
		res[i] = elem
		i++
	}
	return
}

// New builds a new AhoCorasick interface
func New(allStrings []string) AhoCorasick {
	// generate the root node for the trie
	ac := new(node)
	ac.children = make(map[byte]*node)
	// generate the trie
	for i, s := range allStrings {
		ac.add(i, s)
	}
	// generate fail links
	ac.failLink = ac
	setFailLinks(ac, ac, true)
	// generate output links
	setOutputLinks(ac)
	return ac
}

// node is a node in the trie
type node struct {
	// children is a map of children based on it's byte
	children map[byte]*node
	// if this node is the end of a string, then the output points to the string's index
	output *int
	// if no children with a specific byte was found, use the fail link
	failLink *node
	// output links are needed to find substrings
	outputLink *node
	// outputLinkSet is used to prevent multiple calculations of the above output link
	outputLinkSet bool
}

// add adds a new string with an index i to the trie
func (n *node) add(i int, s string) {
	// add the first/next character to the trie
	child := n.children[s[0]]
	if child == nil {
		// if the character was not found yet in the trie's path, add a new node
		child = new(node)
		child.children = make(map[byte]*node)
		n.children[s[0]] = child
	}
	s = s[1:]
	if s == "" {
		// if we add all characters of the string, mark the last node as output node
		child.output = &i
	} else {
		// add the next character to trie recursively
		child.add(i, s)
	}
}

// setFailLinks links all nodes with a fail link
func setFailLinks(root *node, n *node, isRoot bool) {
	// iterate over all children that need a fail link
	for b, child := range n.children {
		if isRoot {
			// the nodes with depth 1 (with root as parent) always have to get the root as fail link
			child.failLink = root
		} else {
			// otherwise the failover is either:
			// - the child node of the fail link of the parent node
			// - the root node
			child.failLink = getFailLink(root, n, b)
		}
		// recursively set fail links of children
		setFailLinks(root, child, false)
	}
}

// getFailLink returns the fail link for a children
func getFailLink(root *node, parent *node, b byte) *node {
	if failChild, ok := parent.failLink.children[b]; ok {
		// if the parent node has a fail link and the fail link has a valid children node for the specific byte,
		// use this child node as fail link
		return failChild
	}
	if failChild, ok := root.children[b]; ok {
		// if the root node has a valid children node for the specific byte, use this child node as fail link
		return failChild
	}
	// otherwise use the root node as fail link
	return root
}

// setOutputLinks sets the output links for all nodes
func setOutputLinks(n *node) {
	for _, child := range n.children {
		// set the output link for the specific child node
		setOutputLink(child)
		// set the output link for the children of the specific child node recursively
		setOutputLinks(child)
	}
}

// setOutputLink sets the output link of the specific node (and returns the set output link for internal usages)
func setOutputLink(n *node) *node {
	if !n.outputLinkSet {
		// only search for a output link if not already set
		n.outputLinkSet = true
		if n.failLink.output != nil {
			// if our fail link has an output, then the fail link is also the output link
			n.outputLink = n.failLink
		} else {
			// otherwise we have to use the output link of our fail link
			n.outputLink = setOutputLink(n.failLink)
		}
	}
	return n.outputLink
}

// find searches for a string in the trie and writes all string indices that were found in the res map
func (n *node) find(s string, res map[int]struct{}) {
	if n.output != nil {
		// if the current node is an output node, use the output as our result
		res[*n.output] = struct{}{}
	}
	// follow the output link to find more results
	n.followOutputLink(res)
	if s == "" {
		// no string left: the end of the recursion
		return
	}
	if child, ok := n.children[s[0]]; ok {
		// recursively search for the string in the child node
		child.find(s[1:], res)
	} else if n.failLink != n {
		// no child node found, search for more results in the fail link instead
		n.failLink.find(s, res)
	} else {
		n.failLink.find(s[1:], res)
	}
	return
}

// followOutputLink follows all output links that are linked with the specific node recursively
func (n *node) followOutputLink(res map[int]struct{}) {
	if n.outputLink != nil {
		// if this node has an output link, add the output of the output link to the result and
		// recursively search for more results
		res[*n.outputLink.output] = struct{}{}
		n.outputLink.followOutputLink(res)
	}
}
