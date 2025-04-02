package data

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

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

func (l *DataConfigListener) OnChange(e *storage.ChangeEvent) {
	value, ok := e.Changes[dbConfigKey]
	if ok {
		str := value.NewValue.(string)
		newCfg := &config.MySQLConfig{}
		err := json.Unmarshal([]byte(str), newCfg)
		if err != nil {
			log.Printf("failed to unmarshal config:%v", err)
		}
		l.NewValue <- newCfg
	}
}

func NewDataProvider() *DataProviderImpl {
	lst := &DataConfigListener{NewValue: make(chan *config.MySQLConfig)}
	return &DataProviderImpl{ConfigListener: lst}
}
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

func (r *DataProviderImpl) ClearEngine() {
	mutex.Lock()
	defer mutex.Unlock()
	engineInstance = nil
	once = sync.Once{} // 重置once以允许重新创建engine
}

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
	log.Printf("MySQL Settings: %+v\n", u)
	engine, ex := xorm.NewEngine("mysql", u.User+":"+u.Password+"@tcp("+u.Host+")/"+u.Database+"?charset=utf8mb4&parseTime=True&loc=Local")
	if ex != nil {
		return nil, fmt.Errorf("fail to create engine: %w", ex)
	}
	//engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	engine.ShowSQL(true)
	return engine, nil
}
