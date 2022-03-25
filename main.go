package main

import "go-language/pipeline"

const (
	protocol = "tcp"
	port     = "localhost:8000"
)

func main() {
	// Fib Example
	//fib.Run(45)

	// Clock Server
	//clock_server.Run(protocol, port, ture)

	//Echo Server
	//echo_server.Run(protocol, port)

	// Unbuffered Channel Example
	//channel.RunUnbufferedChannelExample()

	// Pipeline Example
	pipeline.RunPipelineExample()

	// Unidirectional Channel Example

}
