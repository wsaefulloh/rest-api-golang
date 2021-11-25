package repos

import "github.com/wsaefulloh/rest-api-go/models"

type RepoCategory interface {
	FindAll() (*models.Categories, error)
	Save(cate *models.Category) error
	Remove(id string) error
	Edit(cate *models.Category, id string) error
}
