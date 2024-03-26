// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Azure/ at 2024-03-26T14:17:37.955685
package contractor

import (
	"fmt"
	cinp "github.com/cinp/go"
	"reflect"
	"time"
)

// AzureAzureComplex - Model AzureComplex(/api/v1/Azure/AzureComplex)
/*
AzureComplex(name, site, description, built_percentage, updated, created, complex_ptr, azure_subscription_id, azure_location, azure_resource_group, azure_client_id, azure_password, azure_tenant_id)
*/
type AzureAzureComplex struct {
	cinp.BaseObject
	cinp                *cinp.CInP
	Name                *string    `json:"name"`
	Site                *string    `json:"site"`
	Description         *string    `json:"description"`
	BuiltPercentage     *int       `json:"built_percentage"`
	Updated             *time.Time `json:"updated"`
	Created             *time.Time `json:"created"`
	AzureSubscriptionID *string    `json:"azure_subscription_id"`
	AzureLocation       *string    `json:"azure_location"`
	AzureResourceGroup  *string    `json:"azure_resource_group"`
	AzureClientID       *string    `json:"azure_client_id"`
	AzurePassword       *string    `json:"azure_password"`
	AzureTenantID       *string    `json:"azure_tenant_id"`
	Members             *[]string  `json:"members"`
	State               *string    `json:"state"`
	Type                *string    `json:"type"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *AzureAzureComplex) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{
			"name":                  object.Name,
			"site":                  object.Site,
			"description":           object.Description,
			"built_percentage":      object.BuiltPercentage,
			"azure_subscription_id": object.AzureSubscriptionID,
			"azure_location":        object.AzureLocation,
			"azure_resource_group":  object.AzureResourceGroup,
			"azure_client_id":       object.AzureClientID,
			"azure_password":        object.AzurePassword,
			"azure_tenant_id":       object.AzureTenantID,
			"members":               object.Members,
		}
	}
	return &map[string]interface{}{
		"site":                  object.Site,
		"description":           object.Description,
		"built_percentage":      object.BuiltPercentage,
		"azure_subscription_id": object.AzureSubscriptionID,
		"azure_location":        object.AzureLocation,
		"azure_resource_group":  object.AzureResourceGroup,
		"azure_client_id":       object.AzureClientID,
		"azure_password":        object.AzurePassword,
		"azure_tenant_id":       object.AzureTenantID,
		"members":               object.Members,
	}
}

// AzureAzureComplexNew - Make a new object of Model AzureComplex
func (service *Contractor) AzureAzureComplexNew() *AzureAzureComplex {
	return &AzureAzureComplex{cinp: service.cinp}
}

// AzureAzureComplexNewWithID - Make a new object of Model AzureComplex
func (service *Contractor) AzureAzureComplexNewWithID(id string) *AzureAzureComplex {
	result := AzureAzureComplex{cinp: service.cinp}
	result.SetID("/api/v1/Azure/AzureComplex:" + id + ":")
	return &result
}

// AzureAzureComplexGet - Get function for Model AzureComplex
func (service *Contractor) AzureAzureComplexGet(id string) (*AzureAzureComplex, error) {
	object, err := service.cinp.Get("/api/v1/Azure/AzureComplex:" + id + ":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*AzureAzureComplex)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model AzureComplex
func (object *AzureAzureComplex) Create() error {
	if err := object.cinp.Create("/api/v1/Azure/AzureComplex", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model AzureComplex
func (object *AzureAzureComplex) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model AzureComplex
func (object *AzureAzureComplex) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// AzureAzureComplexListFilters - Return a slice of valid filter names AzureComplex
func (service *Contractor) AzureAzureComplexListFilters() [0]string {
	return [0]string{}
}

// AzureAzureComplexList - List function for Model AzureComplex
func (service *Contractor) AzureAzureComplexList(filterName string, filterValues map[string]interface{}) (<-chan *AzureAzureComplex, error) {
	if filterName != "" {
		for _, item := range service.AzureAzureComplexListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("Filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects("/api/v1/Azure/AzureComplex", reflect.TypeOf(AzureAzureComplex{}), filterName, filterValues, 50)
	out := make(chan *AzureAzureComplex)
	go func() {
		defer close(out)
		for v := range in {
			v.(*AzureAzureComplex).cinp = service.cinp
			out <- v.(*AzureAzureComplex)
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
	cinp              *cinp.CInP
	Locator           *string    `json:"locator"`
	Site              *string    `json:"site"`
	Blueprint         *string    `json:"blueprint"`
	IDMap             *string    `json:"id_map"`
	LocatedAt         *time.Time `json:"located_at"`
	BuiltAt           *time.Time `json:"built_at"`
	Updated           *time.Time `json:"updated"`
	Created           *time.Time `json:"created"`
	AzureComplex      *string    `json:"azure_complex"`
	AzureResourceName *string    `json:"azure_resource_name"`
	State             *string    `json:"state"`
	Type              *string    `json:"type"`
	ClassList         *string    `json:"class_list"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *AzureAzureFoundation) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{
			"locator":             object.Locator,
			"site":                object.Site,
			"blueprint":           object.Blueprint,
			"id_map":              object.IDMap,
			"azure_complex":       object.AzureComplex,
			"azure_resource_name": object.AzureResourceName,
		}
	}
	return &map[string]interface{}{
		"site":                object.Site,
		"blueprint":           object.Blueprint,
		"id_map":              object.IDMap,
		"azure_complex":       object.AzureComplex,
		"azure_resource_name": object.AzureResourceName,
	}
}

