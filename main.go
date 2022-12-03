package main

import (
	"fmt"
	"log"

	"github.com/Davincible/goinsta/v3"
)

func init() {
	insta, err := goinsta.Import("./goinsta-session")
	if err != nil {
		log.Println("Logging again")
		insta = goinsta.New(username, password)
		err := insta.Login()
		if err != nil {
			log.Fatal("login failed: ", err)
		}

		err = insta.Export("./goinsta-session")
		if err != nil {
			log.Fatal("session export failed: ", err)
		}
	}

	err = insta.Inbox.Sync()
	if err != nil {
		log.Printf("syncing inbox failed: %+v", err)
	}

	stories, err := insta.GetMedia("2959723248254424644_25363180311")
	if err != nil {
		log.Print(err)
	}
	err = stories.Sync()
	if err != nil {
		log.Print(err)
	}
	err = stories.GetCommentInfo()
	if err != nil {
		log.Print(err)
	}
	// stories.Items[0].Comments.Sync()
	err = stories.Items[0].Comments.Add("@bolorcln _!_")
	if err != nil {
		log.Print(err)
	}
}

func main() {
	fmt.Println("Hello")
}
