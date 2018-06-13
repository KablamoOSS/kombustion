package tasks

import (
	"fmt"
	"os"
	"strings"

	printer "github.com/KablamoOSS/go-cli-printer"
)

func checkError(err error) {
	if err != nil {
		if strings.Contains(err.Error(), "No updates are to be performed") {
			printer.Error(fmt.Errorf("No updates are to be performed"), "", "")
			os.Exit(0)
		} else if strings.Contains(err.Error(), "Stack with id") && strings.Contains(err.Error(), "does not exist") {
			printer.Error(fmt.Errorf("The stack does not exist"), "", "")
			os.Exit(0)
		} else {
			printer.Fatal(err, "", "")
		}
	}
}

func checkErrorDeletePoll(err error) {
	if err != nil {
		if strings.Contains(err.Error(), "No updates are to be performed") {
			printer.Fatal(fmt.Errorf("No updates are to be performed"), "", "")
			os.Exit(0)
		} else if strings.Contains(err.Error(), "Stack with id") && strings.Contains(err.Error(), "does not exist") {
			os.Exit(0)
		} else {
			printer.Fatal(err, "", "")
		}
	}
}
