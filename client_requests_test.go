package gomo_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/moltin/gomo"
	"github.com/moltin/gomo/entities"
)

type testServer struct {
	expectMethod string
	expectBody   string
	responseCode int
	response     string
	called       bool
}

func testJSON(t *testing.T, name string, got, expected []byte) {
	var g, e interface{}
	var err error
	err = json.Unmarshal(got, &g)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(expected, &e)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(g, e) {
		return
	}
	var gb, eb bytes.Buffer
	err = json.Indent(&gb, got, "", "    ")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Indent(&eb, expected, "", "    ")
	if err != nil {
		t.Fatal(err)
	}
	t.Errorf(
		"unexpected %s\nexpected:\n%s\ngot:\n%s\n",
		name,
		eb.String(),
		gb.String(),
	)
}

func (ts *testServer) Start(t *testing.T) (gomo.Client, func()) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		ts.called = true
		if r.Method != ts.expectMethod {
			t.Errorf("unexpected method: %s", r.Method)
		}
		if ts.expectBody != "" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			testJSON(t, "body", body, []byte(ts.expectBody))
		}
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(ts.responseCode)
		w.Write([]byte(ts.response))
	}
	s := httptest.NewServer(http.HandlerFunc(handler))
	c := gomo.NewClient(
		gomo.ClientCredentials("id", "secret"),
		gomo.Endpoint(s.URL),
	)
	done := func() {
		s.Close()
	}
	return c, done
}

