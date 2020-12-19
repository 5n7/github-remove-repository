package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	ghrm "github.com/skmatz/github-remove-repository"

	"github.com/spf13/cobra"
)

var (
	force   bool
	owner   string
	repo    string
	version bool
)

func tokenPath() (string, error) {
	cfg, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cfg, "ghrm", "ghrm.json"), nil
}

type token struct {
	Token string `json:"token"`
}

func runRoot(cmd *cobra.Command, args []string) error {
	if version {
		return runVersion(cmd, args)
	}

	path, err := tokenPath()
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var tkn token
	if err := json.Unmarshal(bytes, &tkn); err != nil {
		return err
	}

	gh := ghrm.NewGitHubClient(tkn.Token)

	if !force {
		var ok bool
		prompt := &survey.Confirm{
			Message: fmt.Sprintf("Really remove repository %s/%s?", owner, repo),
		}
		if err := survey.AskOne(prompt, &ok); err != nil {
			return err
		}
		if !ok {
			return nil
		}
	}

	if err := ghrm.RemoveRepository(gh, owner, repo); err != nil {
		return err
	}
	return nil
}

var errInvalidArgs = errors.New("invalid args; 'owner/repo' or 'owner repo'")

var rootCmd = &cobra.Command{
	Use:   "ghrm",
	Short: "ghrm is the GitHub repository remover",
	Long:  "ghrm is the GitHub repository remover.",
	Args: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 1:
			ss := strings.Split(args[0], "/")
			if len(ss) != 2 {
				return errInvalidArgs
			}
			owner, repo = ss[0], ss[1]
		case 2:
			owner, repo = args[0], args[1]
		default:
			return errInvalidArgs
		}

		if owner == "" || repo == "" {
			return errInvalidArgs
		}
		return nil
	},
	RunE: runRoot,
}

func init() {
	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "skip confirmation")
	rootCmd.Flags().BoolVarP(&version, "version", "V", false, "show version")
}

func Execute() {
	rootCmd.Execute() //nolint:errcheck
}
