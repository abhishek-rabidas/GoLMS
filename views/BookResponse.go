package views

type BookDTO struct {
	BookName        string `json:"book_name"`
	Author          string `json:"author"`
	PublicationYear string `json:"publication_year"`
}
