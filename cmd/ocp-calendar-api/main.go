package main

import (
	"fmt"
	"github.com/ozoncp/ocp-calendar-api/internal/entities"
	"github.com/ozoncp/ocp-calendar-api/internal/utils"
	"log"
)

func main() {
	if err := utils.ReadConfig("./common.cfg"); err != nil {
		log.Printf("config file can't be open: %v", err)
	}
	fmt.Println((entities.Calendar{
		Id:     1,
		UserId: 1,
		Type:   1,
		Link:   "https://some.io/event",
	}).ToJson())
}
