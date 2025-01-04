package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/0xrawsec/golang-etw/etw"
)

func main() {
	// Define command-line parameters
	provider := flag.String("provider", "", "ETW provider name or GUID")
	eventIds := flag.String("eventIds", "", "Comma-separated list of EventIds to filter")
	complete := flag.Bool("complete", false, "Output complete event information")
	outputFile := flag.String("output", "", "File to write the trace output")

	flag.Parse()

	fmt.Printf("\nPPPPPP                kk     EEEEEEE TTTTTTT WW      WW         tt           hh\n")
	fmt.Printf("PP   PP  oooo    cccc kk  kk EE        TTT   WW      WW   aa aa tt      cccc hh        eee  rr rr\n")
	fmt.Printf("PPPPPP  oo  oo cc     kkkkk  EEEEE     TTT   WW   W  WW  aa aaa tttt  cc     hhhhhh  ee   e rrr  r\n")
	fmt.Printf("PP      oo  oo cc     kk kk  EE        TTT    WW WWW WW aa  aaa tt    cc     hh   hh eeeee  rr\n")
	fmt.Printf("PP       oooo   ccccc kk  kk EEEEEEE   TTT     WW   WW   aaa aa  tttt  ccccc hh   hh  eeeee rr\n\n")

	if *provider == "" {
		log.Fatal("[!] Provider name or GUID is required, use -provider NAME_OR_GUID")
	}

	// Parse EventIds filter
	var eventIdFilter map[uint16]bool
	if *eventIds != "" {
		eventIdFilter = make(map[uint16]bool)
		for _, id := range strings.Split(*eventIds, ",") {
			var eventId uint16
			fmt.Sscanf(id, "%d", &eventId)
			eventIdFilter[eventId] = true
		}
	}

	// Creating the trace (producer part)
	s := etw.NewRealTimeSession("pockETWatcher")
	defer s.Stop()

	// Enable the provider
	if err := s.EnableProvider(etw.MustParseProvider(*provider)); err != nil {
		log.Fatalf("[!] Failed to enable provider: %v", err)
	}
	log.Printf("[i] PocketWatcher started with provider %s\n", *provider)

	// Consuming from the trace
	c := etw.NewRealTimeConsumer(context.Background())
	defer c.Stop()
	log.Printf("[i] Started consuming events ... (Ctrl-c to end)")

	c.FromSessions(s)

	// Open file if outputFile is specified
	var file *os.File
	var err error
	if *outputFile != "" {
		file, err = os.Create(*outputFile)
		if err != nil {
			log.Fatalf("[!] Failed to create output file: %v", err)
		}
		defer file.Close()
	}

	// Process events
	go func() {
		for e := range c.Events {
			// Filter events based on EventIds
			if eventIdFilter != nil {
				if _, ok := eventIdFilter[e.System.EventID]; !ok {
					continue
				}
			}

			var eventJSON []byte
			var err error

			// Convert event to full JSON if requested
			if *complete {
				eventJSON, err = json.MarshalIndent(e, "", "  ")
				if err != nil {
					log.Printf("[!] Failed to marshal event: %v", err)
					continue
				}
			} else {
				eventJSON, err = json.MarshalIndent(map[string]interface{}{
					"Provider":  e.System.Provider.Name,
					"EventID":   e.System.EventID,
					"Time":      e.System.TimeCreated.SystemTime,
					"Level":     e.System.Level.Value,
					"Opcode":    e.System.Opcode.Value,
					"Keywords":  e.System.Keywords.Value,
					"TaskName":  e.System.Task.Name,
					"EventData": e.EventData,
				}, "", "  ")
				if err != nil {
					log.Printf("[!] Failed to marshal event: %v", err)
					continue
				}
			}

			// Output event in JSON to file or console
			if file != nil {
				file.WriteString(string(eventJSON) + "\n")
				file.Sync()

			} else {
				fmt.Println(string(eventJSON))
			}
		}
	}()

	if err := c.Start(); err != nil {
		log.Fatalf("[!] Failed to start consumer: %v", err)
	}

	// Handle exit signals to clean up
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("[~] Shutting down... Bye! 👋")
}
