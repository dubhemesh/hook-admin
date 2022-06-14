package main

import (
	_ "hook-admins/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"hook-admins/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
