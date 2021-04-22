package rabbitsimple

import (
	"flag"
	"fmt"

	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	ConfigEnv  = "GVA_CONFIG"
	ConfigFile = "config.yaml"
)

var (
	Sys System
)

type RabbitMq struct {
	Admin   string `mapstructure:"admin" json:"admin" yaml:"admin"`
	Pwd     string `mapstructure:"pwd" json:"pwd" yaml:"pwd"`
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`
	Ip      string `mapstructure:"ip" json:"ip" yaml:"ip"`
	Verhost string `mapstructure:"verhost" json:"verhost" yaml:"verhost"`
}

type System struct {
	RabbitMq RabbitMq `mapstructure:"rabbitMq" json:"rabbitMq" yaml:"rabbitMq"`
}

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
				config = ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&Sys); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&Sys); err != nil {
		fmt.Println(err)
	}
	return v
}
