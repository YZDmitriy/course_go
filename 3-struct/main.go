package main

import (
    "fmt"
    "main/adapters"
    "main/service"
)

func main() {
    // Создаем адаптеры
    storageAdapter := adapters.NewStorageAdapter()
    fileValidator := adapters.NewFileValidator()

    // Создаем сервис
    binService := service.NewBinService(storageAdapter, fileValidator)

    // Пример использования
    binService.AddBin("1", "Test Bin", false)
    
    // Сохраняем в файл
    err := binService.SaveBinsToFile("bins.json")
    if err != nil {
        fmt.Printf("Error saving bins: %v\n", err)
        return
    }

    // Загружаем из файла
    err = binService.LoadBinsFromFile("bins.json")
    if err != nil {
        fmt.Printf("Error loading bins: %v\n", err)
        return
    }

    // Получаем список bins
    bins := binService.GetBins()
    fmt.Printf("Loaded bins: %+v\n", bins)
}