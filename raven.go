package noopraven

import (
	"github.com/getsentry/raven-go"
)

type RavenClient interface {
	Capture(packet *raven.Packet, captureTags map[string]string) (eventID string, ch chan error)
	CaptureError(err error, tags map[string]string, interfaces ...raven.Interface) string
	CaptureMessage(message string, tags map[string]string, interfaces ...raven.Interface) string
	CapturePanic(f func(), tags map[string]string, interfaces ...raven.Interface)
	Close()
	ProjectID() string
	Release() string
	SetDSN(dsn string) error
	SetRelease(release string)
	URL() string
}

type NoopRavenClient struct {
}

func (client *NoopRavenClient) Capture(packet *raven.Packet, captureTags map[string]string) (eventID string, ch chan error) {
	return "", nil
}

func (client *NoopRavenClient) CaptureError(err error, tags map[string]string, interfaces ...raven.Interface) string {
	return ""
}

func (client *NoopRavenClient) CaptureMessage(message string, tags map[string]string, interfaces ...raven.Interface) string {
	return ""
}

func (client *NoopRavenClient) CapturePanic(f func(), tags map[string]string, interfaces ...raven.Interface) {
	f()
}

func (client *NoopRavenClient) Close() {
}

func (client *NoopRavenClient) ProjectID() string {
	return ""
}

func (client *NoopRavenClient) Release() string {
	return ""
}

func (client *NoopRavenClient) SetDSN(dsn string) error {
	return nil
}

func (client *NoopRavenClient) SetRelease(release string) {
}

func (client *NoopRavenClient) URL() string {
	return ""
}
