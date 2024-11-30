package models

type CreateSection struct {
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	Active        bool            `json:"active"`
	Deleted       bool            `json:"deleted"`
	ProductDetail []SectionDetail `json:"productDetail"`
}

type ListSection struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	Deleted     bool   `json:"deleted"`
}

type SectionDetail struct {
	Code      string `json:"code"`
	Rank      int    `json:"rank"`
	SectionID int64  `json:"section_id"`
}

type Section struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	Deleted     bool      `json:"deleted"`
	Products    []Product `json:"products"`
}

type Product struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Qty         int    `json:"qty"`
	Active      bool   `json:"active"`
	Deleted     bool   `json:"deleted"`
}
