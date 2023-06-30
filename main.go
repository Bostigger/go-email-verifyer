package main

import (
	"bufio"
	"fmt"
	"github.com/bostigger/go-email-verifyer/controller"
	"log"
	"os"
)

func main() {
	fmt.Printf("domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		controller.CheckDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("could not read from input")
	}
}
