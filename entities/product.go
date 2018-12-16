package entities

// Product is a Moltin Product https://docs.moltin.com/catalog/products
type Product struct {
	ID            string         `json:"id,omitempty"`
	Type          string         `json:"type"`
	Name          string         `json:"name"`
	Slug          string         `json:"slug"`
	SKU           string         `json:"sku"`
	Description   string         `json:"description"`
	ManageStock   bool           `json:"manage_stock"`
	Status        string         `json:"status"`
	CommodityType string         `json:"commodity_type"`
	Price         []ProductPrice `json:"price"`
	Meta          struct {
		DisplayPrice DisplayPriceWrapper `json:"display_price"`
		Timestamps   Timestamps          `json:"timestamps,omitempty"`
		Stock        ProductStock        `json:"stock"`
	} `json:"meta,omitempty"`
}

// ProductStock is a stock object for a Products meta
type ProductStock struct {
	Level        int    `json:"level"`
	Availability string `json:"availability"`
}

// ProductPrice is a price for a Products meta
type ProductPrice struct {
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	IncludesTax bool   `json:"includes_tax"`
}

// ProductIncludes is possible includes for a Product
type ProductIncludes struct {
	Brands      []Brand      `json:"brands"`
	Categories  []Category   `json:"categories"`
	Collections []Collection `json:"collections"`
	Files       []File       `json:"files"`
	MainImage   File         `json:"main_image"`
	MainImages  []File       `json:"main_images"`
}

// SetType sets the resource type on the struct
func (p *Product) SetType() {
	p.Type = productType
}