// AzureAzureFoundationNew - Make a new object of Model AzureFoundation
func (service *Contractor) AzureAzureFoundationNew() *AzureAzureFoundation {
	return &AzureAzureFoundation{cinp: service.cinp}
}

// AzureAzureFoundationNewWithID - Make a new object of Model AzureFoundation
func (service *Contractor) AzureAzureFoundationNewWithID(id string) *AzureAzureFoundation {
	result := AzureAzureFoundation{cinp: service.cinp}
	result.SetID("/api/v1/Azure/AzureFoundation:" + id + ":")
	return &result
}

// AzureAzureFoundationGet - Get function for Model AzureFoundation
func (service *Contractor) AzureAzureFoundationGet(id string) (*AzureAzureFoundation, error) {
	object, err := service.cinp.Get("/api/v1/Azure/AzureFoundation:" + id + ":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*AzureAzureFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model AzureFoundation
func (object *AzureAzureFoundation) Create() error {
	if err := object.cinp.Create("/api/v1/Azure/AzureFoundation", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model AzureFoundation
func (object *AzureAzureFoundation) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model AzureFoundation
func (object *AzureAzureFoundation) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// AzureAzureFoundationListFilters - Return a slice of valid filter names AzureFoundation
func (service *Contractor) AzureAzureFoundationListFilters() [1]string {
	return [1]string{"site"}
}

// AzureAzureFoundationList - List function for Model AzureFoundation
func (service *Contractor) AzureAzureFoundationList(filterName string, filterValues map[string]interface{}) (<-chan *AzureAzureFoundation, error) {
	if filterName != "" {
		for _, item := range service.AzureAzureFoundationListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("Filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects("/api/v1/Azure/AzureFoundation", reflect.TypeOf(AzureAzureFoundation{}), filterName, filterValues, 50)
	out := make(chan *AzureAzureFoundation)
	go func() {
		defer close(out)
		for v := range in {
			v.(*AzureAzureFoundation).cinp = service.cinp
			out <- v.(*AzureAzureFoundation)
		}
	}()
	return out, nil
}

func registerAzure(cinp *cinp.CInP) {
	cinp.RegisterType("/api/v1/Azure/AzureComplex", reflect.TypeOf((*AzureAzureComplex)(nil)).Elem())
	cinp.RegisterType("/api/v1/Azure/AzureFoundation", reflect.TypeOf((*AzureAzureFoundation)(nil)).Elem())
}
