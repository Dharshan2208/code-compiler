package sandbox

import (
	"bytes"
	"context"
	"os/exec"
	"time"
)

type Result struct {
	Stdout string
	Stderr string
	Error  error
}

type Sandbox struct{}

func (s *Sandbox) Run(image string, workspace string, command []string) Result {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []string{
		"run",
		"--rm",
		"--memory=256m",
		"--cpus=1",
		"--network=none",
		"-v",
		workspace + ":/workspace",
		"-w",
		"/workspace",
		image,
	}

	args = append(args, command...)

	cmd := exec.CommandContext(ctx, "docker", args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return Result{
			Stderr: "executiion timeout",
			Error:  ctx.Err(),
		}
	}

	return Result{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
		Error:  err,
	}
}
