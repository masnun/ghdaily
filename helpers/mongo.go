package helpers

import (
        //"fmt"
        "log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "time"
)

const MONGOHOST string = "localhost"
const MONGODB string = "ghdaily"
const MONGOCOLLECTION string = "repos"

func InsertRepo(repo Repo) bool {
        session, err := mgo.Dial(MONGOHOST)
        if err != nil {
                panic(err)
        }
        defer session.Close()

        session.SetMode(mgo.Monotonic, true)

        c := session.DB(MONGODB).C(MONGOCOLLECTION)
        err = c.Insert(repo)
        if err != nil {
                log.Fatal(err)
                return false
        } else {
                return true
        }

}

func GetRepos(lang string, day time.Time) []Repo {
        next_day := day.Add(24 * time.Hour)

        session, err := mgo.Dial(MONGOHOST)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(MONGODB).C(MONGOCOLLECTION)

        var repos []Repo
        c.Find(
            bson.M{
                "createdat": bson.M{
                    "$gt": day,
                    "$lt": next_day,
                },

                "language": lang,


        }).All(&repos)

        return repos
        
}

func Exists(title string) bool {
        year, month, day := time.Now().Date()
        today := time.Date(year, month, day, 0, 0, 0, 0, time.UTC )

        session, err := mgo.Dial(MONGOHOST)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(MONGODB).C(MONGOCOLLECTION)

        var repos []Repo
        c.Find(
            bson.M{
                "createdat": bson.M{
                    "$gt": today,
                },

                "name": title,

        }).All(&repos)

        return len(repos) > 0
}



