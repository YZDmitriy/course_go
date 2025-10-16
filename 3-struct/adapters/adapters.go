package adapters

import (
    "main/files"
    "main/storage"
	"main/bins"
)

// StorageAdapter адаптер для storage
type StorageAdapter struct{}

func NewStorageAdapter() *StorageAdapter {
    return &StorageAdapter{}
}

func (a *StorageAdapter) SaveBins(filename string, binList bins.BinList) error {
    return storage.SaveBins(filename, binList)
}

func (a *StorageAdapter) LoadBins(filename string) (bins.BinList, error) {
    return storage.LoadBins(filename)
}

// FileValidator адаптер для files
type FileValidator struct{}

func NewFileValidator() *FileValidator {
    return &FileValidator{}
}

func (v *FileValidator) IsJSONFile(filename string) error {
    return files.IsJSONFile(filename)
}