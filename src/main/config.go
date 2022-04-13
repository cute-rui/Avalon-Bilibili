package main

import (
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
    "log"
    "os"
    "strings"
)

var Conf = viper.New()

func confInit() {
    Conf.SetConfigType("toml")
    Conf.SetConfigName("avalon-bilibili")
    Conf.AddConfigPath(`./soft/avalon/config/`)
    Conf.SetDefault("bilibili.sessdata", "")
    Conf.SetDefault("service.addr", ":6231")
    Conf.SetDefault("proxy.host", "")
    Conf.SetDefault("proxy.region.hk", "")
    Conf.SetDefault("proxy.region.th", "")
    Conf.SetDefault("proxy.region.tw", "")
    replacer := strings.NewReplacer(".", "_")
    Conf.SetEnvKeyReplacer(replacer)
    err := Conf.ReadInConfig()
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
        _, err := os.Create("./soft/avalon/config/avalon-bilibili.toml")
        if err != nil {
            log.Println(err)
            return
        }
    }
    
    err = Conf.WriteConfig()
    if err != nil {
        log.Println(err)
        return
    }
    
    Conf.WatchConfig()
    Conf.OnConfigChange(func(in fsnotify.Event) {
        err := Conf.ReadInConfig()
        if err != nil {
            log.Println(err)
        }
    })
}

func init() {
    confInit()
}
