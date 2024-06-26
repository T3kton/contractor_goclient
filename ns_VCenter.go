// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/VCenter/ at 2024-05-24T15:42:51.144790
package contractor

import (
	"context"
	"fmt"
	cinp "github.com/cinp/go"
	"reflect"
	"strings"
	"time"
)

// VcenterVCenterComplex - Model VCenterComplex(/api/v1/VCenter/VCenterComplex)
/*
VCenterComplex(name, site, description, built_percentage, updated, created, complex_ptr, vcenter_host, vcenter_username, vcenter_password, vcenter_datacenter, vcenter_cluster)
*/
type VcenterVCenterComplex struct {
	cinp.BaseObject
	cinp              cinp.CInPClient `json:"-"`
	Name              *string         `json:"name,omitempty"`
	Site              *string         `json:"site,omitempty"`
	Description       *string         `json:"description,omitempty"`
	BuiltPercentage   *int            `json:"built_percentage,omitempty"`
	Updated           *time.Time      `json:"updated,omitempty"`
	Created           *time.Time      `json:"created,omitempty"`
	VcenterHost       *string         `json:"vcenter_host,omitempty"`
	VcenterUsername   *string         `json:"vcenter_username,omitempty"`
	VcenterPassword   *string         `json:"vcenter_password,omitempty"`
	VcenterDatacenter *string         `json:"vcenter_datacenter,omitempty"`
	VcenterCluster    *string         `json:"vcenter_cluster,omitempty"`
	Members           *[]string       `json:"members,omitempty"`
	State             *string         `json:"state,omitempty"`
	Type              *string         `json:"type,omitempty"`
}

// VcenterVCenterComplexNew - Make a new object of Model VCenterComplex
func (service *Contractor) VcenterVCenterComplexNew() *VcenterVCenterComplex {
	return &VcenterVCenterComplex{cinp: service.cinp}
}

// VcenterVCenterComplexNewWithID - Make a new object of Model VCenterComplex
func (service *Contractor) VcenterVCenterComplexNewWithID(id string) *VcenterVCenterComplex {
	result := VcenterVCenterComplex{cinp: service.cinp}
	result.SetURI("/api/v1/VCenter/VCenterComplex:" + id + ":")
	return &result
}

// VcenterVCenterComplexGet - Get function for Model VCenterComplex
func (service *Contractor) VcenterVCenterComplexGet(ctx context.Context, id string) (*VcenterVCenterComplex, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/VCenter/VCenterComplex:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*VcenterVCenterComplex)
	result.cinp = service.cinp

	return result, nil
}

// VcenterVCenterComplexGetURI - Get function for Model VCenterComplex vi URI
func (service *Contractor) VcenterVCenterComplexGetURI(ctx context.Context, uri string) (*VcenterVCenterComplex, error) {
	if !strings.HasPrefix(uri, "/api/v1/VCenter/VCenterComplex:") {
		return nil, fmt.Errorf("URI is not for a 'VcenterVCenterComplex'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*VcenterVCenterComplex)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model VCenterComplex
func (object *VcenterVCenterComplex) Create(ctx context.Context) (*VcenterVCenterComplex, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/VCenter/VCenterComplex", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*VcenterVCenterComplex), nil
}

// Update - Update function for Model VCenterComplex
func (object *VcenterVCenterComplex) Update(ctx context.Context) (*VcenterVCenterComplex, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*VcenterVCenterComplex), nil
}

// Delete - Delete function for Model VCenterComplex
func (object *VcenterVCenterComplex) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// VcenterVCenterComplexListFilters - Return a slice of valid filter names VCenterComplex
func (service *Contractor) VcenterVCenterComplexListFilters() [0]string {
	return [0]string{}
}

// VcenterVCenterComplexList - List function for Model VCenterComplex
func (service *Contractor) VcenterVCenterComplexList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *VcenterVCenterComplex, error) {
	if filterName != "" {
		for _, item := range service.VcenterVCenterComplexListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/VCenter/VCenterComplex", reflect.TypeOf(VcenterVCenterComplex{}), filterName, filterValues, 50)
	out := make(chan *VcenterVCenterComplex)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*VcenterVCenterComplex).cinp = service.cinp
			out <- (*v).(*VcenterVCenterComplex)
		}
	}()
	return out, nil
}

