package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// mockDownload()
	// mockBarChart()
	// bigText()
	// bulletPoints()
	header()
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func header() {
	pterm.DefaultHeader.
		WithBackgroundStyle(pterm.NewStyle(pterm.BgLightRed)).
		WithFullWidth().
		Println("This is a full-width header.")
}

func bulletPoints() {
	err := pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "Level 0", Bullet: ">"},
		{Level: 1, Text: "Level 1"},
		{Level: 2, Text: "Level 2"},
		{Level: 0, Text: "Level 0", Bullet: ">"},
		{Level: 1, Text: "Level 1"},
		{Level: 2, Text: "Level 2"},
	}).Render()

	handleError(err)
}

func bigText() {
	err := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle(
			"Cimux",pterm.NewStyle(pterm.FgCyan)),
	).Render()
	handleError(err)
}

func mockBarChart() {
	bars := pterm.Bars{
		pterm.Bar{
			Label: "Bar 1",
			Value: 3,
		},
		pterm.Bar{
			Label: "Bar 2",
			Value: 3,
		},
		pterm.Bar{
			Label: "Longer Label",
			Value: 4,
		},
	}

	err := pterm.DefaultBarChart.WithHorizontal().WithBars(bars).Render()
	handleError(err)

	err2 := pterm.DefaultBarChart.WithBars(bars).Render()
	handleError(err2)
}

func mockDownload() {
	packages := []string{
		"Nginx",
		"MySQL",
		"Git",
	}

	progressBar, err := pterm.DefaultProgressbar.WithTotal(len(packages)).Start()
	handleError(err)

	for i := 0; i < progressBar.Total; i++ {
		progressBar.UpdateTitle("Downloading " + packages[i])
		pterm.Success.Println("Downloading " + packages[i])
		progressBar.Increment()
		time.Sleep(time.Millisecond * 350)
	}
}
