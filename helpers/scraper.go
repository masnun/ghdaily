package helpers

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"time"
)


func ScrapeGithub(lang string) []Repo {
	doc, err := goquery.NewDocument("https://github.com/trending?l=" + lang)
	if err != nil {
		log.Fatal(err)
	}

	repos := make([]Repo, 0)

	doc.Find(".repo-list-item").Each(func(index int, element *goquery.Selection) {

		title_parts := strings.Split(element.Find("h3.repo-list-name").Text(), "\n")
		for k, v := range title_parts {
			title_parts[k] = strings.TrimSpace(v)
		}

		title := strings.Join(title_parts, "")
		description := strings.TrimSpace(element.Find("p.repo-list-description").Text())
		repos = append(repos, Repo{title, description, lang, time.Now()})
	})

	return repos
}
