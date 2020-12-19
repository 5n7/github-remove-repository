package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

const url = "https://github.com/settings/tokens/new?description=ghrm&scopes=delete_repo"

func runToken(cmd *cobra.Command, args []string) error {
	fmt.Println(url)

	var ans string
	prompt := &survey.Input{
		Message: "Input your token:",
	}
	if err := survey.AskOne(prompt, &ans); err != nil {
		return err
	}

	tkn := token{Token: ans}
	bytes, err := json.Marshal(tkn)
	if err != nil {
		return err
	}

	path, err := tokenPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, bytes, os.ModePerm); err != nil {
		return err
	}
	return nil
}

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Prompt for the GitHub access token",
	Long:  "Prompt for the GitHub access token.",
	RunE:  runToken,
}

func init() {
	rootCmd.AddCommand(tokenCmd)
}
