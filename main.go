package main

import (
	"fmt"
	"runtime"
	"github.com/SlyMarbo/rss"
	"strconv"
	"github.com/streamrail/concurrent-map"
)

func rssFetch(errorMap cmap.ConcurrentMap, jobs <-chan string, results chan<- rss.Feed){
	for job := range jobs {
		feed, err := rss.Fetch(job)
		if err != nil {
			errorMap.Set("error", err)
			results <- *new(rss.Feed)
			break
		}
		results <- *feed
	}
}

func main() {
	processors := runtime.NumCPU() - 1
	tasks := processors * 2
	//TODO ta inn query i function
	//TODO hit cache før søk
	results := make(chan rss.Feed, 100)
	jobs := make(chan string, 100)
	errMap := cmap.New()
	feedMap := make(map[string]rss.Feed)

	baseUrl := "http://www.nyaa.se/?page=rss&term=shingeki+no+kyojin&offset="

	num := 0

	for worker := 1; worker <= processors; worker++ {
		go rssFetch(errMap, jobs, results)
	}

	for {
		for j := (tasks * num) + 1; j <= (tasks * (num + 1)); j++ {
			jobs <- (baseUrl + strconv.Itoa(j))
		}

		for a := (tasks * num) + 1; a <= (tasks * (num + 1)); a++ {
			feed := <-results
			feedMap[feed.UpdateURL] = feed
		}

		if errMap.Count() > 0 {
			break
		}
		num++
	}
	close(jobs)
	close(results)
	for feed := range feedMap {
		fmt.Println(feed)
		for item := range feedMap[feed].ItemMap{
			fmt.Println(item)
		}
	}
}
//TODO cache search-entries in map[query][map[rss.Feed]][entries]? med TTL?