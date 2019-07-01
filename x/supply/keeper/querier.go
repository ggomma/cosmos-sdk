package keeper

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/supply/types"
)

// NewQuerier creates a querier for supply REST endpoints
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case types.QueryTotalSupply:
			return queryTotalSupply(ctx, req, k)
		case types.QuerySupplyOf:
			return querySupplyOf(ctx, req, k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown supply query endpoint")
		}
	}
}

func queryTotalSupply(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryTotalSupplyParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	totalSupply := k.GetSupply(ctx).Total
	totalSupplyLen := len(totalSupply)

	if params.Limit == 0 {
		params.Limit = totalSupplyLen
	}

	start := (params.Page - 1) * params.Limit
	end := params.Limit + start
	if end >= totalSupplyLen {
		end = totalSupplyLen
	}

	if start >= totalSupplyLen {
		// page is out of bounds
		totalSupply = sdk.Coins{}
	} else {
		totalSupply = totalSupply[start:end]
	}

	res, err := totalSupply.MarshalJSON()
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to JSON marshal result: %s", err.Error()))
	}

	return res, nil
}

func querySupplyOf(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QuerySupplyOfParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	supply := k.GetSupply(ctx).Total.AmountOf(params.Denom)

	res, err := supply.MarshalJSON()
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to JSON marshal result: %s", err.Error()))
	}

	return res, nil
}