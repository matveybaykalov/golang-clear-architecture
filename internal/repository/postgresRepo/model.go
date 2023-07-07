package postgresrepo

type author struct {
	Id      int    `bun:"author_id"`
	Name    string `bun:"name"`
	Surname string `bun:"surname"`
}

type book struct {
	Id        int    `bun:"book_id"`
	AuthorId  int    `bun:"author_id"`
	Name      string `bun:"name"`
	PageCount int    `bun:"page_count"`
}
