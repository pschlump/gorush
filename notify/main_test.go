package notify

import (
	"log"
	"os"
	"testing"

	"github.com/pschlump/gorush/config"
	"github.com/pschlump/gorush/status"
)

func TestMain(m *testing.M) {
	cfg, _ := config.LoadConf()
	if err := status.InitAppStatus(cfg); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
