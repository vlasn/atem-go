package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mraerino/atem-go"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Level = logrus.InfoLevel
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	client := atem.NewClient(log, "192.168.1.156")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := client.Start(ctx)
	if err != nil {
		log.WithError(err).Fatal("Failed to start")
	}
	log.Info("Connected")

	//stateSer, _ := json.MarshalIndent(client.State(), "", "  ")
	stateSer, _ := json.MarshalIndent(client.State().SuperSource, "", "  ")
	fmt.Printf("State:\n%s\n", string(stateSer))

	// allow for some packets to come in
	<-time.After(time.Second * 30)
}
