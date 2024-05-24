// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Azure/ at 2024-05-24T15:42:51.144790
package contractor

import (
	"context"
	"fmt"
	cinp "github.com/cinp/go"
	"reflect"
	"strings"
	"time"
)

// AzureAzureComplex - Model AzureComplex(/api/v1/Azure/AzureComplex)
/*
AzureComplex(name, site, description, built_percentage, updated, created, complex_ptr, azure_subscription_id, azure_location, azure_resource_group, azure_client_id, azure_password, azure_tenant_id)
*/
type AzureAzureComplex struct {
	cinp.BaseObject
	cinp                cinp.CInPClient `json:"-"`
	Name                *string         `json:"name,omitempty"`
	Site                *string         `json:"site,omitempty"`
	Description         *string         `json:"description,omitempty"`
	BuiltPercentage     *int            `json:"built_percentage,omitempty"`
	Updated             *time.Time      `json:"updated,omitempty"`
	Created             *time.Time      `json:"created,omitempty"`
	AzureSubscriptionID *string         `json:"azure_subscription_id,omitempty"`
	AzureLocation       *string         `json:"azure_location,omitempty"`
	AzureResourceGroup  *string         `json:"azure_resource_group,omitempty"`
	AzureClientID       *string         `json:"azure_client_id,omitempty"`
	AzurePassword       *string         `json:"azure_password,omitempty"`
	AzureTenantID       *string         `json:"azure_tenant_id,omitempty"`
	Members             *[]string       `json:"members,omitempty"`
	State               *string         `json:"state,omitempty"`
	Type                *string         `json:"type,omitempty"`
}

// AzureAzureComplexNew - Make a new object of Model AzureComplex
func (service *Contractor) AzureAzureComplexNew() *AzureAzureComplex {
	return &AzureAzureComplex{cinp: service.cinp}
}

// AzureAzureComplexNewWithID - Make a new object of Model AzureComplex
func (service *Contractor) AzureAzureComplexNewWithID(id string) *AzureAzureComplex {
	result := AzureAzureComplex{cinp: service.cinp}
	result.SetURI("/api/v1/Azure/AzureComplex:" + id + ":")
	return &result
}

// AzureAzureComplexGet - Get function for Model AzureComplex
func (service *Contractor) AzureAzureComplexGet(ctx context.Context, id string) (*AzureAzureComplex, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Azure/AzureComplex:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*AzureAzureComplex)
	result.cinp = service.cinp

	return result, nil
}

// AzureAzureComplexGetURI - Get function for Model AzureComplex vi URI
func (service *Contractor) AzureAzureComplexGetURI(ctx context.Context, uri string) (*AzureAzureComplex, error) {
	if !strings.HasPrefix(uri, "/api/v1/Azure/AzureComplex:") {
		return nil, fmt.Errorf("URI is not for a 'AzureAzureComplex'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*AzureAzureComplex)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model AzureComplex
func (object *AzureAzureComplex) Create(ctx context.Context) (*AzureAzureComplex, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Azure/AzureComplex", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*AzureAzureComplex), nil
}

// Update - Update function for Model AzureComplex
func (object *AzureAzureComplex) Update(ctx context.Context) (*AzureAzureComplex, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*AzureAzureComplex), nil
}

// Delete - Delete function for Model AzureComplex
func (object *AzureAzureComplex) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// AzureAzureComplexListFilters - Return a slice of valid filter names AzureComplex
func (service *Contractor) AzureAzureComplexListFilters() [0]string {
	return [0]string{}
}

// AzureAzureComplexList - List function for Model AzureComplex
func (service *Contractor) AzureAzureComplexList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *AzureAzureComplex, error) {
	if filterName != "" {
		for _, item := range service.AzureAzureComplexListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Azure/AzureComplex", reflect.TypeOf(AzureAzureComplex{}), filterName, filterValues, 50)
	out := make(chan *AzureAzureComplex)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*AzureAzureComplex).cinp = service.cinp
			out <- (*v).(*AzureAzureComplex)
		}
	}()
	return out, nil
}

