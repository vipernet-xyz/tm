package v0

import (
	"github.com/vipernet-xyz/tm/abci/example/kvstore"
	"github.com/vipernet-xyz/tm/config"
	mempl "github.com/vipernet-xyz/tm/mempool"
	mempoolv0 "github.com/vipernet-xyz/tm/mempool/v0"
	"github.com/vipernet-xyz/tm/proxy"
)

var mempool mempl.Mempool

func init() {
	app := kvstore.NewApplication()
	cc := proxy.NewLocalClientCreator(app)
	appConnMem, _ := cc.NewABCIClient()
	err := appConnMem.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false
	mempool = mempoolv0.NewCListMempool(cfg, appConnMem, 0)
}

func Fuzz(data []byte) int {
	err := mempool.CheckTx(data, nil, mempl.TxInfo{})
	if err != nil {
		return 0
	}

	return 1
}
