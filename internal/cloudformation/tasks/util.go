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
			printer.Warn(
				fmt.Errorf("No updates are to be performed"), "", "")
			os.Exit(0)
		} else if strings.Contains(err.Error(), "Stack with id") && strings.Contains(err.Error(), "does not exist") {
			printer.Warn(fmt.Errorf("The stack does not exist"), "", "")
			os.Exit(0)
		} else if strings.Contains(err.Error(), "ROLLBACK_COMPLETE") {
			printer.Fatal(
				err,
				"The stack is in a ROLLBACK_COMPLETE state, you need to delete the stack before it can be updated.",
				"",
			)
		} else {
			// TODO: Make this error more helpful
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
			printer.SubStep(
				fmt.Sprintf("Success: Deleted Stack"),
				1,
				true,
				true,
			)
			os.Exit(0)
		} else {
			// TODO: Make this error more helpful
			printer.Fatal(err, "", "")
		}
	}
}
