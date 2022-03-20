package main

import (
	"context"
	"fmt"
	"log"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/tenntenn/connpass"
)

func main() {
	cli := connpass.NewClient()
	params, err := connpass.SearchParam(connpass.KeywordOr("go"), connpass.KeywordOr("golang"))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	r, err := cli.Search(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodPublish)
	for _, e := range r.Events {
		event := cal.AddEvent(fmt.Sprintf("connpass-%d", e.ID))
		event.SetCreatedTime(e.UpdatedAt)
		event.SetDtStampTime(time.Now())
		event.SetModifiedAt(e.UpdatedAt)
		event.SetStartAt(e.StartedAt)
		event.SetEndAt(e.EndedAt)
		event.SetSummary(e.Title)
		event.SetDescription(e.Description)
		event.SetURL(e.URL)
	}

	s := cal.Serialize()
	fmt.Println(s)
}
