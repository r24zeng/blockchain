package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	tttgameTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	tttgameTxCmd.AddCommand(flags.PostCommands(
    // this line is used by starport scaffolding # 1
		GetCmdCreateGame(cdc),
		GetCmdSetGame(cdc),
		GetCmdDeleteGame(cdc),
	)...)

	return tttgameTxCmd
}
