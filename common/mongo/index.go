package mongo

import opts "github.com/qiniu/qmgo/options"

type Indexer interface {
	GetIndexes() []opts.IndexModel
}
