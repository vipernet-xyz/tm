package psql

import (
	"github.com/vipernet-xyz/tm/state/indexer"
	"github.com/vipernet-xyz/tm/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
