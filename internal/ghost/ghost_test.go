package ghost

import (
	"testing"
	"time"

	"github.com/fadilix/couik/database"
)

func TestCursorAt_BeforeAnyKeystroke(t *testing.T) {
	r := Replay{Quote: "hello", KeyTimings: []database.KeyTiming{
		{OffsetMs: 100},
		{OffsetMs: 200},
		{OffsetMs: 300},
	}}
	if got := r.CursorAt(0); got != 0 {
		t.Errorf("at elapsed=0, want cursor 0, got %d", got)
	}
}

func TestCursorAt_AfterNKeystrokes(t *testing.T) {
	r := Replay{Quote: "hello", KeyTimings: []database.KeyTiming{
		{OffsetMs: 100},
		{OffsetMs: 200},
		{OffsetMs: 300},
	}}
	cases := []struct {
		elapsed int64
		want    int
	}{
		{50, 0},
		{100, 1},
		{250, 2},
		{300, 3},
		{500, 3},
	}
	for _, c := range cases {
		if got := r.CursorAt(c.elapsed); got != c.want {
			t.Errorf("at elapsed=%d, want cursor %d, got %d", c.elapsed, c.want, got)
		}
	}
}

func TestCursorAt_WithBackspace(t *testing.T) {
	r := Replay{Quote: "hello world", KeyTimings: []database.KeyTiming{
		{OffsetMs: 100},
		{OffsetMs: 200},
		{OffsetMs: 300, Backspace: true},
		{OffsetMs: 400},
	}}
	cases := []struct {
		elapsed int64
		want    int
	}{
		{150, 1},
		{250, 2},
		{350, 1},
		{450, 2},
	}
	for _, c := range cases {
		if got := r.CursorAt(c.elapsed); got != c.want {
			t.Errorf("at elapsed=%d, want cursor %d, got %d", c.elapsed, c.want, got)
		}
	}
}

func TestCursorAt_ClampsToQuoteLength(t *testing.T) {
	r := Replay{Quote: "ab", KeyTimings: []database.KeyTiming{
		{OffsetMs: 100},
		{OffsetMs: 200},
		{OffsetMs: 300},
		{OffsetMs: 400},
	}}
	if got := r.CursorAt(999); got != 2 {
		t.Errorf("expected clamp to 2, got %d", got)
	}
}

func TestCursorAt_EmptyQuote(t *testing.T) {
	r := Replay{}
	if got := r.CursorAt(100); got != 0 {
		t.Errorf("empty replay should return 0, got %d", got)
	}
}

func TestLoad_PicksMostRecentMatchingQuote(t *testing.T) {
	now := time.Now()
	history := database.History{
		{Quote: "other quote", Date: now.Add(-3 * time.Hour), KeyTimings: []database.KeyTiming{{OffsetMs: 1}}},
		{Quote: "the one", Date: now.Add(-2 * time.Hour), KeyTimings: []database.KeyTiming{{OffsetMs: 1}}},
		{Quote: "the one", Date: now.Add(-1 * time.Hour), KeyTimings: []database.KeyTiming{{OffsetMs: 2}}},
		{Quote: "the one", Date: now.Add(-30 * time.Minute), KeyTimings: []database.KeyTiming{{OffsetMs: 3}}},
	}
	r, ok := Load("the one", history)
	if !ok {
		t.Fatal("expected to find ghost for \"the one\"")
	}
	if len(r.KeyTimings) != 1 || r.KeyTimings[0].OffsetMs != 3 {
		t.Errorf("expected most-recent matching run (OffsetMs=3), got %+v", r.KeyTimings)
	}
}

func TestLoad_SkipsRunsWithNoTimings(t *testing.T) {
	now := time.Now()
	history := database.History{
		{Quote: "old run, has timings", Date: now.Add(-10 * time.Hour), KeyTimings: []database.KeyTiming{{OffsetMs: 5}}},
		{Quote: "newer run, NO timings", Date: now.Add(-1 * time.Hour)},
	}
	r, ok := Load("newer run, NO timings", history)
	if ok {
		t.Errorf("expected ok=false when only no-timings entries exist, got %+v", r)
	}
}

func TestLoad_NoMatch(t *testing.T) {
	history := database.History{
		{Quote: "something else", KeyTimings: []database.KeyTiming{{OffsetMs: 1}}},
	}
	_, ok := Load("the one", history)
	if ok {
		t.Errorf("expected ok=false when no matching quote, got ok=true")
	}
}
