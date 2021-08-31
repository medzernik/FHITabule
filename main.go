package main

import (
	"FHITabule/calendarparser"
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"log"
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

	go runLocalWebpageDisplay()

	calendarparser.Initialization()

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
		var w, _ = a.NewWindow("https://imhd.sk/ba/online-zastavkova-tabula?theme=white&zoom=67&st=66", &astilectron.WindowOptions{
			Center:          astikit.BoolPtr(true),
			Height:          astikit.IntPtr(600),
			Width:           astikit.IntPtr(600),
			Minimizable:     astikit.BoolPtr(false),
			Fullscreen:      astikit.BoolPtr(false),
			BackgroundColor: astikit.StrPtr("black"),
		})
		w.Create()

		//TODO: google API cesta na hopu MHD + auto
		var links []string

		links = append(links, "https://imhd.sk/ba/online-zastavkova-tabula?theme=white&zoom=67&st=66", "http://localhost:3000", "file:///C:/Users/medze/Desktop/workspace/Programovanie/Go/FHITabule/presentation_images/Slide1.PNG")

		imagePath, errDir := os.ReadDir("presentation_images")
		if errDir != nil {
			fmt.Println("ERROR READING THE DIRECTORY", errDir)
		}

		fmt.Println(imagePath[0].Name())

		for _, i := range imagePath {
			links = append(links, "file:///C:/Users/medze/Desktop/workspace/Programovanie/Go/FHITabule/presentation_images/"+i.Name())
		}

		for {
			for i := range links {
				err := w.ExecuteJavaScript("window.location.href = \"" + links[i] + "\";")
				if err != nil {
					fmt.Println("ERROR:", err)
				}
				time.Sleep(3 * time.Second)
			}
		}

		// Blocking pattern
		a.Wait()
	}

}
