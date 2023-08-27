package msurepo

type MsuFormat struct {
	Authors struct {
		Author []struct {
			Id      int    `xml:"id"`
			Name    string `xml:"name"`
			Surname string `xml:"surname"`
		} `xml:"author"`
	} `xml:"authors"`
	Books struct {
		Book []struct {
			Id        int    `xml:"id"`
			AuthorId  int    `xml:"authorId"`
			Name      string `xml:"name"`
			PageCount int    `xml:"pageCount"`
		} `xml:"book"`
	} `xml:"books"`
}
