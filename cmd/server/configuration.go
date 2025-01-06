package main

type mode int

const (
	Join = iota
	BootStrap
)

// KVConfig contains all the configuration for the application.
type KVConfig struct {
	NodeId      string
	Mode        mode
	LogLevel    string
	LogDir      string
	ClusterId   string
	ClusterAddr string
	HttpAddr    string
	GrpcAddr    string
}
