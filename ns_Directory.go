// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Directory/ at 2024-05-23T17:43:11.627675
package contractor

import (
	"context"
	"fmt"
	cinp "github.com/cinp/go"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// DirectoryZone - Model Zone(/api/v1/Directory/Zone)
/*
Zone(id, name, parent, ttl, refresh, retry, expire, minimum, updated, created)
*/
type DirectoryZone struct {
	cinp.BaseObject
	cinp    cinp.CInPClient `json:"-"`
	Name    *string         `json:"name,omitempty"`
	Parent  *string         `json:"parent,omitempty"`
	TTL     *int            `json:"ttl,omitempty"`
	Refresh *int            `json:"refresh,omitempty"`
	Retry   *int            `json:"retry,omitempty"`
	Expire  *int            `json:"expire,omitempty"`
	Minimum *int            `json:"minimum,omitempty"`
	Updated *time.Time      `json:"updated,omitempty"`
	Created *time.Time      `json:"created,omitempty"`
	ID      *int            `json:"id,omitempty"`
	Fqdn    *string         `json:"fqdn,omitempty"`
}

// DirectoryZoneNew - Make a new object of Model Zone
func (service *Contractor) DirectoryZoneNew() *DirectoryZone {
	return &DirectoryZone{cinp: service.cinp}
}

// DirectoryZoneNewWithID - Make a new object of Model Zone
func (service *Contractor) DirectoryZoneNewWithID(id int) *DirectoryZone {
	result := DirectoryZone{cinp: service.cinp}
	result.SetURI("/api/v1/Directory/Zone:" + strconv.FormatInt(int64(id), 10) + ":")
	return &result
}

// DirectoryZoneGet - Get function for Model Zone
func (service *Contractor) DirectoryZoneGet(ctx context.Context, id int) (*DirectoryZone, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Directory/Zone:"+strconv.FormatInt(int64(id), 10)+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*DirectoryZone)
	result.cinp = service.cinp

	return result, nil
}

// DirectoryZoneGetURI - Get function for Model Zone vi URI
func (service *Contractor) DirectoryZoneGetURI(ctx context.Context, uri string) (*DirectoryZone, error) {
	if !strings.HasPrefix(uri, "/api/v1/Directory/Zone:") {
		return nil, fmt.Errorf("URI is not for a 'DirectoryZone'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*DirectoryZone)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Zone
func (object *DirectoryZone) Create(ctx context.Context) (*DirectoryZone, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Directory/Zone", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*DirectoryZone), nil
}

// Update - Update function for Model Zone
func (object *DirectoryZone) Update(ctx context.Context) (*DirectoryZone, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*DirectoryZone), nil
}

// Delete - Delete function for Model Zone
func (object *DirectoryZone) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// DirectoryZoneListFilters - Return a slice of valid filter names Zone
func (service *Contractor) DirectoryZoneListFilters() [0]string {
	return [0]string{}
}

// DirectoryZoneList - List function for Model Zone
func (service *Contractor) DirectoryZoneList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *DirectoryZone, error) {
	if filterName != "" {
		for _, item := range service.DirectoryZoneListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Directory/Zone", reflect.TypeOf(DirectoryZone{}), filterName, filterValues, 50)
	out := make(chan *DirectoryZone)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*DirectoryZone).cinp = service.cinp
			out <- (*v).(*DirectoryZone)
		}
	}()
	return out, nil
}

// DirectoryEntry - Model Entry(/api/v1/Directory/Entry)
/*
Entry(id, zone, type, name, priority, weight, port, target, updated, created)
*/
type DirectoryEntry struct {
	cinp.BaseObject
	cinp     cinp.CInPClient `json:"-"`
	Zone     *string         `json:"zone,omitempty"`
	Type     *string         `json:"type,omitempty"`
	Name     *string         `json:"name,omitempty"`
	Priority *int            `json:"priority,omitempty"`
	Weight   *int            `json:"weight,omitempty"`
	Port     *int            `json:"port,omitempty"`
	Target   *string         `json:"target,omitempty"`
	Updated  *time.Time      `json:"updated,omitempty"`
	Created  *time.Time      `json:"created,omitempty"`
	ID       *int            `json:"id,omitempty"`
}

// DirectoryEntryNew - Make a new object of Model Entry
func (service *Contractor) DirectoryEntryNew() *DirectoryEntry {
	return &DirectoryEntry{cinp: service.cinp}
}

// DirectoryEntryNewWithID - Make a new object of Model Entry
func (service *Contractor) DirectoryEntryNewWithID(id int) *DirectoryEntry {
	result := DirectoryEntry{cinp: service.cinp}
	result.SetURI("/api/v1/Directory/Entry:" + strconv.FormatInt(int64(id), 10) + ":")
	return &result
}

// DirectoryEntryGet - Get function for Model Entry
func (service *Contractor) DirectoryEntryGet(ctx context.Context, id int) (*DirectoryEntry, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Directory/Entry:"+strconv.FormatInt(int64(id), 10)+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*DirectoryEntry)
	result.cinp = service.cinp

	return result, nil
}

// DirectoryEntryGetURI - Get function for Model Entry vi URI
func (service *Contractor) DirectoryEntryGetURI(ctx context.Context, uri string) (*DirectoryEntry, error) {
	if !strings.HasPrefix(uri, "/api/v1/Directory/Entry:") {
		return nil, fmt.Errorf("URI is not for a 'DirectoryEntry'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*DirectoryEntry)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Entry
func (object *DirectoryEntry) Create(ctx context.Context) (*DirectoryEntry, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Directory/Entry", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*DirectoryEntry), nil
}

// Update - Update function for Model Entry
func (object *DirectoryEntry) Update(ctx context.Context) (*DirectoryEntry, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*DirectoryEntry), nil
}

// Delete - Delete function for Model Entry
func (object *DirectoryEntry) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// DirectoryEntryListFilters - Return a slice of valid filter names Entry
func (service *Contractor) DirectoryEntryListFilters() [1]string {
	return [1]string{"zone"}
}

// DirectoryEntryList - List function for Model Entry
func (service *Contractor) DirectoryEntryList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *DirectoryEntry, error) {
	if filterName != "" {
		for _, item := range service.DirectoryEntryListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Directory/Entry", reflect.TypeOf(DirectoryEntry{}), filterName, filterValues, 50)
	out := make(chan *DirectoryEntry)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*DirectoryEntry).cinp = service.cinp
			out <- (*v).(*DirectoryEntry)
		}
	}()
	return out, nil
}

func registerDirectory(cinp cinp.CInPClient) {
	cinp.RegisterType("/api/v1/Directory/Zone", reflect.TypeOf((*DirectoryZone)(nil)).Elem())
	cinp.RegisterType("/api/v1/Directory/Entry", reflect.TypeOf((*DirectoryEntry)(nil)).Elem())
}
