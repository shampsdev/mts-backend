package adapter

import (
	"fmt"

	jsonadapter "api.mts.shamps.dev/external/adapter/json"
)

type AdapterType string

const (
	JsonAdapterType AdapterType = "json"
)

type AdapterFactory struct{}

func (af *AdapterFactory) NewAdapter(adapterType AdapterType) (Adapter, error) {
	switch adapterType {
	case JsonAdapterType:
		return jsonadapter.NewJsonAdapter(), nil
	default:
		return nil, fmt.Errorf("unknown adapter type: %s", adapterType)
	}
}
