package services

import(
	"io"
	"bytes"
	"fmt"
	"os/exec"
	"log"
)

func Execute(outputBuffer *bytes.Buffer, stack []*exec.Cmd) (errorOutput string) {
	var errorBuffer bytes.Buffer
	pipeStack := make([]*io.PipeWriter, len(stack)-1)
	i := 0
	for ; i < len(stack)-1; i++ {
		stdinPipe, stdoutPipe := io.Pipe()
		stack[i].Stdout = stdoutPipe
		stack[i].Stderr = &errorBuffer
		stack[i+1].Stdin = stdinPipe
		pipeStack[i] = stdoutPipe
	}
	stack[i].Stdout = outputBuffer
	stack[i].Stderr = &errorBuffer
	var errStr string
	if err := call(stack, pipeStack); err != nil {
		fmt.Println ("Encounter Error", string(errorBuffer.Bytes()), err)
		errStr = err.Error()
	}

	if errStr != "" && errorBuffer.Bytes() != nil {
		return fmt.Sprintf("%v\n%v", errStr, string(errorBuffer.Bytes()))
	}

	return ""
}

func call(stack []*exec.Cmd, pipes []*io.PipeWriter) (err error) {
	if stack[0].Process == nil {
		if err = stack[0].Start(); err != nil {
			return err
		}
	}
	if len(stack) > 1 {
		if err = stack[1].Start(); err != nil {
			return err
		}
		defer func() {
			if err == nil {
				pipes[0].Close()
				err = call(stack[1:], pipes[1:])
			}
		}()
	}
	return stack[0].Wait()
}

func GetHostNameExec() string {
	out, err := exec.Command("hostname").Output()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", out)
}