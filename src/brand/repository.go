package brand

import (
	"context"

	"github.com/hardyantz/go-clean-arch/model"
)

// Repository interface for brand repository
type Repository interface {
	GetByID(ctx context.Context, ID string) (*model.Brand, error)
}
