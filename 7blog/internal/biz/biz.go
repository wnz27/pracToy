/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/16 7:43 下午
* Description:
 */
package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewArticleUsecase)

