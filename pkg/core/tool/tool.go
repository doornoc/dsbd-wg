package tool

import (
	"github.com/doornoc/dsbd-wg/pkg/core/config"
	"log"
)

func OutputLog(value ...any) {
	if config.Debug {
		log.Println(value)
	}
}
