package v1

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/OddEer0/partners-catalog/internal/domain"
	"github.com/OddEer0/partners-catalog/internal/domain/model"
	"github.com/OddEer0/partners-catalog/internal/domain/repository"
	v1 "github.com/OddEer0/partners-catalog/protogen"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type PartnersCatalogServer struct {
	v1.UnimplementedPartnersCatalogServiceServer
	repo repository.PartnersCatalog
}

func (p *PartnersCatalogServer) GetFilteredDevices(ctx context.Context, req *v1.CatalogDevicesReq) (*v1.FilteredDevices, error) {
	slog.Info("request", slog.Any("req", req))
	return &v1.FilteredDevices{
		Result: []*v1.FilteredDevice{
			{
				Manufacturer: "kek",
				Model:        "kekis",
				Name:         "krutoi",
				Categories:   []string{"sensor"},
				Protocols:    []string{"hub"},
				Image:        "link to image",
				Link:         "link",
				SupportsHub:  false,
				BrandId:      "hehe",
			},
			{
				Manufacturer: "kuzya",
				Model:        "kuzich",
				Name:         "meme",
				Categories:   []string{"sensor_pir"},
				Protocols:    []string{"hub"},
				Image:        "link to image",
				Link:         "link",
				SupportsHub:  true,
				BrandId:      "hoho",
			},
		},
	}, nil
}

func (p *PartnersCatalogServer) GetSearchedPartners(ctx context.Context, req *v1.CatalogPartnersReq) (*v1.Partners, error) {
	return nil, nil
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
			IsViewed:     d.IsView,
			BrandName:    d.BrandName,
		})
	}

	return res
}

func convertTransportEntityToDomain(data []PartnersCatalogEntityDTO) []*model.CatalogEntity {
	res := make([]*model.CatalogEntity, 0, len(data))
	for _, d := range data {
		res = append(res, &model.CatalogEntity{
			Id:   d.Id,
			Name: d.Name,
		})
	}
	return res
}

func convertTransportManufacturerToDomain(data []PartnersCatalogManufacturerDTO) []*model.CatalogManufacturer {
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

func convertTransportPartnersCatalogToDomain(data *PartnersCatalogDTO) *model.PartnersCatalog {
	res := &model.PartnersCatalog{
		Categories:    convertTransportEntityToDomain(data.Categories),
		Protocols:     convertTransportEntityToDomain(data.Protocols),
		Direct:        convertTransportEntityToDomain(data.Direct),
		Cloud:         convertTransportEntityToDomain(data.Cloud),
		Manufacturers: convertTransportManufacturerToDomain(data.Manufacturers),
	}

	for _, d := range data.Devices {
		var (
			isViewed bool
			brand    string
		)

		for _, ex := range d.Exceptions {
			switch ex.Id {
			case "brand_name":
				stringVal, ok := ex.Value.(string)
				if ok {
					brand = stringVal
				}
			case "is_visible_app":
				boolVal, ok := ex.Value.(bool)
				if ok {
					isViewed = boolVal
				}
			}
		}

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
			IsView:       isViewed,
			BrandName:    brand,
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

func (p *PartnersCatalogServer) SetPartnersCatalog(ctx context.Context, catalog *v1.PartnersCatalogUploadRequest) (*empty.Empty, error) {
	var dto *PartnersCatalogDTO
	err := json.Unmarshal([]byte(catalog.Json), &dto)
	if err != nil {
		slog.Error("unmarshal json error", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = dto.Validate()
	if err != nil {
		slog.Warn("bad request", slog.Any("error", err))
		return nil, status.Error(codes.InvalidArgument, "bad request")
	}

	err = p.repo.SetPartnersCatalog(ctx, convertTransportPartnersCatalogToDomain(dto))
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
