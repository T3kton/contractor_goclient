// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/RedFish/ at 2024-05-24T15:42:51.144790
package contractor

import (
	"context"
	"fmt"
	cinp "github.com/cinp/go"
	"reflect"
	"strings"
	"time"
)

// RedfishRedFishFoundation - Model RedFishFoundation(/api/v1/RedFish/RedFishFoundation)
/*
RedFishFoundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr, redfish_username, redfish_password, redfish_ip_address, redfish_sol_port, plot)
*/
type RedfishRedFishFoundation struct {
	cinp.BaseObject
	cinp             cinp.CInPClient `json:"-"`
	Locator          *string         `json:"locator,omitempty"`
	Site             *string         `json:"site,omitempty"`
	Blueprint        *string         `json:"blueprint,omitempty"`
	IDMap            *string         `json:"id_map,omitempty"`
	LocatedAt        *time.Time      `json:"located_at,omitempty"`
	BuiltAt          *time.Time      `json:"built_at,omitempty"`
	Updated          *time.Time      `json:"updated,omitempty"`
	Created          *time.Time      `json:"created,omitempty"`
	RedfishUsername  *string         `json:"redfish_username,omitempty"`
	RedfishPassword  *string         `json:"redfish_password,omitempty"`
	RedfishIPAddress *string         `json:"redfish_ip_address,omitempty"`
	RedfishSolPort   *string         `json:"redfish_sol_port,omitempty"`
	Plot             *string         `json:"plot,omitempty"`
	State            *string         `json:"state,omitempty"`
	Type             *string         `json:"type,omitempty"`
	ClassList        *string         `json:"class_list,omitempty"`
}

// RedfishRedFishFoundationNew - Make a new object of Model RedFishFoundation
func (service *Contractor) RedfishRedFishFoundationNew() *RedfishRedFishFoundation {
	return &RedfishRedFishFoundation{cinp: service.cinp}
}

// RedfishRedFishFoundationNewWithID - Make a new object of Model RedFishFoundation
func (service *Contractor) RedfishRedFishFoundationNewWithID(id string) *RedfishRedFishFoundation {
	result := RedfishRedFishFoundation{cinp: service.cinp}
	result.SetURI("/api/v1/RedFish/RedFishFoundation:" + id + ":")
	return &result
}

// RedfishRedFishFoundationGet - Get function for Model RedFishFoundation
func (service *Contractor) RedfishRedFishFoundationGet(ctx context.Context, id string) (*RedfishRedFishFoundation, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/RedFish/RedFishFoundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*RedfishRedFishFoundation)
	result.cinp = service.cinp

	return result, nil
}

// RedfishRedFishFoundationGetURI - Get function for Model RedFishFoundation vi URI
func (service *Contractor) RedfishRedFishFoundationGetURI(ctx context.Context, uri string) (*RedfishRedFishFoundation, error) {
	if !strings.HasPrefix(uri, "/api/v1/RedFish/RedFishFoundation:") {
		return nil, fmt.Errorf("URI is not for a 'RedfishRedFishFoundation'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*RedfishRedFishFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model RedFishFoundation
func (object *RedfishRedFishFoundation) Create(ctx context.Context) (*RedfishRedFishFoundation, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/RedFish/RedFishFoundation", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*RedfishRedFishFoundation), nil
}

// Update - Update function for Model RedFishFoundation
func (object *RedfishRedFishFoundation) Update(ctx context.Context) (*RedfishRedFishFoundation, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*RedfishRedFishFoundation), nil
}

// Delete - Delete function for Model RedFishFoundation
func (object *RedfishRedFishFoundation) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// RedfishRedFishFoundationListFilters - Return a slice of valid filter names RedFishFoundation
func (service *Contractor) RedfishRedFishFoundationListFilters() [1]string {
	return [1]string{"site"}
}

// RedfishRedFishFoundationList - List function for Model RedFishFoundation
func (service *Contractor) RedfishRedFishFoundationList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *RedfishRedFishFoundation, error) {
	if filterName != "" {
		for _, item := range service.RedfishRedFishFoundationListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/RedFish/RedFishFoundation", reflect.TypeOf(RedfishRedFishFoundation{}), filterName, filterValues, 50)
	out := make(chan *RedfishRedFishFoundation)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*RedfishRedFishFoundation).cinp = service.cinp
			out <- (*v).(*RedfishRedFishFoundation)
		}
	}()
	return out, nil
}

func registerRedFish(cinp cinp.CInPClient) {
	cinp.RegisterType("/api/v1/RedFish/RedFishFoundation", reflect.TypeOf((*RedfishRedFishFoundation)(nil)).Elem())
}
