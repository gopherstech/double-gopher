package storage

import file "double-gopher/internal/File"

// Storage struct key is the hash and it contains Files
type Storage struct {
	Files map[string]file.File
}

// Create new storage
func NewStorage() *Storage {
	return &Storage{
		Files: make(map[string]file.File),
	}
}

// Check if Hash is present in the storage
func (s *Storage) HashExists(hash string) bool {
	_, ok := s.Files[hash]
	if !ok {
		return false
	} else {
		return true
	}

}

// Upload the object with the key as hashvalue, returns bool in order to know if the file was uploaded or no
func (s *Storage) UploadFileData(f file.File) bool {
	hash := f.HashString()
	flag := s.HashExists(hash)
	if !flag {
		s.Files[hash] = f
		return true
	} else {
		return false
	}
}
