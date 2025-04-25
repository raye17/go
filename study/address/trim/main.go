package main

import (
	"fmt"
	"strings"
	"study/address/trim/area"
)

type AddressInfo struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	District      string `json:"district"`
	DetailAddress string `json:"detailAddress"`
}

func main() {
	var address AddressInfo
	addr := ArtistAddressToArtistInfoAddress("")
	fmt.Println(len(addr))
	if addr != nil && len(addr) > 0 {
		addrs := addr[0]
		address.Province = addrs[0]
		address.City = addrs[1]
		address.District = addrs[2]
		address.DetailAddress = addrs[3]
		fmt.Println(address)
	} else {
		fmt.Println("nil")
	}
}
func ArtistAddressToArtistInfoAddress(address string) [][]string {
	if address == "" {
		return [][]string{}
	}
	var addressList [][]string
	if strings.Contains(address, "，") {
		//兼容老版
		p, c, ar, ad := area.AddressParseToFourLevel(address)
		addressList = append(addressList, []string{p, c, ar, ad})
	} else {
		//新版本逻辑
		subAddressList := strings.Split(address, "|")
		for _, v := range subAddressList {
			subAddress := strings.Split(v, ",")
			if len(subAddress) == 4 {
				if strings.Contains(subAddress[3], "省") || strings.Contains(subAddress[3], "市") {
					p, c, ar, ad := area.AddressParseToFourLevel(subAddress[3])
					addressList = append(addressList, []string{p, c, ar, ad})
				} else {
					addressList = append(addressList, subAddress)
				}
			} else if len(subAddress) == 1 {
				p, c, ar, ad := area.AddressParseToFourLevel(subAddress[0])
				addressList = append(addressList, []string{p, c, ar, ad})
			} else {
				maxLen := 4
				var tempSubAddress []string
				for _, sub := range subAddress {
					if sub == "" {
						continue
					}
					if len(tempSubAddress) < maxLen {
						tempSubAddress = append(tempSubAddress, sub)
					} else {
						break
					}
				}
				if len(tempSubAddress) == maxLen {
					if strings.Contains(tempSubAddress[3], "省") || strings.Contains(tempSubAddress[3], "市") {
						p, c, ar, ad := area.AddressParseToFourLevel(tempSubAddress[3])
						addressList = append(addressList, []string{p, c, ar, ad})
					} else {
						addressList = append(addressList, tempSubAddress)
					}
				} else {
					p, c, ar, ad := area.AddressParseToFourLevel(v)
					addressList = append(addressList, []string{p, c, ar, ad})
				}
			}
		}
	}
	return addressList
}
