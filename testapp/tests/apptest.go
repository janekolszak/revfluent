package tests

import (
    log "github.com/janekolszak/revfluent"
    "github.com/revel/revel/testing"
)

type AppTest struct {
    testing.TestSuite
}

func (t *AppTest) TestError() {
    data := map[string]string{"message": "Error"}
    log.Error(data)
}

func (t *AppTest) TestDebug() {
    data := map[string]string{"message": "Debug"}
    log.Debug(data)
}

func (t *AppTest) TestInfo() {
    data := map[string]string{"message": "Info"}
    log.Info(data)
}

func (t *AppTest) TestLog() {
    data := map[string]string{"message": "Log"}
    log.Log("tag", data)
}

func (t *AppTest) TestLogger() {
    data := map[string]string{"message": "Logger"}
    log.Logger.Post("tag", data)
}
