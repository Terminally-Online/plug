package cron

import (
	"fmt"
	"log"
	"os"
	"solver/utils"
)

type AnonymousUserPurgeResponse struct { 
    Result struct {
        Data struct {
            JSON struct {
                Count int `json:"count"`
            } `json:"json"`
        } `json:"data"`
    } `json:"result"`
}

func AnonymousUsers() {
    url := fmt.Sprintf("%s%s", os.Getenv("PLUG_APP_API_URL"), "jobs.anonymous")
    _, err := utils.MakeHTTPRequest(
        url,
        "POST",
        map[string]string{
            "Content-Type": "aplication/json",
            "X-API-Key": os.Getenv("PLUG_APP_API_KEY"),
        },
        nil,
        nil,
        AnonymousUserPurgeResponse{},
    )
    if err != nil {
        log.Println(err.Error())
        return
    }
}
