package book

type Book struct {
	UUID        string
	Name        string
	Photo       string
	Genre       []string
	Year        uint16
	Pages       uint16
	Rating      uint8
	Description string
	PDFLink     string
	Publishing  Publishing
	Authors     []Author
	Awards      []Awards
	Deleted     bool
}

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

type Publishing struct {
	UUID string
	Name string
}

type Awards struct {
	UUID          string
	Name          string
	YearOfReceipt uint16
}
