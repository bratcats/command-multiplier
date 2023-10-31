package multiplier

import (
	"fmt"
	"os/exec"
)

type CmdContext struct {
	CMD            string
	NrOfThreads    int
	NrOfIterations int
}

func Run(ctx CmdContext) {
	fmt.Printf("Will run the '%s' command %d times using %d parallele threads\n", ctx.CMD, ctx.NrOfThreads, ctx.NrOfIterations)

	workload := make(chan string)
	go initWorkload(workload, ctx)

	for work := range workload {
		r, _ := runCmd(work)
		fmt.Printf("\nResult of command '%s' is %q\n", work, r)
	}
}

func initWorkload(workload chan<- string, ctx CmdContext) {
	fmt.Println("Initializing workload")
	for i := 0; i < ctx.NrOfIterations; i++ {
		workload <- ctx.CMD
	}
	close(workload)
}

func runCmd(cmd string) (string, error) {
	c := exec.Command("bash", "-c", cmd)
	output, err := c.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(output), nil
}
