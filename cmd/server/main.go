package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

const (

	// github
	// https://hub.fastgit.xyz/Alvin9999/new-pac/wiki/Goflyway%E5%85%8D%E8%B4%B9%E8%B4%A6%E5%8F%B7
	// 镜像
	url = "https://hub.fastgit.xyz/Alvin9999/new-pac/wiki/Goflyway%E5%85%8D%E8%B4%B9%E8%B4%A6%E5%8F%B7"

	ipGap       = 5
	portGap     = 9
	passwordGap = 9

	saveFileName = "account.env"
)

type goflywayDTO struct {
	IP       string
	Port     string
	Password string
}

func getIP(body string) (ipStr string) {
	ipIndex := strings.Index(body, "IP：")
	for i := ipIndex; i < len(body)-ipGap; i++ {
		if body[i+ipGap] == '<' {
			break
		}
		if body[i+ipGap] != ' ' {
			ipStr += string([]byte{body[i+ipGap]})
		}
	}
	return
}

func getPort(body string) (portStr string) {
	portIndex := strings.Index(body, "端口：")
	for i := portIndex; i < len(body)-portGap; i++ {
		if body[i+portGap] == '<' {
			break
		}
		if body[i+portGap] != ' ' {
			portStr += string([]byte{body[i+portGap]})
		}
	}
	return
}

func getPassword(body string) (passwordStr string) {
	passwordIndex := strings.Index(body, "密码：")
	for i := passwordIndex; i < len(body)-passwordGap; i++ {
		if body[i+passwordGap] == '<' {
			break
		}
		if body[i+passwordGap] != ' ' {
			passwordStr += string([]byte{body[i+passwordGap]})
		}
	}
	return
}

func saveToFile(dto goflywayDTO) {
	fmt.Printf("dto %#v\n", dto)

	viper.SetConfigFile(saveFileName)
	viper.AddConfigPath(".")

	viper.Set("proxy_ip", dto.IP)
	viper.Set("proxy_port", dto.Port)
	viper.Set("proxy_password", dto.Password)

	err := viper.WriteConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	// TODO 调度
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	body := string(bodyByte)
	ip := getIP(body)
	fmt.Println("ip:", ip)
	port := getPort(body)
	fmt.Println("port:", port)
	password := getPassword(body)
	fmt.Println("password:", password)

	saveToFile(goflywayDTO{
		IP:       ip,
		Port:     port,
		Password: password,
	})
}
