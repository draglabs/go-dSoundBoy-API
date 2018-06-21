package main

import (
	"github.com/draglabs/go-dSoundBoy-API/db"
	"github.com/draglabs/go-dSoundBoy-API/routes"

	"context"
	"github.com/MindsightCo/go-mindsight-collector"
)

// https://github.com/MindsightCo/go-mindsight-collector
// https://github.com/MindsightCo/hotpath-agent
func main() {
	// setup for mindsight
	ctx := context.Background()
	collector.StartMindsightCollector(ctx,
		collector.OptionAgentURL("http://localhost:8000/samples/"),
		collector.OptionWatchPackage("github.com/draglabs/go-dSoundBoy-API/"))

	// starting the database and server
	db.EnsureIndices()
	routes.StartServer()
}
