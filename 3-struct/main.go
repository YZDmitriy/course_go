package main

import (
	"main/bins"
)

func main() {
	bin := bins.NewBin("12345", "My First Bin", true)
	binList := bins.NewBinList()
	*binList = append(*binList, *bin)
}
