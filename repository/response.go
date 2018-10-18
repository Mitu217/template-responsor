package repository

import "github.com/Mitu217/template-responsor/repository/model"

type Response interface {
	Gets() ([]*model.Response, error)
	Get(uuid string) (*model.Response, error)
	Create(response *model.Response) error
	Update(response *model.Response) error
	Delete(uuid string) error
}
