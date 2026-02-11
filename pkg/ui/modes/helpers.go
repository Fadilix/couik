package modes

func StringToMode(mode string) ModeStrategy {
	var finalMode ModeStrategy
	switch mode {
	case "quote":
		finalMode = NewQuoteMode()
	case "words":
		finalMode = NewWordMode()
	case "time":
		finalMode = NewTimeMode()
	}
	return finalMode
}
