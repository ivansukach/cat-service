package repositories

type Cat struct {
	Id       string  `db:"id"`
	Name     string  `db:"name"`
	Color    string  `db:"color"`
	Gender   bool    `db:"gender"`
	Weight   float64 `db:"weight"`
	Homeless bool    `db:"homeless"`
}
type Repository interface {
	Create(cat *Cat) error
	Read(id string) (*Cat, error)
	Update(cat *Cat) error
	Delete(id string) error
	Listing() ([]Cat, error)
}
