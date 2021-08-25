package calendarparser

import (
	"fmt"
	"github.com/apognu/gocal"
	"os"
	"strings"
	"time"
)

type CalendarInformation struct {
	EventName    string
	DateStart    time.Time
	DateEnd      time.Time
	Floor        string
	Room         string
	Presentation bool
}

func Initialization() {

	//TODO: make this load more files and also into an array
	f, _ := os.Open("calendar_files/01cal.ics")
	defer f.Close()

	start, end := time.Now(), time.Now().Add(12*30*24*time.Hour)

	c := gocal.NewParser(f)
	c.Start, c.End = &start, &end
	c.Parse()

	ParseIntoVariables(c)

}

func ParseIntoVariables(c *gocal.Gocal) {
	var todaysDate time.Time = time.Now()
	//2D slice array of [X][Y]. X is the room, Y is the subjects within it.
	var CalendarEvents [5][40]CalendarInformation
	var location []string
	var locationProcessed []string
	var floor string
	var room string
	var presentationBool bool
	var summary []string
	var summaryProcessed []string
	var summaryTemp []string

	//Search through the entire file
	//TODO: make it an array
	for i := 0; i < 5; i++ {
		for _, e := range c.Events {
			location = strings.Split(e.Location, ", ")
			locationProcessed = strings.Split(location[0], ".")
			floor = locationProcessed[0]
			room = locationProcessed[1]

			if strings.Contains(e.Summary, " C ") {
				summary = strings.Split(e.Summary, " C ")
				presentationBool = false
			} else if strings.Contains(e.Summary, " P ") {
				summary = strings.Split(e.Summary, " P ")
				presentationBool = true
			}

			summaryTemp = strings.SplitN(summary[0], " ", -1)

			summaryProcessed = summaryTemp[2:]

			fmt.Println("-----\nNESPRACOVANE: ", summary)
			fmt.Println("SPRACOVANE: ", summaryProcessed, "\n------")
			fmt.Println(presentationBool)

			fmt.Println("\nTMP:")
			fmt.Println(summary)

			for j := 0; j < 40; j++ {
				if e.Start.Day() == todaysDate.Day() {

					CalendarEvents[i][j].EventName = e.Summary
					CalendarEvents[i][j].DateStart = *e.Start
					CalendarEvents[i][j].DateEnd = *e.End
					CalendarEvents[i][j].Floor = floor
					CalendarEvents[i][j].Room = room
					CalendarEvents[i][j].Presentation = presentationBool

				}

			}

		}
		//fmt.Println(CalendarEvents[0][0].EventName)
	}

}
