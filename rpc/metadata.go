package rpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/metadata"
)

type MetadataContext struct {
	context.Context
}

func GetMetadataFromContext(ctx context.Context, key string) ([]string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, false
	}
	values := md.Get(key)
	return values, len(values) > 0
}

func NewMetadataContext(ctx context.Context) *MetadataContext {
	return &MetadataContext{ctx}
}

func (c *MetadataContext) GetString(key string) (string, error) {
	values, err := c.GetStringSlice(key)
	if err != nil {
		return "", err
	}
	return values[0], nil
}

func (c *MetadataContext) GetStringSlice(key string) ([]string, error) {
	values, ok := GetMetadataFromContext(c.Context, key)
	if !ok || len(values) == 0 {
		return nil, errors.New("missing token in metadata")
	}
	return values, nil
}

func (c *MetadataContext) GetMetadate() map[string][]string {
	data := make(map[string][]string)
	md, ok := metadata.FromIncomingContext(c.Context)
	if !ok {
		return data
	}
	for k, v := range md {
		data[k] = v
	}
	return data
}
