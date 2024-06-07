package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	vCore "github.com/xinyangme001/V2bX/core"
)

var (
	version  = "TempVersion" //use ldflags replace
	codename = "V2bX"
	intro    = "A V2board backend based on multi core"
)

var versionCommand = cobra.Command{
	Use:   "version",
	Short: "Print version info",
	Run: func(_ *cobra.Command, _ []string) {
		showVersion()
	},
}

func init() {
	command.AddCommand(&versionCommand)
}

func showVersion() {
	fmt.Printf("%s %s (%s) \n", codename, version, intro)
	fmt.Printf("Supported cores: %s\n", strings.Join(vCore.RegisteredCore(), ", "))
	// Warning
	fmt.Println(Warn("This version need V2board version >= 1.7.0."))
	fmt.Println(Warn("The version have many changed for config, please check your config file"))
}
