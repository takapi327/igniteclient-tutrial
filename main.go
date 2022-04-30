package main

import (
	"context"
	"fmt"
	"github.com/ignite-hq/cli/ignite/pkg/cosmosclient"
	"github.com/takapi327/ignite-tutrial/x/ignitetutrial/types"
	"log"
	"os"
	"path/filepath"
)

func main() {

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	homePath := filepath.Join(home, ".ignite-tutrial")

	cosmosOptions := []cosmosclient.Option{
		cosmosclient.WithHome(homePath),
	}

	cosmos, err := cosmosclient.New(context.Background(), cosmosOptions...)
	if err != nil {
		log.Fatal(err)
	}

	accountName := "alice"

	address, err := cosmos.Address(accountName)
	if err != nil {
		log.Fatal(err, cosmos)
	}

	msg := &types.MsgCreatePost{
		Creator: address.String(),
		Title:   "Hello!",
		Body:    "This is the first post",
	}

	txResp, err := cosmos.BroadcastTx(accountName, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("MsgCreatePost:\n\n")
	fmt.Println(txResp)

	queryClient := types.NewQueryClient(cosmos.Context)

	queryResp, err := queryClient.Posts(context.Background(), &types.QueryPostsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\n\nAll posts:\n\n")
	fmt.Println(queryResp)
}
