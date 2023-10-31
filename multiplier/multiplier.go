package multiplier

import (
	"fmt"
)

type CmdContext struct {
	CMD            string
	NrOfThreads    int
	NrOfIterations int
}

func Run(ctx CmdContext) {
	fmt.Printf("Will run the:\n\n\t%s\n\ncommand %d times using %d parallele threads\n", ctx.CMD, ctx.NrOfThreads, ctx.NrOfIterations)
}
