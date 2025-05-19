package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(path string) (*AppConfig, error) {
	viper.SetConfigName("app")  // 配置文件名 (不带扩展名)
	viper.SetConfigType("yaml") // 配置文件类型
	viper.AddConfigPath(path)   // 添加配置文件路径

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// 绑定环境变量
	viper.BindEnv("server.port", "SERVER_PORT")

	//绑定Apollo配置
	viper.BindEnv("agollo.appId", "APOLLO_APP_ID")
	viper.BindEnv("agollo.cluster", "APOLLO_CLUSTER")
	viper.BindEnv("agollo.namespace", "APOLLO_NAMESPACE")
	viper.BindEnv("agollo.meta", "APOLLO_META")
	viper.BindEnv("agollo.secret", "APOLLO_SECRET")
	viper.BindEnv("agollo.isBackup", "APOLLO_IS_BACKUP")

	config := &AppConfig{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
