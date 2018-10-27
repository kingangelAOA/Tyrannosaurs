package component

type Observer interface {
	Receive(interface{}) error
}

type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify() error
}