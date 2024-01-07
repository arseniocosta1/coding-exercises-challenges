package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	Greet(name string) string
	LanguageName() string
}

func SayHello(name string, greeter Greeter) string {
	return greeter.Greet(name)
}

type Italian struct{}

func (i Italian) Greet(name string) string {
	return "I can speak " + i.LanguageName() + ": Ciao " + name + "!"
}

func (i Italian) LanguageName() string {
	return "Italian"
}

type Portuguese struct{}

func (p Portuguese) Greet(name string) string {
	return "I can speak " + p.LanguageName() + ": Olá " + name + "!"
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
