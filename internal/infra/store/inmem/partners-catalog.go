package inmem

import (
	"context"
	"github.com/OddEer0/partners-catalog/internal/domain"
	"github.com/OddEer0/partners-catalog/internal/domain/model"
	"sync"
)

type PartnersCatalog struct {
	catalog *model.PartnersCatalog
	mu      sync.Mutex
}

func (p *PartnersCatalog) RemovePartnersCatalog(ctx context.Context) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.catalog = nil
	return nil
}

func (p *PartnersCatalog) GetPartnersCatalog(ctx context.Context) (*model.PartnersCatalog, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.catalog == nil {
		return nil, domain.ErrNotFound
	}

	return p.catalog, nil
}

func (p *PartnersCatalog) SetPartnersCatalog(ctx context.Context, catalog *model.PartnersCatalog) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.catalog = catalog
	return nil
}

func convertManufacturerToEntities(manufacturers []*model.CatalogManufacturer) []*model.CatalogEntity {
	result := make([]*model.CatalogEntity, 0)
	for _, manufacturer := range manufacturers {
		result = append(result, &model.CatalogEntity{
			Id:   manufacturer.Id,
			Name: manufacturer.Name,
		})
	}
	return result
}

func (p *PartnersCatalog) GetPartnersCatalogFilter(ctx context.Context) (*model.PartnersCatalogFilter, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.catalog == nil {
		return nil, domain.ErrNotFound
	}

	result := &model.PartnersCatalogFilter{
		Categories:   p.catalog.Categories,
		Protocols:    p.catalog.Protocols,
		Manufacturer: convertManufacturerToEntities(p.catalog.Manufacturers),
	}

	return result, nil
}

func NewPartnersCatalog() *PartnersCatalog {
	return &PartnersCatalog{
		mu: sync.Mutex{},
	}
}
