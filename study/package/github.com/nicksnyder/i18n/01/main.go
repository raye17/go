package main

import (
	"fmt"
	"study/package/github.com/nicksnyder/i18n/01/pkg/common"
)

func main() {
	common.InitBundle()
	fmt.Println(common.Translate("zh-TW", "UsernameExists", map[string]interface{}{
		"goods": "商品",
	}))
}
