package repository

import (
	"context"
	"fmt"

	"github.com/hardyantz/go-clean-arch/src/brand"
	"github.com/hardyantz/go-clean-arch/model"
	driver "github.com/arangodb/go-driver"
)

type arangoBrandRepo struct {
	DB driver.Database
}

// NewArangoBrandRepository brand repository arango handler
func NewArangoBrandRepository(db driver.Database) brand.Repository {
	return &arangoBrandRepo{
		DB: db,
	}
}

// GetById get brand by id repository
func (db *arangoBrandRepo) GetByID(ctx context.Context, ID string) (*model.Brand, error) {
	query := fmt.Sprintf(`FOR b IN brands FILTER b._key == "%s" LIMIT 1 RETURN b`, ID)
	cursor, err := db.DB.Query(ctx, query, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var doc *model.Brand

	for {
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}
	}

	return doc, nil
}
