package pkg

import "strings"

type Game struct {
	Name       string
	InternalID string
	RPCID      any
	Endpoint   string
	Event      string
	Icon       string
}

var (
	GenshinImpact = Game{
		Name:       "Genshin Impact",
		InternalID: "genshin_impact",
		Endpoint:   "https://sg-hk4e-api.hoyolab.com/event/sol",
		Event:      "e202102251931481",
		Icon:       "https://webstatic.hoyoverse.com/upload/op-public/2022/02/08/1c77d507474b5a773ef9741ff9d840f0_2195000568465664776.jpeg",
	}

	HonkaiStarRail = Game{
		Name:       "Honkai: Star Rail",
		InternalID: "honkai_star_rail",
		RPCID:      "hkrpg",
		Endpoint:   "https://sg-public-api.hoyolab.com/event/luna/os",
		Event:      "e202303301540311",
		Icon:       "https://webstatic.hoyoverse.com/upload/op-public/2022/02/08/4129763f1dbaacf5f84e6f78c1b8d355_6722088918614682994.jpeg",
	}

	ZenlessZoneZero = Game{
		Name:       "Zenless Zone Zero",
		InternalID: "zenless_zone_zero",
		RPCID:      "zzz",
		Endpoint:   "https://sg-public-api.hoyolab.com/event/luna/zzz/os",
		Event:      "e202406031448091",
		Icon:       "https://webstatic.hoyoverse.com/upload/op-public/2022/05/24/0d4888cf46a385002b1ce63d647069e6_7263531094447256949.png",
	}

	idMap = map[string]Game{
		GenshinImpact.InternalID:   GenshinImpact,
		HonkaiStarRail.InternalID:  HonkaiStarRail,
		ZenlessZoneZero.InternalID: ZenlessZoneZero,
	}
)

func ParseGames(unparsedGames string) []Game {
	var result []Game

	for entry := range strings.SplitSeq(unparsedGames, ",") {
		entry = strings.TrimSpace(entry)
		if game, found := idMap[entry]; found {
			result = append(result, game)
		}
	}

	return result
}
