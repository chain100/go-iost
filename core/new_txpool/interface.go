package txpool

import (
	"github.com/iost-official/Go-IOS-Protocol/core/new_block"
	"github.com/iost-official/Go-IOS-Protocol/core/new_blockcache"
	"github.com/iost-official/Go-IOS-Protocol/core/new_tx"
)

//go:generate mockgen -destination mock/mock_txpool.go -package txpool_mock github.com/iost-official/Go-IOS-Protocol/core/new_txpool TxPool

// TxPool defines all the API of txpool package.
type TxPool interface {
	Start() error
	Stop()
	AddLinkedNode(linkedNode *blockcache.BlockCacheNode, headNode *blockcache.BlockCacheNode) error
	AddTx(tx *tx.Tx) TAddTx
	PendingTxs(maxCnt int) (TxsList, error)
	ExistTxs(hash []byte, chainBlock *block.Block) (FRet, error)
}
