package tasks

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
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

func getParamMap(c *cli.Context) map[string]string {
	paramMap := make(map[string]string)
	params := c.StringSlice("param")
	for _, param := range params {
		parts := strings.Split(param, "=")
		if len(parts) > 1 {
			paramMap[parts[0]] = strings.Join(parts[1:], "=")
		}
	}
	return paramMap
}
