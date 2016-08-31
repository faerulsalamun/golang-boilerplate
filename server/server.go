package server

import (
	"github.com/faerulsalamun/golang-boilerplate/config"
)

func Init() {
	c := config.GetConfig()

	r := NewRouter()
	r.Run(c.GetString("app.port"))
}
