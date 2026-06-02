package contractor

import (
	"context"
	"fmt"
	"log/slog"
)

// NamespaceVersion holds the path and API version for one namespace
type NamespaceVersion struct {
	Path    string
	Version string
}

// GetVersionInfo describes the root namespace and each of its sub-namespaces,
// returning the per-namespace API versions.
func (s *Contractor) GetVersionInfo(ctx context.Context) ([]NamespaceVersion, error) {
	root, t, err := s.cinp.Describe(ctx, "/api/v1/")
	if err != nil {
		return nil, err
	}
	if t != "Namespace" {
		return nil, fmt.Errorf("expected a Namespace got '%s'", t)
	}

	namespaces := []NamespaceVersion{{Path: "/api/v1/", Version: root.APIVersion}}

	for _, nsPath := range root.Namespaces {
		r, t, err := s.cinp.Describe(ctx, nsPath)
		if err != nil {
			return nil, fmt.Errorf("describing '%s': %w", nsPath, err)
		}
		if t != "Namespace" {
			continue
		}
		namespaces = append(namespaces, NamespaceVersion{Path: nsPath, Version: r.APIVersion})
	}

	return namespaces, nil
}

// NewContractor creates a new Contractor and logs it in with the username and password
func NewContractor(ctx context.Context, log *slog.Logger, host string, proxy string, username string, password string) (*Contractor, error) {
	c, err := NewContractorInt(ctx, log, host, proxy)
	if err != nil {
		return nil, err
	}

	err = c.Login(ctx, username, password)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Login calles login and sets the auth headers
func (c *Contractor) Login(ctx context.Context, username string, password string) error {
	token, err := c.AuthUserCallLogin(ctx, username, password)
	if err != nil {
		return err
	}

	c.SetHeader("Auth-Id", username)
	c.SetHeader("Auth-Token", token)

	return nil
}

// Logout calles logout and removes the auth headers
func (c *Contractor) Logout(ctx context.Context) {
	c.AuthUserCallLogout(ctx)

	c.ClearHeader("Auth-Id")
	c.ClearHeader("Auth-Token")
}
