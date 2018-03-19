package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	var token string
	var rep string

	apiURL := "https://api.github.com"

	flag.StringVar(&token, "github_token", "", "Github token")
	flag.StringVar(&rep, "n", "", "Repository Name")
	// flag parse
	flag.VisitAll(func(f *flag.Flag) {
		if s := os.Getenv(strings.ToUpper(f.Name)); s != "" {
			f.Value.Set(s)
		}
	})
	flag.Parse()

	// create context.
	ctx := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tokenClient := oauth2.NewClient(ctx, tokenService)

	// check Env GH_API_URL.
	if len(os.Getenv("GITHUB_API_URL")) != 0 {
		apiURL = os.Getenv("GITHUB_API_URL")
	}

	// create client.
	c, err := github.NewEnterpriseClient(apiURL, "", tokenClient)
	if err != nil {
		fmt.Printf("Problem in TOKEN Set %v\n", err)
		os.Exit(1)
	}

	// set Repository Name
	repo := &github.Repository{
		Name: github.String(rep),
	}

	// create repository
	r, _, err := c.Repositories.Create(ctx, "", repo)
	if err != nil {
		fmt.Printf("Problem Create Repository %v\n", err)
		os.Exit(1)
	}

	// print add command.
	fmt.Println("git remote add origin " + r.GetSSHURL())
}
