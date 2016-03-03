package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/static"
    "net/http"
)


func main(){

    router := gin.Default()

    router.Use(static.Serve("/", static.LocalFile("./webfiles", false)))

    router.LoadHTMLGlob("./webfiles/*.html")

    // set up a redirect for /
    router.GET("/", func (c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/home")
    })

    router.GET("/home", func(c *gin.Context) {
        c.HTML(http.StatusOK, "home.html", nil)
    })

    router.GET("/resume", func(c *gin.Context) {
        c.HTML(http.StatusOK, "resume.html", nil)
    })

    router.GET("/projects", func(c *gin.Context) {
        c.HTML(http.StatusOK, "projects.html", nil)
    })

    router.GET("/login", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", nil)
    })

    router.POST("/login", func(c *gin.Context) {
        Username := c.PostForm("Username")
        Password := c.PostForm("Password")

        fmt.Printf("Username: %s, Password: %s is logged in",
                        Username, Password)
    })


	router.GET("/ping", func(c *gin.Context) {
        c.String(200, "test")
    })


    router.Run(":8000")
}

