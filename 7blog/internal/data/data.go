/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/16 7:51 下午
* Description:
 */
package data

import (
	"7blog/internal/conf"
	"7blog/internal/data/event"
	"github.com/google/wire"
	"log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, event.Events)

// Data .
type Data struct {
	// TODO warpped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Print("message", "closing the data resources")
	}
	return &Data{}, cleanup, nil
}
