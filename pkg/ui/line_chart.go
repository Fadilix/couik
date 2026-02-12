package ui

import (
	"fmt"
	"time"

	"github.com/NimbleMarkets/ntcharts/linechart"
	"github.com/NimbleMarkets/ntcharts/linechart/timeserieslinechart"
	"github.com/charmbracelet/lipgloss"
)

func DisplayChart(wpmSamples []float64, timesSample []time.Time, width, height int) string {
	if len(wpmSamples) == 0 {
		return lipgloss.NewStyle().
			Foreground(CatOverlay).
			Width(width).
			Height(height).
			Align(lipgloss.Center, lipgloss.Center).
			Render("no data")
	}

	axisStyle := lipgloss.NewStyle().Foreground(CatSurface)
	labelStyle := lipgloss.NewStyle().Foreground(CatOverlay)

	yFormatter := func(i int, v float64) string {
		return fmt.Sprintf("%d", int(v))
	}

	tslc := timeserieslinechart.New(width, height,
		timeserieslinechart.WithAxesStyles(axisStyle, labelStyle),
		timeserieslinechart.WithXYSteps(0, 2),
		timeserieslinechart.WithYLabelFormatter(linechart.LabelFormatter(yFormatter)),
		timeserieslinechart.WithStyle(lipgloss.NewStyle().Foreground(CatMauve)),
	)

	for i, v := range wpmSamples {
		tslc.Push(timeserieslinechart.TimePoint{Time: timesSample[i], Value: v})
	}
	tslc.DrawBraille()

	return tslc.View()
}
