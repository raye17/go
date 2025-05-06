package logic

import (
	"chain-dci/internal/model"
	errCommon "chain-dci/pkg/err"
	"encoding/json"
	"github.com/gocolly/colly/v2"
	"net/url"
	"strconv"
	"strings"
)

func getPayUrlAmount(payUrl string) (amount string) {

	u, err := url.ParseQuery(payUrl)
	if err != nil {
		errCommon.NoReturnInfo(err, "从url中获取数登支付金额 失败: ")
	}

	bizContent := u.Get("biz_content")
	if bizContent != "" {
		//fmt.Println("========================================  url 获取 ======================")
		//fmt.Printf("bizContent: %+v\n", bizContent)
		content := new(model.BizContent)

		err = json.Unmarshal([]byte(bizContent), content)
		if err != nil {
			errCommon.NoReturnInfo(err, "从url中获取数登支付金额 json 解析失败: ")
		}

		amount = strconv.FormatFloat(content.TotalAmount, 'f', -1, 64)
	}

	if amount == "" {
		//fmt.Println("========================================  页面 获取 ======================")
		c := colly.NewCollector()

		c.OnHTML(".ft-center.qrcode-header-money", func(e *colly.HTMLElement) {
			amount = strings.TrimSpace(e.Text)
		})

		err = c.Visit(payUrl)
		if err != nil {
			errCommon.NoReturnInfo(err, "从支付页面中获取数登支付金额 失败: ")
		}
	}

	return
}
