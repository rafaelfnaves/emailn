package main

import (
	"emailn/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {
	campaign := campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(campaign)
	if err == nil {
		println("Nada")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {

			// println(v.StructField() + " is invalid:" + v.Tag())
		}
	}
}
