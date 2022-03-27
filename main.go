package main

import (
	_ "embed"
	"errors"
	stats "gocheck/statspkg"
	"os"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails"
)

//go:embed frontend/public/build/bundle.js
var js string

//go:embed frontend/public/build/bundle.css
var svelteCss string

//go:embed frontend/public/build/tailwind.css
var tailwindCss string

//go:embed frontend/public/index.html
var html string

func Exists(name string) (bool, error) {
	fi, err := os.Stat(name)
	if err == nil && !fi.IsDir() {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func main() {
	if exists, _ := Exists(".env"); exists {
		err := godotenv.Load()
		if err != nil {
			println("Error loading .env file")
		}
	}

	cpuStats := &stats.CpuStats{}
	memStats := &stats.MemStats{}

	app := wails.CreateApp(&wails.AppConfig{
		Width:     450,
		Height:    800,
		Title:     "gocheck",
		Resizable: true,
		HTML:      html,
		JS:        js,
		CSS:       svelteCss + "\n\n" + tailwindCss,
		Colour:    "#ffffff",
	})

	app.Bind(cpuStats)
	app.Bind(memStats)

	app.Run()
}
