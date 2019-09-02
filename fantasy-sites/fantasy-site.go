package site

import (
    "github.com/gin-gonic/gin"
)

type FantasySite interface {
    AuthorizeFromSavedToken(context *gin.Context) bool
    GetAuthorizationUrl() string
    Login(code AccessCode, context *gin.Context) string
    GetTeams() map[string]Team
    GetPlayers() map[string]Player
    GetDraftPicks() []DraftPick
}

type Todo struct {
    ID       string `json:"id"`
    Message  string `json:"message"`
    Complete bool   `json:"complete"`
}

type AccessCode struct {
    Id string `json:"id"`
}

type Team struct {
    Key string `json:"team_key"`
    Name string `json:"name"`
    DraftPosition int `json:"draft_position"`
    Manager string `json:"manager"`
    Roster []Player `json:"roster"`
}

type Player struct {
    Key string `json:"key"`
    Name string `json:"name"`
    Position string `json:"position"`
    ByeWeek string `json:"bye_week"`
    DraftRank int `json:"draft_rank"`
    DraftTier int `json:"draft_tier"`
    IsDrafted bool `json:"is_drafted"`
}

type DraftPick struct {
    Pick int `json:"pick"`
    Round int `json:"round"`
    TeamKey string `json:"team_key"`
    PlayerKey string `json:"player_key"`
}
