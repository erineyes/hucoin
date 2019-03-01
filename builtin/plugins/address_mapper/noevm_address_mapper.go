// +build !evm

package address_mapper

import (
	amtypes "github.com/loomnetwork/go-loom/builtin/types/address_mapper"
	"github.com/loomnetwork/go-loom/plugin"
	contract "github.com/loomnetwork/go-loom/plugin/contractpb"
)

type (
	InitRequest               = amtypes.AddressMapperInitRequest
	GetMappingRequest         = amtypes.AddressMapperGetMappingRequest
	GetMappingResponse        = amtypes.AddressMapperGetMappingResponse
	AddIdentityMappingRequest = amtypes.AddressMapperAddIdentityMappingRequest
)

type AddressMapper struct {
}

func (am *AddressMapper) Meta() (plugin.Meta, error) {
	return plugin.Meta{
		Name:    "addressmapper",
		Version: "0.1.0",
	}, nil
}

func (am *AddressMapper) Init(_ contract.Context, _ *InitRequest) error {
	return nil
}

func (am *AddressMapper) GetMapping(_ contract.StaticContext, _ *GetMappingRequest) (*GetMappingResponse, error) {
	return nil, nil
}

func (am *AddressMapper) AddIdentityMapping(_ contract.Context, _ *AddIdentityMappingRequest) error {
	return nil
}

var Contract plugin.Contract = contract.MakePluginContract(&AddressMapper{})
