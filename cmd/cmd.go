package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	_ "github.com/xinyangme001/V2bX/core/imports"
)

var command = &cobra.Command{
	Use: "V2bX",
}

func Run() {
	err := command.Execute()
	if err != nil {
		log.WithField("err", err).Error("Execute command failed")
	}
}
