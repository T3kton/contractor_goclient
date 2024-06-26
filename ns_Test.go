// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Test/ at 2024-05-24T15:42:51.144790
package contractor

import (
	"context"
	"fmt"
	cinp "github.com/cinp/go"
	"reflect"
	"strings"
	"time"
)

// TestTestComplex - Model TestComplex(/api/v1/Test/TestComplex)
/*
TestComplex(name, site, description, built_percentage, updated, created, complex_ptr)
*/
type TestTestComplex struct {
	cinp.BaseObject
	cinp            cinp.CInPClient `json:"-"`
	Name            *string         `json:"name,omitempty"`
	Site            *string         `json:"site,omitempty"`
	Description     *string         `json:"description,omitempty"`
	BuiltPercentage *int            `json:"built_percentage,omitempty"`
	Updated         *time.Time      `json:"updated,omitempty"`
	Created         *time.Time      `json:"created,omitempty"`
	Members         *[]string       `json:"members,omitempty"`
	State           *string         `json:"state,omitempty"`
	Type            *string         `json:"type,omitempty"`
}

// TestTestComplexNew - Make a new object of Model TestComplex
func (service *Contractor) TestTestComplexNew() *TestTestComplex {
	return &TestTestComplex{cinp: service.cinp}
}

// TestTestComplexNewWithID - Make a new object of Model TestComplex
func (service *Contractor) TestTestComplexNewWithID(id string) *TestTestComplex {
	result := TestTestComplex{cinp: service.cinp}
	result.SetURI("/api/v1/Test/TestComplex:" + id + ":")
	return &result
}

// TestTestComplexGet - Get function for Model TestComplex
func (service *Contractor) TestTestComplexGet(ctx context.Context, id string) (*TestTestComplex, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Test/TestComplex:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*TestTestComplex)
	result.cinp = service.cinp

	return result, nil
}

// TestTestComplexGetURI - Get function for Model TestComplex vi URI
func (service *Contractor) TestTestComplexGetURI(ctx context.Context, uri string) (*TestTestComplex, error) {
	if !strings.HasPrefix(uri, "/api/v1/Test/TestComplex:") {
		return nil, fmt.Errorf("URI is not for a 'TestTestComplex'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*TestTestComplex)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model TestComplex
func (object *TestTestComplex) Create(ctx context.Context) (*TestTestComplex, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Test/TestComplex", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*TestTestComplex), nil
}

// Update - Update function for Model TestComplex
func (object *TestTestComplex) Update(ctx context.Context) (*TestTestComplex, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*TestTestComplex), nil
}

// Delete - Delete function for Model TestComplex
func (object *TestTestComplex) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// TestTestComplexListFilters - Return a slice of valid filter names TestComplex
func (service *Contractor) TestTestComplexListFilters() [0]string {
	return [0]string{}
}

// TestTestComplexList - List function for Model TestComplex
func (service *Contractor) TestTestComplexList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *TestTestComplex, error) {
	if filterName != "" {
		for _, item := range service.TestTestComplexListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Test/TestComplex", reflect.TypeOf(TestTestComplex{}), filterName, filterValues, 50)
	out := make(chan *TestTestComplex)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*TestTestComplex).cinp = service.cinp
			out <- (*v).(*TestTestComplex)
		}
	}()
	return out, nil
}

// TestTestFoundation - Model TestFoundation(/api/v1/Test/TestFoundation)
/*
TestFoundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr, test_delay_variance, test_fail_likelihood)
*/
type TestTestFoundation struct {
	cinp.BaseObject
	cinp               cinp.CInPClient `json:"-"`
	Locator            *string         `json:"locator,omitempty"`
	Site               *string         `json:"site,omitempty"`
	Blueprint          *string         `json:"blueprint,omitempty"`
	IDMap              *string         `json:"id_map,omitempty"`
	LocatedAt          *time.Time      `json:"located_at,omitempty"`
	BuiltAt            *time.Time      `json:"built_at,omitempty"`
	Updated            *time.Time      `json:"updated,omitempty"`
	Created            *time.Time      `json:"created,omitempty"`
	TestDelayVariance  *int            `json:"test_delay_variance,omitempty"`
	TestFailLikelihood *int            `json:"test_fail_likelihood,omitempty"`
	State              *string         `json:"state,omitempty"`
	Type               *string         `json:"type,omitempty"`
	ClassList          *string         `json:"class_list,omitempty"`
}

// TestTestFoundationNew - Make a new object of Model TestFoundation
func (service *Contractor) TestTestFoundationNew() *TestTestFoundation {
	return &TestTestFoundation{cinp: service.cinp}
}

// TestTestFoundationNewWithID - Make a new object of Model TestFoundation
func (service *Contractor) TestTestFoundationNewWithID(id string) *TestTestFoundation {
	result := TestTestFoundation{cinp: service.cinp}
	result.SetURI("/api/v1/Test/TestFoundation:" + id + ":")
	return &result
}

// TestTestFoundationGet - Get function for Model TestFoundation
func (service *Contractor) TestTestFoundationGet(ctx context.Context, id string) (*TestTestFoundation, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Test/TestFoundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*TestTestFoundation)
	result.cinp = service.cinp

	return result, nil
}

