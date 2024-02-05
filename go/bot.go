package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/gocolly/colly"
)

func main() {
	version := "v1.2.2"
	fmt.Println("Starting EMatDystenBot " + version)
	// Checking for updates
	fmt.Println("Checking for updates...")
	latest := getLatestVersion()
	if version == latest {
		fmt.Println("You're running the latest version of EMatDystenBot.")
		fmt.Println("(Press enter to start)")
	} else {
		latestname := getLatestVersionName()
		fmt.Println("There is a new update available for EMatDystenBot.")
		fmt.Println("Version: " + latest)
		fmt.Println("Name: " + latestname)
		fmt.Println("(Press enter to start anyways)")
	}

	fmt.Scanln()

	// Start chrome
	openBrowser("www.emat.dk")

	time.Sleep(1 * time.Second)

	// Open Devtools
	robotgo.KeyTap("f12")

	fmt.Println("Log into website and open console.")
	fmt.Println("(Press enter to start)")
	fmt.Scanln()

	fmt.Println("Return to console.")
	fmt.Println("Starting in 3")
	time.Sleep(1 * time.Second)
	fmt.Println("Starting in 2")
	time.Sleep(1 * time.Second)
	fmt.Println("Starting in 1")
	time.Sleep(1 * time.Second)
	fmt.Println("Starting to harvest. You can stop at any time by")

	GoToStartingPage()

	time.Sleep(250 * time.Millisecond)

	// Write text in console

	fmt.Scanln()
}

func StartExercise() {

}

func SolveExercise() {

}

func GoToStartingPage() {
	robotgo.TypeStr("ctl00_ImageButton_Dysten.click()")
	robotgo.KeyTap("enter")
	time.Sleep(250 * time.Millisecond)
	robotgo.TypeStr("ctl00_ContentPlaceHolder1_ImageButton_StartDysten.click()")
	robotgo.KeyTap("enter")
	time.Sleep(250 * time.Millisecond)
	robotgo.TypeStr("ctl00_ContentPlaceHolder1_ImageButton_StartNormal.click()")
	robotgo.KeyTap("enter")
}

func write(text string) {
	for _, char := range text {
		robotgo.KeyTap(string(char))
	}
}

func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

func getLatestVersion() string {
	version := "ERROR"
	c := colly.NewCollector()
	c.OnHTML("#version", func(e *colly.HTMLElement) {
		version = e.Text
	})
	c.Visit("https://sv3ks.github.io/EMatDystenBot/web/releases/latest/")
	return version
}

func getLatestVersionName() string {
	name := "ERROR"
	c := colly.NewCollector()
	c.OnHTML("#name", func(e *colly.HTMLElement) {
		name = e.Text
	})
	c.Visit("https://sv3ks.github.io/EMatDystenBot/web/releases/latest/")
	return name
}
