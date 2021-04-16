package server

import "github.com/aodpi/hrm-go-graphql/v2/config"

func Init() {
	config := config.GetConfig()
	r := SetupRouter()
	addr := config.GetString("server.address")
	r.Run(addr)
}
