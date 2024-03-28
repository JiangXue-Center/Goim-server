package mongo

import (
	"Goim-server/common/utils"
	"context"
	"github.com/qiniu/qmgo"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoCollectionConf struct {
	Uri             string
	Database        string
	Collection      string
	MaxPoolSize     uint64 `json:",default=100"`
	MinPoolSize     uint64 `json:",optional"`
	SocketTimeoutMS int64  `json:",default=300000"`
	ReadPreference  int    `json:",optional"`
}

func MustNewMongoCollection(config MongoCollectionConf, model any) *qmgo.QmgoClient {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	var readPreference *qmgo.ReadPref
	if config.ReadPreference > 0 && config.ReadPreference < 6 {
		readPreference = &qmgo.ReadPref{
			Mode: readpref.Mode(config.ReadPreference),
		}
	}
	qmgoClient, err := qmgo.Open(ctx, &qmgo.Config{
		Uri:              config.Uri,
		Database:         config.Database,
		Coll:             config.Collection,
		ConnectTimeoutMS: utils.AnyPtr(int64(5000)),
		MaxPoolSize:      utils.AnyPtr(config.MaxPoolSize),
		MinPoolSize:      utils.AnyPtr(config.MinPoolSize),
		SocketTimeoutMS:  utils.AnyPtr(config.SocketTimeoutMS),
		ReadPreference:   readPreference,
	})
	if err != nil {
		logx.Errorf("open mongo error: %v", err)
		//os.Exit(1)
		return nil
	}
	if indexer, ok := model.(Indexer); ok {
		err = CreateIndex(ctx, qmgoClient, indexer)
		if err != nil {
			logx.Errorf("create index error: %v", err)
			return nil
		}
	}
	return qmgoClient
}

func CreateIndex(ctx context.Context, qmgoClient *qmgo.QmgoClient, indexer Indexer) error {
	indexes := indexer.GetIndexes()
	if len(indexes) == 0 {
		return nil
	}
	return qmgoClient.CreateIndexes(ctx, indexes)
}
