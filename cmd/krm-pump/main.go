package main

import (
	"github.com/costa92/krm/cmd/krm-pump/app"
	_ "k8s.io/component-base/metrics/prometheus/clientgo" // load all the prometheus client-go plugins
	_ "k8s.io/component-base/metrics/prometheus/version"  // for version metric registration
)

func main() {
	app.NewApp().Run()
}
