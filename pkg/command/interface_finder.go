package command

type Finder interface {
	Find() Command
}
