package main

import (
	"runtime/debug"
	"strings"

	"github.com/nexi-intra/nexi-booking/magicapp"
	"github.com/nexi-intra/nexi-booking/utils"
)

func main() {
	info, _ := debug.ReadBuildInfo()

	// split info.Main.Path by / and get the last element
	s1 := strings.Split(info.Main.Path, "/")
	name := s1[len(s1)-1]

	utils.Setup(".env")

	magicapp.RegisterServiceCmd()

	utils.RootCmd.PersistentFlags().BoolVarP(&utils.Verbose, "verbose", "v", false, "verbose output")

	magicapp.Execute(name, "magic-people", "")
}
