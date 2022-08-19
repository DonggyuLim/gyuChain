package blockchain

type block struct {
	Hash         string
	PrevHash     string
	Height       int
	Difficulty   int
	Nonce        int
	Timestamp    int
	Transactions []*Tx
}
