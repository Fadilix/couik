package modes

type StaticMode struct {
	Target string
}

func (s StaticMode) GetTarget() string {
	return s.Target
}

func (s StaticMode) GetInitialTime() int {
	return 0
}
