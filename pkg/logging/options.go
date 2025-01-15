package logging

type Option func(*Options) *Options

type Options struct {
	Format            string
	Level             string
	Name              string
	DisableCaller     bool
	DisableStacktrace bool
	EnableColor       bool
	Outputs           []string
	ErrorOutputs      []string
}
