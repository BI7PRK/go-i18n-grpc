package rpc

import (
	"context"
	"i18n-service/config"
	"i18n-service/data/entity"
	"i18n-service/data/repository"
	"i18n-service/proto"
	"log"

	"github.com/jinzhu/copier"
)

type CulturesRpc struct {
	proto.UnimplementedI18NServiceServer
	repo repository.CulturesRepository
}

// NewCulturesRpc 创建并初始化一个新的 CulturesRpc 实例。
// 该函数接收一个配置管理器指针 cfgManager，用于访问配置信息。
// 返回值是一个指向 CulturesRpc 结构的指针，该结构包含了
// 用于操作文化的 RPC（远程过程调用）相关功能。
func NewCulturesRpc(cfgManager *config.ConfigManager) *CulturesRpc {
	// 使用配置管理器创建一个新的文化仓库实例。
	// 这里忽略了错误处理，因为示例代码没有提供错误处理的逻辑。
	var _repo, _ = repository.NewCulturesRepository(cfgManager)

	// 创建并返回一个新的 CulturesRpc 实例，将上面创建的仓库实例传递给它。
	// 这表示该 Rpc 实例将使用这个仓库实例来进行数据操作。
	return &CulturesRpc{
		repo: _repo,
	}
}

// CultureFeature 处理文化特征的相关请求。
// 该方法根据传入的Action类型执行不同的操作，支持列出文化特征和添加或更新文化特征。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的值。
//	req - 包含请求信息的数据结构，包括Action类型和参数数据。
//
// 返回值:
//
//	*proto.CulturesReply - 包含响应信息的数据结构，包括操作结果、消息和代码。
//	error - 如果在处理请求过程中遇到错误，则返回该错误。
func (c *CulturesRpc) CultureFeature(ctx context.Context, req *proto.CulturesRequest) (*proto.CulturesReply, error) {
	// md := NewMetadataContext(ctx)
	// for k, v := range md.GetMetadata() {
	// 	log.Println("metadata:", k, v)
	// }

	// 根据请求的Action类型执行相应的逻辑。
	switch req.Action {
	case proto.ActionTypes_List:
		// 处理列出文化特征的请求。
		cultures, err := c.repo.GetCultures()
		if err != nil {
			// 如果从数据库获取文化特征时发生错误，返回错误响应。
			return &proto.CulturesReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
		var data []*proto.CultureItem
		for _, culture := range cultures {
			var item proto.CultureItem
			copier.Copy(&item, culture) // 自动映射字段
			data = append(data, &item)
		}
		// 返回成功响应，包含文化特征列表。
		return &proto.CulturesReply{Items: data, Code: proto.ReplyCode_Success}, nil

	case proto.ActionTypes_AddOrUpdate:
		// 处理添加或更新文化特征的请求。
		if req.ParamData == nil {
			// 如果请求参数为空，返回错误响应。
			return &proto.CulturesReply{Message: "param data is null", Code: proto.ReplyCode_InvalidParam}, nil
		}
		culture := &entity.CulturesResources{}
		if err := copier.Copy(culture, req.ParamData); err != nil {
			// 如果参数数据复制到文化特征资源时发生错误，返回错误响应。
			return &proto.CulturesReply{Message: err.Error(), Code: proto.ReplyCode_Error}, nil
		}
		if err := c.repo.AddOrUpdateCultures(*culture); err != nil {
			// 如果添加或更新数据库中的文化特征时发生错误，返回错误响应。
			return &proto.CulturesReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
		// 返回成功响应，表示操作成功。
		return &proto.CulturesReply{Code: proto.ReplyCode_Success}, nil

	}
	// 如果请求的Action类型不受支持，返回错误响应。
	return &proto.CulturesReply{Message: "not support action " + req.Action.String(), Code: proto.ReplyCode_InvalidAction}, nil
}

// CulturesResourceTypeFeature 根据请求处理文化资源类型的功能。
// 该方法根据不同的操作类型（列表、添加或更新、删除）处理文化资源类型的请求。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的数据、取消信号等。
//	req - 包含操作类型和其他请求数据的结构体。
//
// 返回值:
//
//	*proto.CulturesTypesReply - 包含操作结果或其他相关信息的回复结构体。
//	error - 错误信息，如果操作成功则为nil。
func (c *CulturesRpc) CulturesResourceTypeFeature(ctx context.Context, req *proto.CultureTypesRequest) (*proto.CulturesTypesReply, error) {
	switch req.Action {
	case proto.ActionTypes_List:
		// 处理列表操作，根据请求的条件获取文化资源类型列表和总数。
		var cultures []entity.CulturesResourceTypes
		var total int64
		var err error
		if len(req.CultureIds) > 0 {
			// 如果请求中包含文化ID列表，则根据这些ID获取文化资源类型。
			cultures, err = c.repo.GetCulturesResourceTypeByIds(req.CultureIds)
		} else {
			// 否则，根据请求的分页信息和查找关键词获取文化资源类型列表和总数。
			var findKey string
			if req.ParamData != nil {
				findKey = req.ParamData.Name
			}
			cultures, total, err = c.repo.GetCulturesResourceTypePager(int(req.Index), int(req.Size), findKey)
		}
		if err != nil {
			// 如果发生错误，返回错误信息和数据库错误代码。
			return &proto.CulturesTypesReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
		var data []*proto.CultureTypeItem
		for _, culture := range cultures {
			var item proto.CultureTypeItem
			// 将查询到的文化资源类型数据复制到回复项中。
			if err := copier.Copy(&item, culture); err != nil {
				log.Printf("Failed to copy CultureTypeItem: %v", err)
				continue
			}
			data = append(data, &item)
		}
		// 返回成功回复，包含文化资源类型列表和总数。
		return &proto.CulturesTypesReply{Items: data, Total: total, Code: proto.ReplyCode_Success}, nil

	case proto.ActionTypes_AddOrUpdate:
		// 处理添加或更新操作，根据请求数据添加或更新文化资源类型。
		if req.ParamData == nil {
			// 如果请求数据为空，返回无效参数错误。
			return &proto.CulturesTypesReply{Message: "param data is null", Code: proto.ReplyCode_InvalidParam}, nil
		}
		culture := &entity.CulturesResourceTypes{}
		// 将请求数据复制到文化资源类型实体中。
		if err := copier.Copy(culture, req.ParamData); err != nil {
			return &proto.CulturesTypesReply{Message: err.Error(), Code: proto.ReplyCode_Error}, nil
		}
		// 调用仓库方法添加或更新文化资源类型。
		if err := c.repo.AddOrUpdateCulturesResourceType(*culture); err != nil {
			return &proto.CulturesTypesReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
		// 返回成功回复。
		return &proto.CulturesTypesReply{Code: proto.ReplyCode_Success}, nil

	case proto.ActionTypes_Delete:
		// 处理删除操作，根据请求的ID删除文化资源类型。
		if req.ParamData == nil || req.ParamData.Id <= 0 {
			// 如果请求数据为空或ID无效，返回无效参数错误。
			return &proto.CulturesTypesReply{Message: "param data is null", Code: proto.ReplyCode_InvalidParam}, nil
		}
		// 调用仓库方法删除文化资源类型。
		if err := c.repo.DeleteCulturesResourceType(req.ParamData.Id); err != nil {
			return &proto.CulturesTypesReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
		// 返回成功回复。
		return &proto.CulturesTypesReply{Code: proto.ReplyCode_Success}, nil

	}
	// 如果操作类型不受支持，返回无效操作错误。
	return &proto.CulturesTypesReply{
		Message: "not support action " + req.Action.String(),
		Code:    proto.ReplyCode_InvalidAction,
	}, nil
}

// CulturesResourceKeyFeature 根据不同的操作类型处理文化资源键相关的请求。
// 该函数支持三种操作类型：List（列出资源）、AddOrUpdate（添加或更新资源）、Delete（删除资源）。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的数据、取消信号等。
//	req - 包含操作类型和请求数据的结构体。
//
// 返回值:
//
//	*proto.CultureKeysReply - 包含操作结果的响应对象。
//	error - 错误对象，如果操作成功则为nil。
func (c *CulturesRpc) CulturesResourceKeyFeature(ctx context.Context, req *proto.CultureKeysRequest) (*proto.CultureKeysReply, error) {
	switch req.Action {
	case proto.ActionTypes_List:
		// 处理列出文化资源键的请求。
		var findKey string
		if req.ParamData != nil {
			findKey = req.ParamData.Name
		}
		cultures, total, err := c.repo.GetCulturesResourceKeyPager(int(req.Index), int(req.Size), findKey)
		if err != nil {
			return &proto.CultureKeysReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
		var tids []int32
		for _, v := range cultures {
			tids = append(tids, int32(v.TypeID))
		}
		types := make(map[int32]string)
		if len(tids) > 0 {
			cultureTypes, _ := c.repo.GetCulturesResourceTypeByIds(tids)
			for _, item := range cultureTypes {
				types[item.ID] = item.Name
			}
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
		// 处理添加或更新文化资源键的请求。
		if req.ParamData == nil {
			return &proto.CultureKeysReply{Message: "param data is null", Code: proto.ReplyCode_InvalidParam}, nil
		}
		culture := &entity.CulturesResourceKeys{}
		if err := copier.Copy(culture, req.ParamData); err != nil {
			return &proto.CultureKeysReply{Message: err.Error(), Code: proto.ReplyCode_Error}, nil
		}
		if _, err := c.repo.AddOrUpdateCulturesResourceKey(*culture); err != nil {
			return &proto.CultureKeysReply{Message: err.Error()}, nil
		}
		return &proto.CultureKeysReply{Code: proto.ReplyCode_Success, Message: "ok"}, nil
	case proto.ActionTypes_Delete:
		// 处理删除文化资源键的请求。
		if req.ParamData == nil || req.ParamData.Id <= 0 {
			return &proto.CultureKeysReply{Message: "param data is null", Code: proto.ReplyCode_InvalidParam}, nil
		}
		if err := c.repo.DeleteCulturesResourceKey(req.ParamData.Id); err != nil {
			return &proto.CultureKeysReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
	}
	// 如果操作类型不匹配任何已知操作，返回错误响应。
	return &proto.CultureKeysReply{Code: proto.ReplyCode_InvalidAction, Message: "not support action " + req.Action.String()}, nil
}

func (c *CulturesRpc) AddResourceKeyValue(ctx context.Context, req *proto.AddCultureKeyValueRequest) (*proto.CultureBaseReply, error) {
	var cultureLang []entity.CulturesResourceLangs
	for _, v := range req.Values {
		cultureLang = append(cultureLang, entity.CulturesResourceLangs{CultureID: v.CultureId, Text: v.Text})
	}
	if err := c.repo.AddCulturesResourceLangs(req.Key, req.TypeId, cultureLang); err != nil {
		return &proto.CultureBaseReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
	}
	return &proto.CultureBaseReply{Code: proto.ReplyCode_Success, Message: "ok"}, nil
}

// GetCultureResources 获取特定文化的资源。
// 该方法根据文化代码请求从数据库中提取相应的资源信息，并构建文化资源响应对象返回。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的 deadline、取消信号等。
//	req - 包含文化代码的请求对象。
//
// 返回值:
//
//	*proto.CultureResourcesReply - 包含文化资源信息的响应对象，包括响应码、消息和资源项列表。
//	error - 错误对象，如果在处理请求过程中遇到错误，则可能不为 nil。
func (c *CulturesRpc) GetCultureResources(ctx context.Context, req *proto.CultureCodeRequest) (*proto.CultureResourcesReply, error) {
	// 检查请求参数是否为空，如果为空则返回错误响应。
	if req.Code == "" {
		return &proto.CultureResourcesReply{Message: "param data is null", Code: proto.ReplyCode_InvalidParam}, nil
	}

	// 根据文化代码从数据库中获取资源。
	resource, err := c.repo.GetResourcesByCode(req.Code)
	if err != nil {
		// 如果数据库操作失败，则返回错误响应。
		return &proto.CultureResourcesReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
	}

	// 获取文化资源的键值对数据。
	keyData, ex := c.repo.GetCulturesResourceKeys()
	if ex != nil {
		// 如果获取键值对数据失败，则返回错误响应。
		return &proto.CultureResourcesReply{Message: ex.Error(), Code: proto.ReplyCode_DataBaseError}, nil
	}

	// 初始化一个映射，用于存储资源的键值对。
	langs := make(map[int32]string)
	for _, v := range resource {
		langs[v.KeyID] = v.Text
	}

	// 初始化文化资源项列表。
	var culture []*proto.CultureResourceItem
	var text = ""
	for id, v := range keyData {
		if langs[id] != "" {
			text = langs[id]
		} else {
			text = v
		}
		// 构建文化资源项并添加到列表中。
		culture = append(culture, &proto.CultureResourceItem{Key: v, Text: text})
	}

	// 返回成功响应，包含文化资源项列表。
	return &proto.CultureResourcesReply{Code: proto.ReplyCode_Success, Message: "ok", Items: culture}, nil
}

// CulturesResourceKeyValueFeature 处理文化资源的键值对特征请求。
// 该方法根据请求的动作类型来执行相应的操作，目前只支持列表操作。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的数据、取消信号等。
//	req - 包含请求信息的数据结构，包括动作类型、页码、每页大小等。
//
// 返回值:
//
//	*proto.CultureKeyValuesReply - 包含操作结果或数据的响应对象。
//	error - 错误对象，如果发生错误。
func (c *CulturesRpc) CulturesResourceKeyValueFeature(ctx context.Context, req *proto.CultureKeyValuesRequest) (*proto.CultureKeyValuesReply, error) {
	// 根据请求的动作类型执行相应的操作。
	switch req.Action {
	case proto.ActionTypes_List:
		// 初始化查询键，如果请求中包含参数数据且有搜索键，则使用之。
		var findKey string
		if req.ParamData != nil {
			findKey = req.SearchKey
		}
		// 初始化文化ID，如果请求中包含参数数据且有文化ID，则使用之。
		cultureId := int32(0)
		if req.ParamData != nil {
			cultureId = req.ParamData.CultureId
		}
		// 调用仓库方法获取分页的文化资源数据。
		cultures, total, err := c.repo.GetCulturesResourceLangPager(int(req.Index), int(req.Size), int(cultureId), findKey)
		if err != nil {
			// 如果发生错误，返回错误响应。
			return &proto.CultureKeyValuesReply{Message: err.Error(), Code: proto.ReplyCode_DataBaseError}, nil
		}
		// 遍历查询结果，转换为响应所需的格式。
		var data []*proto.CultureKeyValueItem
		for _, culture := range cultures {
			var item proto.CultureKeyValueItem
			copier.Copy(&item, culture)
			data = append(data, &item)
		}
		// 返回成功响应，包含查询到的数据和总记录数。
		return &proto.CultureKeyValuesReply{Items: data, Total: total}, nil
	}

	// 如果请求的动作类型不受支持，返回错误响应。
	return &proto.CultureKeyValuesReply{Code: proto.ReplyCode_InvalidAction, Message: "not support action " + req.Action.String()}, nil
}
