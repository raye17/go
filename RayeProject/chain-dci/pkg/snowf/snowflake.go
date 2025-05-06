package snowf

import (
	dciConfig "chain-dci/config"
	"chain-dci/pkg/app"
	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
	"strconv"
	"time"
)

var Provider = wire.NewSet(NewSf)

func NewSf() *snowflake.Node {
	var err error
	var st time.Time
	nodeNum, _ := strconv.Atoi(dciConfig.Data.SnowFlake.NodeNum)
	st, err = time.Parse("2006-01-02", dciConfig.Data.SnowFlake.StartTime)
	if err != nil {
		panic(err)
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, errS := snowflake.NewNode(int64(nodeNum))
	if errS != nil {
		panic(errS)
	}
	return node
}

func GenIDInt64() int64 {
	return app.ModuleClients.SfNode.Generate().Int64()
}

func GetIDBase64() string {
	return app.ModuleClients.SfNode.Generate().Base64()
}
