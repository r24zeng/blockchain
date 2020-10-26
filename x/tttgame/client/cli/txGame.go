package cli

import (
	"bufio"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	crkeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/r24zeng/tttgame/x/tttgame/types"
)

func GetCmdInviteGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "open-game [gameID]",
		Short: "open a new game, waiting for another player",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string, key crkeys.Keybase) error {
			argsGame := string(args[0])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// get key Info from key databse
			fromName := cliCtx.GetFromName()
			keyInfo, err := key.Get(fromName)
			if err != nil {
				keyInfo, _, _ := key.CreateMnemonic(fromName, crkeys.English, "123", crkeys.Secp256k1)
			}
			// build message with gameID and key Address
			msg := types.NewMsgInviteGame(keyInfo.GetAddress(), argsGame)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			// sign message by tx builder
			txBytes, _ := txBldr.BuildAndSign(fromName, "123", []sdk.Msg{msg})

			// broadcast to a Tendermint node
			res, err := cliCtx.BroadcastTx(txBytes)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(res)
		},
	}
}

func GetCmdAcceptGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "accept-game [gameID]",
		Short: "accept invitation, start game",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, key crkeys.Keybase, args []string) error {
			argsGame := string(args[0])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// get key Info from key databse
			fromName := cliCtx.GetFromName()
			keyInfo, err := key.Get(fromName)
			if err != nil {
				keyInfo, _, _ := key.CreateMnemonic(fromName, crkeys.English, "123", crkeys.Secp256k1)
			}
			// build message with gameID and key Address
			msg := types.NewMsgAcceptGame(keyInfo.GetAddress(), argsGame)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			// sign message by tx builder
			txBytes, _ := txBldr.Sign(fromName, "123", []sdk.Msg{msg})

			// broadcast to a Tendermint node
			res, err := cliCtx.BroadcastTx(txBytes)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(res)
		},
	}
}

func GetCmdPlayGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "play-game [coordinate-X] [coordinate-Y]",
		Short: "play a piece to board",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, key crkeys.Keybase, args []string) error {
			argsGame := string(args[0])
			argsX, _ := strconv.Atoi(args[1])
			argsY, _ := strconv.Atoi(args[2])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// get key Info from key databse
			fromName := cliCtx.GetFromName()
			keyInfo, err := key.Get(fromName)
			if err != nil {
				keyInfo, _, _ := key.CreateMnemonic(fromName, crkeys.English, "123", crkeys.Secp256k1)
			}
			// build message with gameID and key Address
			msg := types.NewMsgPlayGame(keyInfo.GetAddress(), argsGame, argsX, argsY)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			// sign message by tx builder
			txBytes, _ := txBldr.Sign(fromName, "123", []sdk.Msg{msg})

			// broadcast to a Tendermint node
			res, err := cliCtx.BroadcastTx(txBytes)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(res)
		},
	}
}
