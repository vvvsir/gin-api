package config

var App struct {
	AppName        string
	AppPort        int
	Pprof          bool
	ContextTimeout int
	ReadTimeout    int
	WriteTimeout   int
}
