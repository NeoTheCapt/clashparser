package api

import (
	"clashparser/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
)

var Parser ParserConfig

func r_checkToken(c *gin.Context) {
	log.Printf("token: %s \n", c.Param("token"))
	if !utils.IfElementExistInArray(c.Params.ByName("token"), Parser.Tokens) {
		c.String(403, "forbidden")
		return
	}
	c.String(200, "token authenticated")
}

func r_profile(c *gin.Context) {
	log.Printf("token: %s \n", c.Param("token"))
	if !utils.IfElementExistInArray(c.Params.ByName("token"), Parser.Tokens) {
		c.String(403, "forbidden")
		return
	}
	profileUrl, err := utils.Base64Decode(c.Params.ByName("b64url"))
	if err != nil {
		c.String(500, err.Error())
		return
	}
	profile, err := utils.HttpGet(string(profileUrl))
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, profile)
}

func r_parser(c *gin.Context) {
	log.Printf("token: %s \n", c.Param("token"))
	if !utils.IfElementExistInArray(c.Params.ByName("token"), Parser.Tokens) {
		c.String(403, "forbidden")
		return
	}
	profileUrl, err := utils.Base64Decode(c.Params.ByName("b64url"))
	if err != nil {
		c.String(500, err.Error())
		return
	}
	profile, err := utils.HttpGet(string(profileUrl))
	profile = utils.GetStringBetween(profile, "---", "...")
	if profile == "" {
		c.String(500, "profile is empty")
		return
	}
	var config ClashConfig
	larkProxy := Proxy{
		Name:           "lark-beijing",
		Type:           "socks5",
		Server:         Parser.Socks5IP,
		Port:           Parser.Socks5Port,
		TLS:            true,
		Username:       Parser.Socks5User,
		Password:       Parser.Socks5Pass,
		SkipCertVerify: true,
	}
	larkProxyGroup := ProxyGroup{
		Name:    "lark",
		Type:    "select",
		Proxies: []string{"lark-beijing"},
	}
	err = yaml.Unmarshal([]byte(profile), &config)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	config.Proxies = append(config.Proxies, larkProxy)
	config.ProxyGroups = append(config.ProxyGroups, larkProxyGroup)
	os := c.Params.ByName("os")
	switch os {
	case "win":
		config.Rules = win(config.Rules)
		break
	case "mac":
		config.Rules = mac(config.Rules)
		break
	case "ios":
		config.Rules = ios(config.Rules)
		break
	default:
		c.String(500, "os not supported")
		return
	}
	yamlConfig, err := yaml.Marshal(config)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	rst := fmt.Sprintf("---\n%s\n...", yamlConfig)
	c.String(200, rst)
}
