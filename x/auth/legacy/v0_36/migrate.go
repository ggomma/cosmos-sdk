// DONTCOVER
// nolint
package v0_36

import (
	v034auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/v0_34"
)

// Migrate accepts exported genesis state from v0.34 and migrates it to v0.36
// genesis state. All entries are identical except for validator slashing events
// which now include the period.
func Migrate(oldGenState v034auth.GenesisState) GenesisState {
	params := Params{
		MaxMemoCharacters: oldGenState.Params.MaxMemoCharacters,
		TxSizeCostPerByte: oldGenState.Params.TxSizeCostPerByte,
		SigVerifyCostED25519: oldGenState.Params.SigVerifyCostED25519,
		SigVerifyCostSecp256k1: oldGenState.Params.SigVerifyCostSecp256k1,
	}
	
	return NewGenesisState(params)
}