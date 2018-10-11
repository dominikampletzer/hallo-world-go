package hello

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
	return "TODO implement Strin gfor BookIndex"
}
