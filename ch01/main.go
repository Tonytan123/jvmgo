package main

import "fmt"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n\n", cmd.cpOption, cmd.class, cmd.args)
}
