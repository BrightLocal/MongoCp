package copier

import (
	"github.com/BrightLocal/MongoCp/dsn"
	"go.mongodb.org/mongo-driver/mongo"
)

type Copier struct {
	srcConn *mongo.Client
	dstConn *mongo.Client
}

func New(src, dst *mongo.Client) *Copier {
	return &Copier{
		srcConn: src,
		dstConn: dst,
	}
}

func (c *Copier) Copy(src, dst dsn.DSN) error {
	panic("implement me")
	return nil
}
