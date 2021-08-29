package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type keelApproval struct {
	ID         string
	Archived   bool      `json:"archived"`
	Identifier string    `json:"identifier"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"createdAt"`
}

type ApprovalListener struct {
	approvalChannel chan keelApproval
	emailChannel    chan Email
	seen            map[string]keelApproval
}

func (a ApprovalListener) listener() {
	for {
		select {
		case approval := <-a.approvalChannel:
			// Check if this approval has been seen before, if no, send e-mail about it.
			if _, exists := a.seen[approval.Identifier]; !exists {
				a.seen[approval.Identifier] = approval

				log.Printf("Saw new approval %s", approval)
				subject := "Keel: Approval required for " + approval.Identifier
				body := fmt.Sprintf("Hi,\n\nThere is a new approval required in Keel.\n\n%s\n\nBest,\nKeel Mailer", approval.Message)
				a.emailChannel <- Email{subject, body}
			}
		}
	}
}

func (a ApprovalListener) checkForNewApprovals() error {
	hostname := os.Getenv("KEEL_HOST")
	path := "/v1/approvals"

	client := &http.Client{}
	req, err := http.NewRequest("GET", hostname+path, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(os.Getenv("KEEL_USER"), os.Getenv("KEEL_PASS"))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var approvals []keelApproval
	if err := json.NewDecoder(resp.Body).Decode(&approvals); err != nil {
		return err
	}

	// Send all non-archived approvals to the listener.
	for _, approval := range approvals {
		if !approval.Archived {
			a.approvalChannel <- approval
		}
	}

	return nil
}
