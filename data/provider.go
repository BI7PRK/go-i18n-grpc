package data

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"i18n-service/config"

	"github.com/apolloconfig/agollo/v4/storage"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	engineInstance *xorm.Engine
	once           sync.Once
	mutex          sync.Mutex
	dbConfigKey    = "I18ndb"
)

type DataConfigListener struct {
	NewValue chan *config.MySQLConfig
}

type DataProviderImpl struct {
	ConfigListener *DataConfigListener
}

// OnChange handles configuration changes.
func (l *DataConfigListener) OnChange(e *storage.ChangeEvent) {
	value, ok := e.Changes[dbConfigKey]
	if ok {
		str := value.NewValue.(string)
		newCfg := &config.MySQLConfig{}
		err := json.Unmarshal([]byte(str), newCfg)
		if err != nil {
			log.Printf("failed to unmarshal config: %v", err)
			return // 处理错误后返回
		}
		select {
		case l.NewValue <- newCfg:
		case <-time.After(time.Second * 10):
			log.Printf("发送配置更新到通道 l.NewValue 超时 for key '%s'", dbConfigKey)
		default:
			log.Printf("通道 l.NewValue 阻塞，丢弃配置更新 for key '%s'", dbConfigKey)
		}
	}
}

// NewDataProvider creates a new DataProviderImpl instance.
func NewDataProvider() *DataProviderImpl {
	lst := &DataConfigListener{NewValue: make(chan *config.MySQLConfig)}
	return &DataProviderImpl{ConfigListener: lst}
}

// GetEngine retrieves the database engine.
func (r *DataProviderImpl) GetEngine() (*xorm.Engine, error) {
	var err error
	once.Do(func() {
		engineInstance, err = createEngine(r.ConfigListener)
		if err == nil {
			log.Println("database engine created")
		}
	})
	if err != nil {
		return nil, err
	}
	return engineInstance, nil
}

// ClearEngine resets the database engine.
func (r *DataProviderImpl) ClearEngine() {
	mutex.Lock()
	defer mutex.Unlock()
	engineInstance = nil
	once = sync.Once{} // 重置once以允许重新创建engine
}

// createEngine creates a new database engine.
func createEngine(lst *DataConfigListener) (*xorm.Engine, error) {
	ag, e := config.NewAgolloClient(lst)
	if e != nil {
		return nil, fmt.Errorf("fail to get apollo config: %w", e)
	}
	u := config.MySQLConfig{}
	err := ag.GetCacheValue(dbConfigKey, &u)
	if err != nil {
		return nil, fmt.Errorf("database config not found")
	}
	engine, ex := xorm.NewEngine("mysql", buildDSN(u))
	if ex != nil {
		return nil, fmt.Errorf("fail to create engine: %w", ex)
	}
	//engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	engine.ShowSQL(true)
	return engine, nil
}

// buildDSN constructs the Data Source Name for the database connection.
func buildDSN(u config.MySQLConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", u.User, u.Password, u.Host, u.Database)
}
