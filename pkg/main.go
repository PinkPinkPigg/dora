package main

import (
	"context"
	"github.com/PinkPinkPigg/dora/pkg/kits"
	"time"
)

func main() {
	//print(kits.GenerateInstanceID(123, 0, 120301203))
	ctx, _ := context.WithCancel(context.Background())
	//go kits.RunShellScript(ctx, "/Users/bytedance/Desktop/projects/dora/pkg/kits/test_shell.sh", "test_shell_log", "/Users/bytedance/Desktop/projects/dora/pkg/kits")
	go kits.RunPythonScript(ctx, "/Users/bytedance/Desktop/projects/dora/pkg/kits/test_py3.py", "test_py3_log", "/Users/bytedance/Desktop/projects/dora/pkg/kits", "python3", "arg1", "arg2", "arg3")
	time.Sleep(2 * time.Second)
	ctx.Done()
}
