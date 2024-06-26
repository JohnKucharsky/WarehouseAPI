package shared

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type (
	Pagination struct {
		Total  int32 `json:"total"`
		Limit  int   `json:"limit"`
		Offset int   `json:"offset"`
	}

	PaginationParams struct {
		Limit  string `json:"limit" validate:"omitempty,numeric"`
		Offset string `json:"offset" validate:"omitempty,numeric"`
	}

	ParsedPaginationParams struct {
		Limit  int
		Offset int
	}
)

func GetPaginationParams(c *fiber.Ctx) (*ParsedPaginationParams, error) {
	var pp PaginationParams
	if err := BindQueries(c, &pp); err != nil {
		return nil, err
	}

	limit, _ := strconv.Atoi(pp.Limit)
	offset, _ := strconv.Atoi(pp.Offset)

	if limit < 5 || limit > 100 {
		limit = 5
	}

	return &ParsedPaginationParams{Limit: limit, Offset: offset}, nil
}
