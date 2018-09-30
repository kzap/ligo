// Package client initializes the Linode Client and provides related helpers.
package ligo

import (
	"context"
	"fmt"

	"github.com/linode/linodego"
	"golang.org/x/oauth2"

	"log"
	"net/http"
	"os"
)

// NewClient creates a new Linode Client using the given LINODE_TOKEN
func NewClient() {
	apiKey, ok := os.LookupEnv("LINODE_TOKEN")
	if !ok {
		log.Fatal("Could not find LINODE_TOKEN, please assert it is set.")
	}
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: apiKey})

	oauth2Client := &http.Client{
		Transport: &oauth2.Transport{
			Source: tokenSource,
		},
	}

	linodeClient := linodego.NewClient(oauth2Client)

	if _, ok := os.LookupEnv("LINODE_DEBUG"); ok {
		linodeClient.SetDebug(true)
	}

	res, err := linodeClient.GetInstance(context.Background(), 4090913)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", res)
}
