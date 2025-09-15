package database

import (
	"net/http"

	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (pS *PGService) GetCategories() ([]models.Category, utils.CashTrackError) {
	query := "SELECT id, name, color, icon, created_at, updated_at FROM categories"
	rows, err := pS.db.Query(query)
	if err != nil {
		return nil, utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Unable to fetch categories",
			Error:   err,
		}
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Id, &category.Name, &category.Color, &category.Icon, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, utils.CashTrackError{
				Code:    http.StatusInternalServerError,
				Message: "Unable to read category values",
				Error:   err,
			}
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Error while fetching category values",
			Error:   err,
		}
	}

	return categories, utils.NoErr
}
