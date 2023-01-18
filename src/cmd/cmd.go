package cmd

type Cmd interface {
	Execute() error
}
