package global

import (
	ut "github.com/go-playground/universal-translator"
	"mxshop_api/config"
	"mxshop_api/proto"
)

var (
	Trans          ut.Translator
	ServerConfig   = &config.ServerConfig{}
	GoodsSrvClient proto.GoodsClient
	NacosConfig    = &config.NacosConfig{}

	OrderSrvClient     proto.OrderClient
	InventorySrvClient proto.InventoryClient
)
