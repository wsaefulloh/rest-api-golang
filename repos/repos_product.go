package repos

import "github.com/wsaefulloh/rest-api-go/models"

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
