package v1

import (
	"pAppServer/api/v1/business"
	"pAppServer/api/v1/common"
	"pAppServer/api/v1/consumer"
)

type ApiGroup struct {
	Common   common.CommonApiGroup
	Business business.BusinessApiGroup
	Consumer consumer.ConsumerApiGroup
}

var Group = new(ApiGroup)
