package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime/debug"
	"xlib/log"

	"github.com/XzavierLuo/mahjong/config"
)

var (
	help       bool
	configFile string
	pidFile    string
	cfg        *config.Config
)

func parseFlag() {
	flag.StringVar(&configFile, "svrConfFile", "config/config.json", "config file")
	flag.StringVar(&pidFile, "pidFile", "/tmp/mahjong.pid", "pid file")
	flag.BoolVar(&help, "version", false, "help")
	flag.Parse()
}

func writePid() {
	pid := os.Getpid()
	f, err := os.OpenFile(pidFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%d\n", pid))
}

func InitLog() {
	err := log.SetLogger(cfg.LogConfig.RollType, cfg.LogConfig.Dir, cfg.LogConfig.File, cfg.LogConfig.Count, cfg.LogConfig.Size, cfg.LogConfig.Unit, cfg.LogConfig.Level, cfg.LogConfig.Compress)
	if err != nil {
		panic(fmt.Sprintf("init log config failed,err: %s", err.Error()))
	}
	go log.HandleSignalChangeLogLevel()
}

func getConfig(exit bool) *config.Config {
	bs, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Errorf("can't load config file. the path is %s. %s", configFile, err)
		if exit {
			os.Exit(1)
		} else {
			return nil
		}
	}
	cfg := &config.Config{}
	err = json.Unmarshal(bs, cfg)
	if err != nil {
		if exit {
			os.Exit(1)
		} else {
			return nil
		}
	}
	return cfg
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("\r\n"+
				"                      |\r\n"+
				"                  \\       /\r\n"+
				"                    .---. \r\n"+
				"               '-.  |   |  .-'\r\n"+
				"                 ___|   |___\r\n"+
				"            -=  [FATAL ERROR]  =-\r\n"+
				"                `---.   .---' \r\n"+
				"             __||__ |   | __||__\r\n"+
				"             '-..-' |   | '-..-'\r\n"+
				"               ||   |   |   ||\r\n"+
				"               ||_.-|   |-,_||\r\n"+
				"             .-\"`   `\"`'`   `\"-.\r\n"+
				"           .'                   '.\r\n",
				err,
				"\r\n"+string(debug.Stack())+"\r\n")
			StopServer()
			os.Exit(1)
		}
	}()

	parseFlag()
	if help {
		flag.PrintDefaults()
	}

	writePid()

	//解析
	cfg = getConfig(true)

	StartServer()
}
