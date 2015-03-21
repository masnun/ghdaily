#GHDaily

Scrapes the Github daily trending page for several programming languages and stores them in a mongo db collection. 

### Setting Up

First, we need to install the dependencies:

```
go get -u github.com/PuerkitoBio/goquery
go get -u gopkg.in/mgo.v2
go get -u github.com/zenazn/goji
```

Then we can run the server: 

	go run main.go

Or you can use `go build` to create a single binary and run it. 

### Configuration

By default, it tries to connect to a Mongodb using the url `localhost`. The default database name is `ghdaily` and the collection is named `repos`. These are configurable in the `helpers/mongo.go` file. There are 3 constants for each of these settings, update them to your liking/setup. 

### URL Patterns

##### Cron Job

You can start processing the repos for the day by visiting the url: 

	http://localhost:8000/cron

That would start a background process to scrape Github and store the results. 

You would usually want to setup a cron job to hit this url daily. I could write a CLI to do this but this is just a hobby project and I'm lazy.

##### Repos List

The trending repos for a particular date can be viewed from: 

	http://localhost:8000/list/language/dd-mm-yyyy

For example:

	http://localhost:8000/list/php/21-03-2015

This URL would contain the trending PHP repos for 21st March, 2015. 

### Selecting The Languages

Find the `ProcessLanguages` function in `main.go` file. There's a list named `languages` inside a go routine. Update the list as required. 

### Demos? 

No, too lazy to set one up. 



