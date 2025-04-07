package tests

import (
	"context"
	"i18n-service/proto"
	"i18n-service/rpc"
	"testing"
)

func TestCulturesRpc_CultureList(t *testing.T) {

	rpcServer := rpc.NewCulturesRpc()

	// 测试列表功能
	request := &proto.CulturesRequest{
		Action: proto.ActionTypes_List,
	}
	response, err := rpcServer.CultureFeature(context.Background(), request)
	if err != nil {
		t.Fatalf("CultureFeature failed: %v", err)
	}
	if response.Code != proto.ReplyCode_Success {
		t.Fatalf("CultureFeature failed: %v", response.Message)
	}
	if len(response.Items) == 0 {
		t.Fatalf("CultureFeature failed: no items returned")
	}

	t.Logf("CultureFeature success: %v", response.Items)

}

func TestCulturesRpc_CulturesResourceTypeList(t *testing.T) {

	rpcServer := rpc.NewCulturesRpc()

	// 测试列表功能
	request := &proto.CultureTypesRequest{
		Action: proto.ActionTypes_List,
		Index:  0,
		Size:   10,
	}
	response, err := rpcServer.CulturesResourceTypeFeature(context.Background(), request)
	if err != nil {
		t.Fatalf("CulturesResourceTypeFeature failed: %v", err)
	}
	if response.Code != proto.ReplyCode_Success {
		t.Fatalf("CulturesResourceTypeFeature failed: %v", response.Message)
	}
	if len(response.Items) == 0 {
		t.Fatalf("CulturesResourceTypeFeature failed: no items returned")
	}

	t.Logf("CulturesResourceTypeFeature success: %v", response.Items)
}

func TestCulturesRpc_CulturesResourceKeyList(t *testing.T) {

	rpcServer := rpc.NewCulturesRpc()

	// 测试列表功能
	request := &proto.CultureKeysRequest{
		Action: proto.ActionTypes_List,
		Index:  0,
		Size:   10,
	}
	response, err := rpcServer.CulturesResourceKeyFeature(context.Background(), request)
	if err != nil {
		t.Fatalf("CulturesResourceKeyFeature failed: %v", err)
	}
	if response.Code != proto.ReplyCode_Success {
		t.Fatalf("CulturesResourceKeyFeature failed: %v", response.Message)
	}
	if len(response.Items) == 0 {
		t.Fatalf("CulturesResourceKeyFeature failed: no items returned")
	}
	t.Logf("CulturesResourceKeyFeature success: %v", response.Items)
}
