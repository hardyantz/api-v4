package usecase

import (
	"context"

	"github.com/hardyantz/go-clean-arch/src/brand"
	"github.com/hardyantz/go-clean-arch/model"
)

type brandUsecase struct {
	brandRepos brand.Repository
}

// NewBrandUsecase use case for brand
func NewBrandUsecase(b brand.Repository) brand.UseCase {
	return &brandUsecase{
		brandRepos: b,
	}
}

func (b *brandUsecase) GetByID(ctx context.Context, ID string) (*model.Brand, error) {
	res, err := b.brandRepos.GetByID(ctx, ID)
	return res, err
}
