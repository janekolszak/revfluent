package tests

import (
    "github.com/janekolszak/revfluent"
    "github.com/revel/revel/testing"
)

type AppTest struct {
    testing.TestSuite
}

func (t *AppTest) TestLOGE() {
    data := map[string]string{"message": "Error"}
    revfluent.LOGE(data)
}

func (t *AppTest) TestLOGD() {
    data := map[string]string{"message": "Debug"}
    revfluent.LOGD(data)
}

func (t *AppTest) TestLOGI() {
    data := map[string]string{"message": "Info"}
    revfluent.LOGI(data)
}
