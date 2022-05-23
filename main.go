package main

import (
	"BE_WEB_BEM_Proker/model"
	"BE_WEB_BEM_Proker/service"
	"fmt"
)

func main() {
	var bph = model.EntitasBPH{
		Kementrian:          "Kementrian Sosling",
		Kontak:              "0812-3456-7890",
		Password:            "12345678",
		NamaBPH:             "BEM",
		DeskripsiKementrian: "Test",
	}
	admin := service.NewAdminService()
	_, err := admin.Create(&bph)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success create new bph")
	fmt.Println(admin.FindByID(1))
}
