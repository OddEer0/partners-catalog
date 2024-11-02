package repository

import (
	"context"
	"github.com/OddEer0/partners-catalog/internal/domain/model"
)

type PartnersCatalog interface {
	GetPartnersCatalog(ctx context.Context) (*model.PartnersCatalog, error)
	SetPartnersCatalog(ctx context.Context, catalog *model.PartnersCatalog) error
	GetPartnersCatalogFilter(ctx context.Context) (*model.PartnersCatalogFilter, error)
}
