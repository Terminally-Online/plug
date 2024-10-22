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
            JSON []struct {
                Success bool `json:"success"`
                Token   struct {
                    Address string `json:"address"`
                    Chain   string `json:"chain"`
                    TokenId string `json:"tokenId"`
                } `json:"token"`
            } `json:"json"`
        } `json:"data"`
    } `json:"result"`
}

func CollectibleMetadata() {
    url := fmt.Sprintf("%s%s", os.Getenv("PLUG_APP_API_URL"), "jobs.collectibleMetadata")
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
