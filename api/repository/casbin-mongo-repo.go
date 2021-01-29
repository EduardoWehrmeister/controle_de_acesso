package repository

import (
	"github.com/casbin/casbin/v2/persist"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
)

type repoCasbinMongo struct{}

var (
	connectionString string = "localhost:27017"
)

func NewCasbinMongoRepository() ICasbinRepository {
	return &repoCasbinMongo{}
}

func (*repoCasbinMongo) GetTheAdapter() persist.BatchAdapter {
	adapter, err1 := mongodbadapter.NewAdapter(connectionString)
	if err1 != nil {
		panic(err1)
	}
	return adapter
}
