package components

type Selector interface {
	Increment()
	Decrement()
	Selected() string
	GetChoices() []string
	GetCursor() int
}
