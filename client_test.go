package ethclient

import (
	"context"
	"fmt"
	"testing"
)

func TestEthereumClient(t *testing.T) {
	cli := NewEthereumClient("http://127.0.0.1:8545")
	height, err := cli.BlockNumber(context.Background())
	if err != nil {
		fmt.Printf("error %s\n", err)
		return
	}
	fmt.Printf("blockchain height %v\n", height)
}
