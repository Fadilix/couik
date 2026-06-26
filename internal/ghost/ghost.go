package ghost

import (
	"github.com/fadilix/couik/database"
)

type Replay struct {
	Quote      string
	KeyTimings []database.KeyTiming
	TotalMs    int64
}

func Load(quote string, history database.History) (Replay, bool) {
	var best *database.TestResult
	for i := range history {
		tr := &history[i]
		if tr.Quote != quote {
			continue
		}
		if len(tr.KeyTimings) == 0 {
			continue
		}
		if best == nil || tr.Date.After(best.Date) {
			best = tr
		}
	}
	if best == nil {
		return Replay{}, false
	}
	var total int64
	for _, k := range best.KeyTimings {
		if k.OffsetMs > total {
			total = k.OffsetMs
		}
	}
	return Replay{
		Quote:      best.Quote,
		KeyTimings: best.KeyTimings,
		TotalMs:    total,
	}, true
}

func (r Replay) CursorAt(elapsedMs int64) int {
	if r.Quote == "" {
		return 0
	}
	max := len([]rune(r.Quote))
	cursor := 0
	for _, k := range r.KeyTimings {
		if k.OffsetMs > elapsedMs {
			break
		}
		if k.Backspace {
			if cursor > 0 {
				cursor--
			}
		} else {
			if cursor < max {
				cursor++
			}
		}
	}
	if cursor < 0 {
		cursor = 0
	}
	return cursor
}
