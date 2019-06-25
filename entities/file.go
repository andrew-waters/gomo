package entities

// File is a Moltin File - https://docs.moltin.com/advanced/files
type File struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	FileName string `json:"file_name"`
	Public   bool   `json:"public"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
	Meta     struct {
		Dimensions struct {
			Width  int32 `json:"width"`
			Height int32 `json:"height"`
		} `json:"dimensions"`
		Timestamps Timestamps `json:"timestamps"`
	}
	Link struct {
		Href string `json:"href"`
	} `json:"link"`
	Links Links `json:"links"`
}

// SetType sets the resource type on the struct
func (f *File) SetType() {
	f.Type = fileType
}
