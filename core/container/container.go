package container

import (
	"ginchat/global/config"
	"ginchat/global/consts"
	"sync"
)

var sMap sync.Map

type containers struct{}

func CreateFactoryContainers() *containers {
	return &containers{}
}

func (c *containers) Get(key string) interface{} {
	if value, exists := c.keyIsExist(key); exists {
		return value
	}
	return nil
}

func (c *containers) Set(key string, value interface{}) (res bool) {
	if _, exists := c.keyIsExist(key); !exists {
		sMap.Store(key, value)
		res = true
	} else {
		config.ZapLog.Warn(consts.ErrorsContainerKeyAlreadyExists + ", key name：" + key)
	}

	return
}

func (c *containers) keyIsExist(key string) (interface{}, bool) {
	return sMap.Load(key)
}
