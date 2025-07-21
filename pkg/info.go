package pkg

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"tomwmth.dev/hoyo-daily/internal"
)

type Date time.Time

type InfoResponse struct {
	BaseResponse
	Data struct {
		SignedDays  int  `json:"total_sign_day"`
		CurrentDate Date `json:"today"`
		LastDay     bool `json:"month_last_day"`
	} `json:"data"`
}

func (g *Game) Info(credentials HoyoCredentials) (*InfoResponse, error) {
	res, err := internal.RequestJSONFromURL[InfoResponse]("GET", fmt.Sprintf("%s/info?act_id=%s", g.Endpoint, g.Event), nil, func(req *http.Request) {
		g.configureRequest(req, credentials)
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	time, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	*d = Date(time)
	return nil
}

func (d Date) ToTime() time.Time {
	return time.Time(d)
}
