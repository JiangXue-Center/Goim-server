package svc

import (
	"Goim-server/app/user/internal/config"
	"Goim-server/app/user/internal/models"
	"Goim-server/common/cache"
	"Goim-server/common/mongo"
	"Goim-server/common/utils"
	"github.com/qiniu/qmgo"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis

	UserCollection        *qmgo.QmgoClient
	UserSettingCollection *qmgo.QmgoClient

	Jwt *utils.Jwt
}

func NewServiceContext(c config.Config) *ServiceContext {
	s := &ServiceContext{
		Config:         c,
		Redis:          cache.MustNewRedis(c.RedisConf),
		UserCollection: mongo.MustNewMongoCollection(c.MongoCollection.User, &models.UserBasic{}),
		//UserSettingCollection: mongo.MustNewMongoCollection(c.MongoCollection.UserSetting, nil),
	}
	s.Jwt = utils.NewJwt(s.Config.Account.JwtConfig, s.Redis)
	return s
}
