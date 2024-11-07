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

	CatalogDeviceException struct {
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
		BrandName    string
		IsView       bool
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
