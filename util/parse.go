package util

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/Forget-C/clash-provider-converter/model/clash"
)

func ReadFile(f string) ([]byte, error) {
	return os.ReadFile(f)
}

func ParseClashConfig(b []byte) (clash.Clash, error) {
	c := clash.Clash{}
	err := yaml.Unmarshal(b, &c)
	return c, err
}

// ReadConfig 配置文件
func ReadConfig(f string) (*clash.Clash, error) {
	b, err := ReadFile(f)
	if err != nil {
		return nil, err
	}
	c, err := ParseClashConfig(b)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// GetYamlFiles 遍历目录，获取yaml配置文件
func GetYamlFiles(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var yamlFiles []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if len(file.Name()) > 4 && file.Name()[len(file.Name())-4:] == "yaml" {
			yamlFiles = append(yamlFiles, dir+"/"+file.Name())
		}
	}
	return yamlFiles, nil
}

// ExtractProxies 提取配置文件中的节点
func ExtractProxies(c ...*clash.Clash) []*clash.Proxies {
	var proxies []*clash.Proxies
	for _, v := range c {
		proxies = append(proxies, v.Proxies...)
	}
	return proxies
}

// GetFileNameWithoutExt 获取文件面，不包含文件扩展后缀
func GetFileNameWithoutExt(f string) string {
	return strings.TrimRight(filepath.Base(f), filepath.Ext(f))
}
