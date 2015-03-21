package main

import (
	"fmt"
	helpers "github.com/masnun/ghdaily/helpers"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
	"time"
	//"encoding/json"
)

func main() {

	goji.Get("/list/:lang/:date", ListRepos)
	goji.Get("/cron", ProcessLanguages)
	goji.Serve()

}

func ProcessLanguages(c web.C, w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%s", "Processing started")

	go func() {
		languages := []string{
			"php",
			"python",
			"javascript",
			"java",
			"go",
		}
		for _, lang := range languages {
			var repos []helpers.Repo
			repos = helpers.ScrapeGithub(lang)
			for _, repo := range repos {
				if !helpers.Exists(repo.Title) {
					helpers.InsertRepo(repo)
				}

			}
		}
	}()

}

func ListRepos(c web.C, w http.ResponseWriter, r *http.Request) {
	lang := c.URLParams["lang"]
	date, _ := time.Parse("02-01-2006", c.URLParams["date"])

	repos := helpers.GetRepos(lang, date)
	//response, _ := json.Marshal(repos)
	if len(repos) > 0 {
		for _, repo := range repos {
			fmt.Fprintf(w, "<a target='_blank' href='https://github.com/%s'>%s</a> - %s <br/>", repo.Title, repo.Title, repo.Desc)
		}
	} else {
		fmt.Fprintf(w, "There is no data for this language for this date!")
	}

}
