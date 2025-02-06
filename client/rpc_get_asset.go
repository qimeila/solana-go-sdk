package client

import (
	"context"

	"github.com/qimeila/solana-go-sdk/rpc"
)

type Asset struct {
	Interface   string
	Id          string
	Content     *AssetContent
	Authorities []AssetAuthority
	Compression *AssetCompression
	Grouping    []AssetGrouping
	Royalty     *AssetRoyalty
	Creators    []AssetCreator
	Ownership   AssetOwnership
	Uses        *AssetUses
	Supply      *AssetSupply
	Mutable     bool
	Burnt       bool
	TokenInfo   *AssetTokenInfo
}

type AssetContent struct {
	Schema   string
	JsonUri  string
	Files    []AssetFile
	Metadata AssetMetadata
	Links    *AssetLinks
}

type AssetFile struct {
	Uri      string
	Mime     string
	Quality  *AssetQuality
	Contexts []string
}

type AssetQuality struct {
	// Add quality fields as needed
}

type AssetMetadata struct {
	Name          string
	Description   string
	Symbol        string
	TokenStandard string
	Attributes    []AssetAttribute
}

type AssetAttribute struct {
	Value     interface{} // can be integer or string
	TraitType string
}

type AssetLinks struct {
	ExternalUrl string
	Image       string
}

type AssetAuthority struct {
	Address string
	Scopes  []string // "full", "royalty", "metadata", "extension"
}

type AssetCompression struct {
	AssetHash   string
	Compressed  bool
	CreatorHash string
	DataHash    string
	Eligible    bool
	LeafId      int64
	Seq         int64
	Tree        string
}

type AssetGrouping struct {
	GroupKey   string // e.g., "collection"
	GroupValue string
}

type AssetRoyalty struct {
	BasisPoints         int
	Locked              bool
	Percent             float64
	PrimarySaleHappened bool
	RoyaltyModel        string // "creators", "fanout", "single"
	Target              string
}

type AssetCreator struct {
	Address  string
	Share    int
	Verified bool
}

type AssetOwnership struct {
	Delegate       string
	Delegated      bool
	Frozen         bool
	Owner          string
	OwnershipModel string // "single", "token"
}

type AssetUses struct {
	Remaining int64
	Total     int64
	UseMethod string // "burn", "multiple", "single"
}

type AssetSupply struct {
	EditionNonce       *int64
	PrintCurrentSupply *int64
	PrintMaxSupply     *int64
}

type AssetTokenInfo struct {
	Supply          int64
	Decimals        int
	TokenProgram    string
	MintAuthority   string
	FreezeAuthority string
	MintExtensions  map[string]interface{}
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
		Interface:   v.Interface,
		Id:          v.Id,
		Content:     convertAssetContent(v.Content),
		Authorities: convertAssetAuthorities(v.Authorities),
		Compression: convertAssetCompression(v.Compression),
		Grouping:    convertAssetGrouping(v.Grouping),
		Royalty:     convertAssetRoyalty(v.Royalty),
		Creators:    convertAssetCreators(v.Creators),
		Ownership:   convertAssetOwnership(v.Ownership),
		Uses:        convertAssetUses(v.Uses),
		Supply:      convertAssetSupply(v.Supply),
		Mutable:     v.Mutable,
		Burnt:       v.Burnt,
		TokenInfo:   convertAssetTokenInfo(v.TokenInfo),
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
			Uri:      file.Uri,
			Mime:     file.Mime,
			Quality:  convertAssetQuality(file.Quality),
			Contexts: file.Contexts,
		}
	}
	return files
}

func convertAssetMetadata(v rpc.AssetMetadata) AssetMetadata {
	return AssetMetadata{
		Name:          v.Name,
		Description:   v.Description,
		Symbol:        v.Symbol,
		TokenStandard: v.TokenStandard,
		Attributes:    convertAssetAttributes(v.Attributes),
	}
}

func convertAssetLinks(v *rpc.AssetLinks) *AssetLinks {
	if v == nil {
		return nil
	}
	return &AssetLinks{
		ExternalUrl: v.ExternalUrl,
		Image:       v.Image,
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
		AssetHash:   v.AssetHash,
		Compressed:  v.Compressed,
		CreatorHash: v.CreatorHash,
		DataHash:    v.DataHash,
		Eligible:    v.Eligible,
		LeafId:      v.LeafId,
		Seq:         v.Seq,
		Tree:        v.Tree,
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
		}
	}
	return groupings
}

func convertAssetRoyalty(v *rpc.AssetRoyalty) *AssetRoyalty {
	if v == nil {
		return nil
	}
	return &AssetRoyalty{
		BasisPoints:         v.BasisPoints,
		Locked:              v.Locked,
		Percent:             v.Percent,
		PrimarySaleHappened: v.PrimarySaleHappened,
		RoyaltyModel:        v.RoyaltyModel,
		Target:              v.Target,
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
		Delegate:       v.Delegate,
		Delegated:      v.Delegated,
		Frozen:         v.Frozen,
		Owner:          v.Owner,
		OwnershipModel: v.OwnershipModel,
	}
}

func convertAssetUses(v *rpc.AssetUses) *AssetUses {
	if v == nil {
		return nil
	}
	return &AssetUses{
		Remaining: v.Remaining,
		Total:     v.Total,
		UseMethod: v.UseMethod,
	}
}

func convertAssetSupply(v *rpc.AssetSupply) *AssetSupply {
	if v == nil {
		return nil
	}
	return &AssetSupply{
		EditionNonce:       v.EditionNonce,
		PrintCurrentSupply: v.PrintCurrentSupply,
		PrintMaxSupply:     v.PrintMaxSupply,
	}
}

func convertAssetTokenInfo(v *rpc.AssetTokenInfo) *AssetTokenInfo {
	if v == nil {
		return nil
	}
	return &AssetTokenInfo{
		Supply:          v.Supply,
		Decimals:        v.Decimals,
		TokenProgram:    v.TokenProgram,
		MintAuthority:   v.MintAuthority,
		FreezeAuthority: v.FreezeAuthority,
		MintExtensions:  v.MintExtensions,
	}
}

func convertAssetAttributes(v []rpc.AssetAttribute) []AssetAttribute {
	if v == nil {
		return nil
	}
	attributes := make([]AssetAttribute, len(v))
	for i, attr := range v {
		attributes[i] = AssetAttribute{
			Value:     attr.Value,
			TraitType: attr.TraitType,
		}
	}
	return attributes
}

func convertAssetQuality(v *rpc.AssetQuality) *AssetQuality {
	if v == nil {
		return nil
	}
	return &AssetQuality{}
}
