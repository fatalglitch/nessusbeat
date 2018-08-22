package cmd

import (
    cmd "github.com/elastic/beats/libbeat/cmd"

	"github.com/fatalglitch/nessusbeat/beater"
)

// Name of this beat
var Name = "nessusbeat"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmd(Name, "", beater.New)
