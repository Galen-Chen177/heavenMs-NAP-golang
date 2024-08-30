package db

import (
	"testing"

	"heavenMs-NAP-golang/config"
)

func setup() {
	config.ReadConfig("../../local.yaml")
}

func TestConn(t *testing.T) {
	setup()
	NewGormConn()
}
