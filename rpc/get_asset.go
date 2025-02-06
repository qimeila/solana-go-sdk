package rpc

import "context"

type GetAssetResponse JsonRpcResponse[Asset]

type Asset struct {
	Interface   string            `json:"interface"` // V1_NFT, V1_PRINT, LEGACY_NFT, etc.
	Id          string            `json:"id"`
	Content     *AssetContent     `json:"content,omitempty"`
	Authorities []AssetAuthority  `json:"authorities,omitempty"`
	Compression *AssetCompression `json:"compression,omitempty"`
	Grouping    []AssetGrouping   `json:"grouping,omitempty"`
	Royalty     *AssetRoyalty     `json:"royalty,omitempty"`
	Creators    []AssetCreator    `json:"creators,omitempty"`
	Ownership   AssetOwnership    `json:"ownership"`
	Uses        *AssetUses        `json:"uses,omitempty"`
	Supply      *AssetSupply      `json:"supply,omitempty"`
	Mutable     bool              `json:"mutable"`
	Burnt       bool              `json:"burnt"`
	TokenInfo   *AssetTokenInfo   `json:"token_info,omitempty"`
}

type AssetContent struct {
	Schema   string        `json:"$schema"`
	JsonUri  string        `json:"json_uri"`
	Files    []AssetFile   `json:"files,omitempty"`
	Metadata AssetMetadata `json:"metadata"`
	Links    *AssetLinks   `json:"links,omitempty"`
}

type AssetFile struct {
	Uri      string        `json:"uri"`
	Mime     string        `json:"mime"`
	Quality  *AssetQuality `json:"quality,omitempty"`
	Contexts []string      `json:"contexts,omitempty"`
}

type AssetQuality struct {
	// Add quality fields as needed
}

type AssetMetadata struct {
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	Symbol        string           `json:"symbol"`
	TokenStandard string           `json:"token_standard"`
	Attributes    []AssetAttribute `json:"attributes,omitempty"`
}

type AssetAttribute struct {
	Value     interface{} `json:"value"` // can be integer or string
	TraitType string      `json:"trait_type"`
}

type AssetLinks struct {
	ExternalUrl string `json:"external_url"`
	Image       string `json:"image"`
}

type AssetAuthority struct {
	Address string   `json:"address"`
	Scopes  []string `json:"scopes"` // "full", "royalty", "metadata", "extension"
}

type AssetCompression struct {
	AssetHash   string `json:"asset_hash"`
	Compressed  bool   `json:"compressed"`
	CreatorHash string `json:"creator_hash"`
	DataHash    string `json:"data_hash"`
	Eligible    bool   `json:"eligible"`
	LeafId      int64  `json:"leaf_id"`
	Seq         int64  `json:"seq"`
	Tree        string `json:"tree"`
}

type AssetGrouping struct {
	GroupKey   string `json:"group_key"` // e.g., "collection"
	GroupValue string `json:"group_value"`
}

type AssetRoyalty struct {
	BasisPoints         int     `json:"basis_points"`
	Locked              bool    `json:"locked"`
	Percent             float64 `json:"percent"`
	PrimarySaleHappened bool    `json:"primary_sale_happened"`
	RoyaltyModel        string  `json:"royalty_model"` // "creators", "fanout", "single"
	Target              string  `json:"target,omitempty"`
}

type AssetCreator struct {
	Address  string `json:"address"`
	Share    int    `json:"share"`
	Verified bool   `json:"verified"`
}

type AssetOwnership struct {
	Delegate       string `json:"delegate,omitempty"`
	Delegated      bool   `json:"delegated"`
	Frozen         bool   `json:"frozen"`
	Owner          string `json:"owner"`
	OwnershipModel string `json:"ownership_model"` // "single", "token"
}

type AssetUses struct {
	Remaining int64  `json:"remaining"`
	Total     int64  `json:"total"`
	UseMethod string `json:"use_method"` // "burn", "multiple", "single"
}

type AssetSupply struct {
	EditionNonce       *int64 `json:"edition_nonce,omitempty"`
	PrintCurrentSupply *int64 `json:"print_current_supply,omitempty"`
	PrintMaxSupply     *int64 `json:"print_max_supply,omitempty"`
}

type AssetTokenInfo struct {
	Supply          int64                  `json:"supply"`
	Decimals        int                    `json:"decimals"`
	TokenProgram    string                 `json:"token_program"`
	MintAuthority   string                 `json:"mint_authority,omitempty"`
	FreezeAuthority string                 `json:"freeze_authority,omitempty"`
	MintExtensions  map[string]interface{} `json:"mint_extensions,omitempty"`
}

func (c *RpcClient) GetAsset(ctx context.Context, assetId string) (JsonRpcResponse[Asset], error) {
	return call[JsonRpcResponse[Asset]](c, ctx, "getAsset", assetId)
}
