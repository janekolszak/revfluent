package revfluent

import (
    "fmt"
    "github.com/fluent/fluent-logger-golang/fluent"
    "github.com/revel/revel"
)

var (
    Logger *fluent.Fluent // Global logger
)

// Reads configuration and connects to fluentd
func Init() {
    timeout, err := time.ParseDuration(revel.Config.StringDefault("revfluent.timeout", fluent.defaultTimeout.String()))
    if err != nil {
        revel.ERROR.Panic(err)
    }

    port := revel.Config.IntDefault("revfluent.port", fluent.defaultPort)
    host := revel.Config.StringDefault("revfluent.host", fluent.defaultHost)
    network := revel.Config.StringDefault("revfluent.network", fluent.defaultNetwork)
    socketPath := revel.Config.StringDefault("revfluent.socketPath", fluent.defaultSocketPath)
    bufferLimit := revel.Config.IntDefault("revfluent.bufferLimit", fluent.defaultBufferLimit)
    retryWait := revel.Config.IntDefault("revfluent.retryWait", fluent.defaultRetryWait)
    maxRetry := revel.Config.IntDefault("revfluent.maxRetry", fluent.defaultMaxRetry)
    tagPrefix := revel.Config.StringDefault("revfluent.tagPrefix", revel.Config.String("app.name"))

    config := fluent.Config{
        FluentPort:       port,
        FluentHost:       host,
        FluentNetwork:    network,
        FluentSocketPath: socketPath,
        Timeout:          timeout,
        BufferLimit:      bufferLimit,
        RetryWait:        retryWait,
        MaxRetry:         maxRetry,
        TagPrefix:        tagPrefix,
    }

    revel.INFO.Printf("Connecting Fluentd: %s:%i", host, port)

    Logger, err := fluent.New(config)
    if err != nil {
        revel.ERROR.Panic("Failed to connect Fluentd: %s", err)
    }
}
