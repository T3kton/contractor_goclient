// Package contractor - Automatically generated by cinp-codegen from http://127.0.0.1:8888/api/v1/ at 2024-05-24T15:42:51.144790
package contractor

import (
	"context"
	"fmt"
	"log/slog"

	cinp "github.com/cinp/go"
)

// Contractor from http://127.0.0.1:8888/api/v1/
type Contractor struct {
	cinp cinp.CInPClient
}

// NewContractorInt creates and returns a new Contractor
func NewContractorInt(ctx context.Context, log *slog.Logger, host string, proxy string) (*Contractor, error) {
	var err error
	s := Contractor{}
	s.cinp, err = cinp.NewCInP(log, host, "/api/v1/", proxy)
	if err != nil {
		return nil, err
	}

	registerAuth(s.cinp)
	registerBluePrint(s.cinp)
	registerSite(s.cinp)
	registerSurvey(s.cinp)
	registerDirectory(s.cinp)
	registerUtilities(s.cinp)
	registerBuilding(s.cinp)
	registerForeman(s.cinp)
	registerSubContractor(s.cinp)
	registerPostOffice(s.cinp)
	registerRecords(s.cinp)
	registerVirtualBox(s.cinp)
	registerVCenter(s.cinp)
	registerDocker(s.cinp)
	registerIPMI(s.cinp)
	registerManual(s.cinp)
	registerAzure(s.cinp)
	registerAMT(s.cinp)
	registerTest(s.cinp)
	registerAWS(s.cinp)
	registerRedFish(s.cinp)

	APIVersion, err := s.GetAPIVersion(ctx, "/api/v1/")
	if err != nil {
		return nil, err
	}

	if APIVersion != "1.0" {
		return nil, fmt.Errorf("API version mismatch.  Got '%s', expected '1.0'", APIVersion)
	}

	return &s, nil
}

// GetAPIVersion Get the API version number for the Namespace at the URI
func (s *Contractor) GetAPIVersion(ctx context.Context, uri string) (string, error) {
	r, t, err := s.cinp.Describe(ctx, uri)
	if err != nil {
		return "", err
	}

	if t != "Namespace" {
		return "", fmt.Errorf("excpected a Namespace got '%s'", t)
	}

	return r.APIVersion, nil
}

// SetHeader sets a request header
func (s *Contractor) SetHeader(name string, value string) {
	s.cinp.SetHeader(name, value)
}

// ClearHeader clears a request header
func (s *Contractor) ClearHeader(name string) {
	s.cinp.ClearHeader(name)
}

// OverrideCINPClient is for use during testing to replace the cinp clinet with a mocked out client
func (s *Contractor) OverrideCINPClient(cinp cinp.CInPClient) {
	s.cinp = cinp
}
