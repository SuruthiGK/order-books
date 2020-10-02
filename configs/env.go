package configs

import (
	//"os"
	//"strconv"
	"time"

	//"github.com/golang/glog"
)

//Config The important stuff
type env struct {
	OrderBookRefreshTime          time.Duration
	DSRefreshTime          time.Duration
	OBSRefreshTime          time.Duration
}

//Env an instance to manage environment
var Env *env

func init() {
	//Get the catalog refresh time
	Env = &env{
		OrderBookRefreshTime:     5 * time.Second,
		DSRefreshTime:     5 * time.Second,
		OBSRefreshTime:     5 * time.Second,
	}
}
