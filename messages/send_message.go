package messages

import "github.com/slack-go/slack"

func SendMsg(payload string, api *slack.Client) error {
	_, _, err := api.PostMessage("alert", slack.MsgOptionText(payload, false), slack.MsgOptionAsUser(false))

	if err != nil {
		return err
	}

	return nil
}
