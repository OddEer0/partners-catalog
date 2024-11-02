package v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/OddEer0/partners-catalog/internal/domain"
	"github.com/OddEer0/partners-catalog/internal/domain/model"
	"github.com/OddEer0/partners-catalog/internal/domain/repository"
	v1 "github.com/OddEer0/partners-catalog/protogen"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PartnersCatalogServer struct {
	v1.UnimplementedPartnersCatalogServiceServer
	repo repository.PartnersCatalog
}

func convertDomainEntityToTransport(data []*model.CatalogEntity) []*v1.CatalogEntity {
	res := make([]*v1.CatalogEntity, 0, len(data))
	for _, d := range data {
		res = append(res, &v1.CatalogEntity{
			Id:   d.Id,
			Name: d.Name,
		})
	}
	return res
}

func convertDomainManufacturerToTransport(data []*model.CatalogManufacturer) []*v1.Manufacturer {
	res := make([]*v1.Manufacturer, 0, len(data))
	for _, v := range data {
		res = append(res, &v1.Manufacturer{
			Id:        v.Id,
			Name:      v.Name,
			BrandName: v.BrandName,
		})
	}
	return res
}

func convertDomainPartnersCatalogToTransport(data *model.PartnersCatalog) *v1.PartnersCatalog {
	res := &v1.PartnersCatalog{
		Categories:    convertDomainEntityToTransport(data.Categories),
		Protocols:     convertDomainEntityToTransport(data.Protocols),
		Direct:        convertDomainEntityToTransport(data.Direct),
		Cloud:         convertDomainEntityToTransport(data.Cloud),
		Manufacturers: convertDomainManufacturerToTransport(data.Manufacturers),
	}

	for _, d := range data.Devices {
		res.Devices = append(res.Devices, &v1.CatalogDevice{
			Manufacturer: d.Manufacturer,
			Model:        d.Model,
			Name:         d.Name,
			Categories:   d.Categories,
			Protocols:    d.Protocols,
			Direct:       d.Direct,
			Cloud:        d.Cloud,
			Image:        d.Image,
			Link:         d.Link,
		})
	}

	return res
}

func convertTransportEntityToDomain(data []*v1.CatalogEntity) []*model.CatalogEntity {
	res := make([]*model.CatalogEntity, 0, len(data))
	for _, d := range data {
		res = append(res, &model.CatalogEntity{
			Id:   d.Id,
			Name: d.Name,
		})
	}
	return res
}

func convertTransportManufacturerToDomain(data []*v1.Manufacturer) []*model.CatalogManufacturer {
	res := make([]*model.CatalogManufacturer, 0, len(data))
	for _, v := range data {
		res = append(res, &model.CatalogManufacturer{
			Id:        v.Id,
			Name:      v.Name,
			BrandName: v.BrandName,
		})
	}
	return res
}

func convertTransportPartnersCatalogToDomain(data *v1.PartnersCatalog) *model.PartnersCatalog {
	res := &model.PartnersCatalog{
		Categories:    convertTransportEntityToDomain(data.Categories),
		Protocols:     convertTransportEntityToDomain(data.Protocols),
		Direct:        convertTransportEntityToDomain(data.Direct),
		Cloud:         convertTransportEntityToDomain(data.Cloud),
		Manufacturers: convertTransportManufacturerToDomain(data.Manufacturers),
	}

	for _, d := range data.Devices {
		res.Devices = append(res.Devices, &model.CatalogDevice{
			Manufacturer: d.Manufacturer,
			Model:        d.Model,
			Name:         d.Name,
			Categories:   d.Categories,
			Protocols:    d.Protocols,
			Direct:       d.Direct,
			Cloud:        d.Cloud,
			Image:        d.Image,
			Link:         d.Link,
		})
	}

	return res
}

func (p *PartnersCatalogServer) GetPartnersCatalog(ctx context.Context, _ *empty.Empty) (*v1.PartnersCatalog, error) {
	res, err := p.repo.GetPartnersCatalog(ctx)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return convertDomainPartnersCatalogToTransport(res), nil
}

func (p *PartnersCatalogServer) SetPartnersCatalog(ctx context.Context, catalog *v1.PartnersCatalog) (*empty.Empty, error) {
	err := catalog.ValidateAll()
	if err != nil {
		fmt.Println(err)
		return nil, status.Error(codes.InvalidArgument, "Bad Request")
	}
	err = p.repo.SetPartnersCatalog(ctx, convertTransportPartnersCatalogToDomain(catalog))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}

func (p *PartnersCatalogServer) GetPartnersCatalogFilters(ctx context.Context, empty *empty.Empty) (*v1.PartnersCatalogFilter, error) {
	res, err := p.repo.GetPartnersCatalogFilter(ctx)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	filter := &v1.PartnersCatalogFilter{
		Categories:    convertDomainEntityToTransport(res.Categories),
		Protocols:     convertDomainEntityToTransport(res.Protocols),
		Manufacturers: convertDomainEntityToTransport(res.Manufacturer),
	}

	return filter, nil
}

func NewPartnersCatalogServer(repo repository.PartnersCatalog) *PartnersCatalogServer {
	return &PartnersCatalogServer{
		repo: repo,
	}
}
