package config

import (
	"os"
	"sync"

	"github.com/apolloconfig/agollo/v4"
	agocfg "github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
)

type ConfigUpdateListener interface {
	OnConfigUpdate(namespace string, key string, newValue interface{})
}

// ConfigManager manages the Agollo client and dispatches configuration updates to listeners.
type ConfigManager struct {
	client    agollo.Client
	listeners map[string]map[string][]ConfigUpdateListener // namespace -> key -> []listeners
	appConfig *AgolloConfig
	sync.RWMutex
}

// NewConfigManager creates a new ConfigManager instance.
func NewConfigManager(apolloConfig *AgolloConfig) (*ConfigManager, error) {
	return &ConfigManager{
		listeners: make(map[string]map[string][]ConfigUpdateListener),
		appConfig: apolloConfig,
	}, nil
}

// Start starts the Agollo client.
func (cm *ConfigManager) Start() error {
	client, err := agollo.StartWithConfig(func() (*agocfg.AppConfig, error) {
		return &agocfg.AppConfig{
			AppID:            cm.appConfig.AppId,
			Cluster:          cm.appConfig.Cluster,
			NamespaceName:    cm.appConfig.Namespace,
			IP:               cm.appConfig.Meta,
			Secret:           cm.appConfig.Secret,
			IsBackupConfig:   cm.appConfig.IsBackup,
			MustStart:        true,
			BackupConfigPath: os.TempDir(),
		}, nil
	})
	if err != nil {
		return err
	}
	cm.client = client
	client.AddChangeListener(cm) // Register ConfigManager as a change listener
	return nil
}

func (cm *ConfigManager) RegisterListener(namespace string, key string, listener ConfigUpdateListener) {
	cm.Lock()
	defer cm.Unlock()
	if _, ok := cm.listeners[namespace]; !ok {
		cm.listeners[namespace] = make(map[string][]ConfigUpdateListener)
	}
	cm.listeners[namespace][key] = append(cm.listeners[namespace][key], listener)
}

// OnChange implements the agollo.ChangeListener interface.
func (cm *ConfigManager) OnChange(changeEvent *storage.ChangeEvent) {
	cm.RLock()
	defer cm.RUnlock()

	namespace := changeEvent.Namespace
	if keyListeners, ok := cm.listeners[namespace]; ok {
		for key, change := range changeEvent.Changes {
			// Notify listeners registered for this specific key
			if listeners, ok := keyListeners[key]; ok {
				for _, listener := range listeners {
					listener.OnConfigUpdate(namespace, key, change.NewValue)
				}
			}

		}
	}
}

// OnNewestChange implements the storage.ChangeListener interface.
func (cm *ConfigManager) OnNewestChange(changes *storage.FullChangeEvent) {
	// We are handling changes in the OnChange method, so we can leave this empty.
	// This method is part of the storage.ChangeListener interface that Agollo expects.
}

func (cm *ConfigManager) GetConfig(namespace string) *storage.Config {
	if cm.client == nil {
		return nil
	}
	return cm.client.GetConfig(namespace)
}

func (cm *ConfigManager) GetValue(key string) string {
	if cm.client == nil {
		return ""
	}
	return cm.client.GetValue(key)
}

func (cm *ConfigManager) GetValueCache(namespace string, key string) interface{} {
	if cm.client == nil {
		return nil
	}
	value, err := cm.client.GetConfigCache(namespace).Get(key)
	if err != nil {
		return nil
	}
	return value
}

var _ storage.ChangeListener = (*ConfigManager)(nil)
