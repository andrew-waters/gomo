package entities

// File is a Moltin File - https://docs.moltin.com/advanced/files
type File struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	FileName string `json:"file_name"`
	Public   bool   `json:"public"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
	Link     struct {
		Href string `json:"href"`
	} `json:"link"`
}

// SetType sets the resource type on the struct
func (f *File) SetType() {
	f.Type = fileType
}
