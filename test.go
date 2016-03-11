package main

import "fmt"
import "github.com/SlyMarbo/rss"

func main() {
	feed, err := rss.Fetch("http://www.nyaa.se/?page=rss&term=shingeki+no+kyojin&offset=1")
	if err != nil || feed == nil {
		fmt.Println("Couldn't fetch rss-feed")
		return
	}

	// do something
	fmt.Println(feed.String())

	err = feed.Update()
	if err != nil {
		// handle error.
		fmt.Println(err.Error())
	}
}
