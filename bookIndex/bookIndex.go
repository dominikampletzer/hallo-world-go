package bookIndex

import "fmt"

type Page []string

type Book []Page

type Index map[string][]int

func CreateIndex(book Book) Index {
	index := make(Index)

	for idx, page := range book {
		for _, word := range page {
			index[word] = append(index[word], idx)

		}
	}
	return index
}

func (i Index) String() string {
	result := ""
	for k, v := range i {
		result += fmt.Sprintf("\n\tWord: %v : Pages: %v", k, v)
	}
	return result + "\n"
}
