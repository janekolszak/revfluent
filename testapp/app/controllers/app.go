package controllers

import (
    log "github.com/janekolszak/revfluent"
    "github.com/revel/revel"
)

type App struct {
    *revel.Controller
}

func (c App) Index() revel.Result {

    data := map[string]string{"src": "App", "message": "Info message from the App"}
    log.Info(data)
    return c.Render()
}
