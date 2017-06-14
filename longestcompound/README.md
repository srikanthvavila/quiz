My first Go language program to find the longest compound word in the given list of words. Plz note that my primary focus now is to familarize with various Go language concepts. The solution used in this program is defintely not novel and may not be best optimized for compute and memory. 

The program uses trie based solution to find the longest compound word. The algorithm is:
- Create a trie with all the words from the list
- for each word in list {
-    Using trie, find if the word starts with any existing words as prefixes
-    if only one full match word is found, then it is not a compound word
-    if more than one prefix found {
-         for each such prefix, find recursively if the remaining string of 
          the word is a complete word or starts with some other existing words.
-         if last remaining string of word is a complete existing word in the 
          trie, then the initial word is a compound word
-    }
- } 

How to run this program:
1. Set your GOPATH to the absolute path of .../quiz/longestcompound folder
2. Build trie package: `go install trie`
3. Build main program: `go install longestcompound`
4. Run main program: `bin/longestcompound`
5. Input the absolute path of wordlist filename
6. The program outputs longest compound word from the list
