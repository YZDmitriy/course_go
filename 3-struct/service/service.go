package service

import (
	"main/bins"
)

// FileStorage определяет интерфейс для работы с хранилищем
type FileStorage interface {
	SaveBins(filename string, binList bins.BinList) error
	LoadBins(filename string) (bins.BinList, error)
}

// FileValidator определяет интерфейс для валидации файлов
type FileValidator interface {
	IsJSONFile(filename string) error
}

// BinService представляет основной сервис приложения
type BinService struct {
	storage   FileStorage
	validator FileValidator
	bins      bins.BinList
}

// NewBinService создает новый экземпляр сервиса
func NewBinService(storage FileStorage, validator FileValidator) *BinService {
	return &BinService{
		storage:   storage,
		validator: validator,
		bins:      make(bins.BinList, 0),
	}
}

// SaveBinsToFile сохраняет bins в файл после валидации
func (s *BinService) SaveBinsToFile(filename string) error {
	if err := s.validator.IsJSONFile(filename); err != nil {
		return err
	}
	return s.storage.SaveBins(filename, s.bins)
}

// LoadBinsFromFile загружает bins из файла после валидации
func (s *BinService) LoadBinsFromFile(filename string) error {
	if err := s.validator.IsJSONFile(filename); err != nil {
		return err
	}

	loadedBins, err := s.storage.LoadBins(filename)
	if err != nil {
		return err
	}

	s.bins = loadedBins
	return nil
}

// AddBin добавляет новый bin в список
func (s *BinService) AddBin(id, name string, private bool) {
	newBin := bins.NewBin(id, name, private)
	s.bins = append(s.bins, *newBin)
}

// GetBins возвращает текущий список bins
func (s *BinService) GetBins() bins.BinList {
	return s.bins
}
