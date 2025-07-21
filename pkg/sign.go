package pkg

import (
	"bytes"
	"net/http"

	"tomwmth.dev/hoyo-daily/internal"
)

type SignResponse struct {
	BaseResponse
}

func (g *Game) Sign(credentials HoyoCredentials) (*SignResponse, error) {
	data, err := internal.ToJSON(&struct {
		Event string `json:"act_id"`
	}{Event: g.Event}, false)
	if err != nil {
		return nil, err
	}

	res, err := internal.RequestJSONFromURL[SignResponse]("POST", g.Endpoint+"/sign", bytes.NewBuffer(data), func(req *http.Request) {
		g.configureRequest(req, credentials)
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *SignResponse) WasAlreadySigned() bool {
	return r.Code == -5003
}
