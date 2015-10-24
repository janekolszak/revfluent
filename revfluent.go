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
    var err error
    var timeout time.Duration = 0
    timeStr, exists := revel.Config.String("revfluent.timeout")
    if exists {
        timeout, err = time.ParseDuration(timeStr)
        if err != nil {
            revel.ERROR.Panic(err)
        }
    }

    var retryWait int = 0
    retryWaitStr, exists := revel.Config.String("revfluent.retryWait")
    if exists {
        tmp, err := time.ParseDuration(retryWaitStr)
        if err != nil {
            revel.ERROR.Panic(err)
        }

        retryWait = int(tmp.Nanoseconds() / 1e6)
    }

    port, _ := revel.Config.Int("revfluent.port")
    host, _ := revel.Config.String("revfluent.host")
    network, _ := revel.Config.String("revfluent.network")
    socketPath, _ := revel.Config.String("revfluent.socketPath")
    bufferLimit, _ := revel.Config.Int("revfluent.bufferLimit")
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

    var hostUsed string
    if hostUsed = host; host == "" {
        hostUsed = "127.0.0.1"
    }

    var portUsed int
    if portUsed = port; port == 0 {
        portUsed = 24224
    }
    revel.INFO.Printf("Connecting Fluentd: %s:%d", hostUsed, portUsed)

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
