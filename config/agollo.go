package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"

	"github.com/apolloconfig/agollo/v4"
	agocfg "github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
)

var (
	once           sync.Once
	client         agollo.Client
	_innerListener innerListener
)

type (
	AgolloClientImpl struct {
		Client agollo.Client
	}
)

// 重新定义一个监听器，以简化外部监听器的实现
type AgolloChangeListener interface {
	OnChange(event *storage.ChangeEvent)
}

// 内部监听器，用于处理Apollo配置变化的逻辑
type innerListener struct {
	OuterChange map[string]func(event *storage.ChangeEvent)
}

// 初始化Apollo客户端
func NewAgolloClient(listener AgolloChangeListener) (*AgolloClientImpl, error) {
	var err error
	once.Do(func() {
		cfg, e := LoadAppConfig()
		if e != nil {
			err = e
			log.Println("Load AppConfig error: ", e)
			return
		}
		agoConfig := cfg.Apollo
		client, err = agollo.StartWithConfig(func() (*agocfg.AppConfig, error) {
			// 配置Apollo客户端
			return &agocfg.AppConfig{
				AppID:            agoConfig.Appid,     // 应用ID
				Cluster:          agoConfig.Env,       // 集群环境
				NamespaceName:    agoConfig.Namespace, // 命名空间
				IP:               agoConfig.Host,      // 主机
				Secret:           agoConfig.Secret,    // 密钥
				MustStart:        true,                // 是否必须启动
				BackupConfigPath: os.TempDir(),        // 备份配置路径
				IsBackupConfig:   true,                // 是否备份配置
			}, nil
		})
		log.Println("Apollo Client Started")

		if err == nil {
			_innerListener = innerListener{OuterChange: make(map[string]func(event *storage.ChangeEvent))}
			// 为客户端添加一个配置变化监听器
			client.AddChangeListener(&_innerListener)
		}
	})
	if err != nil {
		return nil, err
	}
	if listener != nil {
		name := reflect.TypeOf(listener).String()
		_innerListener.OuterChange[name] = listener.OnChange

	}
	return &AgolloClientImpl{
		Client: client,
	}, nil
}

// 根据Key获取配置对象
// 参数：
//   - key：配置对象的Key
//
// 返回值：
//   - outValue：配置对象的值
//
// 错误：
//   - error：如果发生错误，则返回错误信息，否则返回nil
func (ag *AgolloClientImpl) GetValue(key string, outValue any) error {
	value := ag.Client.GetValue(key)
	if len(value) == 0 {
		return fmt.Errorf("key %s not found", key)
	}
	if e := json.Unmarshal([]byte(value), &outValue); e != nil {
		return fmt.Errorf("unmarshal %s error: %w", key, e)
	}
	return nil
}

// 获取配置
// 参数：
//   - namespace：配置对象的命名空间
//   - key：配置对象的Key
//
// 返回值：
//   - string：配置对象的值
func (ag *AgolloClientImpl) GetConfig(namespace string, key string) string {
	value := ag.Client.GetConfig(namespace)
	return value.GetValue(key)
}

// 获取缓存中的值
// 参数：
//   - namespace：配置对象的命名空间
//   - key：配置对象的Key
//   - outValue：配置对象的值
//
// 返回值：
//   - error：如果发生错误，则返回错误信息，否则返回nil
func (ag *AgolloClientImpl) GetCacheConfig(namespace string, key string, outValue any) error {
	value := ag.Client.GetConfigCache(namespace)
	if value == nil {
		return fmt.Errorf("namespace %s not found", namespace)
	}
	obj, e := value.Get(key)
	if e != nil {
		return fmt.Errorf("key %s not found", key)
	}
	if e := json.Unmarshal([]byte(obj.(string)), &outValue); e != nil {
		return fmt.Errorf("unmarshal %s error: %w", key, e)
	}
	return nil
}

// 获取缓存中的值
// 参数：
//   - key：配置对象的Key
//   - outValue：配置对象的值
//
// 返回值：
//   - error：如果发生错误，则返回错误信息，否则返回nil
func (ag *AgolloClientImpl) GetCacheValue(key string, outValue any) error {
	cache := ag.Client.GetDefaultConfigCache()
	value, err := cache.Get(key)
	if err != nil {
		return fmt.Errorf("key %s not found", key)
	}
	if e := json.Unmarshal([]byte(value.(string)), &outValue); e != nil {
		return fmt.Errorf("unmarshal %s error: %w", key, e)
	}
	return nil
}

// 配置变化监听器
func (l *innerListener) OnChange(e *storage.ChangeEvent) {
	log.Println("Apollo Config Changed: ", e.Namespace)
	for k, action := range l.OuterChange {
		action(e)
		log.Println("action listener: ", k)
	}
}

// 配置变化监听器
func (l *innerListener) OnNewestChange(e *storage.FullChangeEvent) {
	//log.Println("Apollo Newest Config: ", e.Namespace)
	// for k, v := range e.Changes {
	// 	fmt.Printf("Key: %s, Value: %s\n", k, v)
	// }
}