// AzureAzureFoundation - Model AzureFoundation(/api/v1/Azure/AzureFoundation)
/*
AzureFoundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr, azure_complex, azure_resource_name)
*/
type AzureAzureFoundation struct {
	cinp.BaseObject
	cinp              cinp.CInPClient `json:"-"`
	Locator           *string         `json:"locator,omitempty"`
	Site              *string         `json:"site,omitempty"`
	Blueprint         *string         `json:"blueprint,omitempty"`
	IDMap             *string         `json:"id_map,omitempty"`
	LocatedAt         *time.Time      `json:"located_at,omitempty"`
	BuiltAt           *time.Time      `json:"built_at,omitempty"`
	Updated           *time.Time      `json:"updated,omitempty"`
	Created           *time.Time      `json:"created,omitempty"`
	AzureComplex      *string         `json:"azure_complex,omitempty"`
	AzureResourceName *string         `json:"azure_resource_name,omitempty"`
	State             *string         `json:"state,omitempty"`
	Type              *string         `json:"type,omitempty"`
	ClassList         *string         `json:"class_list,omitempty"`
}

// AzureAzureFoundationNew - Make a new object of Model AzureFoundation
func (service *Contractor) AzureAzureFoundationNew() *AzureAzureFoundation {
	return &AzureAzureFoundation{cinp: service.cinp}
}

// AzureAzureFoundationNewWithID - Make a new object of Model AzureFoundation
func (service *Contractor) AzureAzureFoundationNewWithID(id string) *AzureAzureFoundation {
	result := AzureAzureFoundation{cinp: service.cinp}
	result.SetURI("/api/v1/Azure/AzureFoundation:" + id + ":")
	return &result
}

// AzureAzureFoundationGet - Get function for Model AzureFoundation
func (service *Contractor) AzureAzureFoundationGet(ctx context.Context, id string) (*AzureAzureFoundation, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Azure/AzureFoundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*AzureAzureFoundation)
	result.cinp = service.cinp

	return result, nil
}

// AzureAzureFoundationGetURI - Get function for Model AzureFoundation vi URI
func (service *Contractor) AzureAzureFoundationGetURI(ctx context.Context, uri string) (*AzureAzureFoundation, error) {
	if !strings.HasPrefix(uri, "/api/v1/Azure/AzureFoundation:") {
		return nil, fmt.Errorf("URI is not for a 'AzureAzureFoundation'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*AzureAzureFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model AzureFoundation
func (object *AzureAzureFoundation) Create(ctx context.Context) (*AzureAzureFoundation, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Azure/AzureFoundation", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*AzureAzureFoundation), nil
}

// Update - Update function for Model AzureFoundation
func (object *AzureAzureFoundation) Update(ctx context.Context) (*AzureAzureFoundation, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*AzureAzureFoundation), nil
}

// Delete - Delete function for Model AzureFoundation
func (object *AzureAzureFoundation) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// AzureAzureFoundationListFilters - Return a slice of valid filter names AzureFoundation
func (service *Contractor) AzureAzureFoundationListFilters() [1]string {
	return [1]string{"site"}
}

// AzureAzureFoundationList - List function for Model AzureFoundation
func (service *Contractor) AzureAzureFoundationList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *AzureAzureFoundation, error) {
	if filterName != "" {
		for _, item := range service.AzureAzureFoundationListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Azure/AzureFoundation", reflect.TypeOf(AzureAzureFoundation{}), filterName, filterValues, 50)
	out := make(chan *AzureAzureFoundation)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*AzureAzureFoundation).cinp = service.cinp
			out <- (*v).(*AzureAzureFoundation)
		}
	}()
	return out, nil
}

func registerAzure(cinp cinp.CInPClient) {
	cinp.RegisterType("/api/v1/Azure/AzureComplex", reflect.TypeOf((*AzureAzureComplex)(nil)).Elem())
	cinp.RegisterType("/api/v1/Azure/AzureFoundation", reflect.TypeOf((*AzureAzureFoundation)(nil)).Elem())
}
