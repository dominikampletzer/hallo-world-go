package hello

import (
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	var page1 = []string{"a", "b", "c"}
	var page2 = []string{"b", "c"}
	var page3 = []string{"a"}
	var page4 = []string{"c"}
	var page5 = []string{"a", "c"}
	var page6 = []string{"a", "b", "c"}

	var book = Book{page1, page2, page3, page4, page5, page6}
	var index = CreateIndex(book)
	fmt.Println(index)
}