// VcenterVCenterFoundation - Model VCenterFoundation(/api/v1/VCenter/VCenterFoundation)
/*
VCenterFoundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr, vcenter_complex, vcenter_uuid)
*/
type VcenterVCenterFoundation struct {
	cinp.BaseObject
	cinp           cinp.CInPClient `json:"-"`
	Locator        *string         `json:"locator,omitempty"`
	Site           *string         `json:"site,omitempty"`
	Blueprint      *string         `json:"blueprint,omitempty"`
	IDMap          *string         `json:"id_map,omitempty"`
	LocatedAt      *time.Time      `json:"located_at,omitempty"`
	BuiltAt        *time.Time      `json:"built_at,omitempty"`
	Updated        *time.Time      `json:"updated,omitempty"`
	Created        *time.Time      `json:"created,omitempty"`
	VcenterComplex *string         `json:"vcenter_complex,omitempty"`
	VcenterUUID    *string         `json:"vcenter_uuid,omitempty"`
	State          *string         `json:"state,omitempty"`
	Type           *string         `json:"type,omitempty"`
	ClassList      *string         `json:"class_list,omitempty"`
}

// VcenterVCenterFoundationNew - Make a new object of Model VCenterFoundation
func (service *Contractor) VcenterVCenterFoundationNew() *VcenterVCenterFoundation {
	return &VcenterVCenterFoundation{cinp: service.cinp}
}

// VcenterVCenterFoundationNewWithID - Make a new object of Model VCenterFoundation
func (service *Contractor) VcenterVCenterFoundationNewWithID(id string) *VcenterVCenterFoundation {
	result := VcenterVCenterFoundation{cinp: service.cinp}
	result.SetURI("/api/v1/VCenter/VCenterFoundation:" + id + ":")
	return &result
}

// VcenterVCenterFoundationGet - Get function for Model VCenterFoundation
func (service *Contractor) VcenterVCenterFoundationGet(ctx context.Context, id string) (*VcenterVCenterFoundation, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/VCenter/VCenterFoundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*VcenterVCenterFoundation)
	result.cinp = service.cinp

	return result, nil
}

// VcenterVCenterFoundationGetURI - Get function for Model VCenterFoundation vi URI
func (service *Contractor) VcenterVCenterFoundationGetURI(ctx context.Context, uri string) (*VcenterVCenterFoundation, error) {
	if !strings.HasPrefix(uri, "/api/v1/VCenter/VCenterFoundation:") {
		return nil, fmt.Errorf("URI is not for a 'VcenterVCenterFoundation'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*VcenterVCenterFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model VCenterFoundation
func (object *VcenterVCenterFoundation) Create(ctx context.Context) (*VcenterVCenterFoundation, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/VCenter/VCenterFoundation", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*VcenterVCenterFoundation), nil
}

// Update - Update function for Model VCenterFoundation
func (object *VcenterVCenterFoundation) Update(ctx context.Context) (*VcenterVCenterFoundation, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*VcenterVCenterFoundation), nil
}

// Delete - Delete function for Model VCenterFoundation
func (object *VcenterVCenterFoundation) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// VcenterVCenterFoundationListFilters - Return a slice of valid filter names VCenterFoundation
func (service *Contractor) VcenterVCenterFoundationListFilters() [1]string {
	return [1]string{"site"}
}

// VcenterVCenterFoundationList - List function for Model VCenterFoundation
func (service *Contractor) VcenterVCenterFoundationList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *VcenterVCenterFoundation, error) {
	if filterName != "" {
		for _, item := range service.VcenterVCenterFoundationListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/VCenter/VCenterFoundation", reflect.TypeOf(VcenterVCenterFoundation{}), filterName, filterValues, 50)
	out := make(chan *VcenterVCenterFoundation)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*VcenterVCenterFoundation).cinp = service.cinp
			out <- (*v).(*VcenterVCenterFoundation)
		}
	}()
	return out, nil
}

func registerVCenter(cinp cinp.CInPClient) {
	cinp.RegisterType("/api/v1/VCenter/VCenterComplex", reflect.TypeOf((*VcenterVCenterComplex)(nil)).Elem())
	cinp.RegisterType("/api/v1/VCenter/VCenterFoundation", reflect.TypeOf((*VcenterVCenterFoundation)(nil)).Elem())
}
