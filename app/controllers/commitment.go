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
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	User        string        `json:"user"`
	Date        time.Time     `json:"date"`
	Week        int           `json:"week"`
	Description string        `json:"description"`
	Status      string        `json:"status"`
}

func (c Commitment) Create(user string, description string, date string) revel.Result {
	c.Validation.Required(user)
	c.Validation.Required(description)
	c.Validation.Required(date)

	_, week := parseDate(date).ISOWeek() //calculate week
	commit := Commit{bson.NewObjectId(), user, parseDate(date), week, description, "created"}
	session, _ := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	collection(session).Insert(&commit)
	defer session.Close()

	return c.RenderJson(commit)
}

func (c Commitment) Update(user string, status string, id string) revel.Result {
	c.Validation.Required(user)
	c.Validation.Required(id)
	c.Validation.Required(status)

	session, _ := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	err := collection(session).Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		panic(err)
	}

	defer session.Close()
	return c.RenderJson("ok")
}

func (c Commitment) Get(user string) revel.Result {
	result := Commit{}
	session, _ := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	err := collection(session).Find(bson.M{"user": user, "status": "created"}).One(&result)
	if err != nil {
		return c.NotFound("No current commitments")
	}

	defer session.Close()
	return c.RenderJson(result)
}

func (c Commitment) Percent(user string) revel.Result {
	session, _ := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	_, week := time.Now().ISOWeek()
	max, _ := collection(session).Find(bson.M{"user": user, "week": week}).Count()
	p, _ := collection(session).Find(bson.M{"user": user, "status": "completed", "week": week}).Count()

	result := map[string]int{"max": max, "p": p}

	defer session.Close()
	return c.RenderJson(result)
}

func parseDate(date string) time.Time {
	parsedDate, _ := time.Parse(time.RFC3339, date)
	return parsedDate
}

func collection(session *mgo.Session) *mgo.Collection {
	// TODO: Add error handling
	return session.DB("heroku_jl9vx3v3").C("commitments")
}
