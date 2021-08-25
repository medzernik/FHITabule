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
	//var todaysDate time.Time = time.Now()
	//2D slice array of [X][Y]. X is the room, Y is the subjects within it.

	var location []string
	var locationProcessed []string
	var floor string
	var room string
	var presentationBool bool
	var summary []string
	var summaryProcessed []string
	var summaryTemp []string
	var CalendarEvents []CalendarInformation = make([]CalendarInformation, len(c.Events))
	//Search through the entire file
	//TODO: make it an array

	for _, e := range c.Events {
		for i := 0; i < len(c.Events); i++ {

			//Parsing the location and getting a room number and floor number
			location = strings.Split(e.Location, ", ")
			locationProcessed = strings.Split(location[0], ".")
			floor = locationProcessed[0]
			room = locationProcessed[1]

			//Parsing the ridiculous and messy Summary line. First we get the information about whether it's a course or presentation
			if strings.Contains(e.Summary, " C ") {
				summary = strings.Split(e.Summary, " C ")
				presentationBool = false
			} else if strings.Contains(e.Summary, " P ") {
				summary = strings.Split(e.Summary, " P ")
				presentationBool = true
			}

			//Then after splitting the line in a slice we further split the slice by every space
			summaryTemp = strings.SplitN(summary[0], " ", -1)

			//We know that the name of the subject can be multiple words but it always starts after the first--
			//--2 space members. We get that until the end (end was split by C or P, since its always before that)
			summaryProcessed = summaryTemp[2:]

			//Important debug statements
			/*
				fmt.Println("-----\nNESPRACOVANE: ", summary)
				fmt.Println("SPRACOVANE: ", summaryProcessed, "\n------")
				fmt.Println(presentationBool)

				fmt.Println("\nTMP:")
				fmt.Println(summary)

			*/

			CalendarEvents[i].EventName = strings.Join(summaryProcessed[0:], " ")
			CalendarEvents[i].DateStart = *e.Start
			CalendarEvents[i].DateEnd = *e.End
			CalendarEvents[i].Floor = floor
			CalendarEvents[i].Room = room
			CalendarEvents[i].Presentation = presentationBool

		}
		fmt.Printf("%+v\n", CalendarEvents)
		fmt.Println(len(CalendarEvents))
	}

}
