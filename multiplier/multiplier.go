package multiplier

import (
	"fmt"
	"os/exec"
	"sync"
)

type CmdContext struct {
	CMD            string
	NrOfThreads    int
	NrOfIterations int
}

func Run(ctx CmdContext) {
	fmt.Printf("Will run the '%s' command %d times using %d parallel threads\n", ctx.CMD, ctx.NrOfIterations, ctx.NrOfThreads)

	workload := make(chan string, ctx.NrOfIterations)

	var wg sync.WaitGroup
	for i := 0; i < ctx.NrOfThreads; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, workload)
		}(i)
	}

	initWorkload(workload, ctx)
	wg.Wait()

	fmt.Println("Program done")
}

func initWorkload(workload chan<- string, ctx CmdContext) {
	fmt.Println("Initializing workload")
	for i := 0; i < ctx.NrOfIterations; i++ {
		workload <- ctx.CMD
	}
	close(workload)
}

func worker(id int, workload <-chan string) {
	for cmd := range workload {
		r, _ := runCmd(cmd)
		fmt.Printf("\n[Worker-%d] Result of command '%s' is %q\n", id, cmd, r)
	}
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
