# noop-raven

This is a interface and noop implementation for [Raven](https://github.com/getsentry/raven-go).

This is useful if you want to drop a raven client into some code without having to always check if it's enabled or not when using it.

## Installation

    go get github.com/bigcommerce/noopraven-go

## Usage

Import:

    import "github.com/bigcommerce/noopraven-go"
    import "github.com/getsentry/raven-go"

    ...

    var ravenClient noopraven.RavenClient
    if ravenEnabled {
        var err error
        ravenClient, err = raven.NewClient(dsn, nil)
        // ...
    } else {
        ravenClient = &noopraven.NoopRavenClient{}
    }

    // use ravenClient whether it's enabled or not

    // use the noopraven.RavenClient interface type for arguments in your
    // functions which expect a raven client
