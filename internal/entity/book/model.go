package book

import "github.com/alexPavlikov/Library-Manegement-System/internal/entity/genre"

type Book struct {
	UUID        string
	Name        string
	Photo       string
	Genre       []genre.Genre
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

type Author struct {
	UUID        string
	Firstname   string
	Lastname    string
	Patronymic  string
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

type Comment struct {
	Id        string
	Book_id   string
	User_id   string
	Text      string
	Time      string
	Firstname string
	Lastname  string
}
