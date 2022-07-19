package docker

import (
	"bytes"
	"os/exec"
	"testing"
)

// StartContainer runs a db container to execute commands.
func StartPgContainer(t *testing.T) {
	t.Helper()

	StopPgContainer(t)

	cmd := exec.Command("docker", "run", "-d", "--name", "postgres_test", "--publish", "54320:5432", "--env", "POSTGRES_PASSWORD=1234", "bitnami/postgresql:11")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		t.Fatalf("could not start docker : %v", err)
	}

}

// StopContainer stops and removes the specified container.
func StopPgContainer(t *testing.T) {
	t.Helper()

	if err := exec.Command("docker", "container", "rm", "-f", "-v", "postgres_test").Run(); err != nil {
		t.Fatalf("could not stop db container: %v", err)
	}
}
