package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Keel Mailer started")

	// Channels
	mailChannel := make(chan Email)
	approvalChannel := make(chan keelApproval)

	// Setup goroutine to listen for new approvals, and send e-mail when something happens.
	approvals := ApprovalListener{approvalChannel, mailChannel, make(map[string]keelApproval)}
	go approvals.listener()

	// Setup periodic check for new approvals
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := approvals.checkForNewApprovals(); err != nil {
					log.Printf("Failed to check for new approvals. Got error: %s", err)
				}
			}
		}
	}()

	// Listen for and send emails
	emailListener(mailChannel)
}
