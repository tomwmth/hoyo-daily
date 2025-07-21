package main

import (
	"errors"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"tomwmth.dev/hoyo-daily/pkg"
)

func main() {
	godotenv.Load()

	if err := run(); err != nil {
		log.Fatalf("[Error] %v", err)
	}
}

func run() error {
	uid, found := os.LookupEnv("USER_ID")
	if !found {
		return errors.New("USER_ID environment variable was not defined")
	}

	token, found := os.LookupEnv("USER_TOKEN")
	if !found {
		return errors.New("USER_TOKEN environment variable was not defined")
	}

	unparsedGames, found := os.LookupEnv("GAMES")
	if !found {
		return errors.New("GAMES environment variable was not defined")
	}

	webhookURL, hasWebhookURL := os.LookupEnv("DISCORD_WEBHOOK_URL")

	games := pkg.ParseGames(unparsedGames)
	if len(games) == 0 {
		return errors.New("GAMES environment variable did not contain any valid game IDs")
	}

	credentials := pkg.HoyoCredentials{
		UID:   uid,
		Token: token,
	}

	if hasWebhookURL {
		if err := pkg.InitWebhook(webhookURL); err != nil {
			log.Printf("[Error] Failed to initialize webhook client: %v", err)
		}
	}

	var wg sync.WaitGroup
	for _, game := range games {
		wg.Add(1)
		go func() {
			runGame(game, credentials)
			wg.Done()
		}()
	}
	wg.Wait()

	return nil
}

func runGame(game pkg.Game, credentials pkg.HoyoCredentials) {
	signResponse, err := game.Sign(credentials)
	if err != nil {
		log.Printf("[Error] %v", err)
	}

	infoResponse, err := game.Info(credentials)
	if err != nil {
		log.Printf("[Error] %v", err)
	}

	rewardsResponse, err := game.Rewards(credentials)
	if err != nil {
		log.Printf("[Error] %v", err)
	}

	log.Printf("[%s] Completed with code %d: %s", game.Name, signResponse.Code, signResponse.Message)
	pkg.SendWebhookMessage(game, signResponse, infoResponse, rewardsResponse)
}
