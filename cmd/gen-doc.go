package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
)

var swaggoInitCmd = &cobra.Command{
	Use:   "gen-doc",
	Short: "Initializes swagger documentation",
	Long:  "Initializes swagger documentation for Templater API",
	Run: func(cmd *cobra.Command, args []string) {
		// runs command swag init --parseDependency --parseInternal
		swag := exec.Command("swag", "init", "--parseDependency", "--parseInternal")

		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		swag.Stdout = mw
		swag.Stderr = mw

		// Execute the command
		if err := swag.Run(); err != nil {
			panic(err)
		}
		fmt.Println(stdBuffer.String())
	},
}

func init() {
	RootCmd.AddCommand(swaggoInitCmd)
}
