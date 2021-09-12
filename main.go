package main

import (
	"FHITabule/calendarparser"
	"FHITabule/config"
	"FHITabule/distance"
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func runLocalWebpageDisplay() {
	command := "/home/medzernik/.nvm/versions/node/v16.8.0/bin/node app.js"
	parts := strings.Fields(command)
	_, err := exec.Command(parts[0], parts[1:]...).Output()
	if err != nil {
		fmt.Println("ERROR LOADING THE WEBSITE FORMATTER: ", err)
	}

}

func main() {
	config.Initialization()
	//go runLocalWebpageDisplay()
	go StartCustomServer()

	calendarparser.Initialization()
	minutesTraveled := distance.GetTimeToHOPA("University of Economics in Bratislava, Dolnozemská cesta 1, 852 35 Petržalka, Slovakia", "Študentský domov Prokopa Veľkého, \"HOPA\", Prokopa Veľkého 41, 811 04 Bratislava, Slovakia")
	fmt.Println("Travel time to HOPA: ", minutesTraveled, " minut")

	var runElectron = true

	if runElectron == true {

		// Make HTTP GET request
		//response, err := http.Get("https://imhd.sk/ba/online-zastavkova-tabula?theme=white&zoom=67&st=66")
		var a, _ = astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
			AppName:            "DAVID MANCA - EUBA TABULE",
			AppIconDefaultPath: "", // If path is relative, it must be relative to the data directory
			AppIconDarwinPath:  "", // Same here
			VersionAstilectron: "",
			VersionElectron:    "",
			SingleInstance:     true,
		})
		defer a.Close()

		// Start astilectron
		a.Start()
		var w, _ = a.NewWindow("http://localhost:3000", &astilectron.WindowOptions{
			Center:          astikit.BoolPtr(true),
			Height:          astikit.IntPtr(600),
			Width:           astikit.IntPtr(600),
			Minimizable:     astikit.BoolPtr(false),
			Fullscreen:      astikit.BoolPtr(false),
			BackgroundColor: astikit.StrPtr("black"),
		})
		w.Create()
		w.OpenDevTools()

		//TODO: google API cesta na hopu MHD + auto
		var linksMain []string
		var linksPresentation []string

		linksMain = append(linksMain, "https://imhd.sk/ba/online-zastavkova-tabula?theme=white&zoom=67&st=66", "http://localhost:3000")

		imagePath, errDir := os.ReadDir("presentation_images")
		if errDir != nil {
			fmt.Println("ERROR READING THE DIRECTORY", errDir)
		}

		fmt.Println(imagePath[0].Name())

		for _, i := range imagePath {
			linksPresentation = append(linksPresentation, "http://localhost:3000/presentation_images/"+i.Name())
		}

		for {
			for i := range linksMain {
				err := w.ExecuteJavaScript("window.location.href = \"" + linksMain[i] + "\";")
				if err != nil {
					fmt.Println("ERROR:", err)
				}
				time.Sleep(3 * time.Second)
				err1 := w.ExecuteJavaScript("window.location.href = \"" + linksPresentation[i] + "\";")
				if err1 != nil {
					fmt.Println("ERROR:", err1)
				}
				time.Sleep(3 * time.Second)
			}

		}

		// Blocking pattern
		a.Wait()
	}

}

func StartCustomServer() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}
