package storage

import (
	"encoding/json"
	"main/bins"
	"os"
)

// SaveBins сохраняет список Bin в файл в формате JSON.
func SaveBins(filename string, binList bins.BinList) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(binList); err != nil {
		return err
	}
	return nil
}

// LoadBins читает список Bin из файла JSON.
func LoadBins(filename string) (bins.BinList, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var binList bins.BinList
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&binList); err != nil {
		return nil, err
	}
	return binList, nil
}
