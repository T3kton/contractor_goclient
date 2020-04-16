/*Package contractor(version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Manual/ at 2020-04-16T20:18:01.673062
 */
package contractor

import (
	"reflect"
	"time"
	cinp "github.com/cinp/go"
)

//ManualManualComplex - Model ManualComplex(/api/v1/Manual/ManualComplex)
/*
ManualComplex(name, site, description, built_percentage, updated, created, complex_ptr)
 */
type ManualManualComplex struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Name string `json:"name"`
	Site string `json:"site"`
	Description string `json:"description"`
	BuiltPercentage int `json:"built_percentage"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	Members []string `json:"members"`
	State string `json:"state"`
	Type string `json:"type"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *ManualManualComplex) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"name": object.Name,
			"site": object.Site,
			"description": object.Description,
			"built_percentage": object.BuiltPercentage,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"description": object.Description,
		"built_percentage": object.BuiltPercentage,
	}
}

// ManualManualComplexNew - Make a new object of Model ManualComplex
func (service *Contractor) ManualManualComplexNew() *ManualManualComplex {
	return &ManualManualComplex{cinp: service.cinp}
}

// ManualManualComplexNewWithID - Make a new object of Model ManualComplex
func (service *Contractor) ManualManualComplexNewWithID(id string) *ManualManualComplex {
	result := ManualManualComplex{cinp: service.cinp}
	result.SetID("/api/v1/Manual/ManualComplex:"+id+":")
	return &result
}

// ManualManualComplexGet - Get function for Model ManualComplex
func (service *Contractor) ManualManualComplexGet(id string) (*ManualManualComplex, error) {
	object, err := service.cinp.Get("/api/v1/Manual/ManualComplex:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*ManualManualComplex)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model ManualComplex
func (object *ManualManualComplex) Create() error {
	if err := object.cinp.Create("/api/v1/Manual/ManualComplex", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model ManualComplex
func (object *ManualManualComplex) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model ManualComplex
func (object *ManualManualComplex) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// ManualManualComplexList - List function for Model ManualComplex
func (service *Contractor) ManualManualComplexList(filterName string, filterValues map[string]interface{}) <-chan *ManualManualComplex {
	in := service.cinp.ListObjects("/api/v1/Manual/ManualComplex", reflect.TypeOf(ManualManualComplex{}), filterName, filterValues, 50)
	out := make(chan *ManualManualComplex)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*ManualManualComplex)
		}
	}()
	return out
}


//ManualManualFoundation - Model ManualFoundation(/api/v1/Manual/ManualFoundation)
/*
ManualFoundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr)
 */
type ManualManualFoundation struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Locator string `json:"locator"`
	Site string `json:"site"`
	Blueprint string `json:"blueprint"`
	IDMap string `json:"id_map"`
	LocatedAt time.Time `json:"located_at"`
	BuiltAt time.Time `json:"built_at"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	State string `json:"state"`
	Type string `json:"type"`
	ClassList string `json:"class_list"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *ManualManualFoundation) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"locator": object.Locator,
			"site": object.Site,
			"blueprint": object.Blueprint,
			"id_map": object.IDMap,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"blueprint": object.Blueprint,
		"id_map": object.IDMap,
	}
}

// ManualManualFoundationNew - Make a new object of Model ManualFoundation
func (service *Contractor) ManualManualFoundationNew() *ManualManualFoundation {
	return &ManualManualFoundation{cinp: service.cinp}
}

// ManualManualFoundationNewWithID - Make a new object of Model ManualFoundation
func (service *Contractor) ManualManualFoundationNewWithID(id string) *ManualManualFoundation {
	result := ManualManualFoundation{cinp: service.cinp}
	result.SetID("/api/v1/Manual/ManualFoundation:"+id+":")
	return &result
}

