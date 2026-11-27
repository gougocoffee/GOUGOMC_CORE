// GOUGOMC_CORE 程序入口。
package main

import (
	"fmt"
	"runtime"
)

const (
	// Version 是当前服务端版本。
	Version = "0.1.0"
	// TargetGame 是目标基岩版版本。
	TargetGame = "1.26.60"
)

func main() {
	fmt.Printf("GOUGOMC_CORE %s (bedrock %s, %s %s/%s)\n",
		Version, TargetGame, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
