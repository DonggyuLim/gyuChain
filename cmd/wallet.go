package cmd

import (
	"github.com/gyuChain/wallet"
	"github.com/spf13/cobra"
)

var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Make wallet",
	Long:  "Blockchain use Wallet",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		value := args[0]
		switch value {
		case "wallet":
			wallet.NewWallet()
		case "mnemonic":
			wallet.NewMnemonic()
		}

	},
}

func init() {
	rootCmd.AddCommand(walletCmd)
}
