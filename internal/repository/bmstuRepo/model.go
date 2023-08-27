package bmsturepo

type BmstuFormat struct {
	Books []struct {
		Id            int    `yaml:"id"`
		Name          string `yaml:"name"`
		AuthorName    string `yaml:"authorName"`
		AuthorSurname string `yaml:"authorSurname"`
		PageCount     int    `yaml:"pageCount"`
	} `yaml:"books"`
}
