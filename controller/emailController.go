package controller

import (
	"fmt"
	"net"
	"strings"
)

func CheckDomain(url string) {
	var hasMx, hasSPF, hasDMARC bool
	var spfRecords, dmarcRecords []string

	mxRecords, err := net.LookupMX(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	hasMx = len(mxRecords) > 0

	txtRecords, err := net.LookupTXT(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecords = append(spfRecords, record)
		}
	}

	dmarcTxtRecords, err := net.LookupTXT("_dmarc." + url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, record := range dmarcTxtRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecords = append(dmarcRecords, record)
		}
	}

	fmt.Printf("Domain: %s\n", url)
	fmt.Printf("Has MX: %v\n", hasMx)
	fmt.Printf("Has SPF: %v\n", hasSPF)
	if hasSPF {
		fmt.Printf("SPF Records: %v\n", spfRecords)
	}
	fmt.Printf("Has DMARC: %v\n", hasDMARC)
	if hasDMARC {
		fmt.Printf("DMARC Records: %v\n", dmarcRecords)
	}
}
