package revfluent

import (
    "github.com/fluent/fluent-logger-golang/fluent"
    "github.com/revel/revel"
    "time"
)

var (
    Logger *fluent.Fluent // Global logger
)

// Reads configuration and connects to fluentd
func Init() {
    timeString, _ := revel.Config.String("revfluent.timeout")
    timeout, err := time.ParseDuration(timeString)
    if err != nil {
        revel.ERROR.Panic(err)
    }

    port, _ := revel.Config.Int("revfluent.port")
    host, _ := revel.Config.String("revfluent.host")
    network, _ := revel.Config.String("revfluent.network")
    socketPath, _ := revel.Config.String("revfluent.socketPath")
    bufferLimit, _ := revel.Config.Int("revfluent.bufferLimit")
    retryWait, _ := revel.Config.Int("revfluent.retryWait")
    maxRetry, _ := revel.Config.Int("revfluent.maxRetry")

    appName, _ := revel.Config.String("app.name")
    tagPrefix := revel.Config.StringDefault("revfluent.tagPrefix", appName)

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

    Logger, err = fluent.New(config)
    if err != nil {
        revel.ERROR.Panic("Failed to connect Fluentd: %s", err)
    }
}

func LOGE(message interface{}) {
    Logger.Post("E", message)
}

func LOGD(message interface{}) {
    Logger.Post("D", message)
}

func LOGI(message interface{}) {
    Logger.Post("I", message)
}
