package main

import (
	"os"

	"github.com/dpacheconr/newrelic-python-paketo-buildpack/newrelic"
	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
)

func main() {
	libpak.Main(
		newrelic.Detect{},
		newrelic.Build{Logger: bard.NewLogger(os.Stdout)},
	)
}
