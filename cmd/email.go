package cmd

import (
	"fmt"
	"github.com/resend/resend-go/v2"
	"github.com/spf13/cobra"
)

var (
	fromEmail string
	toEmail   []string
	subject   string
	texts     string
)
var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "send emails ",
	Long:  "use this command to send your email to the direction email",
	Run: func(cmd *cobra.Command, args []string) {
		api := "re_e8jpJkD6_Pu9kKeDHdA2jYM613NNjSA2M"
		client := resend.NewClient(api)
		sendrequest := &resend.SendEmailRequest{
			From:    fromEmail,
			To:      toEmail,
			Subject: subject,
			Text:    texts,
		}
		sendResponse, err := client.Emails.Send(sendrequest)
		if err != nil {
			fmt.Println("send email failed")
			return
		}
		fmt.Println("the id is : ", sendResponse.Id)

	},
}

func init() {
	emailCmd.Flags().StringVarP(&fromEmail, "from", "f", "123456789@qq.com", "show the message from which email")
	emailCmd.Flags().StringSliceVarP(&toEmail, "to", "t", []string{"123456789@qq.com"}, "show the message send to which email")
	emailCmd.Flags().StringVarP(&subject, "subject", "s", "hello world", "the subject of this email")
	emailCmd.Flags().StringVarP(&texts, "body", "b", "hello world...", "the message you want to convey")
}
