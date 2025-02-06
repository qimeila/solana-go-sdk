package client

import (
	"context"

	"github.com/qimeila/solana-go-sdk/rpc"
)

type Asset struct {
	Interface     string
	Id            string
	Content       *AssetContent
	Authorities   []AssetAuthority
	Compression   *AssetCompression
	Grouping      []AssetGrouping
	Royalty       *AssetRoyalty
	Creators      []AssetCreator
	Ownership     AssetOwnership
	Uses          *AssetUses
	Supply        *int64
	Mutable       bool
	Burnt         bool
	Lamports      *uint64
	Executable    *bool
	RentEpoch     *uint64
	MetadataOwner *string
}

type AssetContent struct {
	Schema   string
	JsonUri  string
	Files    []AssetFile
	Metadata AssetMetadata
	Links    *AssetLinks
}

type AssetFile struct {
	Uri  string
	Mime string
}

type AssetMetadata struct {
	Description   string
	Name          string
	Symbol        string
	TokenStandard string
}

type AssetLinks struct {
	Image       string
	ExternalUrl string
}

type AssetAuthority struct {
	Address string
	Scopes  []string
}

type AssetCompression struct {
	Eligible    bool
	Compressed  bool
	DataHash    string
	CreatorHash string
	AssetHash   string
	Tree        string
	Seq         int64
	LeafId      int64
}

type AssetGrouping struct {
	GroupKey   string
	GroupValue string
	Verified   *bool
}

type AssetRoyalty struct {
	RoyaltyModel        string
	Target              *string
	Percent             float64
	BasisPoints         int
	PrimarySaleHappened bool
	Locked              bool
}

type AssetCreator struct {
	Address  string
	Share    int
	Verified bool
}

type AssetOwnership struct {
	Frozen         bool
	Delegated      bool
	Delegate       *string
	OwnershipModel string
	Owner          string
}

type AssetUses struct {
	UseMethod string
	Remaining int64
	Total     int64
}

func (c *Client) GetAsset(ctx context.Context, assetId string) (*Asset, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.Asset], error) {
			return c.RpcClient.GetAsset(ctx, assetId)
		},
		convertAsset,
	)
}

func convertAsset(v rpc.Asset) (*Asset, error) {
	if v.Id == "" {
		return nil, nil
	}

	return &Asset{
		Interface:     v.Interface,
		Id:            v.Id,
		Content:       convertAssetContent(v.Content),
		Authorities:   convertAssetAuthorities(v.Authorities),
		Compression:   convertAssetCompression(v.Compression),
		Grouping:      convertAssetGrouping(v.Grouping),
		Royalty:       convertAssetRoyalty(v.Royalty),
		Creators:      convertAssetCreators(v.Creators),
		Ownership:     convertAssetOwnership(v.Ownership),
		Uses:          convertAssetUses(v.Uses),
		Supply:        v.Supply,
		Mutable:       v.Mutable,
		Burnt:         v.Burnt,
		Lamports:      v.Lamports,
		Executable:    v.Executable,
		RentEpoch:     v.RentEpoch,
		MetadataOwner: v.MetadataOwner,
	}, nil
}

func convertAssetContent(v *rpc.AssetContent) *AssetContent {
	if v == nil {
		return nil
	}
	return &AssetContent{
		Schema:   v.Schema,
		JsonUri:  v.JsonUri,
		Files:    convertAssetFiles(v.Files),
		Metadata: convertAssetMetadata(v.Metadata),
		Links:    convertAssetLinks(v.Links),
	}
}

func convertAssetFiles(v []rpc.AssetFile) []AssetFile {
	if v == nil {
		return nil
	}
	files := make([]AssetFile, len(v))
	for i, file := range v {
		files[i] = AssetFile{
			Uri:  file.Uri,
			Mime: file.Mime,
		}
	}
	return files
}

func convertAssetMetadata(v rpc.AssetMetadata) AssetMetadata {
	return AssetMetadata{
		Description:   v.Description,
		Name:          v.Name,
		Symbol:        v.Symbol,
		TokenStandard: v.TokenStandard,
	}
}

func convertAssetLinks(v *rpc.AssetLinks) *AssetLinks {
	if v == nil {
		return nil
	}
	return &AssetLinks{
		Image:       v.Image,
		ExternalUrl: v.ExternalUrl,
	}
}

func convertAssetAuthorities(v []rpc.AssetAuthority) []AssetAuthority {
	if v == nil {
		return nil
	}
	authorities := make([]AssetAuthority, len(v))
	for i, auth := range v {
		authorities[i] = AssetAuthority{
			Address: auth.Address,
			Scopes:  auth.Scopes,
		}
	}
	return authorities
}

func convertAssetCompression(v *rpc.AssetCompression) *AssetCompression {
	if v == nil {
		return nil
	}
	return &AssetCompression{
		Eligible:    v.Eligible,
		Compressed:  v.Compressed,
		DataHash:    v.DataHash,
		CreatorHash: v.CreatorHash,
		AssetHash:   v.AssetHash,
		Tree:        v.Tree,
		Seq:         v.Seq,
		LeafId:      v.LeafId,
	}
}

func convertAssetGrouping(v []rpc.AssetGrouping) []AssetGrouping {
	if v == nil {
		return nil
	}
	groupings := make([]AssetGrouping, len(v))
	for i, group := range v {
		groupings[i] = AssetGrouping{
			GroupKey:   group.GroupKey,
			GroupValue: group.GroupValue,
			Verified:   group.Verified,
		}
	}
	return groupings
}

func convertAssetRoyalty(v *rpc.AssetRoyalty) *AssetRoyalty {
	if v == nil {
		return nil
	}
	return &AssetRoyalty{
		RoyaltyModel:        v.RoyaltyModel,
		Target:              v.Target,
		Percent:             v.Percent,
		BasisPoints:         v.BasisPoints,
		PrimarySaleHappened: v.PrimarySaleHappened,
		Locked:              v.Locked,
	}
}

func convertAssetCreators(v []rpc.AssetCreator) []AssetCreator {
	if v == nil {
		return nil
	}
	creators := make([]AssetCreator, len(v))
	for i, creator := range v {
		creators[i] = AssetCreator{
			Address:  creator.Address,
			Share:    creator.Share,
			Verified: creator.Verified,
		}
	}
	return creators
}

func convertAssetOwnership(v rpc.AssetOwnership) AssetOwnership {
	return AssetOwnership{
		Frozen:         v.Frozen,
		Delegated:      v.Delegated,
		Delegate:       v.Delegate,
		OwnershipModel: v.OwnershipModel,
		Owner:          v.Owner,
	}
}

func convertAssetUses(v *rpc.AssetUses) *AssetUses {
	if v == nil {
		return nil
	}
	return &AssetUses{
		UseMethod: v.UseMethod,
		Remaining: v.Remaining,
		Total:     v.Total,
	}
}
