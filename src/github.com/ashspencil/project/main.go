package main

import (
        "fmt"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "github.com/gin-gonic/gin"
        "net/http"
)

type employee struct {
        Account string
        Passwd  string
}

func main(){
        r := gin.Default()
        r.GET("/login", func(c *gin.Context){
                // html here
        })

        r.POST("/login_post", func(c *gin.Context){
                message := c.PostForm("message")
                employee_name := c.DefaultPostForm("em_name", "anonymous")
                employee_passwd := c.PostForm("em_passwd")

                /*connect to db*/
                session, err := mgo.Dial("")
                if err != nil{
                        panic(err)
                }
                defer session.Close()

                connect := session.DB("employee").C("data")

                result := employee{}
                err = connect.Find(bson.M{"account": employee_name}).One(&result)

                if err != nil{
                        panic(err)
                }

                //fmt.Println("Passwd", result.Passwd)

                c.JSON(http.StatusOK, gin.H{
                        "status":  gin.H{
                                "status_code": http.StatusOK,
                                "status":      "ok",
                        },
                        "message": message,
                        "em_name": employee_name,
                        "password": result.Passwd,
                })
                if employee_passwd == result.Passwd{
                        fmt.Println("got")
                }else{
                        fmt.Println("you are denied to access")
                }
        })
        r.Run()
}
