package main

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"strings"
	"words/model"
)

// ByLength implements sort.Interface for []WordLen based on
// the Length field.
type ByLength []model.WordLen

func (a ByLength) Len() int           { return len(a) }
func (a ByLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLength) Less(i, j int) bool { return a[i].Length < a[j].Length }

func main() {
	content, err := os.ReadFile("input/words.json")
	if err != nil {
		log.Fatal(err)
	}
	words := []model.WordsStruct{}
	err = json.Unmarshal(content, &words)

	if err != nil {
		log.Fatal(err)
	}

	wordLens := []model.WordLen{}

	for i := 0; i < len(words); i++ {
		list_words := strings.Split(words[i].Words, " ")
		for j := 0; j < len(list_words); j++ {
			w_len := len(string(list_words[j]))
			log.Default().Println("Word: ", list_words[j], " Num_Of_Characters: ", w_len)
			wordLens = append(wordLens, model.WordLen{Word: list_words[j], Length: w_len})
			if w_len%2 == 1 {
				log.Default().Println("Word: ", list_words[j], " Num_Of_Characters: ", w_len)
				wordLens = append(wordLens, model.WordLen{Word: list_words[j], Length: w_len})
			}
		}

	}
	sort.Sort(ByLength(wordLens))
	log.Default().Println("Ascending order:", wordLens)
}
