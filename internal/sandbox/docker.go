package sandbox

import (
	"bytes"
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
)

type Result struct {
	Stdout string
	Stderr string
	Error  error
}

type Sandbox struct {
	Container *WarmContainer
	Manager   *PoolManager
}

func (s *Sandbox) Execute(ctx context.Context, command []string) Result {
	execConfig := container.ExecOptions{
		User:         "1000", // Execute as the standard user
		Cmd:          command,
		AttachStdout: true,
		AttachStderr: true,
		WorkingDir:   "/workspace",
	}

	exec, err := s.Manager.cli.ContainerExecCreate(ctx, s.Container.ID, execConfig)
	if err != nil {
		return Result{Error: err}
	}

	resp, err := s.Manager.cli.ContainerExecAttach(ctx, exec.ID, container.ExecStartOptions{})
	if err != nil {
		return Result{Error: err}
	}
	defer resp.Close()

	var stdout, stderr bytes.Buffer
	// stdcopy helps split the multiplexed stream from Docker back into stdout and stderr
	_, err = stdcopy.StdCopy(&stdout, &stderr, resp.Reader)

	return Result{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
		Error:  err,
	}
}
