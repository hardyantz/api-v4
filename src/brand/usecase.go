package brand

import (
	"context"

	"github.com/hardyantz/go-clean-arch/model"
)

// UseCase use case for brand
type UseCase interface {
	GetByID(ctx context.Context, ID string) (*model.Brand, error)
}
