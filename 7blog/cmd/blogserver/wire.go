/**
 * @project pracToy
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/14 01:09 5月
 **/
package main

import (
	"7blog/internal/biz"
	"7blog/internal/conf"
	"7blog/internal/data"
	"github.com/google/wire"
	"log"
)



//initApp init 7blog application.  *conf.Server, *conf.Data, log.Logger
func initApp(*conf.Data, log.Logger) (*TApp, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, newApp))
}

