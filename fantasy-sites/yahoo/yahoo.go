package yahoo

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/bob.chevalier/fantasy-drafter/fantasy-sites"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/yahoo"
)

type Yahoo struct {
    config oauth2.Config
    client *http.Client
}

const (
//    GAME_ID string = "380"
    GAME_ID string = "nfl"
//    LEAGUE_ID string = "671449" // MHI 2018
    LEAGUE_ID string = "878768" // MHI 2019
    BASE_URL string = "https://fantasysports.yahooapis.com/fantasy/v2"
    LEAGUE_URL string = BASE_URL + "/league/" + GAME_ID + ".l." + LEAGUE_ID
    TEAMS_URL string = LEAGUE_URL + "/teams?format=json"
    DRAFT_PICKS_URL string = LEAGUE_URL + "/draftresults?format=json"
    TOKEN_FILENAME string = "oauth2-token.json"
    //TODO move this up to controller? or into config
    MAX_PLAYERS int = 500
)

func New() *Yahoo {
    config := oauth2.Config {
        ClientID: "dj0yJmk9SVFQS0c0VFdYZGs5JmQ9WVdrOWRIaFZNMnBMTTJNbWNHbzlNQS0tJnM9Y29uc3VtZXJzZWNyZXQmeD00Zg--",
        ClientSecret: "bcf5fe96f49a18d33d133adbda5e74378029d6fc",
        Scopes: []string{ "fspt-w" },
        RedirectURL: "oob",
        Endpoint: yahoo.Endpoint,
    }

    return &Yahoo{ config: config, client: nil }
}

func (yahoo *Yahoo) AuthorizeFromSavedToken(context *gin.Context) bool {
    if _, err := os.Stat(TOKEN_FILENAME); err == nil {
        // token file exists, so load token from file
        jsonBlob, err := ioutil.ReadFile(TOKEN_FILENAME)
        if err != nil {
            log.Panic("Unable to read token file")
        }
        var token oauth2.Token
        err = json.Unmarshal(jsonBlob, &token)
        if err != nil {
            log.Panic("Unable to unmarshal token blob")
        }

        // create a client using token
        yahoo.client = yahoo.config.Client(context, &token)
        return true
    }
    return false
}

func (yahoo *Yahoo) GetAuthorizationUrl() string {
    return yahoo.config.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func (yahoo *Yahoo) Login(code site.AccessCode, context *gin.Context) string {
    token, err := yahoo.config.Exchange(context, code.Id)
    if err != nil {
        log.Fatal(err)
    }
    yahoo.client = yahoo.config.Client(context, token)

    // save token to disk
    jsonBlob, err := json.MarshalIndent(token, "", "    ")
    if err != nil {
        log.Panic("Unable to marshal token")
    }
    err = ioutil.WriteFile(TOKEN_FILENAME, jsonBlob, 0644)
    if err != nil {
        log.Panic("Unable to write token to file")
    }

    return "success"
}

func (yahoo *Yahoo) GetTeams() map[string]site.Team {
    resp := validateResponse(yahoo.client.Get(TEAMS_URL))
    return responseToTeams(resp)
}

func (yahoo *Yahoo) GetPlayers() map[string]site.Player {
    allPlayers := make(map[string]site.Player)

    start := 0
    numRemaining := MAX_PLAYERS
    outOfPlayers := false
    for numRemaining > 0 && !outOfPlayers {
        //url := fmt.Sprintf("%s/players;position=%s;sort=OR;start=%d;count=%d?format=json", LEAGUE_URL, position, start, numRemaining)
        url := fmt.Sprintf("%s/players;sort=OR;start=%d;count=%d?format=json", LEAGUE_URL, start, numRemaining)
        resp := validateResponse(yahoo.client.Get(url))
        batch := responseToPlayers(resp, start)
        batchSize := len(batch)
        if batchSize < 1 {
            outOfPlayers = true
        } else {
            for k, v := range batch {
                allPlayers[k] = v
            }
            start += batchSize
            numRemaining -= batchSize
        }
        fmt.Printf("num players in batch: %d\n", batchSize)
    }

    return allPlayers
}

func (yahoo *Yahoo) GetDraftPicks() []site.DraftPick {
    resp := validateResponse(yahoo.client.Get(DRAFT_PICKS_URL))
    return responseToDraftPicks(resp)
}

func validateResponse(response *http.Response, err error) *http.Response {
    if err != nil {
        log.Panic(err)
    }
    if response.StatusCode != 200 {
        dump, err := httputil.DumpRequest(response.Request, true)
        if err == nil {
            fmt.Println(string(dump))
        }
        log.Panicf("Invalid response: [%s]", response.Status)
    }
    return response
}

func unmarshalResponse(response *http.Response, body interface{}) {
    buf, err := ioutil.ReadAll(response.Body)
    response.Body.Close()
    if err != nil {
        log.Panic(err)
    }

    err = json.Unmarshal(buf, body)
    if err != nil {
        log.Panic(err)
    }
}
