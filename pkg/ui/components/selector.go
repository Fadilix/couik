package components

type Selector interface {
	Increment()
	Decrement()
	Selected() string
}
