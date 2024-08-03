package entity

type Author struct {
	Id          uint
	Name        string
	PhotoUrl    string
	Gender      string
	AuthorGenre *AuthorGenre
}
