// repository/cultures.go
package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"i18n-service/config"
	"i18n-service/data/entity"
	"log"
	"reflect"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type CulturesRepositoryImpl struct {
	sync.RWMutex
	db *xorm.Engine
}

type CulturesRepository interface {
	GetCultures() ([]entity.CulturesResources, error)
	GetResourcesByCode(code string) ([]entity.CulturesResourceLangs, error)
	AddOrUpdateCultures(culture entity.CulturesResources) error
	AddOrUpdateCulturesResourceType(data entity.CulturesResourceTypes) error
	DeleteCulturesResourceType(id int64) error
	AddOrUpdateCulturesResourceKey(data entity.CulturesResourceKeys) (*entity.CulturesResourceKeys, error)
	AddOrUpdateCulturesResourceLang(data entity.CulturesResourceLangs) error
	GetCulturesResourceLangPager(i int, param2 int, cultureId int, findKey string) ([]entity.CulturesResourceLangs, int64, error)
	GetCulturesResourceTypeByIds(ids []int32) ([]entity.CulturesResourceTypes, error)
	GetCulturesResourceKeyPager(i int, param2 int, text string) ([]entity.CulturesResourceKeys, int64, error)
	GetCulturesResourceKeyByIds(ids []int32) (map[int32]string, error)
	AddCulturesResourceLangs(key string, tid int32, cultureLang []entity.CulturesResourceLangs) error
	DeleteCulturesResourceKey(id int32) error
	GetCulturesResourceTypePager(i int, param2 int, text string) ([]entity.CulturesResourceTypes, int64, error)
	GetCulturesResourceLangByKeyId(keyId int) ([]entity.CulturesResourceLangs, error)
}

// 确保 CulturesRepository 实现了接口 (编译时检查)
var _ CulturesRepository = (*CulturesRepositoryImpl)(nil)

func NewCulturesRepository(configManager *config.ConfigManager) (*CulturesRepositoryImpl, error) {
	str := configManager.GetValue(dbConfigKey)
	if str == "" {
		return nil, errors.New("database config is empty")
	}
	db, err := createEngine(str)
	if err != nil {
		return nil, err
	}
	obj := &CulturesRepositoryImpl{
		db: db,
	}
	configManager.RegisterListener("application", dbConfigKey, obj)

	// 自动建表
	// if err := r.db.Sync2(new(entity.CulturesResources),
	// 	new(entity.CulturesResourceTypes),
	// 	new(entity.CulturesResourceKeys),
	// 	new(entity.CulturesResourceLangs)); err != nil {
	// 	panic(err)
	// }
	return obj, nil
}

var dbConfigKey = "I18ndb"

func createEngine(str string) (*xorm.Engine, error) {
	var cfg config.MySQLConfig
	err := json.Unmarshal([]byte(str), &cfg)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Printf("create engine error: %v", err)
		return nil, err
	}
	engine.ShowSQL(true)
	return engine, nil
}

func (r *CulturesRepositoryImpl) OnConfigUpdate(namespace string, key string, newValue interface{}) {
	r.Lock()
	defer r.Unlock()
	if v, ok := newValue.(string); ok {
		fmt.Printf("database config updated to: %+v\n", v)
		r.db, _ = createEngine(v)
	} else {
		fmt.Printf("database config Invalid type for key: %s, expected string, got: %v\n", key, reflect.TypeOf(newValue))
	}
}

// 获取支持的语言列表
func (r *CulturesRepositoryImpl) GetCultures() ([]entity.CulturesResources, error) {
	var cultures []entity.CulturesResources
	err := r.db.Find(&cultures)
	return cultures, err
}

// 根据 Code 获取语言
func (r *CulturesRepositoryImpl) GetResourcesByCode(code string) ([]entity.CulturesResourceLangs, error) {
	culture := &entity.CulturesResources{
		Code: code,
	}
	has, err := r.db.Get(culture)
	if !has {
		return nil, errors.New("culture not exists")
	}
	if err != nil {
		return nil, err
	}
	var langs []entity.CulturesResourceLangs
	err = r.db.Where("culture_id = ?", culture.ID).Find(&langs)
	return langs, err
}

// 添加或更新语言
func (r *CulturesRepositoryImpl) AddOrUpdateCultures(culture entity.CulturesResources) error {
	source := entity.CulturesResources{
		Code: culture.Code,
	}
	has, err := r.db.Get(&source)
	if err != nil {
		return err
	}
	if has && source.ID != culture.ID {
		return errors.New("culture already exists")
	}
	if culture.ID > 0 {
		_, err := r.db.ID(culture.ID).Update(&culture)
		return err
	} else {
		_, err := r.db.Insert(&culture)
		return err
	}
}

// 添加或更新资源类型
func (r *CulturesRepositoryImpl) AddOrUpdateCulturesResourceType(data entity.CulturesResourceTypes) error {
	source := entity.CulturesResourceTypes{Name: data.Name}
	has, err := r.db.Get(&source)
	if err != nil {
		return err
	}
	if has && source.ID != data.ID {
		return errors.New("culture type already exists")
	}
	if data.ID > 0 {
		_, err := r.db.ID(data.ID).Update(&data)
		return err
	} else {
		_, err := r.db.Insert(&data)
		return err
	}
}

func (r *CulturesRepositoryImpl) DeleteCulturesResourceType(id int64) error {
	_, err := r.db.ID(id).Delete(&entity.CulturesResourceTypes{})
	return err
}

