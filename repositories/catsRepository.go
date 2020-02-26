package repositories

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func New(db *sqlx.DB) Repository {
	return &catsRepository{db: db}
}

type catsRepository struct {
	db *sqlx.DB
}

func (br *catsRepository) Create(cat *Cat) error {
	_, err := br.db.NamedExec("INSERT INTO cats VALUES (:id, :name, :color, :gender, :weight, :homeless)", cat)
	return err
}
func (br *catsRepository) Read(id string) (*Cat, error) {
	c := Cat{}
	err := br.db.QueryRowx("SELECT * FROM cats WHERE id=$1", id).StructScan(&c)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &c, err
}
func (br *catsRepository) Update(cat *Cat) error {
	_, err := br.db.NamedExec("UPDATE cats SET Name=:name, Color=:color, Gender=:gender, Weight=:weight, "+
		"Homeless=:homeless WHERE Id=:id", cat)
	return err
}
func (br *catsRepository) Delete(id string) error {
	_, err := br.db.Exec("DELETE FROM cats WHERE id=$1", id)
	return err
}
func (br *catsRepository) Listing() ([]Cat, error) {
	rows, err := br.db.Queryx("SELECT Id, Name, Color, Gender, Weight, Homeless FROM cats") //SELECT *?
	if err != nil {
		log.Warning(err)
		return nil, err
	}
	c := make([]Cat, 0)
	for rows.Next() {
		cat := new(Cat)
		err = rows.StructScan(&cat)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		c = append(c, *cat)
	}
	return c, err
}
