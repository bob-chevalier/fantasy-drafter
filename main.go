package main

import (
//    "fmt"
    "log"
//    "net/http"
//    "os"
    "path"
    "path/filepath"
//    "github.com/auth0-community/go-auth0"
    "github.com/gin-gonic/gin"
 //   jose "gopkg.in/square/go-jose.v2"
    "github.com/bob.chevalier/fantasy-drafter/controller"
//    "github.com/bob.chevalier/fantasy-drafter/fantasy-sites"
    "github.com/bob.chevalier/fantasy-drafter/fantasy-sites/yahoo"
)

var CONTROLLER = "controller"
//TODO move this into middleware function?
//var controller = controllers.New(yahoo.New())

func ControllerMiddleware() gin.HandlerFunc {
    controller := controller.New(yahoo.New())
    return func(context *gin.Context) {
        context.Set(CONTROLLER, &controller)
        context.Next()
    }
}

func CorsMiddleware() gin.HandlerFunc {
    return func(context *gin.Context) {
        context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        context.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

        if context.Request.Method == "OPTIONS" {
            context.AbortWithStatus(204)
            return
        }

        context.Next()
    }
}

func main() {
    r := gin.Default()
    r.NoRoute(func(c *gin.Context) {
        dir, file := path.Split(c.Request.RequestURI)
        ext := filepath.Ext(file)
        if file == "" || ext == "" {
            c.File("./ui/dist/ui/index.html")
        } else {
            c.File("./ui/dist/ui/" + path.Join(dir, file))
        }
    })

    r.Use(ControllerMiddleware())
    r.Use(CorsMiddleware())

    r.GET("/code", func(context *gin.Context) {
        controller, ok := context.MustGet(CONTROLLER).(*controller.Controller)
        if !ok {
            log.Fatal("Unable to retrieve Controller from Context")
        }
        controller.GetAccessCode(context)
    })
    r.POST("/login", func(context *gin.Context) {
        controller, ok := context.MustGet(CONTROLLER).(*controller.Controller)
        if !ok {
            log.Fatal("Unable to retrieve Controller from Context")
        }
        controller.Login(context)
    })
    r.GET("/teams", func(context *gin.Context) {
        controller, ok := context.MustGet(CONTROLLER).(*controller.Controller)
        if !ok {
            log.Fatal("Unable to retrieve Controller from Context")
        }
        controller.GetTeams(context)
    })
    r.GET("/players", func(context *gin.Context) {
        controller, ok := context.MustGet(CONTROLLER).(*controller.Controller)
        if !ok {
            log.Fatal("Unable to retrieve Controller from Context")
        }
        controller.GetPlayers(context)
    })
    r.GET("/draftpicks", func(context *gin.Context) {
        controller, ok := context.MustGet(CONTROLLER).(*controller.Controller)
        if !ok {
            log.Fatal("Unable to retrieve Controller from Context")
        }
        controller.GetDraftPicks(context)
    })
    r.GET("/startpolling", func(context *gin.Context) {
        controller, ok := context.MustGet(CONTROLLER).(*controller.Controller)
        if !ok {
            log.Fatal("Unable to retrieve Controller from Context")
        }
        controller.StartPolling(context)
    })
    r.GET("/stoppolling", func(context *gin.Context) {
        controller, ok := context.MustGet(CONTROLLER).(*controller.Controller)
        if !ok {
            log.Fatal("Unable to retrieve Controller from Context")
        }
        controller.StopPolling(context)
    })
//    r.POST("/todo", teams.AddTodoHandler)
//    r.DELETE("/todo/:id", teams.DeleteTodoHandler)
//    r.PUT("/todo", teams.CompleteTodoHandler)

    err := r.Run(":3000")
    if err != nil {
        panic(err)
    }
}
