package controllers

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"
)

type Commitment struct {
	*revel.Controller
}

type Commit struct {
	User        string
	Date        time.Time
	Description string
	Status      string
}

func (c Commitment) Create(user string, description string, date string) revel.Result {
	c.Validation.Required(user)
	c.Validation.Required(description)
	c.Validation.Required(date)

	commit := Commit{user, parseDate(date), description, "created"}

	session, _ := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	collection(session).Insert(&commit)
	defer session.Close()

	return c.RenderJson(commit)
}

func (c Commitment) Get(user string) revel.Result {
	result := Commit{}

	session, _ := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	err := collection(session).Find(bson.M{"user": user, "status": "expired"}).One(&result)
	if err != nil {
		return c.NotFound("No expired commitments")
	}
	defer session.Close()

	return c.RenderJson(result)
}

func parseDate(date string) time.Time {
	// TODO: Parse date with time too in correct timezone
	const shortForm = "2006-01-15"
	parsedDate, _ := time.Parse(shortForm, date)
	return parsedDate
}

func collection(session *mgo.Session) *mgo.Collection {
	// TODO: Add error handling
	return session.DB("commit").C("commitments")
}
