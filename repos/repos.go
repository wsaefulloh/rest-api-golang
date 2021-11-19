package repos

import "github.com/wsaefulloh/rest-api-go/models"

type RepoCategory interface {
	FindAll() (*models.Categories, error)
	Save(cate *models.Category) error
	Remove(id string) error
	Edit(cate *models.Category, id string) error
}

type RepoUser interface {
	FindAll() (*models.Users, error)
	Save(user *models.User) error
}

type RepoProduct interface {
	FindAll() (*models.Products, error)
	Save(prod *models.Product) error
	Remove(id string) error
	Edit(prod *models.Product, id string) error
	SearchProductName(name string) (*models.Products, error)
	SearchProductCategory(name string) (*models.Products, error)
	FindbyCategory() (*models.Products, error)
	FindbyDateDESC() (*models.Products, error)
	FindbyDateASC() (*models.Products, error)
	FindbyPriceDESC() (*models.Products, error)
	FindbyPriceASC() (*models.Products, error)
}

type RepoHistory interface {
	FindAll() (*models.Histories, error)
	Save(histo *models.History) error
	Remove(id string) error
	Edit(histo *models.History, id string) error
}
