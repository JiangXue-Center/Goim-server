package config

import (
	"Goim-server/common/mongo"
	"Goim-server/common/pb"
	"Goim-server/common/utils"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConf       redis.RedisConf
	MongoCollection struct {
		User        mongo.MongoCollectionConf
		UserSetting mongo.MongoCollectionConf
	}
	Account struct {
		JwtConfig      utils.JwtConfig
		UsernameUnique bool     `json:",optional"`
		UserRegex      string   `json:",optional"`
		PhoneUnique    bool     `json:",optional"`
		PhoneRegex     string   `json:",optional"`
		PhoneCode      []string `json:",optional"`
		EmailUnique    bool     `json:",optional"`
		EmailRegex     string   `json:",optional"`
		Register       struct {
			AllowPlatform        []pb.Platform `json:",optional"`
			RequirePassword      bool          `json:",optional"`
			RequireNickname      bool          `json:",optional"`
			DefaultNicknameRule  string        `json:",options=random|fixed"`
			FixedNickname        string        `json:",default=用户"`
			RandomNicknamePrefix string        `json:",default=用户"`
			RequireAvatar        bool          `json:",optional"`
			DefaultAvatarRule    string        `json:",options=byName|fixed"`
			ByNameAvatarBgColors []string
			ByNameAvatarFgColors []string
			FixedAvatar          string `json:",default=group_avatar.png"`
			RequireBindPhone     bool   `json:".optional"`
			RequireBindEmail     bool   `json:",optional"`
			RequireCaptcha       bool   `json:",optional"`
		}
		Login struct {
			AllowPlatform  []pb.Platform `json:",optional"`
			RequireCaptcha bool          `json:",optional"`
		}
		Robot struct {
			AllowCreate     bool   `json:",optional"`
			RequireNickname bool   `json:",optional"`
			DefaultNickname string `json:",default=Robot"`
			RequireAvatar   bool   `json:",optional"`
		}
	}
	RpcClientConf struct {
		Dispatch     zrpc.RpcClientConf
		User         zrpc.RpcClientConf
		Conversation zrpc.RpcClientConf
		Third        zrpc.RpcClientConf
		Message      zrpc.RpcClientConf
	}
}
