package rpc

import "context"

type GetAssetResponse JsonRpcResponse[Asset]

type Asset struct {
	Interface     string            `json:"interface"`
	Id            string            `json:"id"`
	Content       *AssetContent     `json:"content,omitempty"`
	Authorities   []AssetAuthority  `json:"authorities,omitempty"`
	Compression   *AssetCompression `json:"compression,omitempty"`
	Grouping      []AssetGrouping   `json:"grouping,omitempty"`
	Royalty       *AssetRoyalty     `json:"royalty,omitempty"`
	Creators      []AssetCreator    `json:"creators,omitempty"`
	Ownership     AssetOwnership    `json:"ownership"`
	Uses          *AssetUses        `json:"uses,omitempty"`
	Supply        *int64            `json:"supply"`
	Mutable       bool              `json:"mutable"`
	Burnt         bool              `json:"burnt"`
	Lamports      *uint64           `json:"lamports,omitempty"`
	Executable    *bool             `json:"executable,omitempty"`
	RentEpoch     *uint64           `json:"rent_epoch,omitempty"`
	MetadataOwner *string           `json:"metadata_owner,omitempty"`
}

type AssetContent struct {
	Schema   string        `json:"$schema"`
	JsonUri  string        `json:"json_uri"`
	Files    []AssetFile   `json:"files,omitempty"`
	Metadata AssetMetadata `json:"metadata"`
	Links    *AssetLinks   `json:"links,omitempty"`
}

type AssetFile struct {
	Uri  string `json:"uri"`
	Mime string `json:"mime"`
}

type AssetMetadata struct {
	Description   string `json:"description"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	TokenStandard string `json:"token_standard"`
}

type AssetLinks struct {
	Image       string `json:"image"`
	ExternalUrl string `json:"external_url"`
}

type AssetAuthority struct {
	Address string   `json:"address"`
	Scopes  []string `json:"scopes"`
}

type AssetCompression struct {
	Eligible    bool   `json:"eligible"`
	Compressed  bool   `json:"compressed"`
	DataHash    string `json:"data_hash"`
	CreatorHash string `json:"creator_hash"`
	AssetHash   string `json:"asset_hash"`
	Tree        string `json:"tree"`
	Seq         int64  `json:"seq"`
	LeafId      int64  `json:"leaf_id"`
}

type AssetGrouping struct {
	GroupKey   string `json:"group_key"`
	GroupValue string `json:"group_value"`
	Verified   *bool  `json:"verified,omitempty"`
}

type AssetRoyalty struct {
	RoyaltyModel        string  `json:"royalty_model"`
	Target              *string `json:"target"`
	Percent             float64 `json:"percent"`
	BasisPoints         int     `json:"basis_points"`
	PrimarySaleHappened bool    `json:"primary_sale_happened"`
	Locked              bool    `json:"locked"`
}

type AssetCreator struct {
	Address  string `json:"address"`
	Share    int    `json:"share"`
	Verified bool   `json:"verified"`
}

type AssetOwnership struct {
	Frozen         bool    `json:"frozen"`
	Delegated      bool    `json:"delegated"`
	Delegate       *string `json:"delegate"`
	OwnershipModel string  `json:"ownership_model"`
	Owner          string  `json:"owner"`
}

type AssetUses struct {
	UseMethod string `json:"use_method"`
	Remaining int64  `json:"remaining"`
	Total     int64  `json:"total"`
}

func (c *RpcClient) GetAsset(ctx context.Context, assetId string) (JsonRpcResponse[Asset], error) {
	return call[JsonRpcResponse[Asset]](c, ctx, "getAsset", assetId)
}
