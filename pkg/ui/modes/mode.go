package modes

type ModeStrategy interface {
	GetTarget() string
	GetInitialTime() int
}
