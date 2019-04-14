package lib

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func SendJoin(){
	config := tls.Config{
		InsecureSkipVerify:true,
	}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig: &config,
	}
	client := &http.Client{Transport: tr}
	_, err := client.PostForm(
		"https://"+os.Getenv("MasterUrl")+"/nodes/join",
		url.Values{"Host": {os.Getenv("hostname")}, "Role": {"agent"}})
	if err != nil {
		log.Printf("Send Join message failed.")
		panic(err)
	}
}