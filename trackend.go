package main

import (
	"github.com/austinletson/trackend/api"
	"github.com/austinletson/trackend/common"
	"github.com/austinletson/trackend/world"
)

func main() {
	sqlLiteData, err := world.NewSqlLiteDB()
	common.LogErr(err)
	env := &api.Env{
		Data: sqlLiteData,
	}
	env.ApiIndex()

}
