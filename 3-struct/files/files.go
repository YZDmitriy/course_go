package files

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// ReadFile читает содержимое любого файла.
func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// IsJSONFile проверяет, что файл имеет расширение .json (без учета регистра).
func IsJSONFile(filename string) error {
	if strings.ToLower(filepath.Ext(filename)) != ".json" {
		return errors.New("file is not a JSON file")
	}
	return nil
}
