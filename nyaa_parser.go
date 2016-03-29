package main

import "strings"
import "strconv"
import (
	"regexp"
	"fmt"
)

func StripSeriesTitle(text string) string {
	whiteSpaceRegex, _ := regexp.Compile("_")
	titleRegex, _ := regexp.Compile("\\w+")
	subGroupEnd := strings.Index(text, "]") + 1
	text = whiteSpaceRegex.ReplaceAllString(text, " ")
	text = text[subGroupEnd:]
	title := titleRegex.FindString(text)
	title = strings.TrimSpace(title)
	fmt.Println(title)
	return title

}

func stripSubgroup(text string) string {
	pre := strings.Index(text, "[")
	post := strings.Index(text, "]")
	if pre < 0 || post < 0 {
		return ""
	}
	text = text[pre + 1 : post]
	return strings.TrimSpace(text)
}

func StripEpisodeNum(text string) int {
	regex, _ := regexp.Compile("\\d+")
	pre := strings.Index(text, "-")
	if pre < 0 {
	}else {
		text = strings.TrimSpace(text[pre + 1:])
	}
	text = regex.FindString(text)
	i, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return -2
	}
	return i
}

// whole series containing multiple episodes
type Series struct {
	title    string
	episodes []Episode
	seeders  int
}

// a single episode
type Episode struct {
	title       string
	subGroup    string
	episodeNum  int
	quality     string
	downloadURL string
	aPlus       bool
	FilterType  string
	seeders     int
	size        string
}