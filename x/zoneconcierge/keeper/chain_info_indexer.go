package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	"github.com/babylonchain/babylon/x/zoneconcierge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) setChainInfo(ctx sdk.Context, chainInfo *types.ChainInfo) {
	store := k.chainInfoStore(ctx)
	store.Set([]byte(chainInfo.ChainId), k.cdc.MustMarshal(chainInfo))
}

// GetChainInfo returns the ChainInfo struct for a chain with a given ID
// Since IBC does not provide API that allows to initialise chain info right before creating an IBC connection,
// we can only check its existence every time, and return an empty one if it's not initialised yet.
func (k Keeper) GetChainInfo(ctx sdk.Context, chainID string) *types.ChainInfo {
	store := k.chainInfoStore(ctx)

	if !store.Has([]byte(chainID)) {
		return &types.ChainInfo{
			ChainId:      chainID,
			LatestHeader: nil,
			LatestForks: &types.Forks{
				Headers: []*types.IndexedHeader{},
			},
		}
	}
	chainInfoBytes := store.Get([]byte(chainID))
	var chainInfo types.ChainInfo
	k.cdc.MustUnmarshal(chainInfoBytes, &chainInfo)
	return &chainInfo
}

func (k Keeper) UpdateLatestHeader(ctx sdk.Context, chainID string, header *types.IndexedHeader) error {
	if header == nil {
		return sdkerrors.Wrapf(types.ErrInvalidHeader, "header is nil")
	}
	// NOTE: we can accept header without ancestor since IBC connection can be established at any height
	chainInfo := k.GetChainInfo(ctx, chainID)
	chainInfo.LatestHeader = header
	k.setChainInfo(ctx, chainInfo)
	return nil
}

func (k Keeper) UpdateLatestForkHeader(ctx sdk.Context, chainID string, header *types.IndexedHeader) error {
	if header == nil {
		return sdkerrors.Wrapf(types.ErrInvalidHeader, "header is nil")
	}
	chainInfo := k.GetChainInfo(ctx, chainID)
	chainInfo.TryToUpdateForkHeader(header)
	k.setChainInfo(ctx, chainInfo)
	return nil
}

// msgChainInfoStore stores the information of canonical chains and forks for CZs
// prefix: ChainInfoKey
// key: chainID
// value: ChainInfo
func (k Keeper) chainInfoStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.ChainInfoKey)
}