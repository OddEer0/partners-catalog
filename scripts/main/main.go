package main

import "github.com/OddEer0/partners-catalog/scripts"

func main() {
	err := scripts.SetPartnerCatalog()
	if err != nil {
		panic(err)
	}
}
