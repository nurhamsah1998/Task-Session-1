package services

import (
	"task-session-1/models"
	"task-session-1/repositories"
)

type CategoryService interface {
	GetAll() ([]models.Category, error)
	GetByID(id int) (*models.Category, error)
	Create(product *models.Category) error
	Update(product *models.Category) error
	Delete(id int) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

// Create implements CategoryService.
func (c *categoryService) Create(product *models.Category) error {
	return c.Create(product)
}

// Delete implements CategoryService.
func (c *categoryService) Delete(id int) error {
	return c.Delete(id)
}

// GetAll implements CategoryService.
func (c *categoryService) GetAll() ([]models.Category, error) {
	return c.GetAll()
}

// GetByID implements CategoryService.
func (c *categoryService) GetByID(id int) (*models.Category, error) {
	return c.GetByID(id)
}

// Update implements CategoryService.
func (c *categoryService) Update(product *models.Category) error {
	return c.Update(product)
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}
