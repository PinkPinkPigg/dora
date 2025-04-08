package kits

import (
	"encoding/base64"
	"fmt"
)

//task id--->实例生成器

// 全局唯一、可逆验证、不可重复
func GetInstanceID(taskID uint64, frequency int32, timestamp uint64) string {
	// 构造原始字符串
	raw := fmt.Sprintf("%d|%d|%d", taskID, frequency, timestamp)

	//// SHA256哈希
	//hash := sha256.Sum256([]byte(raw))

	// Base64 URL-safe 编码
	encoded := base64.URLEncoding.EncodeToString([]byte(raw))

	// 截取前24字符（32 bytes -> 24 base64 chars）
	return encoded
}
