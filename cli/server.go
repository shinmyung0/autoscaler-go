package cli

import (
	"github.com/shinmyung0/autoscaler/api"
	log "github.com/shinmyung0/loglite"
)

func RunServer(args []string) int {

	// if len(args) > 0 && args[0] == "-v" {
	// 	log.SetDebug(true)
	// 	log.Info("Logging to verbose mode")
	// }

	log.SetDebug(true)
	log.Info("Logging to verbose mode")

	api.ListenAndServe()

	return 0
}
