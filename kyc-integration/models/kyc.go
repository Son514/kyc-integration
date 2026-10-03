package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://verification.didit.me"

type KYC struct {
	apiKey string
}

func NewKYC(apiKey string) *KYC {
	return &KYC{apiKey: apiKey}
}

type CreateSessionRequest struct {
	WorkflowID string `json:"workflow_id"`
	VendorData string `json:"vendor_data,omitempty"`
	Callback   string `json:"callback,omitempty"`
}

type Session struct {
	SessionID       string `json:"session_id"`
	SessionKind     string `json:"session_kind"`
	SessionNumber   int    `json:"session_number"`
	SessionToken    string `json:"session_token"`
	URL             string `json:"url"`
	VendorData      string `json:"vendor_data"`
	Status          string `json:"status"`
	WorkflowID      string `json:"workflow_id"`
	WorkflowVersion int    `json:"workflow_version"`
	Callback        string `json:"callback"`
}

func (k *KYC) CreateSession(req CreateSessionRequest) (*Session, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, baseURL+"/v3/session/", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("x-api-key", k.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("didit: %d: %s", httpResp.StatusCode, respBody)
	}

	var session Session
	if err := json.Unmarshal(respBody, &session); err != nil {
		return nil, err
	}

	return &session, nil
}
