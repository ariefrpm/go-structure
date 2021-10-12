package api

type Movie struct {
	Title    string	`json:"title"`
	Year     string	`json:"year"`
	Rated    string	`json:"rated"`
	Released string	`json:"released"`
	Runtime  string	`json:"runtime"`
	Genre    string	`json:"genre"`
	Director string	`json:"director"`
	Writer   string	`json:"writer"`
	Actors   string	`json:"actors"`
	Plot     string	`json:"plot"`
	Language string	`json:"language"`
	Country  string	`json:"country"`
	Awards   string	`json:"awards"`
	Poster   string	`json:"poster"`
	Response string	`json:"response"`
}
