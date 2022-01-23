package internal

import (
	"github.com/mmcdole/gofeed"
	"github.com/myl7/bibak/internal/config"
	"log"
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
	p := gofeed.NewParser()
	feed, err := p.ParseURL(config.Cfg.RsshubUrl)
	if err != nil {
		panic(err)
	}

	for i := range feed.Items {
		handleItem(feed.Items[i])
	}
}
