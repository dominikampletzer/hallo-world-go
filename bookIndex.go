package hello

type Page []string

type Book []Page

type Index map[string][]int

func CreateIndex(book Book) Index {
	var index = make(Index)

	for idx, page := range book {
		for _, str := range page {
			index[str] = append(index[str], idx)

		}
	}
	return index
}
