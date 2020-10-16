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

func GetCmdInviteGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "open-game [playerID] [gameID]",
		Short: "open a new game, waiting for another player",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsPlayer := string(args[0])
			argsGame := string(args[1])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgInviteGame(cliCtx.GetFromAddress(), argsPlayer, argsGame)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdAcceptGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "accept-game [playerID] [gameID]",
		Short: "accept invitation, start game",
		Args:  cobra.ExtractArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsPlayer := string(args[0])
			argsGame := string(args[1])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgAcceptGame(cliCtx.GetFromAddress(), argsPlayer, argsGame)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdPlayGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "play-game [playerID] [coordinate-X] [coordinate-Y]",
		Short: "play a piece to board",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsPlayer := string(args[0])
			argsX := int(args[1])
			argsY := int(args[2])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgPlayGame(cliCtx.GetFromAddress(), argsPlayer, argsX, argsY)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
