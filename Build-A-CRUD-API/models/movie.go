package models 

type Movie struct {
	ID		int		`json:"id"`
	Title		string		`json:"title"`
	Director *director 	`json:"director"`
}

type director struct {
	FirstName	string		`json:"firstname"`
	LastName	string		`json:"lastname"`
}