package repository

import "github.com/jackc/pgx/v5"

type CategoryRepository struct {
	DB *pgx.Conn
}

func (db *pgx.Conn) NewCategoryRepository() {
	return &CategoryRepository{DB: db}
}

// sort
// limit
// page
func (r *CategoryRepository) GetCategories() {
	// 1. GET AND PARSE ALL QUERIES
	// | id | parent_id | type     | name       |
	// | -- | --------- | -------- | ---------- |
	// | 1  | null      | country  | Kazakhstan |
	// | 2  | 1         | region   | Astana     |
	// | 3  | 2         | city     | Astana     |
	// | 4  | 3         | district | Esil       |
}
