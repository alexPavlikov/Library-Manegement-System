package user

type User struct {
	Id           string
	Firstname    string
	Lastname     string
	Age          uint8
	DateOfBirth  string
	Gender       string
	Login        string
	PasswordHash string
}
