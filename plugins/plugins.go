package plugins

import (
	"github.com/kuuland/kuu"
	"github.com/kuuland/kuu/plugins/mongo"
)

// Mongo MongoDB插件别名
func Mongo() *kuu.Plugin {
	return mongo.Plugin()
}