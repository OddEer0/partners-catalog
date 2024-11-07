package v1

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	PartnersCatalogEntityDTO struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	PartnersCatalogManufacturerDTO struct {
		Id        string   `json:"id"`
		Name      string   `json:"name"`
		BrandName string   `json:"brand_name"`
		Aliases   []string `json:"aliases,omitempty"`
	}

	PartnersCatalogPartnerDTO struct {
		BrandName string   `json:"brand_name"`
		Aliases   []string `json:"aliases,omitempty"`
	}

	PartnersCatalogDeviceExceptionDTO struct {
		Id    string `json:"id"`
		Type  string `json:"type"`
		Value any    `json:"value"`
	}

	PartnersCatalogDeviceDTO struct {
		Manufacturer string                              `json:"manufacturer"`
		Model        string                              `json:"model,omitempty"`
		Name         string                              `json:"name"`
		Categories   []string                            `json:"categories"`
		Protocols    []string                            `json:"protocols"`
		Image        string                              `json:"image"`
		Link         string                              `json:"link,omitempty"`
		Direct       []string                            `json:"direct,omitempty"`
		Cloud        []string                            `json:"cloud,omitempty"`
		Exceptions   []PartnersCatalogDeviceExceptionDTO `json:"exceptions,omitempty"`
	}

	PartnersCatalogDTO struct {
		Categories    []PartnersCatalogEntityDTO       `json:"categories"`
		Protocols     []PartnersCatalogEntityDTO       `json:"protocols"`
		Manufacturers []PartnersCatalogManufacturerDTO `json:"manufacturers"`
		Partners      []PartnersCatalogPartnerDTO      `json:"partners"`
		Direct        []PartnersCatalogEntityDTO       `json:"direct"`
		Cloud         []PartnersCatalogEntityDTO       `json:"cloud"`
		Devices       []PartnersCatalogDeviceDTO       `json:"devices"`
	}
)

func (p PartnersCatalogDTO) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Devices, validation.Required),
		validation.Field(&p.Categories, validation.Required),
		validation.Field(&p.Protocols, validation.Required),
		validation.Field(&p.Partners, validation.Required),
		validation.Field(&p.Direct, validation.Required),
		validation.Field(&p.Cloud, validation.Required),
		validation.Field(&p.Manufacturers, validation.Required),
	)
}