// ManualManualFoundationGet - Get function for Model ManualFoundation
func (service *Contractor) ManualManualFoundationGet(id string) (*ManualManualFoundation, error) {
	object, err := service.cinp.Get("/api/v1/Manual/ManualFoundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*ManualManualFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model ManualFoundation
func (object *ManualManualFoundation) Create() error {
	if err := object.cinp.Create("/api/v1/Manual/ManualFoundation", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model ManualFoundation
func (object *ManualManualFoundation) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model ManualFoundation
func (object *ManualManualFoundation) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// ManualManualFoundationList - List function for Model ManualFoundation
func (service *Contractor) ManualManualFoundationList(filterName string, filterValues map[string]interface{}) <-chan *ManualManualFoundation {
	in := service.cinp.ListObjects("/api/v1/Manual/ManualFoundation", reflect.TypeOf(ManualManualFoundation{}), filterName, filterValues, 50)
	out := make(chan *ManualManualFoundation)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*ManualManualFoundation)
		}
	}()
	return out
}


//ManualManualComplexedFoundation - Model ManualComplexedFoundation(/api/v1/Manual/ManualComplexedFoundation)
/*
ManualComplexedFoundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr, complex_host)
 */
type ManualManualComplexedFoundation struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Locator string `json:"locator"`
	Site string `json:"site"`
	Blueprint string `json:"blueprint"`
	IDMap string `json:"id_map"`
	LocatedAt time.Time `json:"located_at"`
	BuiltAt time.Time `json:"built_at"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	ComplexHost string `json:"complex_host"`
	State string `json:"state"`
	Type string `json:"type"`
	ClassList string `json:"class_list"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *ManualManualComplexedFoundation) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"locator": object.Locator,
			"site": object.Site,
			"blueprint": object.Blueprint,
			"id_map": object.IDMap,
			"complex_host": object.ComplexHost,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"blueprint": object.Blueprint,
		"id_map": object.IDMap,
		"complex_host": object.ComplexHost,
	}
}

// ManualManualComplexedFoundationNew - Make a new object of Model ManualComplexedFoundation
func (service *Contractor) ManualManualComplexedFoundationNew() *ManualManualComplexedFoundation {
	return &ManualManualComplexedFoundation{cinp: service.cinp}
}

// ManualManualComplexedFoundationNewWithID - Make a new object of Model ManualComplexedFoundation
func (service *Contractor) ManualManualComplexedFoundationNewWithID(id string) *ManualManualComplexedFoundation {
	result := ManualManualComplexedFoundation{cinp: service.cinp}
	result.SetID("/api/v1/Manual/ManualComplexedFoundation:"+id+":")
	return &result
}

// ManualManualComplexedFoundationGet - Get function for Model ManualComplexedFoundation
func (service *Contractor) ManualManualComplexedFoundationGet(id string) (*ManualManualComplexedFoundation, error) {
	object, err := service.cinp.Get("/api/v1/Manual/ManualComplexedFoundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*ManualManualComplexedFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model ManualComplexedFoundation
func (object *ManualManualComplexedFoundation) Create() error {
	if err := object.cinp.Create("/api/v1/Manual/ManualComplexedFoundation", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model ManualComplexedFoundation
func (object *ManualManualComplexedFoundation) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model ManualComplexedFoundation
func (object *ManualManualComplexedFoundation) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// ManualManualComplexedFoundationList - List function for Model ManualComplexedFoundation
func (service *Contractor) ManualManualComplexedFoundationList(filterName string, filterValues map[string]interface{}) <-chan *ManualManualComplexedFoundation {
	in := service.cinp.ListObjects("/api/v1/Manual/ManualComplexedFoundation", reflect.TypeOf(ManualManualComplexedFoundation{}), filterName, filterValues, 50)
	out := make(chan *ManualManualComplexedFoundation)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*ManualManualComplexedFoundation)
		}
	}()
	return out
}

func registerManual(cinp *cinp.CInP) { 
	cinp.RegisterType("/api/v1/Manual/ManualComplex", reflect.TypeOf((*ManualManualComplex)(nil)).Elem())
	cinp.RegisterType("/api/v1/Manual/ManualFoundation", reflect.TypeOf((*ManualManualFoundation)(nil)).Elem())
	cinp.RegisterType("/api/v1/Manual/ManualComplexedFoundation", reflect.TypeOf((*ManualManualComplexedFoundation)(nil)).Elem())
}