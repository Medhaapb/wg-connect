package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func main() {
	client, err := wgctrl.New()
	if err != nil {
		log.Fatalf("Failed to create WireGuard client: %v", err)
	}
	defer client.Close()

	// Generate server keys
	privateKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}
	publicKey := privateKey.PublicKey()

	fmt.Println("Server Public Key:", publicKey)

	//  server configuration
	config := wgtypes.Config{
		PrivateKey:   &privateKey,
		ListenPort:   new(int),
		ReplacePeers: true,
		Peers: []wgtypes.PeerConfig{
			{
				PublicKey:  publicKey,
				AllowedIPs: []net.IPNet{{IP: net.ParseIP("10.200.200.2"), Mask: net.CIDRMask(24, 32)}},
			},
		},
	}

	// Apply configuration
	err = client.ConfigureDevice("wg0", config)
	if err != nil {
		log.Fatalf("Failed to configure WireGuard server: %v", err)
	}

	fmt.Println("WireGuard Server started on wg0 (10.200.200.1)")

	// Start UDP server for receiving messages
	go startUDPServer("10.200.200.1:51820")

	// Send messages to the client
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Server: Enter message to send (or type 'exit' to quit): ")
		message, _ := reader.ReadString('\n')
		if message == "exit\n" {
			break
		}
		sendUDPMessage("10.200.200.2:51821", "Server: "+message)
	}
}

func startUDPServer(address string) {
	conn, err := net.ListenPacket("udp", address)
	if err != nil {
		log.Fatalf("Failed to start UDP server: %v", err)
	}
	

	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Printf("Read error: %v", err)
			continue
		}
		fmt.Printf("Received from %s: %s", addr, string(buffer[:n]))
	}
}

func sendUDPMessage(address, message string) {
	conn, err := net.Dial("udp", address)
	if err != nil {
		log.Fatalf("Failed to send UDP message: %v", err)
	}

	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("Failed to write to UDP socket: %v", err)
	}
}
