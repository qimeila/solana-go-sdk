package main

import (
	"context"
	"fmt"
	"log"

	"github.com/qimeila/solana-go-sdk/client"
	"github.com/qimeila/solana-go-sdk/common"
	"github.com/qimeila/solana-go-sdk/pkg/pointer"
	"github.com/qimeila/solana-go-sdk/program/metaplex/token_metadata"
	"github.com/qimeila/solana-go-sdk/rpc"
	"github.com/qimeila/solana-go-sdk/types"
)

// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz
var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

// 9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde
var alice, _ = types.AccountFromBase58("4voSPg3tYuWbKzimpQK9EbXHmuyy5fUrtXvpLDMLkmY6TRncaTHAKGD8jUg3maB5Jbrd9CkQg4qjJMyN6sQvnEF2")

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// NFT
	nft := common.PublicKeyFromString("FK8eFRgmqewUzr7R6pYUxSPnSNusYmkhAV4e3pUJYdCd")

	// get the metadata account address
	tokenMetadataPubkey, err := token_metadata.GetTokenMetaPubkey(nft)
	if err != nil {
		log.Fatalf("failed to find a valid token metadata, err: %v", err)
	}

	recentBlockhashResponse, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get recent blockhash, err: %v", err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: recentBlockhashResponse.Blockhash,
			Instructions: []types.Instruction{
				token_metadata.UpdateMetadataAccountV2(token_metadata.UpdateMetadataAccountV2Param{
					MetadataAccount: tokenMetadataPubkey,
					UpdateAuthority: feePayer.PublicKey,
					Data: &token_metadata.DataV2{
						Name:                 "Fake Fake SMS #1355",
						Symbol:               "FFSMB",
						Uri:                  "https://34c7ef24f4v2aejh75xhxy5z6ars4xv47gpsdrei6fiowptk2nqq.arweave.net/3wXyF1wvK6ARJ_9ue-O58CMuXrz5nyHEiPFQ6z5q02E",
						SellerFeeBasisPoints: 10000,
						Creators: &[]token_metadata.Creator{
							{
								Address:  alice.PublicKey,
								Verified: false,
								Share:    100,
							},
						},
					},
					NewUpdateAuthority:  &alice.PublicKey,
					PrimarySaleHappened: pointer.Get[bool](true),
					IsMutable:           pointer.Get[bool](true),
				}),
			},
		}),
	})
	if err != nil {
		log.Fatalf("failed to new a tx, err: %v", err)
	}

	sig, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}

	fmt.Println(sig)
}
