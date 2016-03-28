package main

import "fmt"
import "strings"
import "strconv"
import "github.com/SlyMarbo/rss"

var seriesList []Series

func main() {
	feed, err := rss.Fetch("http://www.nyaa.se/?page=rss&cats=1_37&term=boku+dake+ga+inai+machi&sort=2")
	if err != nil || feed == nil {
		fmt.Println("Couldn't fetch rss-feed")
		return
	}

	// do something

	for _, item := range feed.Items {
		//fmt.Println(item.String())

		episodeTitle, subgroup, episodenr := stripEpisodeTitle(item.Title), stripSubgroup(item.Title), stripEpisodeNr(item.Title)
		episode := Episode{episodeTitle, subgroup, episodenr}

		// episode not valid
		if episode.Title == "" {
			continue
		}

		// find series
		seriesTitle := stripSeriesTitle(item.Title)
		series := findSeries(seriesTitle)

		// add to series if exists, create new if it doesn't
		if series.Title != "" {
			episodeList := series.Episodes
			episodeList = append(episodeList, episode)
		} else {
			episodeList := []Episode{episode}
			series = Series{stripSeriesTitle(item.Title), episodeList}
			seriesList = append(seriesList, series)
		}
	}

	// for testing
	for _, series := range seriesList {
		fmt.Println(series.Title)
		for _, episode := range series.Episodes {
			fmt.Println("-", episode.Title)
			fmt.Println("-", episode.Subgroup)
			fmt.Println("-", episode.EpisodeNr)
		}
	}

	// errors
	err = feed.Update()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func stripSeriesTitle(text string) string {
	pre := strings.Index(text, "]")
	post := strings.LastIndex(text, "[")
	post2 := strings.LastIndex(text, "-")
	if post2 < post && post2 > 0 {
		post = post2
	}
	if pre < 0 || post < 0 {
		return ""
	}
	text = text[pre+1 : post]
	return strings.TrimSpace(text)
}

/*
func stripEpisode(text string) string {
	subgroup := strings.Index(text, "]")+1

	return ""
}*/

func stripEpisodeTitle(text string) string {
	pre := strings.Index(text, "]")
	post := strings.LastIndex(text, "[")
	if pre < 0 || post < 0 {
		return ""
	}
	text = text[pre+1 : post]
	return strings.TrimSpace(text)
}

func stripSubgroup(text string) string {
	pre := strings.Index(text, "[")
	post := strings.Index(text, "]")
	if pre < 0 || post < 0 {
		return ""
	}
	text = text[pre+1 : post]
	return strings.TrimSpace(text)
}

func stripEpisodeNr(text string) int {
	fmt.Println(text)
	pre := strings.Index(text, "-")
	text = strings.TrimSpace(text[pre+1:])
	post := strings.Index(text, " ")
	if pre < 0 || post < 0 {
		return -1
	}
	text = text[:post]
	fmt.Println(strings.TrimSpace(text))
	i, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return -2
	}
	return i
}

func findSeries(title string) Series {
	for _, series := range seriesList {
		if series.Title == title {
			return series
		}
	}
	return Series{"", nil}
}

// whole series containing multiple episodes
type Series struct {
	Title    string
	Episodes []Episode
}

// a single episode
type Episode struct {
	Title     string
	Subgroup  string
	EpisodeNr int
	/*Quality     string
	DownloadURL string
	FilterType  string
	Seeders     int
	Size        string
	Downloads   int*/
}
