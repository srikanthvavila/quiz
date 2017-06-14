package trie

import (
	"fmt"
	"strings"
)

type node struct {
	isWord   bool
	charNext [26]*node
}

type Trie struct {
	root *node
}

func (t *Trie) Init() {
	//Create and init root node
	if t.root == nil {
		t.root = &node{}
		t.initNode(t.root)
	}
}

func (t *Trie) AddWordsToTrie(wordList []string) {
	if t.root == nil {
		t.Init()
	}

	//for each word, call addWordToTrie()
	for s := range wordList {
		t.addWordToTrie(t.root, wordList[s])
	}

	return
}

//Function to find a given word and any prefixes
func (t *Trie) FindWords(word string) (bool, []string) {
	if t.root == nil {
		fmt.Println("No words in trie")
		return false, nil
	}
	foundWords := make([]string, 0, 5)
	return t.findWords(t.root, word, foundWords[:], "")
}

func (t *Trie) PrintTrie() {
	if t.root == nil {
		fmt.Println("No words in trie")
		return
	}
	t.printWords(t.root, "")
}

func (t *Trie) initNode(n *node) {
	n.isWord = false
	for i := 0; i < 26; i++ {
		n.charNext[i] = nil
	}
}

func (t *Trie) addWordToTrie(n *node, word string) {
	if len(word) == 0 {
		n.isWord = true
		return
	}
	first := strings.ToLower(word)[0]
	if n.charNext[first-'a'] == nil {
		var node2 node
		t.initNode(&node2)
		n.charNext[first-'a'] = &node2
	}
	t.addWordToTrie(n.charNext[first-'a'], word[1:len(word)])
}

func (t *Trie) printWords(n *node, word string) {
	if n.isWord == true {
		fmt.Println("word=", word)
	}
	for i := range n.charNext {
		if n.charNext[i] == nil {
			continue
		}
		alphabet := 'a' + i
		new_word := word + string(alphabet)
		t.printWords(n.charNext[i], new_word)
	}
}

func (t *Trie) findWords(n *node, word string, foundWords []string, partial string) (bool, []string) {
	if (n == nil) || (n.isWord == true) {
		foundWords = append(foundWords, partial)
	}
	if len(word) == 0 {
		return true, foundWords
	}
	first := strings.ToLower(word)[0]
	if n.charNext[first-'a'] == nil {
		return false, foundWords
	}
	partial = partial + string(first)
	found, foundWords := t.findWords(n.charNext[first-'a'], word[1:len(word)], foundWords, partial)
	return found, foundWords
}
