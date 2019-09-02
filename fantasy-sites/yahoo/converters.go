package yahoo

import (
    "strconv"
    "sort"
    "net/http"
    "github.com/bob.chevalier/fantasy-drafter/fantasy-sites"
)

func responseToTeams(response *http.Response) map[string]site.Team {
    var result = make(map[string]site.Team)

    var json interface{}
    unmarshalResponse(response, &json)

    body := json.(map[string]interface{})
    fantasyContent := body["fantasy_content"].(map[string]interface{})
    league := fantasyContent["league"].([]interface{})
    leagueMembers := league[1].(map[string]interface{})
    teams := leagueMembers["teams"].(map[string]interface{})
    for _, j := range teams {
        if teamEntry, ok := j.(map[string]interface{}); ok {
            team := site.Team{}

            outer := teamEntry["team"].([]interface{})
            inner := outer[0].([]interface{})
            for _, k := range inner {
                if info, ok := k.(map[string]interface{}); ok {
                    if val, ok := info["team_key"]; ok {
                        team.Key = val.(string)
                    } else if val, ok := info["name"]; ok {
                        team.Name = val.(string)
                    } else if val, ok := info["draft_position"]; ok {
                        team.DraftPosition = int(val.(float64))
                    } else if val, ok := info["managers"]; ok {
                        mgrList := val.([]interface{})
                        mgrEntry := mgrList[0].(map[string]interface{})
                        mgrInfo := mgrEntry["manager"].(map[string]interface{})
                        mgrName := mgrInfo["nickname"]
                        team.Manager = mgrName.(string)
                    }
                }
            }
            result[team.Key] = team
        }
    }

    return result
}

func responseToPlayers(response *http.Response, startingRank int) map[string]site.Player {
    var result = make(map[string]site.Player)

    var json interface{}
    unmarshalResponse(response, &json)

    body := json.(map[string]interface{})
    fantasyContent := body["fantasy_content"].(map[string]interface{})
    league := fantasyContent["league"].([]interface{})
    leaguePlayers := league[1].(map[string]interface{})
    players := leaguePlayers["players"].(map[string]interface{})
    for i, j := range players {
        if playerEntry, ok := j.(map[string]interface{}); ok {
            batchRank, _ := strconv.Atoi(i)
            player := site.Player{ DraftRank: startingRank + batchRank }

            outer := playerEntry["player"].([]interface{})
            inner := outer[0].([]interface{})
            for _, k := range inner {
                if info, ok := k.(map[string]interface{}); ok {
                    if val, ok := info["player_key"]; ok {
                        player.Key = val.(string)
                    } else if val, ok := info["name"]; ok {
                        nameInfo := val.(map[string]interface{})
                        player.Name = nameInfo["full"].(string)
                    } else if val, ok := info["primary_position"]; ok {
                        player.Position = val.(string)
                    } else if val, ok := info["bye_weeks"]; ok {
                        byeInfo := val.(map[string]interface{})
                        player.ByeWeek = byeInfo["week"].(string)
                    }
                }
            }
            result[player.Key] = player
        }
    }

    return result
}

func responseToDraftPicks(response *http.Response) []site.DraftPick {
    var result []site.DraftPick

    var json interface{}
    unmarshalResponse(response, &json)

    body := json.(map[string]interface{})
    fantasyContent := body["fantasy_content"].(map[string]interface{})
    league := fantasyContent["league"].([]interface{})
    leaguePicks := league[1].(map[string]interface{})
    if picks, ok := leaguePicks["draft_results"].(map[string]interface{}); ok {
        for _, j := range picks {
            if draftEntry, ok := j.(map[string]interface{}); ok {
                pick := site.DraftPick{}

                info := draftEntry["draft_result"].(map[string]interface{})
                if val, ok := info["pick"]; ok {
                    pick.Pick = int(val.(float64))
                } else if val, ok := info["round"]; ok {
                    pick.Round = int(val.(float64))
                } else if val, ok := info["team_key"]; ok {
                    pick.TeamKey = val.(string)
                } else if val, ok := info["player_key"]; ok {
                    pick.PlayerKey = val.(string)
                }
                result = insertSortedDraftPick(result, pick)
            }
        }
    }

    return result
}

func insertSortedDraftPick(picks []site.DraftPick, item site.DraftPick) []site.DraftPick {
    idx := sort.Search(len(picks), func(i int) bool { return picks[i].Pick > item.Pick })
    picks = append(picks, site.DraftPick{})
    copy(picks[idx+1:], picks[idx:])
    picks[idx] = item
    return picks
}
