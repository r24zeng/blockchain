package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgInviteGame{}, "tttgame/InviteGame", nil)
	cdc.RegisterConcrete(MsgAcceptGame{}, "tttgame/AcceptGame", nil)
	cdc.RegisterConcrete(MsgPlayGame{}, "tttgame/PlayGame", nil)
	cdc.RegisterConcrete(Game{}, "tttgame/Game", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
