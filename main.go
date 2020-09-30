package main

import (
	"github.com/spf13/cobra"
	"github.com/thefuriouscoder/golang-exercise/internal/cli"
	"github.com/thefuriouscoder/golang-exercise/internal/model"
	"github.com/thefuriouscoder/golang-exercise/internal/storage/punk"
)

func main() {

	var repo model.PunkRepo
	repo = punk.NewPunkRepository()

	rootCmd := &cobra.Command{Use: "punk-cli"}
	rootCmd.AddCommand(cli.InitBeerCmd(repo))
	rootCmd.AddCommand(cli.InitSearchCmd(repo))
	rootCmd.Execute()
}
