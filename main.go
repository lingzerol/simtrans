package main

import (
	"flag"
)

func main() {
	configFile := flag.String("-config", "~/.simtrans.toml", "Path to configuration file")
	mode := flag.String("-mode", "", "running mode (server|client)")
	copyContent := flag.String("-copy", "", "copy text content")
	putFilePath := flag.String("-put", "", "put the file path to the server")
	pasteContent := flag.String("--paste", "", "paste the content to the local storage")

	if configFile == nil || *configFile == "" {
		panic("config file can not be empty.")
	}
	if mode == nil || *mode == "" {
		panic("running mode can not be empty.")
	}
	if *mode != "server" && *mode != "client" {
		panic("running mode can only be server or client.")
	}

}
