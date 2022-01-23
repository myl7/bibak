package internal

import (
	"github.com/myl7/bibak/internal/config"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

func Loop() {
	for true {
		func() {
			defer func() {
				err := recover()
				if err != nil {
					log.Println("error:", err.(error), "stacktrace:\n", string(debug.Stack()))
				}
			}()

			round()
		}()

		time.Sleep(config.Cfg.LoopInterval)
	}
}

func round() {
	res, err := http.Get(config.Cfg.RsshubUrl)
	if err != nil {
		panic(err)
	}
}
