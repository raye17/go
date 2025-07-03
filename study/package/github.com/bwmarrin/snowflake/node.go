package snowflakeNode

import (
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
)

func NewSf() *snowflake.Node {
	var err error
	var st time.Time
	nodeNum, _ := strconv.Atoi("5")
	st, err = time.Parse("2006-01-02", "2023-05-31")
	if err != nil {
		panic(err)
	}
	snowflake.Epoch = st.UnixNano() / 1e6
	node, errs := snowflake.NewNode(int64(nodeNum))
	if errs != nil {
		panic(errs)
	}
	return node
}
