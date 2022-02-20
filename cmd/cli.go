package cmd

import (
	"os"

	"github.com/marco-souza/bashbot/pkg/usecases"
)

func CommandHandler() {
	args := os.Args
	if len(args) <= 1 {
		panic("you must specify a command")
	}

	switch (args[1]) {
	case "report":
		usecases.SendWalletReport()
	}
}
