package noopraven

import (
	"testing"

	"github.com/getsentry/raven-go"
)

func TestNoopRavenClientImplementsRavenClient(t *testing.T) {
	var _ RavenClient = &NoopRavenClient{}
}

func TestRavenImplementsRavenClient(t *testing.T) {
	var _ RavenClient = &raven.Client{}
}
