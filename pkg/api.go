package pkg

import (
	"net/http"
)

type HoyoCredentials struct {
	UID   string
	Token string
}

type BaseResponse struct {
	Code    int    `json:"retcode"`
	Message string `json:"message"`
}

func (r *BaseResponse) WasSuccess() bool {
	return r.Code == 0
}

func (g *Game) configureRequest(req *http.Request, credentials HoyoCredentials) {
	// These headers aren't required, just trying to avoid any potential future anti-automation measures
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:140.0) Gecko/20100101 Firefox/140.0")
	req.Header.Set("Origin", "https://act.hoyolab.com")
	req.Header.Set("Referer", "https://act.hoyolab.com/")

	uidCookie := &http.Cookie{
		Name:  "ltuid_v2",
		Value: credentials.UID,
	}
	req.AddCookie(uidCookie)

	tokenCookie := &http.Cookie{
		Name:  "ltoken_v2",
		Value: credentials.Token,
	}
	req.AddCookie(tokenCookie)

	// ZZZ doesn't work without this
	// HSR has it implemented but still works without
	// GI doesn't have it implemented
	if g.RPCID != nil {
		req.Header.Set("x-rpc-signgame", g.RPCID.(string))
	}
}
