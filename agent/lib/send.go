package lib

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Send Token to the master using TLS
func SendToken(token string) {
	config := tls.Config{
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &config,
	}
	client := &http.Client{Transport: tr}
	_, err := client.PostForm(
		"https://"+os.Getenv("MasterUrl")+"/tokens/send",
		url.Values{"Host": {os.Getenv("hostname")}, "Token": {token}})
	if err != nil {
		log.Printf("Send Token message failed.")
		panic(err)
	}
}

// 	Update the agent node information
func SendUpdate() {
	info, err := DockerInfo()
	if err != nil {
		panic(err)
	}
	config := tls.Config{
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &config,
	}
	client := &http.Client{Transport: tr}
	_, err = client.PostForm(
		"https://"+os.Getenv("MasterUrl")+"/nodes/update",
		url.Values{"Host": {os.Getenv("hostname")},
			"Role": {"agent"},
			"kv":   {info.KernelVersion},
			"os":   {info.OperatingSystem},
			"dv":   {info.ServerVersion},
		})
	if err != nil {
		log.Printf("Send Join message failed.")
		panic(err)
	}
}

// Send the join package to the master
func SendJoin() {
	info, err := DockerInfo()
	if err != nil {
		panic(err)
	}
	config := tls.Config{
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &config,
	}
	client := &http.Client{Transport: tr}
	_, err = client.PostForm(
		"https://"+os.Getenv("MasterUrl")+"/nodes/join",
		url.Values{"Host": {os.Getenv("hostname")},
			"Role": {"agent"},
			"kv":   {info.KernelVersion},
			"os":   {info.OperatingSystem},
			"dv":   {info.ServerVersion},
		})
	if err != nil {
		log.Printf("Send Join message failed.")
		panic(err)
	}
}
