package calendarparser

import (
	"fmt"
	"github.com/apognu/gocal"
	"os"
	"time"
)

type CalendarInformation struct {
	EventName string
	DateStart time.Time
	DateEnd   time.Time
}

func Initialization() {
	f, _ := os.Open("calendar_files/01cal.ics")
	defer f.Close()

	start, end := time.Now(), time.Now().Add(12*30*24*time.Hour)

	c := gocal.NewParser(f)
	c.Start, c.End = &start, &end
	c.Parse()

	//2D slice array of [X][Y]. X is the room, Y is the subjects within it.
	var CalendarEvents [][]CalendarInformation

	ParseIntoVariables(c, &CalendarEvents)

}

func ParseIntoVariables(c *gocal.Gocal, calendarEvents *[][]CalendarInformation) {
	//var i int
	//var j int

	//Search through the entire file
	for _, e := range c.Events {
		fmt.Println(e.Summary)
		//if e.
	}

}
