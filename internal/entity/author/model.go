package author

type Author struct {
	UUID        string
	Firstname   string
	Lastname    string
	Photo       string
	BirthPlace  string
	Age         uint8
	DateOfBirth string
	DateOfDeath string
	Gender      string
	Books       []Book
	Awards      []Awards
}

type Book struct {
	UUID        string
	Name        string
	Photo       string
	Genre       []Genre
	Year        uint16
	Pages       uint16
	Rating      []Rating
	Description string
	PDFLink     string
	Publishing  Publishing
	Authors     []Author
	Awards      []Awards
	Deleted     bool
}

type Awards struct {
	UUID          string
	Name          string
	YearOfReceipt uint16
}

type Publishing struct {
	UUID string
	Name string
}
type Rating struct {
	Id      string
	Book_id string
	User_id string
	Rating  uint8
}

type Genre struct {
	Id   string
	Name string
	Link string
}
