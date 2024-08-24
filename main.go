package main

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"

	"github.com/Forget-C/clash-provider-converter/util"
)

var (
	homedir               string   = os.Getenv("HOME")                       // 用户目录
	subscribedConfigDir   string   = filepath.Join(homedir, ".config/clash") // clash订阅后生成的文件所在目录
	subscribedConfigFiles []string                                           // clash订阅后生成的文件列表， 不填则去目录中读取所有
	convertedConfigDir    string   = subscribedConfigDir + "/proxies"        // 转换后的文件存放目录， 默认与 subscribedConfigDir/proxies/{name}_convert_proxies.yaml
)

func init() {
	flag.StringVar(&subscribedConfigDir, "subscribed-config-dir", subscribedConfigDir, "clash订阅后生成的文件所在目录")
	flag.StringVar(&convertedConfigDir, "converted-config-dir", convertedConfigDir, "转换后的文件存放目录")
	flag.StringSliceVar(&subscribedConfigFiles, "subscribed-config-files", subscribedConfigFiles, "clash订阅后生成的文件列表， 不填则去目录中读取所有")
	flag.Parse()
	convertedConfigDir = filepath.Join(subscribedConfigDir, "proxies")
}

func main() {
	files := subscribedConfigFiles
	dirFiles, err := util.GetYamlFiles(subscribedConfigDir)
	if err != nil {
		fmt.Printf("获取配置文件失败: %v", err)
		return
	}
	files = append(files, dirFiles...)
	if len(files) == 0 {
		fmt.Printf("没有找到配置文件")
		return
	}
	count := 0
	var convertFiles []string
	for _, file := range files {
		if filepath.Base(file) == "config.yaml" {
			continue
		}
		data, err := util.ReadConfig(file)
		if err != nil {
			fmt.Printf("读取配置文件失败: %v", err)
			continue
		}
		proxyName := util.GetFileNameWithoutExt(file)
		proxies := util.ExtractProxies(data)
		if len(proxies) == 0 {
			continue
		}
		proxiesPath := filepath.Join(convertedConfigDir, fmt.Sprintf("%s_convert_proxies.yaml", proxyName))
		err = util.WriteProxiesFile(proxiesPath, proxies)
		if err != nil {
			fmt.Printf("写入文件%s失败: %v", proxiesPath, err)
			continue
		}
		count++
		convertFiles = append(convertFiles, proxiesPath)
	}
	fmt.Printf("成功转换%d个文件\n", count)
	fmt.Println("转换后的文件列表:")
	for _, file := range convertFiles {
		fmt.Println(file)
	}
}
