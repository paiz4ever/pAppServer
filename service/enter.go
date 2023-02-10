package service

import (
	"pAppServer/service/business"
	"pAppServer/service/common"
	"pAppServer/service/consumer"
)

type ServiceGroup struct {
	Common   common.CommonServiceGroup
	Business business.BusinessServiceGroup
	Consumer consumer.ConsumerServiceGroup
}

var Group = new(ServiceGroup)
