package model

type (
	CatalogEntity struct {
		Id   string
		Name string
	}

	CatalogManufacturer struct {
		Id        string
		Name      string
		BrandName string
	}

	CatalogDevice struct {
		Manufacturer string
		Model        string
		Name         string
		Categories   []string
		Protocols    []string
		Image        string
		Link         string
		Direct       []string
		Cloud        []string
	}

	PartnersCatalog struct {
		Categories    []*CatalogEntity
		Protocols     []*CatalogEntity
		Direct        []*CatalogEntity
		Cloud         []*CatalogEntity
		Manufacturers []*CatalogManufacturer
		Devices       []*CatalogDevice
	}

	PartnersCatalogFilter struct {
		Manufacturer []*CatalogEntity
		Categories   []*CatalogEntity
		Protocols    []*CatalogEntity
	}
)
