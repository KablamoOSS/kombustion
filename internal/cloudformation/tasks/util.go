package tasks

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func checkError(err error) {
	if err != nil {
		if strings.Contains(err.Error(), "No updates are to be performed") {
			log.Warn("No updates are to be performed.")
			os.Exit(0)
		} else if strings.Contains(err.Error(), "Stack with id") && strings.Contains(err.Error(), "does not exist") {
			log.Warn("The stack does not exist.")
			os.Exit(0)
		} else {
			log.Fatal(err)
		}
	}
}

func checkErrorDeletePoll(err error) {
	if err != nil {
		if strings.Contains(err.Error(), "No updates are to be performed") {
			log.Warn("No updates are to be performed.")
			os.Exit(0)
		} else if strings.Contains(err.Error(), "Stack with id") && strings.Contains(err.Error(), "does not exist") {
			os.Exit(0)
		} else {
			log.Fatal(err)
		}
	}
}
