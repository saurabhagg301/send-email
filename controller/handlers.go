package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/saurabh-arch/send-email/common"
	"github.com/saurabh-arch/send-email/store"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// SendEmail to send email using provided mentioned in the query parameter
// if no query param is present, default will be considered as 'Mailjet'
// Handler for POST /sendemail
func SendEmail(w http.ResponseWriter, r *http.Request) {
	var err error
	queryParams := r.URL.Query()
	// trim space (if present) and convert it to lowercase
	provider := strings.TrimSpace(strings.ToLower(queryParams.Get("provider")))
	if provider == "" {
		provider = "mailjet"
	}

	var payload map[string]interface{}

	errDecode := json.NewDecoder(r.Body).Decode(&payload)
	if errDecode != nil {
		errMsg := fmt.Sprintf("Invalid JSON. Error: %v", errDecode)
		common.WebJSONResponse(w, http.StatusBadRequest, map[string]string{"error": errMsg})
		return
	}

	mInfo, errValidatePayload := validatePayload(payload)
	if errValidatePayload != nil {
		// log error and return
		errMsg := "Invalid payload. Error: " + errValidatePayload.Error()
		log.Error(errMsg)
		common.WebJSONResponse(w, http.StatusBadRequest, map[string]string{"error": errMsg})
		return
	}

	switch provider {
	case "mailjet":
		err = SendMailViaMailjet(mInfo)
	default:
		err = errors.New("Not a valid provider")

	}

	if err != nil {
		common.WebJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// success
	common.WebJSONResponse(w, http.StatusOK, map[string]string{"success": "Mail sent"})
}

func validatePayload(payload map[string]interface{}) (*store.MessageInfo, error) {
	mInfo := store.MessageInfo{}
	var ok bool
	// check 'to' key is present in payload
	to, toExists := payload["to"]
	if !toExists {
		return nil, errors.New("'to' key missing")
	}

	// populate 'to' list
	// mInfo.To = []store.Recipient{}
	recipients := to.([]interface{})
	for _, v := range recipients {
		var receiver store.Recipient
		var mailIDExists bool
		r, okk := v.(map[string]interface{})
		if !okk {
			return nil, errors.New("Invalid recipient list")
		}

		email, mailIDExists := r["email"]
		if !mailIDExists {
			return nil, errors.New("email id missing")
		}
		receiver.Email = email.(string)

		name, nameExists := r["name"]
		if nameExists {
			receiver.Name = name.(string)
		}

		mInfo.To = append(mInfo.To, receiver)

	}
	// check 'cc' key is present in payload
	cc, ccExists := payload["cc"]
	if ccExists {
		// populate 'cc' list
		ccList := cc.([]interface{})
		for _, v := range ccList {
			var receiver store.Recipient
			var mailIDExists bool
			r, ok := v.(map[string]interface{})
			if !ok {
				return nil, errors.New("Invalid recipient CC list")
			}

			email, mailIDExists := r["email"]
			if !mailIDExists {
				return nil, errors.New("email id missing")
			}
			receiver.Email = email.(string)

			name, nameExists := r["name"]
			if nameExists {
				receiver.Name = name.(string)
			}

			mInfo.CC = append(mInfo.CC, receiver)
		}
	}

	// check 'bcc' key is present in payload
	bcc, bccExists := payload["bcc"]
	if bccExists {
		// populate 'bcc' list
		bccList := bcc.([]interface{})
		for _, v := range bccList {
			var receiver store.Recipient
			var mailIDExists bool
			r, ok := v.(map[string]interface{})
			if !ok {
				return nil, errors.New("Invalid recipient BCC list")
			}
			email, mailIDExists := r["email"]
			if !mailIDExists {
				return nil, errors.New("email id missing")
			}
			receiver.Email = email.(string)

			name, nameExists := r["name"]
			if nameExists {
				receiver.Name = name.(string)
			}

			mInfo.BCC = append(mInfo.BCC, receiver)
		}
	}

	// check subject key present
	subject, subExists := payload["subject"]
	if !subExists {
		return nil, errors.New("'subject' key missing")
	}
	mInfo.Subject = subject.(string)

	// check htmlPart key present
	htmlPart, htmlPartExists := payload["htmlPart"]
	if htmlPartExists {
		mInfo.HTMLPart, ok = htmlPart.(string)
		if !ok {
			return nil, errors.New("'htmlPart' should be a string")
		}

	}

	// check textPart key present
	textPart, textPartExists := payload["textPart"]
	if textPartExists {
		mInfo.TextPart, ok = textPart.(string)
		if !ok {
			return nil, errors.New("'textPart' should be a string")
		}
	}

	// success
	return &mInfo, nil
}
