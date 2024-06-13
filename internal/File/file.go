package file

import (
	"fmt"
)

// File struct contains name and hash value
type File struct {
	Name string
	Hash [16]byte
}

// Convert hash value to string
func (f File) HashString() string {
	return fmt.Sprintf("%x", f.Hash)
}

// Simplify the output of File values
func (f File) String() string {
	return fmt.Sprintf("Name: %s, Hash: %x ", f.Name, f.Hash)
}
