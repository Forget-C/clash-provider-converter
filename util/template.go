package util

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/Forget-C/clash-provider-converter/model/clash"
)

func WriteConfig(f string, c *clash.Clash) error {
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return WriteFile(f, b)
}

func WriteFile(f string, b []byte) error {
	dir := filepath.Dir(f)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return os.WriteFile(f, b, 0644)
}

// writeProxiesFile 将节点配置写入文件
func WriteProxiesFile(f string, proxies []*clash.Proxies) error {
	data := map[string][]*clash.Proxies{
		"proxies": proxies,
	}
	b, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return WriteFile(f, b)
}

// FindGroupByName 在给定配置中寻找是否有指定的组名
func FindGroupByName(groupName string, template *clash.Clash) (*clash.ProxyGroup, bool) {
	for _, group := range template.ProxyGroup {
		if group.Name == groupName {
			return group, true
		}
	}
	return nil, false
}

// FindProviderByName 在给定配置中寻找是否有指定的提供者名
func FindProviderByName(providerName string, template *clash.Clash) (*clash.ProxyProvider, bool) {
	g, ok := template.ProxyProviders[providerName]
	return g, ok
}
