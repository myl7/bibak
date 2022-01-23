package internal

import (
	"github.com/myl7/bibak/internal/config"
	"time"
)

func Loop() error {
	for true {
		err := round()
		if err != nil {
			return err
		}

		time.Sleep(config.Cfg.LoopInterval)
	}
	return nil
}

func round() error {
	return nil
}
