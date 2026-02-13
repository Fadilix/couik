package modes

import (
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/ui/core"
)

func StringToMode(mode string, ctx core.ModeConfig) ModeStrategy {
	var finalMode ModeStrategy
	switch mode {
	case "quote":
		finalMode = NewQuoteMode(
			WithLanguageQ(ctx.Language),
			WithTargetQ(ctx.Target),
			WithCategoryQ(ctx.Category),
		)
	case "words":
		finalMode = NewWordMode(
			WithInitialWords(ctx.InitialWords),
			WithLanguageW(ctx.Language),
			WithTargetW(ctx.Target),
		)
	case "time":
		finalMode = NewTimeMode(
			WithTargetT(ctx.Target),
			WithLanguageT(ctx.Language),
			WithInitialTimeT(ctx.InitialTime),
		)
	default: // default to quote mode mid like the cmd/couik/main.go
		finalMode = NewQuoteMode(
			WithLanguageQ(ctx.Language),
			WithTargetQ(ctx.Target),
			WithCategoryQ(database.Mid),
		)
	}
	return finalMode
}