func TestGetProductWithError(t *testing.T) {
	test := testServer{
		expectMethod: "GET",
		responseCode: http.StatusInternalServerError,
		response: `
{
    "errors": [
        {
	    "title": "Oops",
            "detail": "Something is wrong"
        }
    ]
}
		`,
	}
	client, done := test.Start(t)
	defer done()

	_, err := client.Get(
		"products/9eda5ba0-4f4a-4074-8547-ccb05d1b5981",
	)
	if err == nil {
		t.Fatal("expected an error")
	}
	if !test.called {
		t.Fatal("server not called")
	}
	if err.Error() != "Oops: Something is wrong" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetProductWithFlows(t *testing.T) {
	test := testServer{
		expectMethod: "GET",
		responseCode: http.StatusOK,
		response: `
{
    "data": {
        "type": "product",
        "id": "9eda5ba0-4f4a-4074-8547-ccb05d1b5981",
        "name": "Crown",
        "slug": "crown",
        "sku": "CWLP100BLK",
        "manage_stock": true,
        "description": "Abstract, sculptural, refined and edgy with a modern twist. Its symmetrical, spoked structure generates a clever geometric presence, which works well in a contemporary environment.",
        "price": [
            {
                "amount": 47500,
                "currency": "USD",
                "includes_tax": true
            }
        ],
        "status": "live",
        "commodity_type": "physical",
        "meta": {
            "timestamps": {
                "created_at": "2017-06-19T14:58:42+00:00",
                "updated_at": "2018-04-10T09:12:05+00:00"
            },
            "display_price": {
                "with_tax": {
                    "amount": 47500,
                    "currency": "USD",
                    "formatted": "$475.00"
                },
                "without_tax": {
                    "amount": 47500,
                    "currency": "USD",
                    "formatted": "$475.00"
                }
            },
            "stock": {
                "level": 500,
                "availability": "in-stock"
            },
            "variation_matrix": []
        },
        "relationships": {
            "files": {
                "data": [
                    {
                        "type": "file",
                        "id": "7cc08cbb-256e-4271-9b01-d03a9fac9f0a"
                    }
                ]
            },
            "categories": {
                "data": [
                    {
                        "type": "category",
                        "id": "a636c261-0259-4975-ac8e-77246ec9cfe0"
                    }
                ]
            },
            "main_image": {
                "data": {
                    "type": "main_image",
                    "id": "7cc08cbb-256e-4271-9b01-d03a9fac9f0a"
                }
            }
        },
        "material": null,
        "max_watt": null,
        "bulb_qty": null,
        "bulb": null,
        "new": null,
        "on_sale": null,
        "background_colour": "#d9d9d9",
        "finish": "test"
    }
}
		`,
	}
	client, done := test.Start(t)
	defer done()

	type MyProduct struct {
		entities.Product
		Material         string `json:"material"`
		MaxWatt          int    `json:"max_watt"`
		BulbQty          int    `json:"bulb_qty"`
		Bulb             bool   `json:"bulb"`
		New              string `json:"new"`
		OnSale           bool   `json:"on_sale"`
		BackgroundColour string `json:"background_colour"`
		Finish           string `json:"finish"`
	}
	var product MyProduct
	_, err := client.Get(
		"products/9eda5ba0-4f4a-4074-8547-ccb05d1b5981",
		gomo.Data(&product),
	)
	if err != nil {
		t.Fatal(err)
	}
	if !test.called {
		t.Fatal("server not called")
	}
	expected := MyProduct{
		Product: entities.Product{
			ID:            "9eda5ba0-4f4a-4074-8547-ccb05d1b5981",
			Type:          "product",
			Name:          "Crown",
			Slug:          "crown",
			SKU:           "CWLP100BLK",
			Description:   "Abstract, sculptural, refined and edgy with a modern twist. Its symmetrical, spoked structure generates a clever geometric presence, which works well in a contemporary environment.",
			ManageStock:   true,
			Status:        "live",
			CommodityType: "physical",
			Price: []entities.ProductPrice{
				entities.ProductPrice{
					Amount:      47500,
					Currency:    "USD",
					IncludesTax: true,
				},
			},
			Meta: entities.ProductMeta{
				DisplayPrice: entities.DisplayPriceWrapper{
					WithTax: entities.DisplayPrice{
						Amount:    47500,
						Currency:  "USD",
						Formatted: "$475.00",
					},
					WithoutTax: entities.DisplayPrice{
						Amount:    47500,
						Currency:  "USD",
						Formatted: "$475.00",
					},
				},
				Timestamps: entities.Timestamps{
					CreatedAt: "2017-06-19T14:58:42+00:00",
				},
				Stock: entities.ProductStock{
					Level:        500,
					Availability: "in-stock",
				},
				Variations:      []entities.ProductVariation(nil),
				VariationMatrix: []interface{}{},
			},
			Relationships: map[string]interface{}{
				"categories": map[string]interface{}{
					"data": []interface{}{
						map[string]interface{}{
							"id":   "a636c261-0259-4975-ac8e-77246ec9cfe0",
							"type": "category",
						},
					},
				},
				"files": map[string]interface{}{
					"data": []interface{}{
						map[string]interface{}{
							"id":   "7cc08cbb-256e-4271-9b01-d03a9fac9f0a",
							"type": "file",
						},
					},
				},
				"main_image": map[string]interface{}{
					"data": map[string]interface{}{
						"id":   "7cc08cbb-256e-4271-9b01-d03a9fac9f0a",
						"type": "main_image",
					},
				},
			},
			Included: entities.ProductIncludes{
				Brands:      []entities.Brand(nil),
				Categories:  []entities.Category(nil),
				Collections: []entities.Collection(nil),
				Files:       []entities.File(nil),
				MainImage: entities.File{
					ID:       "",
					Type:     "",
					FileName: "",
					Public:   false,
					MimeType: "",
					FileSize: 0,
					Meta: entities.FileMeta{
						Dimensions: struct {
							Width  int32 "json:\"width\""
							Height int32 "json:\"height\""
						}{
							Width:  0,
							Height: 0,
						},
						Timestamps: entities.Timestamps{
							CreatedAt: "",
						},
					},
					Link: struct {
						Href string "json:\"href\""
					}{
						Href: "",
					},
					Links: entities.Links{
						Self: "",
					},
				},
				MainImages: []entities.File(nil),
			},
		},
		Material:         "",
		MaxWatt:          0,
		BulbQty:          0,
		Bulb:             false,
		New:              "",
		OnSale:           false,
		BackgroundColour: "#d9d9d9",
		Finish:           "test",
	}
	if !reflect.DeepEqual(product, expected) {
		t.Errorf("\nunexpected response, expected:\n%#v\ngot:\n%#v\n", expected, product)
	}
}

func TestPostProductWithFlows(t *testing.T) {
	test := testServer{
		expectMethod: "POST",
		expectBody: `
{
    "data": {
	"type": "product",
	"name": "Crown",
	"slug": "crown",
	"sku": "CWLP100BLK",
	"description": "Abstract, sculptural, refined and edgy with a modern twist. Its symmetrical, spoked structure generates a clever geometric presence, which works well in a contemporary environment.",
	"manage_stock": true,
	"status": "live",
	"commodity_type": "physical",
	"price": [
	    {
		"amount": 47500,
		"currency": "USD",
		"includes_tax": true
	    }
	],
	"meta": {
	    "display_price": {
		"with_tax": {
		    "amount": 0,
		    "currency": "",
		    "formatted": ""
		},
		"without_tax": {
		    "amount": 0,
		    "currency": "",
		    "formatted": ""
		}
	    },
	    "timestamps": {},
	    "stock": {
		"level": 0,
		"availability": ""
	    },
	    "variation_matrix": null
	},
	"included": {
	    "brands": null,
	    "categories": null,
	    "collections": null,
	    "files": null,
	    "main_image": {
		"type": "",
		"file_name": "",
		"public": false,
		"mime_type": "",
		"file_size": 0,
		"meta": {
		    "dimensions": {
			"width": 0,
			"height": 0
		    },
		    "timestamps": {}
		},
		"link": {
		    "href": ""
		},
		"links": {
		    "self": ""
		}
	    },
	    "main_images": null
	},
	"background_colour": "#d9d9d9",
	"finish": "test"
    }
}
		`,
		responseCode: http.StatusOK,
		response: `
{
    "data": {
        "type": "product",
        "id": "9eda5ba0-4f4a-4074-8547-ccb05d1b5981",
        "name": "Crown",
        "slug": "crown",
        "sku": "CWLP100BLK",
        "manage_stock": true,
        "description": "Abstract, sculptural, refined and edgy with a modern twist. Its symmetrical, spoked structure generates a clever geometric presence, which works well in a contemporary environment.",
        "price": [
            {
                "amount": 47500,
                "currency": "USD",
                "includes_tax": true
            }
        ],
        "status": "live",
        "commodity_type": "physical",
        "meta": {
            "timestamps": {
                "created_at": "2017-06-19T14:58:42+00:00",
                "updated_at": "2018-04-10T09:12:05+00:00"
            },
            "display_price": {
                "with_tax": {
                    "amount": 47500,
                    "currency": "USD",
                    "formatted": "$475.00"
                },
                "without_tax": {
                    "amount": 47500,
                    "currency": "USD",
                    "formatted": "$475.00"
                }
            },
            "stock": {
                "level": 500,
                "availability": "in-stock"
            },
            "variation_matrix": []
        },
        "relationships": {
            "files": {
                "data": [
                    {
                        "type": "file",
                        "id": "7cc08cbb-256e-4271-9b01-d03a9fac9f0a"
                    }
                ]
            },
            "categories": {
                "data": [
                    {
                        "type": "category",
                        "id": "a636c261-0259-4975-ac8e-77246ec9cfe0"
                    }
                ]
            },
            "main_image": {
                "data": {
                    "type": "main_image",
                    "id": "7cc08cbb-256e-4271-9b01-d03a9fac9f0a"
                }
            }
        },
        "material": null,
        "max_watt": null,
        "bulb_qty": null,
        "bulb": null,
        "new": null,
        "on_sale": null,
        "background_colour": "#d9d9d9",
        "finish": "test"
    }
}
		`,
	}
	client, done := test.Start(t)
	defer done()

	type MyProduct struct {
		entities.Product
		Material         string `json:"material,omitempty"`
		MaxWatt          int    `json:"max_watt,omitempty"`
		BulbQty          int    `json:"bulb_qty,omitempty"`
		Bulb             bool   `json:"bulb,omitempty"`
		New              string `json:"new,omitempty"`
		OnSale           bool   `json:"on_sale,omitempty"`
		BackgroundColour string `json:"background_colour,omitempty"`
		Finish           string `json:"finish,omitempty"`
	}
	product := MyProduct{
		Product: entities.Product{
			Type:          "product",
			Name:          "Crown",
			Slug:          "crown",
			SKU:           "CWLP100BLK",
			Description:   "Abstract, sculptural, refined and edgy with a modern twist. Its symmetrical, spoked structure generates a clever geometric presence, which works well in a contemporary environment.",
			ManageStock:   true,
			Status:        "live",
			CommodityType: "physical",
			Price: []entities.ProductPrice{
				entities.ProductPrice{
					Amount:      47500,
					Currency:    "USD",
					IncludesTax: true,
				},
			},
		},
		BackgroundColour: "#d9d9d9",
		Finish:           "test",
	}
	_, err := client.Post(
		"products",
		gomo.Body(&product),
		gomo.Data(&product),
	)
	if err != nil {
		t.Fatal(err)
	}
	if !test.called {
		t.Fatal("server not called")
	}
	expected := MyProduct{
		Product: entities.Product{
			ID:            "9eda5ba0-4f4a-4074-8547-ccb05d1b5981",
			Type:          "product",
			Name:          "Crown",
			Slug:          "crown",
			SKU:           "CWLP100BLK",
			Description:   "Abstract, sculptural, refined and edgy with a modern twist. Its symmetrical, spoked structure generates a clever geometric presence, which works well in a contemporary environment.",
			ManageStock:   true,
			Status:        "live",
			CommodityType: "physical",
			Price: []entities.ProductPrice{
				entities.ProductPrice{
					Amount:      47500,
					Currency:    "USD",
					IncludesTax: true,
				},
			},
			Meta: entities.ProductMeta{
				DisplayPrice: entities.DisplayPriceWrapper{
					WithTax: entities.DisplayPrice{
						Amount:    47500,
						Currency:  "USD",
						Formatted: "$475.00",
					},
					WithoutTax: entities.DisplayPrice{
						Amount:    47500,
						Currency:  "USD",
						Formatted: "$475.00",
					},
				},
				Timestamps: entities.Timestamps{
					CreatedAt: "2017-06-19T14:58:42+00:00",
				},
				Stock: entities.ProductStock{
					Level:        500,
					Availability: "in-stock",
				},
				Variations:      []entities.ProductVariation(nil),
				VariationMatrix: []interface{}{},
			},
			Relationships: map[string]interface{}{
				"categories": map[string]interface{}{
					"data": []interface{}{
						map[string]interface{}{
							"id":   "a636c261-0259-4975-ac8e-77246ec9cfe0",
							"type": "category",
						},
					},
				},
				"files": map[string]interface{}{
					"data": []interface{}{
						map[string]interface{}{
							"id":   "7cc08cbb-256e-4271-9b01-d03a9fac9f0a",
							"type": "file",
						},
					},
				},
				"main_image": map[string]interface{}{
					"data": map[string]interface{}{
						"id":   "7cc08cbb-256e-4271-9b01-d03a9fac9f0a",
						"type": "main_image",
					},
				},
			},
			Included: entities.ProductIncludes{
				Brands:      []entities.Brand(nil),
				Categories:  []entities.Category(nil),
				Collections: []entities.Collection(nil),
				Files:       []entities.File(nil),
				MainImage: entities.File{
					ID:       "",
					Type:     "",
					FileName: "",
					Public:   false,
					MimeType: "",
					FileSize: 0,
					Meta: entities.FileMeta{
						Dimensions: struct {
							Width  int32 "json:\"width\""
							Height int32 "json:\"height\""
						}{
							Width:  0,
							Height: 0,
						},
						Timestamps: entities.Timestamps{
							CreatedAt: "",
						},
					},
					Link: struct {
						Href string "json:\"href\""
					}{
						Href: "",
					},
					Links: entities.Links{
						Self: "",
					},
				},
				MainImages: []entities.File(nil),
			},
		},
		Material:         "",
		MaxWatt:          0,
		BulbQty:          0,
		Bulb:             false,
		New:              "",
		OnSale:           false,
		BackgroundColour: "#d9d9d9",
		Finish:           "test",
	}
	if !reflect.DeepEqual(product, expected) {
		t.Errorf("\nunexpected response, expected:\n%#v\ngot:\n%#v\n", expected, product)
	}
}

func TestDeleteProduct(t *testing.T) {
	test := testServer{
		expectMethod: "DELETE",
		responseCode: http.StatusNoContent,
		response:     "",
	}
	client, done := test.Start(t)
	defer done()

	_, err := client.Delete(
		"products/9eda5ba0-4f4a-4074-8547-ccb05d1b5981",
	)
	if err != nil {
		t.Fatal(err)
	}
	if !test.called {
		t.Fatal("server not called")
	}
}

func TestPutProduct(t *testing.T) {
	test := testServer{
		expectMethod: "PUT",
		expectBody: `
{
    "data": {
        "type": "product",
	"name": "Foo"
    }
}
		`,
		responseCode: http.StatusOK,
		response: `
{
    "data": {
        "type": "product",
        "id": "b47372eb-6f13-4bcb-ad06-329f4ffee69d",
        "name": "Foo",
        "slug": "a-product-1",
        "sku": "SP01-{COLOUR}-new-2",
        "manage_stock": false,
        "description": "Some description",
        "price": [
            {
                "amount": 6999,
                "currency": "GBP",
                "includes_tax": true
            },
            {
                "amount": 7499,
                "currency": "USD",
                "includes_tax": true
            }
        ],
        "status": "live",
        "commodity_type": "physical",
        "meta": {
            "timestamps": {
                "created_at": "2018-05-11T20:07:56+00:00",
                "updated_at": "2018-05-12T00:50:11+00:00"
            },
            "display_price": {
                "with_tax": {
                    "amount": 7499,
                    "currency": "USD",
                    "formatted": "$74.99"
                },
                "without_tax": {
                    "amount": 7499,
                    "currency": "USD",
                    "formatted": "$74.99"
                }
            },
            "stock": {
                "level": 0,
                "availability": "out-stock"
            }
        },
        "relationships": {
            "categories": {
                "data": [
                    {
                        "type": "category",
                        "id": "3acf1815-ef09-44df-a6d3-e3e702663524"
                    }
                ]
            },
            "collections": {
                "data": [
                    {
                        "type": "collection",
                        "id": "10229491-7dbb-424d-be08-80c5735795cc"
                    }
                ]
            }
        }
    }
}
		`,
	}
	client, done := test.Start(t)
	defer done()

	var product entities.Product
	_, err := client.Put(
		"products/b47372eb-6f13-4bcb-ad06-329f4ffee69d",
		gomo.Body(struct {
			Type string `json:"type"`
			Name string `json:"name"`
		}{
			Type: "product",
			Name: "Foo",
		}),
		gomo.Data(&product),
	)
	if err != nil {
		t.Fatal(err)
	}
	if !test.called {
		t.Fatal("server not called")
	}

	expected := entities.Product{
		ID:            "b47372eb-6f13-4bcb-ad06-329f4ffee69d",
		Type:          "product",
		Name:          "Foo",
		Slug:          "a-product-1",
		SKU:           "SP01-{COLOUR}-new-2",
		Description:   "Some description",
		ManageStock:   false,
		Status:        "live",
		CommodityType: "physical",
		Price: []entities.ProductPrice{
			entities.ProductPrice{
				Amount:      6999,
				Currency:    "GBP",
				IncludesTax: true,
			},
			entities.ProductPrice{
				Amount:      7499,
				Currency:    "USD",
				IncludesTax: true,
			},
		},
		Meta: entities.ProductMeta{
			DisplayPrice: entities.DisplayPriceWrapper{
				WithTax: entities.DisplayPrice{
					Amount:    7499,
					Currency:  "USD",
					Formatted: "$74.99",
				},
				WithoutTax: entities.DisplayPrice{
					Amount:    7499,
					Currency:  "USD",
					Formatted: "$74.99",
				},
			},
			Timestamps: entities.Timestamps{
				CreatedAt: "2018-05-11T20:07:56+00:00",
			},
			Stock: entities.ProductStock{
				Level:        0,
				Availability: "out-stock",
			},
			Variations:      []entities.ProductVariation(nil),
			VariationMatrix: interface{}(nil),
		},
		Relationships: map[string]interface{}{
			"categories": map[string]interface{}{
				"data": []interface{}{
					map[string]interface{}{
						"id":   "3acf1815-ef09-44df-a6d3-e3e702663524",
						"type": "category",
					},
				},
			},
			"collections": map[string]interface{}{
				"data": []interface{}{
					map[string]interface{}{
						"id":   "10229491-7dbb-424d-be08-80c5735795cc",
						"type": "collection",
					},
				},
			},
		},
		Included: entities.ProductIncludes{
			Brands:      []entities.Brand(nil),
			Categories:  []entities.Category(nil),
			Collections: []entities.Collection(nil),
			Files:       []entities.File(nil),
			MainImage: entities.File{
				ID:       "",
				Type:     "",
				FileName: "",
				Public:   false,
				MimeType: "",
				FileSize: 0,
				Meta: entities.FileMeta{
					Dimensions: struct {
						Width  int32 "json:\"width\""
						Height int32 "json:\"height\""
					}{
						Width:  0,
						Height: 0,
					},
					Timestamps: entities.Timestamps{
						CreatedAt: "",
					},
				},
				Link: struct {
					Href string "json:\"href\""
				}{
					Href: "",
				},
				Links: entities.Links{
					Self: "",
				},
			},
			MainImages: []entities.File(nil),
		},
	}
	if !reflect.DeepEqual(product, expected) {
		t.Errorf("\nunexpected response, expected:\n%#v\ngot:\n%#v\n", expected, product)
	}
}
