package service

import (
	"i18n-service/data/entity"
	"i18n-service/data/repository"
)

type CulturesService struct {
	culturesRepository *repository.CulturesRepository
}

func NewCulturesService() *CulturesService {
	return &CulturesService{culturesRepository: repository.NewCulturesRepository()}
}

func (s *CulturesService) GetCultures() ([]entity.CulturesResources, error) {
	return s.culturesRepository.GetCultures()
}
func (s *CulturesService) GetResourcesByCode(code string) ([]entity.CulturesResourceLangs, error) {
	return s.culturesRepository.GetResourcesByCode(code)
}
func (s *CulturesService) AddOrUpdateCultures(culture entity.CulturesResources) error {
	return s.culturesRepository.AddOrUpdateCultures(culture)
}
func (s *CulturesService) AddOrUpdateCulturesResourceType(data entity.CulturesResourceTypes) error {
	return s.culturesRepository.AddOrUpdateCulturesResourceType(data)
}
func (s *CulturesService) DeleteCulturesResourceType(id int64) error {
	return s.culturesRepository.DeleteCulturesResourceType(id)
}

func (s *CulturesService) AddOrUpdateCulturesResourceKey(data entity.CulturesResourceKeys) (*entity.CulturesResourceKeys, error) {
	return s.culturesRepository.AddOrUpdateCulturesResourceKey(data)
}
func (s *CulturesService) AddOrUpdateCulturesResourceLang(data entity.CulturesResourceLangs) error {
	return s.culturesRepository.AddOrUpdateCulturesResourceLang(data)
}

func (s *CulturesService) AddCulturesResourceLangs(key string, tid int32, cultureLang []entity.CulturesResourceLangs) error {
	return s.culturesRepository.AddCulturesResourceLangs(key, tid, cultureLang)
}

func (s *CulturesService) GetCulturesResourceTypePager(index, size int, text string) ([]entity.CulturesResourceTypes, int64, error) {
	return s.culturesRepository.GetCulturesResourceTypePager(index, size, text)
}
func (s *CulturesService) GetCulturesResourceTypeByIds(ids []int32) ([]entity.CulturesResourceTypes, error) {
	return s.culturesRepository.GetCulturesResourceTypeByIds(ids)
}
func (s *CulturesService) GetCulturesResourceKeyPager(index, size int, text string) ([]entity.CulturesResourceKeys, int64, error) {
	return s.culturesRepository.GetCulturesResourceKeyPager(index, size, text)
}
func (s *CulturesService) GetCulturesResourceKeyByIds(ids []int32) (map[int32]string, error) {
	return s.culturesRepository.GetCulturesResourceKeyByIds(ids)
}
func (s *CulturesService) GetCulturesResourceLangPager(index, size, cultureId int, text string) ([]entity.CulturesResourceLangs, int64, error) {
	return s.culturesRepository.GetCulturesResourceLangPager(index, size, cultureId, text)
}
func (s *CulturesService) GetCulturesResourceLangByKeyId(keyId int) ([]entity.CulturesResourceLangs, error) {
	return s.culturesRepository.GetCulturesResourceLangByKeyId(keyId)
}
