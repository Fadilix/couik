package cli

import "github.com/charmbracelet/bubbles/table"

var (
	Columns = []table.Column{
		{Title: "WPM", Width: 10},
		{Title: "RawWPM", Width: 10},
		{Title: "Acc", Width: 10},
		{Title: "Duration", Width: 10},
		// {Title: "Date", Width: 10},
		// {Title: "Quote", Width: 10},
	}

	Rows = []table.Row{}
)
