package spider

import (
	"github.com/metacubex/mihomo/common/convert"
	"github.com/metacubex/mihomo/log"
	"pandora-box/backend/constant"
	"sync"
)

func init() {
	Register(constant.CollectV2ray, NewV2rayCollect)
}

type V2ray struct {
	Url string
}

func (c *V2ray) Get() []map[string]any {
	proxies := make([]map[string]any, 0)

	all := GetBytes(c.Url)
	if all != nil {
		v2ray, err := convert.ConvertsV2Ray(all)
		if err == nil && v2ray != nil {
			proxies = v2ray
		}
	}

	return proxies
}

func (c *V2ray) Get2ChanWG(pc chan []map[string]any, wg *sync.WaitGroup) {
	defer wg.Done()
	nodes := c.Get()
	log.Infoln("STATISTIC: V2ray count=%d url=%s", len(nodes), c.Url)
	if len(nodes) > 0 {
		pc <- nodes
	}
}

func NewV2rayCollect(getter Getter) Collect {
	return &V2ray{Url: getter.Url}
}
