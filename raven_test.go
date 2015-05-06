package noopraven

import (
	"testing"
)

func TestNoopRavenClientImplementsRavenClient(t *testing.T) {
	var _ RavenClient = &NoopRavenClient{}
}
