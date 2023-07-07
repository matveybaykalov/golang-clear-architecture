package clickhouserepo

type author struct {
	Id      int    `ch:"author_id"`
	Name    string `ch:"name"`
	Surname string `ch:"surname"`
}

type book struct {
	Id        int    `ch:"book_id"`
	AuthorId  int    `ch:"author_id"`
	Name      string `ch:"name"`
	PageCount int    `ch:"page_count"`
}
