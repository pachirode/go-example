package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func readSetDefault() {
	viper.SetDefault("test", map[string]string{"test": "test"})

	fmt.Printf("test: %v", viper.Get("test"))
}

func readConfigFile() {

	var (
		cfg = flag.String("c", "", "config file")
	)

	flag.Parse()

	if *cfg != "" {
		viper.SetConfigFile(*cfg)
		viper.SetConfigType("yaml")
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("cfg")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			println("config not found")
		} else {
			println("other error")
		}
	}

	fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())

	fmt.Printf("test: %s\n", viper.Get("test"))
}

func readWatchConfigFile() {
	viper.SetConfigFile("./cfg.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件出错:", err)
		return
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("配置文件发生变化: %s\n", e.Name)
		// 可以输出更新后的配置值
		fmt.Printf("更新后的 test 值: %s\n", viper.GetString("test"))
	})

	viper.WatchConfig()

	select {}
}

func readIOReader() {
	viper.SetConfigType("yaml")

	var yamExample = []byte(`
test: test
`)
	viper.ReadConfig(bytes.NewBuffer(yamExample))

	fmt.Printf("test: %v", viper.Get("test"))
}

func readEnv() {
	viper.SetEnvPrefix("Pre")
	// 空环境变量视为已设置
	viper.AllowEmptyEnv(true)

	// 绑定全部环境变量
	viper.AutomaticEnv()
	// 绑定单个
	viper.BindEnv("tmp")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// 只能获取当个绑定的环境变量，自动获取的无法使用
	fmt.Printf("%v", viper.AllSettings())
}

func readPFlag() {
	pflag.StringP("username", "u", "1", "help message for username")
	pflag.Parse()
	viper.BindPFlag("username", pflag.Lookup("username"))
	viper.BindPFlags(pflag.CommandLine)

	fmt.Printf("username: %s\n", viper.Get("username"))
}

func readRemote() {
	viper.AddRemoteProvider("consul", "localhost:8500", "test")
	viper.SetConfigType("YAML")
	viper.ReadRemoteConfig()

	fmt.Printf("test: %v", viper.Get("test"))
}

func main() {
	//readSetDefault()
	//readConfigFile()
	//readWatchConfigFile()
	//readIOReader()
	//readEnv()
	//readPFlag()
	readRemote()
}
