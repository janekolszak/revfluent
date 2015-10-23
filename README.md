# revfluentd
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
- **revfluent.host** =
- **revfluent.port** =
- **revfluent.network** =
- **revfluent.socketPath** =
- **revfluent.bufferLimit** =
- **revfluent.retryWait** =
- **revfluent.maxRetry** =
- **revfluent.tagPrefix** =
- **revfluent.timeout** =

### Initialization
- In app.init() in `app/init.go` add:
``` go
    revel.OnAppStart(revfluent.Init)
```