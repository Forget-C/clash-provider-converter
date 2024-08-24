package util

import (
	"strings"
	"testing"

	"github.com/Forget-C/clash-provider-converter/model/clash"
)

func TestParseClashConfig(t *testing.T) {
	b, err := ReadFile("template.yaml")
	if err != nil {
		t.Error(err)
	}
	c, err := ParseClashConfig(b)
	if err != nil {
		t.Error(err)
	}
	t.Log(c)
}

func TestGetYamlFiles(t *testing.T) {
	files, err := GetYamlFiles("tests/configs")
	if err != nil {
		t.Error(err)
	}
	t.Log(files)
}

func TestExtractProxies(t *testing.T) {
	files, err := GetYamlFiles("tests/configs")
	if err != nil {
		t.Error(err)
	}
	var configs []*clash.Clash
	for _, file := range files {
		if strings.Contains(file, "template.yaml") {
			continue
		}
		b, err := ReadFile(file)
		if err != nil {
			t.Error(err)
		}
		c, err := ParseClashConfig(b)
		if err != nil {
			t.Error(err)
		}
		configs = append(configs, &c)
	}
	proxies := ExtractProxies(configs...)
	t.Log(proxies)
}
