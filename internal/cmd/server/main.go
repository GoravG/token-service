package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"time"
	"vms-go-token-service/internal/config"
	"vms-go-token-service/internal/utils"
)

func main() {
	cfg := config.FromEnv()
	host := cfg.RedisHost
	port := cfg.RedisPort
	addr := host + ":" + port
	log.Printf("Address: %s", addr)
	channel := cfg.Channel
	publishInterval := cfg.PublishInterval

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	defer conn.Close()

	log.Printf("Connected to redis at %s", addr)
	log.Printf("Using channel %s", channel)

	ticker := time.NewTicker(publishInterval)

	defer ticker.Stop()

	for {
		<-ticker.C

		token := generateToken(16)
		fmt.Printf("Generated token: %s\n", token)

		cmd := utils.EncodeCommand("PUBLISH", channel, token)
		conn.Write([]byte(cmd))
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Failed to read response: %v", err)
		} else {
			log.Printf("Redis response: %s", string(buf[:n]))
		}
	}

}

func generateToken(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
