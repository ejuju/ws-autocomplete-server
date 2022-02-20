package suggest

import (
	"fmt"

	"github.com/ejuju/trie"
)

var t = initSuggestor() // t is a Suggestor, it returns suggestions (to complete a text) when given a string and a maximum number of results

//
type Suggestor interface {
	Suggest(str string, maxResults int) []string
}

func initSuggestor() Suggestor {
	strs, err := trie.Load("./internal/suggest/words/", "words_en.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf(">> Loaded %v strings\n", len(strs))

	t, err := trie.New(strs...)
	if err != nil {
		panic(err)
	}

	fmt.Printf(">> Initialized trie\n")

	return t
}

//
func End(str string, maxResults int) []string {
	return t.Suggest(str, maxResults)
}
