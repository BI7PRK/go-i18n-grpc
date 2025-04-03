package rpc

import (
	"context"
	"errors"
	"i18n-service/data/entity"
	"i18n-service/proto"
	"i18n-service/service"
	"log"

	"github.com/jinzhu/copier"
)

type CulturesRpc struct {
	proto.UnimplementedI18NServiceServer
	svc *service.CulturesService
}

func NewCulturesRpc() *CulturesRpc {
	return &CulturesRpc{
		svc: service.NewCulturesService(),
	}
}

func (c *CulturesRpc) CultureFeature(ctx context.Context, req *proto.CulturesRequest) (*proto.CulturesReply, error) {
	md := NewMetadataContext(ctx)
	for k, v := range md.GetMetadate() {
		log.Println("metadata:", k, v)
	}

	switch req.Action {
	case proto.ActionTypes_List:
		cultures, err := c.svc.GetCultures()
		if err != nil {
			return &proto.CulturesReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
		var data []*proto.CultureItem
		for _, culture := range cultures {
			var item proto.CultureItem
			copier.Copy(&item, culture) // 自动映射字段
			data = append(data, &item)
		}
		return &proto.CulturesReply{Items: data, Code: proto.ReplyCode_Success}, nil
	case proto.ActionTypes_AddOrUpdate:
		if req.ParamData == nil {
			return &proto.CulturesReply{Message: "param data is null", Code: proto.ReplyCode_InvalidParam}, nil
		}
		culture := &entity.CulturesResources{}
		if err := copier.Copy(culture, req.ParamData); err != nil {
			return &proto.CulturesReply{Message: err.Error(), Code: proto.ReplyCode_Error}, nil
		}
		if err := c.svc.AddOrUpdateCultures(*culture); err != nil {
			return &proto.CulturesReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
		return &proto.CulturesReply{Code: proto.ReplyCode_Success}, nil
	}
	return &proto.CulturesReply{Message: "not support action " + req.Action.String(), Code: proto.ReplyCode_InvalidAction}, nil
}

func (c *CulturesRpc) CulturesResourceTypeFeature(ctx context.Context, req *proto.CultureTypesRequest) (*proto.CulturesTypesReply, error) {
	switch req.Action {
	case proto.ActionTypes_List:
		var cultures []entity.CulturesResourceTypes
		var total int64
		var err error
		if len(req.CultureIds) > 0 {
			cultures, err = c.svc.GetCulturesResourceTypeByIds(req.CultureIds)
		} else {
			var findKey string
			if req.ParamData != nil {
				findKey = req.ParamData.Name
			}
			cultures, total, err = c.svc.GetCulturesResourceTypePager(int(req.Index), int(req.Size), findKey)
		}
		if err != nil {
			return &proto.CulturesTypesReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
		var data []*proto.CultureTypeItem
		for _, culture := range cultures {
			var item proto.CultureTypeItem
			copier.Copy(&item, culture) // 自动映射字段
		}
		return &proto.CulturesTypesReply{Items: data, Total: total, Code: proto.ReplyCode_Success}, nil

	case proto.ActionTypes_AddOrUpdate:
		if req.ParamData == nil {
			return &proto.CulturesTypesReply{Message: "param data is null", Code: proto.ReplyCode_InvalidParam}, nil
		}
		culture := &entity.CulturesResourceTypes{}
		if err := copier.Copy(culture, req.ParamData); err != nil {
			return &proto.CulturesTypesReply{Message: err.Error(), Code: proto.ReplyCode_Error}, nil
		}
		if err := c.svc.AddOrUpdateCulturesResourceType(*culture); err != nil {
			return &proto.CulturesTypesReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
		return &proto.CulturesTypesReply{}, nil
	case proto.ActionTypes_Delete:
		if req.ParamData == nil || req.ParamData.Id <= 0 {
			return &proto.CulturesTypesReply{Message: "param data is null", Code: proto.ReplyCode_InvalidParam}, nil
		}
		if err := c.svc.DeleteCulturesResourceType(req.ParamData.Id); err != nil {
			return &proto.CulturesTypesReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
		return &proto.CulturesTypesReply{Code: proto.ReplyCode_Success}, nil
	}
	return &proto.CulturesTypesReply{
		Message: "not support action " + req.Action.String(),
		Code:    proto.ReplyCode_InvalidAction,
	}, nil
}

func (c *CulturesRpc) CulturesResourceKeyFeature(ctx context.Context, req *proto.CultureKeysRequest) (*proto.CultureKeysReply, error) {
	switch req.Action {
	case proto.ActionTypes_List:
		var findKey string
		if req.ParamData != nil {
			findKey = req.ParamData.Name
		}
		cultures, total, err := c.svc.GetCulturesResourceKeyPager(int(req.Index), int(req.Size), findKey)
		if err != nil {
			return &proto.CultureKeysReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
		var tids []int32
		for _, v := range cultures {
			tids = append(tids, int32(v.TypeID))
		}
		cultureTypes, _ := c.svc.GetCulturesResourceTypeByIds(tids)
		types := make(map[int32]string)
		for _, item := range cultureTypes {
			types[item.ID] = item.Name
		}
		var data []*proto.CultureKeyItem
		for _, culture := range cultures {
			var item proto.CultureKeyItem
			copier.Copy(&item, culture) // 自动映射字段
			if types[culture.TypeID] != "" {
				item.TypeName = types[culture.TypeID]
			}
			data = append(data, &item)
		}
		return &proto.CultureKeysReply{Items: data, Total: total}, nil
	case proto.ActionTypes_AddOrUpdate:
		if req.ParamData == nil {
			return &proto.CultureKeysReply{Message: "param data is null", Code: proto.ReplyCode_InvalidParam}, nil
		}
		culture := &entity.CulturesResourceKeys{}
		if err := copier.Copy(culture, req.ParamData); err != nil {
			return &proto.CultureKeysReply{Message: err.Error(), Code: proto.ReplyCode_Error}, nil
		}
		if _, err := c.svc.AddOrUpdateCulturesResourceKey(*culture); err != nil {
			return &proto.CultureKeysReply{Message: err.Error()}, nil
		}
		return &proto.CultureKeysReply{Code: proto.ReplyCode_Success, Message: "ok"}, nil
	}
	return nil, errors.New("not support action " + req.Action.String())
}

func (c *CulturesRpc) AddResourceKeyValue(ctx context.Context, req *proto.AddCultureKeyValueRequest) (*proto.CultureBaseReply, error) {
	var cultureLang []entity.CulturesResourceLangs
	for _, v := range req.Values {
		cultureLang = append(cultureLang, entity.CulturesResourceLangs{CultureID: v.CultureId, Text: v.Text})
	}
	if err := c.svc.AddCulturesResourceLangs(req.Key, req.TypeId, cultureLang); err != nil {
		return &proto.CultureBaseReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
	}
	return &proto.CultureBaseReply{Code: proto.ReplyCode_Success, Message: "ok"}, nil
}
