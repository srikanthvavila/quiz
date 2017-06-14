package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
        "log"
	"trie"
)

func main() {
	fmt.Println("Longest compound word search implementation in go")

        //Read the filename from console
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter wordlist filename (full path): ")
	filename, err := reader.ReadString('\n')
        if err != nil {
                log.Fatal(err)
        }

        //Remove any new line characters from the filename
	re := regexp.MustCompile(`\n`)
	filename = re.ReplaceAllString(filename, "")

        //Open the wordlist file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

        //Read the words from wordlist file
	var words []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
                word := scanner.Text()
                isLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(word)
                if isLetter == false {
                         log.Fatalf("Non alphabet words %s are not supported\n",word)
                }
		words = append(words, word)
	}

        //Create a Trie with the words
	var t trie.Trie
	t.Init()
	t.AddWordsToTrie(words)

        maxlen, maxidx := 0,0

	for i := range words {
		compound, _ := isCompoundWord(&t, words[i])
                if (compound == true) && (maxlen < len(words[i])) {
                        maxlen = len(words[i])
                        maxidx = i 
                }
	}
        if maxlen > 0 {
                fmt.Printf("The Longest CompoundWord in the list: %s\n", words[maxidx])
        }
}

//Returns whether a given word exists in the Trie and is a compound word or not
func isCompoundWord(t *trie.Trie, word string) (bool, bool) {
	found1, foundWords := t.FindWords(word)
	if len(foundWords) == 0 {
		return false, false
	}
	if (len(foundWords) == 1) && (strings.Compare(foundWords[0], word) == 0) {
		return false, true
	}
	for s := range foundWords {
		if strings.Compare(foundWords[s], word) == 0 {
			continue
		}
		w2 := strings.TrimPrefix(word, foundWords[s])
		_, found2 := isCompoundWord(t, w2)
		if found2 == true {
			return true, true
		}
	}
	return false, found1
}
