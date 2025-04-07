package config

import (
	"strings"

	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		Server struct {
			Port int `json:"http_port" default:"50001"`
		} `json:"server"`

		Apollo struct {
			Appid     string `json:"appid"`
			Namespace string `json:"namespace"`
			Env       string `json:"env"`
			Cluster   string `json:"cluster"`
			Host      string `json:"host"`
			Secret    string `json:"secret"`
		} `json:"apollo"`
	}

	MySQLConfig struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	}
)

// LoadAppConfig 从 app.yaml 文件中加载应用程序配置
//
// 返回值:
//   - *AppConfig: 包含应用程序配置的 AppConfig 结构体指针
//   - error: 如果加载或解析配置文件过程中出现错误，则返回相应的错误信息
//
// 该方法执行以下步骤:
// 1. 读取 app.yaml 文件的内容。
// 2. 使用 yaml.Unmarshal 将文件内容解析为 AppConfig 结构体。
// 3. 返回解析后的 AppConfig 结构体指针和可能的错误信息。
func LoadAppConfig() (*AppConfig, error) {
	// 1. 创建独立的 Viper 实例，避免全局状态
	v := viper.New()
	// 设置配置文件名和路径
	v.SetConfigName("app")  // 文件名（不含扩展名）
	v.SetConfigType("yaml") // 文件类型
	v.AddConfigPath(".")    // 文件所在路径（当前目录）
	v.AutomaticEnv()
	// v.WatchConfig()

	// v.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("配置文件已更改:", e.Name)
	// })

	// 设置环境变量键的替换规则，例如：DATABASE_USER 映射到 database.user
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var config AppConfig
	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = v.Unmarshal(&config)
	return &config, err
}
