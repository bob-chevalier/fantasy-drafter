package controller

import (
    "encoding/json"
//    "fmt"
    "io"
    "io/ioutil"
    "log"
    "sync"
    "net/http"
//    "github.com/bob.chevalier/fantasy-drafter/fantasy-sites/yahoo"
    "github.com/bob.chevalier/fantasy-drafter/fantasy-sites"
    "github.com/gin-gonic/gin"
//    "golang.org/x/oauth2"
//    "golang.org/x/oauth2/yahoo"
)

var POSITIONS = [...]string{ "QB", "RB", "WR", "TE", "DEF", "K" }

type Controller struct {
    FantasySite site.FantasySite
    teams map[string]site.Team
    players map[string]site.Player
    draftPicks []site.DraftPick
    mutex sync.RWMutex
}

func New(fantasySite site.FantasySite) Controller {
    return Controller { FantasySite: fantasySite }
}

func (controller *Controller) GetAccessCode(context *gin.Context) {
    log.Println("BFC get access code")
    if !controller.FantasySite.AuthorizeFromSavedToken(context) {
        context.Redirect(http.StatusFound, controller.FantasySite.GetAuthorizationUrl())
    }
}

func (controller *Controller) Login(context *gin.Context) {
    log.Println("BFC login")
    code, status, err := convertHttpBodyToAccessCode(context.Request.Body)
    if err != nil {
        context.JSON(status, err)
        return
    }
    result := controller.FantasySite.Login(code, context)
    context.JSON(status, gin.H{"result": result})
}

func (controller *Controller) GetTeams(context *gin.Context) {
    if controller.teams == nil {
        controller.updateTeams()
    }
    context.JSON(http.StatusOK, controller.teams)
}

func (controller *Controller) GetPlayers(context *gin.Context) {
    if controller.players == nil {
        controller.updatePlayers()
    }
    context.JSON(http.StatusOK, controller.players)
}

func (controller *Controller) GetDraftPicks(context *gin.Context) {
    if controller.draftPicks == nil {
        controller.updateDraftPicks()
    }
    context.JSON(http.StatusOK, controller.draftPicks)
}

func (controller *Controller) updateTeams() {
    teams := controller.FantasySite.GetTeams()

    controller.mutex.Lock()
    controller.teams = teams
    controller.mutex.Unlock()
}

func (controller *Controller) updatePlayers() {
//    allPlayers := make(map[string]site.Player)
//    for _, pos := range POSITIONS {
//        fmt.Printf("Fetching %s players\n", pos)
//        players := controller.FantasySite.GetPlayers(pos)
//        for k, v := range players {
//            allPlayers[k] = v
//        }
//    }
    allPlayers := controller.FantasySite.GetPlayers()

    controller.mutex.Lock()
    controller.players = allPlayers
    controller.mutex.Unlock()
}

func (controller *Controller) updateDraftPicks() {
    // we'll be updating teams so ensure we've fetched them
    if controller.teams == nil {
        controller.updateTeams()
    }

    // we'll be updating players so ensure we've fetched them
    if controller.players == nil {
        controller.updatePlayers()
    }

    currentPickIdx := len(controller.draftPicks)
    picks := controller.FantasySite.GetDraftPicks()

    controller.mutex.Lock()
    // iterate over all picks since the last processed pick
    for _, p := range picks[currentPickIdx:] {
        // update player
        player := controller.players[p.PlayerKey]
        player.IsDrafted = true
        controller.players[p.PlayerKey] = player

        // update team
        team := controller.teams[p.TeamKey]
        team.Roster = append(team.Roster, player)
        controller.teams[p.TeamKey] = team
    }
    controller.draftPicks = picks
    controller.mutex.Unlock()
}

//func AddTodoHandler(c *gin.Context) {
//    todoItem, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
//    if err != nil {
//        c.JSON(statusCode, err)
//        return
//    }
//    c.JSON(statusCode, gin.H{"id": yahoo.Add(todoItem.Message)})
//}
//
//func DeleteTodoHandler(c *gin.Context) {
//    todoID := c.Param("id")
//    if err := yahoo.Delete(todoID); err != nil {
//        c.JSON(http.StatusInternalServerError, err)
//        return
//    }
//    c.JSON(http.StatusOK, "")
//}
//
//func CompleteTodoHandler(c *gin.Context) {
//    todoItem, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
//    if err != nil {
//        c.JSON(statusCode, err)
//        return
//    }
//    if yahoo.Complete(todoItem.ID) != nil {
//        c.JSON(http.StatusInternalServerError, err)
//        return
//    }
//    c.JSON(http.StatusOK, "")
//}
//
//func convertHTTPBodyToTodo(httpBody io.ReadCloser) (yahoo.Todo, int, error) {
//    body, err := ioutil.ReadAll(httpBody)
//    if err != nil {
//        return yahoo.Todo{}, http.StatusInternalServerError, err
//    }
//    defer httpBody.Close()
//    return convertJSONBodyToTodo(body)
//}
//
//func convertJSONBodyToTodo(jsonBody []byte) (yahoo.Todo, int, error) {
//    var todoItem yahoo.Todo
//    err := json.Unmarshal(jsonBody, &todoItem)
//    if err != nil {
//        return yahoo.Todo{}, http.StatusBadRequest, err
//    }
//    return todoItem, http.StatusOK, nil
//}

func convertHttpBodyToAccessCode(httpBody io.ReadCloser) (site.AccessCode, int, error) {
    jsonBody, err := ioutil.ReadAll(httpBody)
    if err != nil {
        return site.AccessCode{}, http.StatusInternalServerError, err
    }
    defer httpBody.Close()

    var code site.AccessCode
    err = json.Unmarshal(jsonBody, &code)
    if err != nil {
        return site.AccessCode{}, http.StatusBadRequest, err
    }
    return code, http.StatusOK, nil
}
