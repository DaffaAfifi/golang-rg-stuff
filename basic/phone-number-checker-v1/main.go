package main

import (
	"fmt"
)

func PhoneNumberChecker(number string, result *string) {

	if len(number) >= 11 && number[:3] == "628" {
		if number[:5] == "62811" || number[:5] == "62812" || number[:5] == "62813" || number[:5] == "62814" || number[:5] == "62815" {
			*result = "Telkomsel"
		} else if number[:5] == "62816" || number[:5] == "62817" || number[:5] == "62818" || number[:5] == "62819" {
			*result = "Indosat"
		} else if number[:5] == "62821" || number[:5] == "62822" || number[:5] == "62823" {
			*result = "XL"
		} else if number[:5] == "62827" || number[:5] == "62828" || number[:5] == "62829" {
			*result = "Tri"
		} else if number[:5] == "62852" || number[:5] == "62853" {
			*result = "AS"
		} else if number[:5] == "62881" || number[:5] == "62882" || number[:5] == "62883" || number[:5] == "62884" || number[:5] == "62885" || number[:5] == "62886" || number[:5] == "62887" || number[:5] == "62888" {
			*result = "Smartfren"
		} else {
			*result = "invalid"
		}
	} else if len(number) >= 10 && number[:2] == "08" {
		if number[:4] == "0811" || number[:4] == "0812" || number[:4] == "0813" || number[:4] == "0814" || number[:4] == "0815" {
			*result = "Telkomsel"
		} else if number[:4] == "0816" || number[:4] == "0817" || number[:4] == "0818" || number[:4] == "0819" {
			*result = "Indosat"
		} else if number[:4] == "0821" || number[:4] == "0822" || number[:4] == "0823" {
			*result = "XL"
		} else if number[:4] == "0827" || number[:4] == "0828" || number[:4] == "0829" {
			*result = "Tri"
		} else if number[:4] == "0852" || number[:4] == "0853" {
			*result = "AS"
		} else if number[:4] == "0881" || number[:4] == "0882" || number[:4] == "0883" || number[:4] == "0884" || number[:4] == "0885" || number[:4] == "0886" || number[:4] == "0887" || number[:4] == "0888" {
			*result = "Smartfren"
		} else {
			*result = "invalid"
		}
	} else {
		*result = "invalid"
	}

}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "08993456123"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
