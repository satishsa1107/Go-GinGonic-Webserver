package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "database/sql"
    "github.com/coopernurse/gorp"
    "github.com/gin-gonic/contrib/static"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "time"
    "strconv"
    "net/http"
)

var dbmap = initDb()

func main(){

    defer dbmap.Db.Close()

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

type Article struct {
    Id int64 `db:"article_id"`
    Created int64
    Title string
    Content string
}

func createArticle(title, body string) Article {
    article := Article{
        Created:    time.Now().UnixNano(),
        Title:      title,
        Content:    body,
    }

    err := dbmap.Insert(&article)
    checkErr(err, "Insert failed")
    return article
}

func getArticle(article_id int) Article {
    article := Article{}
    err := dbmap.SelectOne(&article, "select * from articles where article_id=?", article_id)
    checkErr(err, "SelectOne failed")
    return article
}

func ArticlesList(c *gin.Context) {
    var articles []Article
    _, err := dbmap.Select(&articles, "select * from articles order by article_id")
    checkErr(err, "Select failed")
    content := gin.H{}
    for k, v := range articles {
        content[strconv.Itoa(k)] = v
    }
    c.JSON(200, content)
}

func ArticlesDetail(c *gin.Context) {
    article_id := c.Params.ByName("id")
    a_id, _ := strconv.Atoi(article_id)
    article := getArticle(a_id)
    content := gin.H{"title": article.Title, "content": article.Content}
    c.JSON(200, content)
}

func ArticlePost(c *gin.Context) {
    var json Article

    c.Bind(&json) // This will infer what binder to use depending on the content-type header.
    article := createArticle(json.Title, json.Content)
    if article.Title == json.Title {
        content := gin.H{
            "result": "Success",
            "title": article.Title,
            "content": article.Content,
        }
        c.JSON(201, content)
    } else {
        c.JSON(500, gin.H{"result": "An error occured"})
    }
}

func initDb() *gorp.DbMap {
    db, err := sql.Open("sqlite3", "db.sqlite3")
    checkErr(err, "sql.Open failed")

    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

    dbmap.AddTableWithName(Article{}, "articles").SetKeys(true, "Id")

    err = dbmap.CreateTablesIfNotExists()
    checkErr(err, "Create tables failed")

    return dbmap
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}
