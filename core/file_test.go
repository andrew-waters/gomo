package core

import "testing"

func TestFileType(t *testing.T) {
	f := File{}
	f.SetType()
	if f.Type != "file" {
		t.Errorf("File did not return correct type: `%s`", f.Type)
	}
}
