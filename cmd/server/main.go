package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

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

	bashPath    = "/bin/bash"
	restartPath = "/root/zhangxiaofeng/hack/goflyway/restart.sh"
	curlIp      = "/root/zhangxiaofeng/hack/goflyway/curl_ip.sb.sh"

	ipParam       = "IP："
	portParam     = "端口："
	passwordParam = "密码："

	everyTime = 1 * time.Hour
)

type goflywayDTO struct {
	IP       string
	Port     string
	Password string
}

func getParam(body, param string, gap int) (res string) {
	paramIndex := strings.Index(body, param)
	for i := paramIndex; i < len(body)-gap; i++ {
		if body[i+gap] == '<' {
			break
		}
		if body[i+gap] != ' ' {
			res += string([]byte{body[i+gap]})
		}
	}
	return
}

func restartGoflyway() {
	cmd, err := exec.Command(bashPath, restartPath).Output()
	if err != nil {
		panic(err)
	}
	outputString := string(cmd)
	fmt.Println("restart success", outputString)
}

func saveToFile(dto goflywayDTO) {
	fmt.Printf("dto %#v\n", dto)

	viper.SetConfigFile(saveFileName)
	viper.AddConfigPath(".")

	viper.Set("PROXY_IP", dto.IP)
	viper.Set("PROXY_PORT", dto.Port)
	viper.Set("PROXY_PASSWORD", dto.Password)

	err := viper.WriteConfig()
	if err != nil {
		panic(err)
	}

	if dto.IP != "" && dto.Port != "" && dto.Password != "" {
		restartGoflyway()
	}
}

func getBody() (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyByte), nil
}

func getInfoAndRestart() {
	body, err := getBody()
	if err != nil {
		fmt.Printf("from github get body error:%+v\n", err)
	}
	ip := getParam(body, ipParam, ipGap)
	port := getParam(body, portParam, portGap)
	password := getParam(body, passwordParam, passwordGap)

	saveToFile(goflywayDTO{
		IP:       ip,
		Port:     port,
		Password: password,
	})

}

func timerWork() {
	for {
		timer := time.NewTimer(everyTime)
		select {
		case <-timer.C:
			fmt.Println(time.Now())
			out, err := exec.Command(bashPath, curlIp).Output()
			if err != nil {
				log.Printf("proxychain curl ip.sb error:%+v\n", err)
				getInfoAndRestart()
			}
			outStr := string(out)
			if outStr == "" {
				log.Println("ourStr is nil")
			} else {
				// TODO 比较ip和配置ip是否一致
				fmt.Println(curlIp, " : ", outStr)
			}
		}
	}
}

func main() {
	go timerWork()
	for {
	}
}
