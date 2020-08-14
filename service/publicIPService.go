package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//FindMyPublicIP fetches public IP from https://api.ipify.org?format=text
func FindMyPublicIP() (string, error) {
	url := "https://api.ipify.org?format=text"

	fmt.Printf("Getting IP address from  ipify ...\n")
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Printf("My IP is:%s\n", ip)
	return string(ip), nil
}
