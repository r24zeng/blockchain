package cli

import (
	"bufio"
    
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

func GetCmdCreateGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-game [state] [players] [board] [curTurn]",
		Short: "Creates a new game",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsState := string(args[0] )
			argsPlayers := string(args[1] )
			argsBoard := string(args[2] )
			argsCurTurn := string(args[3] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateGame(cliCtx.GetFromAddress(), string(argsState), string(argsPlayers), string(argsBoard), string(argsCurTurn))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-game [id]  [state] [players] [board] [curTurn]",
		Short: "Set a new game",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsState := string(args[1])
			argsPlayers := string(args[2])
			argsBoard := string(args[3])
			argsCurTurn := string(args[4])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetGame(cliCtx.GetFromAddress(), id, string(argsState), string(argsPlayers), string(argsBoard), string(argsCurTurn))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-game [id]",
		Short: "Delete a new game by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteGame(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
