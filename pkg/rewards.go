package pkg

import (
	"fmt"
	"net/http"

	"tomwmth.dev/hoyo-daily/internal"
)

type RewardsResponse struct {
	BaseResponse
	Data struct {
		Month   int `json:"month"`
		Rewards []struct {
			Name  string `json:"name"`
			Icon  string `json:"icon"`
			Count int    `json:"cnt"`
		} `json:"awards"`
	} `json:"data"`
}

func (g *Game) Rewards(credentials HoyoCredentials) (*RewardsResponse, error) {
	res, err := internal.RequestJSONFromURL[RewardsResponse]("GET", fmt.Sprintf("%s/home?act_id=%s", g.Endpoint, g.Event), nil, func(req *http.Request) {
		g.configureRequest(req, credentials)
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
