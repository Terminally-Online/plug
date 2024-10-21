package cron

import (
	"fmt"
	"log"
	"os"
	"solver/utils"
)

type CollectibleMetadataResponse struct { 
    Result struct {
        Data struct {
            JSON struct {
                Count int `json:"count"`
            } `json:"json"`
        } `json:"data"`
    } `json:"result"`
}

func CollectibleMetadata() {
    url := fmt.Sprintf("%s%s", os.Getenv("PLUG_APP_API_URL"), "jobs.collectibleMetadata")
    log.Println(url)
    _, err := utils.MakeHTTPRequest(
        url,
        "POST",
        map[string]string{
            "Content-Type": "application/json",
            "X-API-Key": os.Getenv("PLUG_APP_API_KEY"),
        },
        nil,
        nil,
        CollectibleMetadataResponse{},
    )
    if err != nil {
        log.Println(err.Error())
        return
    }
}
