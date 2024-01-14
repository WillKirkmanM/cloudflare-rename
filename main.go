package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"context"

	"github.com/cloudflare/cloudflare-go"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your API key: ")
	apiKey, _ := reader.ReadString('\n')

	apiKey = strings.TrimSpace(apiKey)
	
	api, err := cloudflare.NewWithAPIToken(apiKey)
	if err != nil {
		fmt.Println("Error creating Cloudflare client:", err)
		return
	}
	
	ctx := context.Background()
	
	zones, err := api.ListZones(ctx)
	if err != nil {
		fmt.Println("Error fetching zones:", err)
		return
	}

	fmt.Println("Here are your domains:")
		for i, zone := range zones {
    fmt.Printf("%d. %s\n", i+1, zone.Name)
	}

	fmt.Print("Enter the number of the domain you want to update: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	index, _ := strconv.Atoi(choice)
	selectedZone := zones[index-1]

	fmt.Print("Enter the desired IP Address: ")
	ip, _ := reader.ReadString('\n')
	ip = strings.TrimSpace(ip)

	records, _, err := api.ListDNSRecords(ctx, cloudflare.ZoneIdentifier(selectedZone.ID), cloudflare.ListDNSRecordsParams{})
	if err != nil {
		fmt.Println("Error fetching DNS records:", err)
		return
	}

	for _, record := range records {
		if record.Type != "A" {
			fmt.Printf("DNS record %s is not an A record. Skipping...\n", record.Name)
			continue
		}

		if record.Content == ip {
			fmt.Printf("DNS record %s already has the IP %s. Skipping...\n", record.Name, ip)
			continue
    }

		_, err := api.UpdateDNSRecord(ctx, cloudflare.ZoneIdentifier(selectedZone.ID), cloudflare.UpdateDNSRecordParams{ID: record.ID, Type: record.Type, Name: record.Name, Content: ip})
		if err != nil {
				fmt.Println("Error updating DNS record:", err)
				return
		}

		fmt.Printf("DNS record %s updated successfully to %s.\n", record.Name, ip)
	}

	fmt.Println("All DNS records updated successfully")
}