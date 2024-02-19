package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/gocolly/colly"
)

type Config struct {
	Version string `json:"version"`
	Delay   int    `json:"delay"`
}

func main() {
	fmt.Println("Loading files...")

	// Create default config
	var defaultConfig Config
	defaultConfig.Delay = 250
	defaultConfig.Version = "v1.2.2"

	var config Config
	configPath := "config.json"

	if FileExists(configPath) {
		config = LoadConfig(configPath)
		fmt.Println(config)
		fmt.Println("Loaded config.json")
	} else {
		// Create config from default config
		config = CreateConfig(defaultConfig, configPath)
		fmt.Println("Created config.json")
	}

	fmt.Println("Loaded files")

	fixedDelay := time.Duration(config.Delay) * time.Millisecond

	fmt.Println("Starting EMatDystenBot " + config.Version)

	// Checking for updates
	fmt.Println("Checking for updates...")
	latest := getLatestVersion()
	if config.Version == latest {
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
	fmt.Println("Starting to harvest. You can stop at any time by closing this window.")

	GoToStartingPage(fixedDelay)

	time.Sleep(fixedDelay)

	for true {
		StartExercise(fixedDelay)
		time.Sleep(fixedDelay)
		SolveExercise(fixedDelay)
		time.Sleep(fixedDelay)
	}
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func CreateConfig(config Config, path string) Config {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return config
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return config
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return config
	}

	return config
}

func LoadConfig(path string) Config {
	var config Config

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return config
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return config
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}

	return config
}

func StartExercise(delay time.Duration) {
	robotgo.TypeStr("ctl00_ContentPlaceHolder1_ImageButton_RegnehierarkietS.click()")
	robotgo.KeyTap("enter")
	time.Sleep(delay)
	robotgo.TypeStr("Link_Start2.click()")
	robotgo.KeyTap("enter")
}

func SolveExercise(delay time.Duration) {
	robotgo.TypeStr("for(let index=0;index<13;index++){Textbox_BrugersvarA.value=FacitA[ValgtSpg];img3.click();}")
	robotgo.KeyTap("enter")
	time.Sleep(delay)
	robotgo.TypeStr("img4.click()")
	robotgo.KeyTap("enter")
	time.Sleep(delay)
	robotgo.TypeStr("ctl00_ContentPlaceHolder1_ImageButton_VSelvStart.click()")
	robotgo.KeyTap("enter")
}

func GoToStartingPage(delay time.Duration) {
	robotgo.TypeStr("ctl00_ImageButton_Dysten.click()")
	robotgo.KeyTap("enter")
	time.Sleep(delay)
	robotgo.TypeStr("ctl00_ContentPlaceHolder1_ImageButton_StartDysten.click()")
	robotgo.KeyTap("enter")
	time.Sleep(delay)
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
