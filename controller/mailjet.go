package controller

import (
	"os"

	"github.com/saurabh-arch/send-email/config"
	"github.com/saurabh-arch/send-email/store"

	mailjet "github.com/mailjet/mailjet-apiv3-go/v3"
)

// SendMailViaMailjet ...
func SendMailViaMailjet(mInfo *store.MessageInfo) error {
	cfg := config.LoadConfig()
	MJ_KEY_PRIVATE := os.Getenv("MJ_APIKEY_PRIVATE")
	if MJ_KEY_PRIVATE == "" {
		log.Info("Private Key not found, hence using from config")
		MJ_KEY_PRIVATE = cfg.MJ_APIKEY_PRIVATE
	}

	MJ_KEY_PUBLIC := os.Getenv("MJ_APIKEY_PUBLIC")
	if MJ_KEY_PUBLIC == "" {
		log.Info("MJ Public Key not found, hence using from config")
		MJ_KEY_PUBLIC = cfg.MJ_APIKEY_PUBLIC
	}

	mailjetClient := mailjet.NewMailjetClient(MJ_KEY_PUBLIC, MJ_KEY_PRIVATE)
	var toList mailjet.RecipientsV31
	var ccList mailjet.RecipientsV31
	var bccList mailjet.RecipientsV31

	for _, v := range mInfo.To {
		var r mailjet.RecipientV31
		r.Name = v.Name
		r.Email = v.Email

		toList = append(toList, r)
	}

	for _, v := range mInfo.CC {
		var r mailjet.RecipientV31
		r.Name = v.Name
		r.Email = v.Email

		ccList = append(ccList, r)
	}

	for _, v := range mInfo.BCC {
		var r mailjet.RecipientV31
		r.Name = v.Name
		r.Email = v.Email

		bccList = append(bccList, r)
	}

	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: cfg.From,
				Name:  cfg.Name,
			},
			To:       &toList,
			Cc:       &ccList,
			Bcc:      &bccList,
			Subject:  mInfo.Subject,
			TextPart: mInfo.TextPart,
			HTMLPart: mInfo.HTMLPart,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("Data: %+v\n", res)

	// success
	return nil
}
