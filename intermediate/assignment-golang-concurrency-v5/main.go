package main

import (
	"fmt"
	"time"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func GetTLD(domain string) (TLD string, IDN_TLD string) {
	var ListIDN_TLD = map[string]string{
		".com": ".co.id",
		".org": ".org.id",
		".gov": ".go.id",
	}

	for i := len(domain) - 1; i >= 0; i-- {
		if domain[i] == '.' {
			TLD = domain[i:]
			break
		}
	}

	if _, ok := ListIDN_TLD[TLD]; ok {
		return TLD, ListIDN_TLD[TLD]
	} else {
		return TLD, TLD
	}
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	TLD, IDN_TLD := GetTLD(website.Domain) // TODO: replace this

	// check if domain is empty & error handling
	if website.Domain == "" {
		chErr <- fmt.Errorf("domain name is empty")
		return
	}

	// check if domain is not valid & error handling
	if !website.Valid {
		chErr <- fmt.Errorf("domain not valid")
		return
	}

	// check if domain RefIPS not valid & error handling
	if website.RefIPs == -1 {
		chErr <- fmt.Errorf("domain RefIPs not valid")
		return
	}

	// complete website property
	website.TLD = TLD
	website.IDN_TLD = IDN_TLD

	// transfer complete website to ch
	ch <- website
	time.Sleep(100 * time.Millisecond)
}

// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)

	// looping input data and run FuncProcessGetTLD for each data
	for _, website := range data {
		go FuncProcessGetTLD(website, ch, errCh)
	}

	// declaration slice filtered for filtered data
	filtered := make([]RowData, 0, len(data))

	// filtering process
	for i := 0; i < len(data); i++ {
		select {
		case finalData := <-ch:
			if finalData.TLD == TLD {
				filtered = append(filtered, finalData)
			}
		case err := <-errCh:
			return nil, err
		}
	}

	return filtered, nil
}

// gunakan untuk melakukan debugging
func main() {
	rows, err := FilterAndFillData(".com", []RowData{
		{1, "google.com", "", "", false, 100},
		{2, "facebook.com", "", "", true, 100},
		{3, "golang.org", "", "", true, 100},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)
}
