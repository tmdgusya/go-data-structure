package trie

type TrieNode struct {
	is_entry bool
	children []*TrieNode
}

type Trie struct {
	root TrieNode
}

func LetterToIndex(letter byte) int {
	return int(letter - 'a')
}

func (t *TrieNode) Insert(value string, index int) {
	if t.children == nil {
		children := make([]*TrieNode, 26)
		t.children = children
	}

	if index == len(value) {
		t.is_entry = true
	} else {
		next_letter := value[index]
		next_index := LetterToIndex(next_letter)
		next_child := t.children[next_index]
		if next_child == nil {
			t.children[next_index] = &TrieNode{}
			t.children[next_index].Insert(value, index+1)
		} else {
			next_child.Insert(value, index+1)
		}
	}
}

func (t *TrieNode) Delete(value string, index int) bool {
	if index == len(value) {
		if t.is_entry {
			t.is_entry = false
		}
	} else {
		next_letter := value[index]
		next_index := LetterToIndex(next_letter)
		next_child := t.children[next_index]

		if next_child != nil {
			if next_child.Delete(value, index+1) {
				t.children[next_index] = nil
			}
		}
	}

	if t.is_entry {
		return false
	}

	for _, c := range t.children {
		if c != nil {
			return false
		}
	}

	return true
}

func (tr *Trie) Insert(value string) {
	tr.root.Insert(value, 0)
}

func (tr *Trie) Delete(value string) {
	tr.root.Delete(value, 0)
}
