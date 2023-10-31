package main

import (
	"command-multiplier/multiplier"
	"fmt"
	"os"
	"strconv"
)

const MAX_ARGS = 3

func main() {
	failIfNotEnoughArguments()

	parallelity, err := strconv.Atoi(os.Args[MAX_ARGS-2])
	if err != nil {
		panic(err)
	}

	nrOfIterations, err := strconv.Atoi(os.Args[MAX_ARGS-1])
	if err != nil {
		panic(err)
	}
	cmd := os.Args[MAX_ARGS]

	failIfParallelityGreterThanNrOfIterations(parallelity, nrOfIterations)

	cmdContext := multiplier.CmdContext{
		CMD:            cmd,
		NrOfIterations: nrOfIterations,
		NrOfThreads:    parallelity,
	}
	multiplier.Run(cmdContext)
}

func failIfNotEnoughArguments() {
	if len(os.Args) != MAX_ARGS+1 {
		fmt.Println("Not enough Arguments provided")
		fmt.Printf("usage: %s <parallelity> <nrOfIterations> <cmd>\n", os.Args[0])
		fmt.Printf("example: %s 2 4 echo \"Hello world\". This will have 2 threads "+
			"that will run in parallele the 'echo \"Hello world\"' 4 times.\n",
			os.Args[0])
		os.Exit(1)
	}
}

func failIfParallelityGreterThanNrOfIterations(parallelity, nrOfIterations int) {
	if parallelity > nrOfIterations {
		fmt.Printf("The given <parallelity> is greater than the <nrOfIterantions>, e.g. %d > %d\n", parallelity, nrOfIterations)
		os.Exit(1)
	}
}
