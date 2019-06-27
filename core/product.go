package core

// Product is a Moltin Product https://docs.moltin.com/catalog/products
type Product struct {
	ID            string                `json:"id,omitempty"`
	Type          string                `json:"type"`
	Name          string                `json:"name"`
	Slug          string                `json:"slug"`
	SKU           string                `json:"sku"`
	Description   string                `json:"description"`
	ManageStock   bool                  `json:"manage_stock"`
	Status        string                `json:"status"`
	CommodityType string                `json:"commodity_type"`
	Price         []ProductPrice        `json:"price"`
	Meta          *ProductMeta          `json:"meta,omitempty"`
	Relationships *ProductRelationships `json:"relationships,omitempty"`
	Included      *ProductIncludes      `json:"included,omitempty"`
}

// ProductMeta represents the meta object of a Moltin product
type ProductMeta struct {
	DisplayPrice    DisplayPriceWrapper `json:"display_price"`
	Timestamps      Timestamps          `json:"timestamps,omitempty"`
	Stock           ProductStock        `json:"stock"`
	Variations      []ProductVariation  `json:"variations,omitempty"`
	VariationMatrix interface{}         `json:"variation_matrix"`
}

// ProductVariation is a variation object for a Products Variations meta
type ProductVariation struct {
	ID      string                    `json:"id"`
	Name    string                    `json:"name"`
	Options []ProductVariationOptions `json:"options"`
}

// ProductVariationOptions is a options object for a Products ProductVariation meta
type ProductVariationOptions struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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

// ProductRelationships is the relationships object for a product
type ProductRelationships struct {
	Files      MultipleRelationship `json:"files"`
	Categories MultipleRelationship `json:"categories"`
	MainImage  SingleRelationship   `json:"main_image"`
}

// SetType sets the resource type on the struct
func (p *Product) SetType() {
	p.Type = productType
}
