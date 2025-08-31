package pkg

import (
	"fmt"
	"strconv"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
)

var client webhook.Client

func InitWebhook(webhookURL string) error {
	_client, err := webhook.NewWithURL(webhookURL)
	if err != nil {
		return err
	}
	client = _client
	return nil
}

func SendWebhookMessage(game Game, signResponse *SignResponse, infoResponse *InfoResponse, rewardsResponse *RewardsResponse) error {
	if client == nil || signResponse == nil {
		return nil
	}

	eb := discord.NewEmbedBuilder()
	eb.SetTitle("📆 Daily Check-In")

	if signResponse.WasSuccess() {
		eb.SetDescription("Successfully completed today's check-in!")
		eb.SetColor(0x26E65C)
	} else if signResponse.WasAlreadySigned() {
		eb.SetDescription("Already completed today's check-in.")
		eb.SetColor(0x878787)
	} else {
		eb.SetDescription("Failed to complete today's check-in.")
		eb.SetColor(0xF42C2C)
	}

	if infoResponse != nil && infoResponse.WasSuccess() {
		daysCompleted := infoResponse.Data.SignedDays
		daysMissed := infoResponse.Data.CurrentDate.ToTime().Day() - daysCompleted

		eb.AddField("Days Completed", strconv.Itoa(daysCompleted), true)
		eb.AddField("Days Missed", strconv.Itoa(daysMissed), true)

		if rewardsResponse != nil && rewardsResponse.WasSuccess() {
			rewardToday := rewardsResponse.Data.Rewards[daysCompleted-1]

			eb.SetThumbnail(rewardToday.Icon)

			eb.AddField("", "", true)
			eb.AddField("Reward Today", fmt.Sprintf("%s x%d", rewardToday.Name, rewardToday.Count), true)

			if len(rewardsResponse.Data.Rewards) > daysCompleted {
				rewardTomorrow := rewardsResponse.Data.Rewards[daysCompleted]

				eb.AddField("Reward Tomorrow", fmt.Sprintf("%s x%d", rewardTomorrow.Name, rewardTomorrow.Count), true)
			}
		}
	}

	eb.SetFooterText(game.Name)
	eb.SetFooterIcon(game.Icon)
	eb.SetTimestamp(time.Now())

	_, err := client.CreateMessage(discord.NewWebhookMessageCreateBuilder().AddEmbeds(eb.Build()).Build())
	if err != nil {
		return err
	}

	return nil
}
