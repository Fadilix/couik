package engine

import "time"

type Session struct {
	Target         []rune
	Results        []bool
	Index          int
	StartTime      time.Time
	EndTime        time.Time
	BackSpaceCount int
	IsError        bool
	Started        bool
	WpmSamples     []float64
	TimesSample    []time.Time
}

func NewSession(target string) *Session {
	targetRune := []rune(target)
	return &Session{
		Target:  targetRune,
		Results: make([]bool, len(targetRune)),
	}
}

func (s *Session) Type(char string) {
	if !s.Started {
		s.StartTime = time.Now()
		s.Started = true
	}

	if s.Index < len(s.Target) {
		isCorrect := char == string(s.Target[s.Index])
		s.IsError = !isCorrect
		s.Results[s.Index] = isCorrect
		s.Index++
	}
}

func (s *Session) BackSpace() {
	if s.Index > 0 {
		s.Index--
		if s.IsError {
			s.BackSpaceCount++
		}
	}
}

func (s *Session) CalculateTypingSpeed() float64 {
	duration := s.EndTime.Sub(s.StartTime)

	correctChars := 0
	for _, r := range s.Results[:s.Index] {
		if r {
			correctChars++
		}
	}

	return (float64(correctChars) / 5.0) / duration.Minutes()
}

func (s *Session) CalculateLiveTypingSpeed() float64 {
	duration := time.Since(s.StartTime)
	if duration.Minutes() == 0 {
		return 0
	}

	correctChars := 0
	for _, r := range s.Results[:s.Index] {
		if r {
			correctChars++
		}
	}

	return (float64(correctChars) / 5.0) / duration.Minutes()
}

func (s *Session) CalculateRawTypingSpeed() float64 {
	duration := s.EndTime.Sub(s.StartTime)

	correctChars := 0
	for _, r := range s.Results[:s.Index] {
		if r {
			correctChars++
		}
	}

	return (float64(s.Index) / 5.0) / duration.Minutes()
}

func (s *Session) CalculateAccuracy() float64 {
	if s.Index == 0 {
		return 0
	}

	correctChars := 0
	for _, r := range s.Results[:s.Index] {
		if r {
			correctChars++
		}
	}

	net := float64(correctChars - s.BackSpaceCount)
	if net < 0 {
		net = 0
	}
	return (net / float64(s.Index)) * 100
}

func (s *Session) IsFinished() bool {
	s.EndTime = time.Now()
	return s.Index >= len(s.Target)
}

func (s *Session) AddWpmSample(sample float64) {
	s.WpmSamples = append(s.WpmSamples, sample)
}

func (s *Session) AddTimesSample(timeSample time.Time) {
	s.TimesSample = append(s.TimesSample, timeSample)
}
