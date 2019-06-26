package core

// File is a Moltin File - https://docs.moltin.com/advanced/files
type File struct {
	ID       string   `json:"id,omitempty"`
	Type     string   `json:"type"`
	FileName string   `json:"file_name"`
	Public   bool     `json:"public"`
	MimeType string   `json:"mime_type"`
	FileSize int      `json:"file_size"`
	Meta     FileMeta `json:"meta,omitempty"`
	Link     struct {
		Href string `json:"href"`
	} `json:"link"`
	Links Links `json:"links,omitempty"`
}

// FileMeta represents the meta object for a moltin file entity
type FileMeta struct {
	Dimensions struct {
		Width  int32 `json:"width"`
		Height int32 `json:"height"`
	} `json:"dimensions"`
	Timestamps Timestamps `json:"timestamps"`
}

// SetType sets the resource type on the struct
func (f *File) SetType() {
	f.Type = fileType
}
