package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vimvek/sendm/mailer"
)

func createSendCommand() *cobra.Command {
	var to, from, message string
	currentCmd := &cobra.Command{
		Use:   "send",
		Short: "send mail",
		Long: `Send user defined mail
		Example:
		sendm send --to "hell@example.com" --from "bye@example.com" --message "Hello world"
		`,
		Run: func(cmd *cobra.Command, args []string) {
			mailer.SendMail(to, from, message)
		},
	}

	currentCmd.Flags().StringVarP(&to, "to", "t", "", "send mail to address")
	currentCmd.Flags().StringVarP(&from, "from", "f", "", "send mail to address")
	currentCmd.Flags().StringVarP(&message, "message", "m", "", "send mail to address")

	//Marking required flags
	currentCmd.MarkFlagRequired("to")
	currentCmd.MarkFlagRequired("from")
	currentCmd.MarkFlagRequired("message")

	return currentCmd
}
