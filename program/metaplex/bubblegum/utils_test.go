package bubblegum

import (
	"testing"

	"github.com/qimeila/solana-go-sdk/common"
)

func TestGetEditionMark(t *testing.T) {
	tree := common.PublicKeyFromString("TrEEuqmjD6XKzRoqWzyPz8DrWFARV33hdhYKr1BCMyP")
	leafIndex := uint64(805306)
	expectedOutput := common.PublicKeyFromString("3RDSyGbEbENEZAnNsgGqNzxJgsLXQf5GdNkgqhJU4193")

	assetId, err := getLeafAssetId(tree, leafIndex)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if assetId != expectedOutput {
		t.Errorf("Expected %v, got %v", expectedOutput, assetId)
	}
}
