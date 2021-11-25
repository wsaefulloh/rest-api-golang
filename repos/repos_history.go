package repos

import "github.com/wsaefulloh/rest-api-go/models"

type RepoHistory interface {
	FindAll() (*models.Histories, error)
	Save(histo *models.History) error
	Remove(id string) error
	Edit(histo *models.History, id string) error
}
