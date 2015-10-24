# revfluentd [![Build Status](https://travis-ci.org/janekolszak/revfluent.svg?branch=master)](https://travis-ci.org/janekolszak/revfluent)
Fluentd logging for Revel framework

### Installation
``` bash
    go get github.com/janekolszak/revfluent
```
### Test
``` bash
    fluentd -c fluent.conf &;
    revel test github.com/janekolszak/revfluent/testapp dev;
```
### Configuration
In app.conf:
```
    revfluent.host = 127.0.0.1
    revfluent.port = 24224
    revfluent.network = tcp
    revfluent.socketPath =
    revfluent.bufferLimit = 8388608
    revfluent.retryWait = 500ms
    revfluent.maxRetry = 13
    revfluent.tagPrefix = RevFluentTest
    revfluent.timeout = 60s
```
Fields retryWait and timeout accept duration ("ms", "s", "m", "h").

### Initialization
In app.init() in `app/init.go` add:
``` go
    revel.OnAppStart(revfluent.Init)
```