package slack

import (
	"strings"

	"github.com/tylerconlee/slab/plugins"
	"github.com/tylerconlee/slack"
)

// parseCommand takes the message that mentions the bot user and identifies
// what the user is asking for.
func parseCommand(text string, user string) {
	var attachment slack.Attachment
	message := "..."
	t := strings.Fields(text)
	if len(t) > 1 {
		switch t[1] {
		case "set":

			attachment = SetMessage()
		case "diag":

			DiagMessage(user)
		case "whois":
			attachment = WhoIsMessage()
		case "status":
			StatusMessage()
		case "help":
			HelpMessage()
		case "unset":
			message = "Triager has been reset. Please use `@slab set` to set Triager."
			attachment = UnsetMessage()
		case "twilio":
			p := plugins.LoadPlugins(c)
			switch t[2] {
			case "set":
				s := plugins.TwilioSet(t[3])
				SendMessage("Plugin message", s)
			case "unset":
				s := plugins.TwilioUnset()
				SendMessage("Plugin message", s)
			case "configure":
				s := plugins.TwilioConfigure(t[3])
				SendMessage("Plugin message", s)
			case "status":
				s := p.TwilioStatus()
				SendMessage("Plugin status", s)
			case "enable":
				p.EnableTwilio()
				a := slack.Attachment{
					Title: "Twilio Plugin",
					Fields: []slack.AttachmentField{
						slack.AttachmentField{
							Title: "Enabled",
							Value: ":white_check_mark:",
						},
					},
				}
				SendMessage("Plugin Twilio has been updated", a)

			case "disable":
				p.DisableTwilio()
				a := slack.Attachment{
					Title: "Twilio Plugin",
					Fields: []slack.AttachmentField{
						slack.AttachmentField{
							Title: "Enabled",
							Value: ":x:",
						},
					},
				}
				SendMessage("Plugin Twilio has been updated", a)
			}

		}
		SendMessage(message, attachment)
	}

}

func parseDMCommand(text string, user string) {
	t := strings.ToLower(text)
	switch t {
	case "start config":
		StartWizard(user)
		ConfirmWizard()
	default:
		UnknownCommandMessage(text, user)
	}

}
