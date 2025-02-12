package bubblegum

import (

	"github.com/qimeila/solana-go-sdk/common"
)

func GetLeafAssetId(tree common.PublicKey, leafIndex uint64) (common.PublicKey, error){
		// Convert leafIndex to a byte array in little-endian format
		leafIndexBytes := make([]byte, 8)
		for i := uint64(0); i < 8; i++ {
			leafIndexBytes[i] = byte(leafIndex >> (8 * i))
		}
	
		// Create the seeds for the program address
		seeds := [][]byte{
			[]byte("asset"),
			tree.Bytes(),
			leafIndexBytes,
		}
	
		// Find the program address
		assetId, _, err := common.FindProgramAddress(seeds, common.MetaplexBubblegumProgramID)
		if err != nil {
			return common.PublicKey{}, err
		}
	
		return assetId, nil

}