// TestTestFoundationGetURI - Get function for Model TestFoundation vi URI
func (service *Contractor) TestTestFoundationGetURI(ctx context.Context, uri string) (*TestTestFoundation, error) {
	if !strings.HasPrefix(uri, "/api/v1/Test/TestFoundation:") {
		return nil, fmt.Errorf("URI is not for a 'TestTestFoundation'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*TestTestFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model TestFoundation
func (object *TestTestFoundation) Create(ctx context.Context) (*TestTestFoundation, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Test/TestFoundation", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*TestTestFoundation), nil
}

// Update - Update function for Model TestFoundation
func (object *TestTestFoundation) Update(ctx context.Context) (*TestTestFoundation, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*TestTestFoundation), nil
}

// Delete - Delete function for Model TestFoundation
func (object *TestTestFoundation) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// TestTestFoundationListFilters - Return a slice of valid filter names TestFoundation
func (service *Contractor) TestTestFoundationListFilters() [1]string {
	return [1]string{"site"}
}

// TestTestFoundationList - List function for Model TestFoundation
func (service *Contractor) TestTestFoundationList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *TestTestFoundation, error) {
	if filterName != "" {
		for _, item := range service.TestTestFoundationListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Test/TestFoundation", reflect.TypeOf(TestTestFoundation{}), filterName, filterValues, 50)
	out := make(chan *TestTestFoundation)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*TestTestFoundation).cinp = service.cinp
			out <- (*v).(*TestTestFoundation)
		}
	}()
	return out, nil
}

// TestTestComplexedFoundation - Model TestComplexedFoundation(/api/v1/Test/TestComplexedFoundation)
/*
TestComplexedFoundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr, complex_host)
*/
type TestTestComplexedFoundation struct {
	cinp.BaseObject
	cinp        cinp.CInPClient `json:"-"`
	Locator     *string         `json:"locator,omitempty"`
	Site        *string         `json:"site,omitempty"`
	Blueprint   *string         `json:"blueprint,omitempty"`
	IDMap       *string         `json:"id_map,omitempty"`
	LocatedAt   *time.Time      `json:"located_at,omitempty"`
	BuiltAt     *time.Time      `json:"built_at,omitempty"`
	Updated     *time.Time      `json:"updated,omitempty"`
	Created     *time.Time      `json:"created,omitempty"`
	ComplexHost *string         `json:"complex_host,omitempty"`
	State       *string         `json:"state,omitempty"`
	Type        *string         `json:"type,omitempty"`
	ClassList   *string         `json:"class_list,omitempty"`
}

// TestTestComplexedFoundationNew - Make a new object of Model TestComplexedFoundation
func (service *Contractor) TestTestComplexedFoundationNew() *TestTestComplexedFoundation {
	return &TestTestComplexedFoundation{cinp: service.cinp}
}

// TestTestComplexedFoundationNewWithID - Make a new object of Model TestComplexedFoundation
func (service *Contractor) TestTestComplexedFoundationNewWithID(id string) *TestTestComplexedFoundation {
	result := TestTestComplexedFoundation{cinp: service.cinp}
	result.SetURI("/api/v1/Test/TestComplexedFoundation:" + id + ":")
	return &result
}

// TestTestComplexedFoundationGet - Get function for Model TestComplexedFoundation
func (service *Contractor) TestTestComplexedFoundationGet(ctx context.Context, id string) (*TestTestComplexedFoundation, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Test/TestComplexedFoundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*TestTestComplexedFoundation)
	result.cinp = service.cinp

	return result, nil
}

// TestTestComplexedFoundationGetURI - Get function for Model TestComplexedFoundation vi URI
func (service *Contractor) TestTestComplexedFoundationGetURI(ctx context.Context, uri string) (*TestTestComplexedFoundation, error) {
	if !strings.HasPrefix(uri, "/api/v1/Test/TestComplexedFoundation:") {
		return nil, fmt.Errorf("URI is not for a 'TestTestComplexedFoundation'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*TestTestComplexedFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model TestComplexedFoundation
func (object *TestTestComplexedFoundation) Create(ctx context.Context) (*TestTestComplexedFoundation, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Test/TestComplexedFoundation", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*TestTestComplexedFoundation), nil
}

// Update - Update function for Model TestComplexedFoundation
func (object *TestTestComplexedFoundation) Update(ctx context.Context) (*TestTestComplexedFoundation, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*TestTestComplexedFoundation), nil
}

// Delete - Delete function for Model TestComplexedFoundation
func (object *TestTestComplexedFoundation) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// TestTestComplexedFoundationListFilters - Return a slice of valid filter names TestComplexedFoundation
func (service *Contractor) TestTestComplexedFoundationListFilters() [1]string {
	return [1]string{"site"}
}

// TestTestComplexedFoundationList - List function for Model TestComplexedFoundation
func (service *Contractor) TestTestComplexedFoundationList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *TestTestComplexedFoundation, error) {
	if filterName != "" {
		for _, item := range service.TestTestComplexedFoundationListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Test/TestComplexedFoundation", reflect.TypeOf(TestTestComplexedFoundation{}), filterName, filterValues, 50)
	out := make(chan *TestTestComplexedFoundation)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*TestTestComplexedFoundation).cinp = service.cinp
			out <- (*v).(*TestTestComplexedFoundation)
		}
	}()
	return out, nil
}

func registerTest(cinp cinp.CInPClient) {
	cinp.RegisterType("/api/v1/Test/TestComplex", reflect.TypeOf((*TestTestComplex)(nil)).Elem())
	cinp.RegisterType("/api/v1/Test/TestFoundation", reflect.TypeOf((*TestTestFoundation)(nil)).Elem())
	cinp.RegisterType("/api/v1/Test/TestComplexedFoundation", reflect.TypeOf((*TestTestComplexedFoundation)(nil)).Elem())
}
