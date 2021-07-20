package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hjldev/go-tools/internal/project"
	"github.com/spf13/cobra"
)

var ProjectNew = &cobra.Command{
	Use:   "new",
	Short: "Create a service template",
	Long:  "Create a service project using the repository template. Example: main new helloworld",
	Run:   run,
}

var repoURL string
var branch string

func init() {
	repoURL = "https://gitee.com/a2437463/go-layout.git"
	ProjectNew.Flags().StringVarP(&repoURL, "repo-url", "r", repoURL, "layout repo")
	ProjectNew.Flags().StringVarP(&branch, "branch", "b", branch, "repo branch")
}

func run(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	name := ""
	if len(args) == 0 {
		prompt := &survey.Input{
			Message: "What is project name ?",
			Help:    "Created project name.",
		}
		survey.AskOne(prompt, &name)
		if name == "" {
			return
		}
	} else {
		name = args[0]
	}
	p := &project.Project{Name: name}
	if err := p.New(ctx, wd, repoURL, branch); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err)
		return
	}
}
