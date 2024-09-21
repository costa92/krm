package main

import (
	"os"

	"github.com/costa92/krm/cmd/krm-apiserver/app"
	"k8s.io/component-base/cli"
)

func main() {
	// Please note that the following WithOptions are all required.
	// zh: 请注意以下 WithOptions 都是必须的。
	command := app.NewAPIServerCommand()

	code := cli.Run(command)
	os.Exit(code)
}
