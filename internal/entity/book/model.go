package book

type Book struct {
	UUID       string
	Name       string
	Genre      []string
	Year       uint16
	Publishing Publishing
	Authors    []Author
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
