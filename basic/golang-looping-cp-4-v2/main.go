package main

import "fmt"

func EmailInfo(email string) string {
	// dapetin @ dan . index ke berapa

	var at int
	var dot int
	var domain string
	var tld string

	for i := 0; i <= len(email)-1; i++ {
		if string(email[i]) == "@" {
			at = i
		} else if string(email[i]) == "." {
			dot = i
			break
		}
	}

	for i := 0; i <= len(email)-1; i++ {
		if i > at && i < dot {
			domain += string(email[i])
		}
	}

	for i := 0; i <= len(email)-1; i++ {
		if i > dot {
			tld += string(email[i])
		}
	}

	domainDetail := "Domain: " + domain + " dan TLD: " + tld

	return domainDetail

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("ptmencaricintasejati@gmail.co.id"))
}
