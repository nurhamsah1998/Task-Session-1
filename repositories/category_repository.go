package repositories

import (
	"database/sql"
	"errors"
	"task-session-1/models"
)

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id int) (*models.Category, error)
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(id int) error
}

type categoryRepository struct {
	db *sql.DB
}

// Create implements CategoryRepository.
func (r *categoryRepository) Create(category *models.Category) error {
	query := `
		INSERT INTO products (name, description)
		VALUES ($1, $2)
		RETURNING id
	`

	return r.db.QueryRow(
		query,
		category.Name,
		category.Description,
	).Scan(&category.ID)
}

// Delete implements CategoryRepository.
func (r *categoryRepository) Delete(id int) error {
	result, err := r.db.Exec(
		`
		DELETE FROM categories
		WHERE id = $1
		`, id,
	)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}

// GetAll implements CategoryRepository.
func (r *categoryRepository) GetAll() ([]models.Category, error) {
	rows, err := r.db.Query(`
	SELECT id, name, description
		FROM categories
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []models.Category
	for rows.Next() {
		var item models.Category
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
		); err != nil {
			return nil, err
		}
		categories = append(categories, item)
	}
	return categories, nil
}

// GetByID implements CategoryRepository.
func (r *categoryRepository) GetByID(id int) (*models.Category, error) {
	var category models.Category
	err := r.db.QueryRow(`
	SELECT id, name, description
		FROM categories
		WHERE id = $1
	`, id).Scan(&category.ID, &category.Name, &category.Description)
	if err == sql.ErrNoRows {
		return nil, errors.New("category not found")
	}
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// Update implements CategoryRepository.
func (r *categoryRepository) Update(category *models.Category) error {
	result, err := r.db.Exec(
		`
		UPDATE categories
		SET name = $1,
		    description = $2
		WHERE id = $3
		`, category.Name, category.Description, category.ID,
	)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}

func NewProductRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}
