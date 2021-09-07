package main

import (
	"encoding/json"
	"os"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type UserDetails struct {
	Captcha string `bson:"captcha,omitempty" json:"captcha,omitempty"`
	Name string `bson:"name,omitempty" json:"name,omitempty"`
}

type GoogleResponse struct {
	Success bool `bson:"success,omitempty" json:"success,omitempty"`
	ErrorCodes []string `bson:"error-codes,omitempty" json:"error-codes,omitempty"`
	Hostname string `bson:"hostname,omitempty" json:"hostname,omitempty"`
}

func main(){
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	server.POST("/ping", func(c *gin.Context){
		secret := os.Getenv("SECRET_KEY")

		var captchaData UserDetails
		var googleData GoogleResponse

		//Get the POST request data
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(jsonData,&captchaData)
		if err != nil {
			fmt.Println(err)
		}

		// Validate the token with Google Recaptcha API
		resp, err := http.Get("https://google.com/recaptcha/api/siteverify?secret="+secret+"&response="+captchaData.Captcha)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		googleResponse, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(googleResponse, &googleData)
		if googleData.Success == true {
			c.JSON(200, gin.H{
				"message":"Captcha successfully validated",
				"data": captchaData.Name,
			})
		} else {
			c.JSON(404, gin.H{
				"message":"Error validating captcha",
			})
		}

	})
	server.Run(":8000")
}