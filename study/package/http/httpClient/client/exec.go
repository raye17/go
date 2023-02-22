package client

import (
	"fmt"
	"net/http"
)

func DoOps(c *http.Client) error {
	resp, err := c.Get("http://www.hao123.com")
	if err != nil {
		return err
	}
	fmt.Println("results of hao123: ", resp.StatusCode)
	return nil
}
func DefaultBaidu() error {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		return err
	}
	fmt.Println("results of defaultBaidu: ", resp.StatusCode)
	return nil

}
