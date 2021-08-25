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

type CurrentTime struct {
	StartTime time.Time
	EndTime   time.Time
}

var Hodina uint

func Initialization() {

	var timeToDisplayHours time.Duration = 6

	//TODO: make this load more files and also into an array
	f, _ := os.Open("calendar_files/01cal.ics")
	defer f.Close()

	//*10000 is a debug
	start, end := time.Date(2021, 10, 11, 10, 0, 0, 0, time.Now().Location()), time.Date(2021, 10, 11, 10, 0, 0, 0, time.Now().Location()).Add(timeToDisplayHours*time.Hour)

	c := gocal.NewParser(f)
	c.Start, c.End = &start, &end
	c.Parse()

	ParseIntoVariables(c)

}

func ParseIntoVariables(c *gocal.Gocal) {
	//var todaysDate time.Time = time.Date(2021, 9, 28, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Now().Location())
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
	var iterator int = 0

	//Search through the entire file
	//TODO: make it an array

	for _, e := range c.Events {
		fmt.Printf("%s on %s - %s\n", e.Summary, e.Start, e.End)

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

		//fmt.Printf("%+v\n", e)

		CalendarEvents[iterator].EventName = strings.Join(summaryProcessed[0:], " ")
		CalendarEvents[iterator].DateStart = *e.Start
		CalendarEvents[iterator].DateEnd = *e.End
		CalendarEvents[iterator].Floor = floor
		CalendarEvents[iterator].Room = room
		CalendarEvents[iterator].Presentation = presentationBool

		iterator += 1

	}

	//fmt.Printf("%+v\n", CalendarEvents)
	//fmt.Println(len(CalendarEvents))
	PrintClasses(CalendarEvents)

}

func PrintClasses(CalendarEvents []CalendarInformation) {

	//TODO: POROVNAT CAS A VYPISAT LEN TO CO TREBA V SUCASNU HODINU
	timeCurrentTemp := time.Date(2021, 10, 11, 11, 46, 0, 0, time.Now().Location())

	for l := range CalendarEvents {
		if AssignTime(timeCurrentTemp) == AssignTime(CalendarEvents[l].DateStart) {
			fmt.Println("---------------------------")
			fmt.Println("POSCHODIE: ", CalendarEvents[l].Floor)
			if CalendarEvents[l].Presentation == true {
				fmt.Println("PREDNASKA")
			} else {
				fmt.Println("CVICENIE")
			}
			fmt.Println("PREDMET: ", CalendarEvents[l].EventName)
			fmt.Println("V MIESTNOSTI", CalendarEvents[l].Floor+"."+CalendarEvents[l].Room)
			fmt.Println("CAS: ", CalendarEvents[l].DateStart.String())

		}
		//fmt.Printf("%+v\n", CalendarEvents[l])

	}
}

func AssignTime(inputTime time.Time) int {
	var sucasnaHodina int

	prvaHodina := time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), 7, 15, 0, 0, time.Now().Location())
	druhaHodina := time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), 9, 20, 0, 0, time.Now().Location())
	tretiaHodina := time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), 10, 45, 0, 0, time.Now().Location())
	stvrtaHodina := time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), 12, 30, 0, 0, time.Now().Location())
	piataHodina := time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), 15, 0, 0, 0, time.Now().Location())
	siestaHodina := time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), 16, 45, 0, 0, time.Now().Location())
	siedmaHodina := time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), 18, 25, 0, 0, time.Now().Location())
	osmaHodina := time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), 20, 00, 0, 0, time.Now().Location())

	if inputTime.Before(prvaHodina) {
		sucasnaHodina = 1
	} else if inputTime.Before(druhaHodina) && inputTime.After(prvaHodina) {
		sucasnaHodina = 2
	} else if inputTime.Before(tretiaHodina) && inputTime.After(druhaHodina) {
		sucasnaHodina = 3
	} else if inputTime.Before(stvrtaHodina) && inputTime.After(tretiaHodina) {
		sucasnaHodina = 4
	} else if inputTime.Before(piataHodina) && inputTime.After(stvrtaHodina) {
		sucasnaHodina = 5
	} else if inputTime.Before(siestaHodina) && inputTime.After(piataHodina) {
		sucasnaHodina = 6
	} else if inputTime.Before(siedmaHodina) && inputTime.After(siestaHodina) {
		sucasnaHodina = 7
	} else if inputTime.Before(osmaHodina) && inputTime.After(siedmaHodina) {
		sucasnaHodina = 8
	}

	fmt.Println("HODINA SUCASNA JE:", sucasnaHodina)
	return sucasnaHodina

}