// 添加或更新资源键
func (r *CulturesRepositoryImpl) AddOrUpdateCulturesResourceKey(data entity.CulturesResourceKeys) (*entity.CulturesResourceKeys, error) {
	source := entity.CulturesResourceKeys{Name: data.Name}
	has, err := r.db.Get(&source)
	if err != nil {
		return nil, err
	}
	if has && source.ID != data.ID {
		return &source, errors.New("culture key already exists")
	}
	if data.ID > 0 {
		_, err = r.db.ID(data.ID).Update(&data)
	} else {
		_, err = r.db.Insert(&data)
	}
	return &data, err
}

// 添加或更新资源
func (r *CulturesRepositoryImpl) AddOrUpdateCulturesResourceLang(data entity.CulturesResourceLangs) error {
	source := entity.CulturesResourceLangs{KeyID: data.KeyID, CultureID: data.CultureID}
	has, err := r.db.Get(&source)
	if err != nil {
		return err
	}
	if has && source.ID != data.ID {
		return errors.New("culture lang already exists")
	}
	if data.ID > 0 {
		_, err = r.db.ID(data.ID).Update(&data)
	} else {
		_, err = r.db.Insert(&data)
	}
	return err
}

func (r *CulturesRepositoryImpl) AddCulturesResourceLangs(key string, tid int32, cultureLang []entity.CulturesResourceLangs) error {
	keyData := &entity.CulturesResourceKeys{Name: key}
	has, ex := r.db.Get(keyData)
	if ex != nil {
		return ex
	}
	// 使用 Transaction 方法执行事务
	_, err := r.db.Transaction(func(s *xorm.Session) (interface{}, error) {
		if !has {
			keyData = &entity.CulturesResourceKeys{Name: key, TypeID: tid}
			s.Insert(keyData)
		}
		for _, v := range cultureLang {
			v.KeyID = keyData.ID
			if h, _ := s.Get(&v); h {
				continue
			}
			if _, ex := s.Insert(&v); ex != nil {
				return nil, ex
			}
		}
		return nil, nil
	})
	return err
}

// 获取资源类型分页
func (r *CulturesRepositoryImpl) GetCulturesResourceTypePager(index, size int, text string) ([]entity.CulturesResourceTypes, int64, error) {
	var types []entity.CulturesResourceTypes
	sess := r.db.NewSession()
	defer sess.Close()
	if text != "" {
		sess.Where("name like ?", "%"+text+"%")
	}
	offset := index*size - size
	total, err := sess.Limit(size, offset).FindAndCount(&types)
	return types, total, err
}

// 根据id获取资源类型
func (r *CulturesRepositoryImpl) GetCulturesResourceTypeByIds(ids []int32) ([]entity.CulturesResourceTypes, error) {
	var types []entity.CulturesResourceTypes
	err := r.db.In("id", ids).Find(&types)
	return types, err
}

// 获取资源键分页
func (r *CulturesRepositoryImpl) GetCulturesResourceKeyPager(index, size int, text string) ([]entity.CulturesResourceKeys, int64, error) {
	var keys []entity.CulturesResourceKeys
	sess := r.db.NewSession()
	defer sess.Close()
	if text != "" {
		sess.Where("name like ?", "%"+text+"%")
	}
	offset := index*size - size
	total, err := sess.Limit(size, offset).FindAndCount(&keys)
	return keys, total, err
}

func (r *CulturesRepositoryImpl) GetCulturesResourceKeyByIds(ids []int32) (map[int32]string, error) {
	var types []entity.CulturesResourceKeys
	err := r.db.In("id", ids).Find(&types)
	if err != nil {
		return nil, err
	}
	data := make(map[int32]string)
	for _, v := range types {
		data[v.ID] = v.Name
	}
	return data, nil
}

// 获取资源分页
func (r *CulturesRepositoryImpl) GetCulturesResourceLangPager(index, size, cultureId int, text string) ([]entity.CulturesResourceLangs, int64, error) {
	var langs []entity.CulturesResourceLangs
	sess := r.db.NewSession()
	defer sess.Close()
	if text != "" {
		keyDatas := &[]entity.CulturesResourceKeys{}
		ex := r.db.Where("name like ?", "%"+text+"%").Find(keyDatas)
		if ex == nil {
			var ids []int32
			for _, v := range *keyDatas {
				ids = append(ids, v.ID)
			}
			if len(ids) > 0 {
				sess.In("key_id", ids)
			}
		}
		sess.Where("text like ?", "%"+text+"%")
	}
	if cultureId > 0 {
		sess.Where("culture_id = ?", cultureId)
	}
	offset := index*size - size
	total, err := sess.Limit(size, offset).FindAndCount(&langs)
	return langs, total, err
}

// 根据keyId获取资源
func (r *CulturesRepositoryImpl) GetCulturesResourceLangByKeyId(keyId int) ([]entity.CulturesResourceLangs, error) {
	var langs []entity.CulturesResourceLangs
	err := r.db.Where("key_id = ?", keyId).Find(&langs)
	return langs, err
}

func (r *CulturesRepositoryImpl) DeleteCulturesResourceKey(id int32) error {
	_, err := r.db.Transaction(func(s *xorm.Session) (interface{}, error) {
		_, err := s.ID(id).Delete(&entity.CulturesResourceKeys{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		_, err = s.Where("key_id = ?", id).Delete(&entity.CulturesResourceLangs{})
		return nil, err
	})
	return err
}
