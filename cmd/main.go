package main

import (
	"clashparser/api"
	"clashparser/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	parserFileContent, err := utils.ReadFile2Str("./config.yaml")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = yaml.Unmarshal([]byte(parserFileContent), &api.Parser)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(api.CORSMiddleware())
	api.R(r)
	s := &http.Server{
		Addr:         ":59999",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		//MaxHeaderBytes: 1 << 20,
	}
	err = s.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error(), "lister failed to start!")
	}
}
