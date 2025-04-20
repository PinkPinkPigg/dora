package kits

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

// 负责对consul的访问
func DiscoverMService(consulAddr, serviceName string) (string, string, error) {
	// 创建 Consul 客户端
	config := api.DefaultConfig()
	config.Address = consulAddr
	client, err := api.NewClient(config)
	if err != nil {
		return "", "", fmt.Errorf("failed to create Consul client: %v", err)
	}

	// 查询服务信息
	services, _, err := client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return "", "", fmt.Errorf("failed to discover service: %v", err)
	}

	// 检查是否有可用实例
	if len(services) == 0 {
		return "", "", fmt.Errorf("no healthy instances found for service: %s", serviceName)
	}

	// 选择第一个健康实例（可以扩展为负载均衡策略）
	service := services[0]
	host := service.Service.Address
	port := service.Service.Port

	return host, fmt.Sprintf("%d", port), nil
}
