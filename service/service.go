package service

import (
	"github.com/ivansukach/cat-service/repositories"
)

type CatService struct {
	r repositories.Repository
}

func (bs *CatService) Create(cat *repositories.Cat) error {
	return bs.r.Create(cat)
}
func (bs *CatService) Update(cat *repositories.Cat) error {
	return bs.r.Update(cat)
}
func (bs *CatService) Read(id string) (*repositories.Cat, error) {
	return bs.r.Read(id)
}
func (bs *CatService) Delete(id string) error {
	return bs.r.Delete(id)
}
func (bs *CatService) Listing() ([]repositories.Cat, error) {
	return bs.r.Listing()
}
func New(repo repositories.Repository) *CatService {
	return &CatService{r: repo}
}
