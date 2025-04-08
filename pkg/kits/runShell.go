package kits

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// RunShellScript 调用指定路径的 Shell 脚本
func RunShellScript(ctx context.Context, scriptPath string, logFileName string, logDir string) error {
	// 检查脚本路径是否存在
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return fmt.Errorf("script not found at path: %s", scriptPath)
	}

	// 确保日志目录存在
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	// 构造日志文件路径,目录+名称
	logFilePath := filepath.Join(logDir, logFileName)

	// 打开日志文件（覆盖写入模式）
	logFile, err := os.OpenFile(logFilePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer logFile.Close()

	// 创建命令
	cmd := exec.CommandContext(ctx, "/bin/bash", scriptPath)

	// 将标准输出和标准错误重定向到日志文件
	cmd.Stdout = logFile
	cmd.Stderr = logFile

	// 执行命令
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start script: %v", err)
	}

	// 等待命令完成或上下文超时/取消
	err = cmd.Wait()
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return fmt.Errorf("script execution timed out")
	} else if errors.Is(ctx.Err(), context.Canceled) {
		return fmt.Errorf("script execution was canceled")
	}

	return err
}
