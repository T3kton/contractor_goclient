// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Building/ at 2024-05-24T15:42:51.144790
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

// BuildingFoundation - Model Foundation(/api/v1/Building/Foundation)
/*
Foundation(locator, site, blueprint, id_map, located_at, built_at, updated, created)
*/
type BuildingFoundation struct {
	cinp.BaseObject
	cinp      cinp.CInPClient `json:"-"`
	Locator   *string         `json:"locator,omitempty"`
	Site      *string         `json:"site,omitempty"`
	Blueprint *string         `json:"blueprint,omitempty"`
	IDMap     *string         `json:"id_map,omitempty"`
	LocatedAt *time.Time      `json:"located_at,omitempty"`
	BuiltAt   *time.Time      `json:"built_at,omitempty"`
	Updated   *time.Time      `json:"updated,omitempty"`
	Created   *time.Time      `json:"created,omitempty"`
	State     *string         `json:"state,omitempty"`
	Type      *string         `json:"type,omitempty"`
	ClassList *string         `json:"class_list,omitempty"`
	Structure *string         `json:"structure,omitempty"`
}

// BuildingFoundationNew - Make a new object of Model Foundation
func (service *Contractor) BuildingFoundationNew() *BuildingFoundation {
	return &BuildingFoundation{cinp: service.cinp}
}

// BuildingFoundationNewWithID - Make a new object of Model Foundation
func (service *Contractor) BuildingFoundationNewWithID(id string) *BuildingFoundation {
	result := BuildingFoundation{cinp: service.cinp}
	result.SetURI("/api/v1/Building/Foundation:" + id + ":")
	return &result
}

// BuildingFoundationGet - Get function for Model Foundation
func (service *Contractor) BuildingFoundationGet(ctx context.Context, id string) (*BuildingFoundation, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Building/Foundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingFoundation)
	result.cinp = service.cinp

	return result, nil
}

// BuildingFoundationGetURI - Get function for Model Foundation vi URI
func (service *Contractor) BuildingFoundationGetURI(ctx context.Context, uri string) (*BuildingFoundation, error) {
	if !strings.HasPrefix(uri, "/api/v1/Building/Foundation:") {
		return nil, fmt.Errorf("URI is not for a 'BuildingFoundation'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Update - Update function for Model Foundation
func (object *BuildingFoundation) Update(ctx context.Context) (*BuildingFoundation, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*BuildingFoundation), nil
}

// Delete - Delete function for Model Foundation
func (object *BuildingFoundation) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// BuildingFoundationListFilters - Return a slice of valid filter names Foundation
func (service *Contractor) BuildingFoundationListFilters() [2]string {
	return [2]string{"site", "todo"}
}

// BuildingFoundationList - List function for Model Foundation
func (service *Contractor) BuildingFoundationList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *BuildingFoundation, error) {
	if filterName != "" {
		for _, item := range service.BuildingFoundationListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Building/Foundation", reflect.TypeOf(BuildingFoundation{}), filterName, filterValues, 50)
	out := make(chan *BuildingFoundation)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*BuildingFoundation).cinp = service.cinp
			out <- (*v).(*BuildingFoundation)
		}
	}()
	return out, nil
}

// CallSetIDMap calls setIdMap
func (object *BuildingFoundation) CallSetIDMap(ctx context.Context, IDMap map[string]interface{}) (string, error) {
	args := map[string]interface{}{
		"id_map": IDMap,
	}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Foundation(setIdMap)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// CallDoCreate calls doCreate
func (object *BuildingFoundation) CallDoCreate(ctx context.Context) (int, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return 0, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Foundation(doCreate)", ids)
	if err != nil {
		return 0, err
	}

	result := 0

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return 0, err
	}

	return result, nil
}

// CallDoDestroy calls doDestroy
func (object *BuildingFoundation) CallDoDestroy(ctx context.Context) (int, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return 0, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Foundation(doDestroy)", ids)
	if err != nil {
		return 0, err
	}

	result := 0

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return 0, err
	}

	return result, nil
}

// CallDoJob calls doJob
func (object *BuildingFoundation) CallDoJob(ctx context.Context, Name string) (int, error) {
	args := map[string]interface{}{
		"name": Name,
	}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return 0, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Foundation(doJob)", ids)
	if err != nil {
		return 0, err
	}

	result := 0

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return 0, err
	}

	return result, nil
}

// CallGetJob calls getJob
func (object *BuildingFoundation) CallGetJob(ctx context.Context) (string, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Foundation(getJob)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// BuildingFoundationCallGetFoundationTypes calls getFoundationTypes
func (service *Contractor) BuildingFoundationCallGetFoundationTypes(ctx context.Context) ([]string, error) {
	args := map[string]interface{}{}
	uri := "/api/v1/Building/Foundation(getFoundationTypes)"

	result := []string{}

	if err := service.cinp.Call(ctx, uri, &args, &result); err != nil {
		return []string{}, err
	}

	return result, nil
}

// CallGetConfig calls getConfig
func (object *BuildingFoundation) CallGetConfig(ctx context.Context) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Foundation(getConfig)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CallGetInterfaceList calls getInterfaceList
func (object *BuildingFoundation) CallGetInterfaceList(ctx context.Context) ([]map[string]interface{}, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Foundation(getInterfaceList)", ids)
	if err != nil {
		return nil, err
	}

	result := []map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// BuildingStructure - Model Structure(/api/v1/Building/Structure)
/*
Structure(id, hostname, site, networked_ptr, blueprint, foundation, config_uuid, config_values, built_at, updated, created)
*/
type BuildingStructure struct {
	cinp.BaseObject
	cinp         cinp.CInPClient         `json:"-"`
	Hostname     *string                 `json:"hostname,omitempty"`
	Site         *string                 `json:"site,omitempty"`
	Blueprint    *string                 `json:"blueprint,omitempty"`
	Foundation   *string                 `json:"foundation,omitempty"`
	ConfigUUID   *string                 `json:"config_uuid,omitempty"`
	ConfigValues *map[string]interface{} `json:"config_values,omitempty"`
	BuiltAt      *time.Time              `json:"built_at,omitempty"`
	Updated      *time.Time              `json:"updated,omitempty"`
	Created      *time.Time              `json:"created,omitempty"`
	ID           *int                    `json:"id,omitempty"`
	State        *string                 `json:"state,omitempty"`
}

// BuildingStructureNew - Make a new object of Model Structure
func (service *Contractor) BuildingStructureNew() *BuildingStructure {
	return &BuildingStructure{cinp: service.cinp}
}

// BuildingStructureNewWithID - Make a new object of Model Structure
func (service *Contractor) BuildingStructureNewWithID(id int) *BuildingStructure {
	result := BuildingStructure{cinp: service.cinp}
	result.SetURI("/api/v1/Building/Structure:" + strconv.FormatInt(int64(id), 10) + ":")
	return &result
}

// BuildingStructureGet - Get function for Model Structure
func (service *Contractor) BuildingStructureGet(ctx context.Context, id int) (*BuildingStructure, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Building/Structure:"+strconv.FormatInt(int64(id), 10)+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingStructure)
	result.cinp = service.cinp

	return result, nil
}

// BuildingStructureGetURI - Get function for Model Structure vi URI
func (service *Contractor) BuildingStructureGetURI(ctx context.Context, uri string) (*BuildingStructure, error) {
	if !strings.HasPrefix(uri, "/api/v1/Building/Structure:") {
		return nil, fmt.Errorf("URI is not for a 'BuildingStructure'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingStructure)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Structure
func (object *BuildingStructure) Create(ctx context.Context) (*BuildingStructure, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Building/Structure", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*BuildingStructure), nil
}

// Update - Update function for Model Structure
func (object *BuildingStructure) Update(ctx context.Context) (*BuildingStructure, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*BuildingStructure), nil
}

// Delete - Delete function for Model Structure
func (object *BuildingStructure) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// BuildingStructureListFilters - Return a slice of valid filter names Structure
func (service *Contractor) BuildingStructureListFilters() [1]string {
	return [1]string{"site"}
}

// BuildingStructureList - List function for Model Structure
func (service *Contractor) BuildingStructureList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *BuildingStructure, error) {
	if filterName != "" {
		for _, item := range service.BuildingStructureListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Building/Structure", reflect.TypeOf(BuildingStructure{}), filterName, filterValues, 50)
	out := make(chan *BuildingStructure)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*BuildingStructure).cinp = service.cinp
			out <- (*v).(*BuildingStructure)
		}
	}()
	return out, nil
}

// CallDoCreate calls doCreate
func (object *BuildingStructure) CallDoCreate(ctx context.Context) (int, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return 0, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Structure(doCreate)", ids)
	if err != nil {
		return 0, err
	}

	result := 0

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return 0, err
	}

	return result, nil
}

// CallDoDestroy calls doDestroy
func (object *BuildingStructure) CallDoDestroy(ctx context.Context) (int, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return 0, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Structure(doDestroy)", ids)
	if err != nil {
		return 0, err
	}

	result := 0

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return 0, err
	}

	return result, nil
}

// CallDoJob calls doJob
func (object *BuildingStructure) CallDoJob(ctx context.Context, Name string) (int, error) {
	args := map[string]interface{}{
		"name": Name,
	}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return 0, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Structure(doJob)", ids)
	if err != nil {
		return 0, err
	}

	result := 0

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return 0, err
	}

	return result, nil
}

// CallGetJob calls getJob
func (object *BuildingStructure) CallGetJob(ctx context.Context) (string, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Structure(getJob)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// CallGetConfig calls getConfig
func (object *BuildingStructure) CallGetConfig(ctx context.Context) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Structure(getConfig)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CallUpdateConfig calls updateConfig
func (object *BuildingStructure) CallUpdateConfig(ctx context.Context, ConfigValueMap map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{
		"config_value_map": ConfigValueMap,
	}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Structure(updateConfig)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// BuildingStructureCallGetWithHostnameSite calls getWithHostnameSite
func (service *Contractor) BuildingStructureCallGetWithHostnameSite(ctx context.Context, Site string, Hostname string) (string, error) {
	args := map[string]interface{}{
		"site":     Site,
		"hostname": Hostname,
	}
	uri := "/api/v1/Building/Structure(getWithHostnameSite)"

	result := ""

	if err := service.cinp.Call(ctx, uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// BuildingComplex - Model Complex(/api/v1/Building/Complex)
/*
Complex(name, site, description, built_percentage, updated, created)
*/
type BuildingComplex struct {
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

// BuildingComplexNew - Make a new object of Model Complex
func (service *Contractor) BuildingComplexNew() *BuildingComplex {
	return &BuildingComplex{cinp: service.cinp}
}

// BuildingComplexNewWithID - Make a new object of Model Complex
func (service *Contractor) BuildingComplexNewWithID(id string) *BuildingComplex {
	result := BuildingComplex{cinp: service.cinp}
	result.SetURI("/api/v1/Building/Complex:" + id + ":")
	return &result
}

// BuildingComplexGet - Get function for Model Complex
func (service *Contractor) BuildingComplexGet(ctx context.Context, id string) (*BuildingComplex, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Building/Complex:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingComplex)
	result.cinp = service.cinp

	return result, nil
}

// BuildingComplexGetURI - Get function for Model Complex vi URI
func (service *Contractor) BuildingComplexGetURI(ctx context.Context, uri string) (*BuildingComplex, error) {
	if !strings.HasPrefix(uri, "/api/v1/Building/Complex:") {
		return nil, fmt.Errorf("URI is not for a 'BuildingComplex'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingComplex)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Complex
func (object *BuildingComplex) Create(ctx context.Context) (*BuildingComplex, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Building/Complex", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*BuildingComplex), nil
}

// Update - Update function for Model Complex
func (object *BuildingComplex) Update(ctx context.Context) (*BuildingComplex, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*BuildingComplex), nil
}

// Delete - Delete function for Model Complex
func (object *BuildingComplex) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// BuildingComplexListFilters - Return a slice of valid filter names Complex
func (service *Contractor) BuildingComplexListFilters() [1]string {
	return [1]string{"site"}
}

// BuildingComplexList - List function for Model Complex
func (service *Contractor) BuildingComplexList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *BuildingComplex, error) {
	if filterName != "" {
		for _, item := range service.BuildingComplexListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Building/Complex", reflect.TypeOf(BuildingComplex{}), filterName, filterValues, 50)
	out := make(chan *BuildingComplex)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*BuildingComplex).cinp = service.cinp
			out <- (*v).(*BuildingComplex)
		}
	}()
	return out, nil
}

// CallCreateFoundation calls createFoundation
func (object *BuildingComplex) CallCreateFoundation(ctx context.Context, Hostname string, Site string, InterfaceMapList []map[string]interface{}) (string, error) {
	args := map[string]interface{}{
		"hostname":           Hostname,
		"site":               Site,
		"interface_map_list": InterfaceMapList,
	}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Complex(createFoundation)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// BuildingComplexStructure - Model ComplexStructure(/api/v1/Building/ComplexStructure)
/*
ComplexStructure(id, complex, structure, updated, created)
*/
type BuildingComplexStructure struct {
	cinp.BaseObject
	cinp      cinp.CInPClient `json:"-"`
	Complex   *string         `json:"complex,omitempty"`
	Structure *string         `json:"structure,omitempty"`
	Updated   *time.Time      `json:"updated,omitempty"`
	Created   *time.Time      `json:"created,omitempty"`
	ID        *int            `json:"id,omitempty"`
}

// BuildingComplexStructureNew - Make a new object of Model ComplexStructure
func (service *Contractor) BuildingComplexStructureNew() *BuildingComplexStructure {
	return &BuildingComplexStructure{cinp: service.cinp}
}

// BuildingComplexStructureNewWithID - Make a new object of Model ComplexStructure
func (service *Contractor) BuildingComplexStructureNewWithID(id int) *BuildingComplexStructure {
	result := BuildingComplexStructure{cinp: service.cinp}
	result.SetURI("/api/v1/Building/ComplexStructure:" + strconv.FormatInt(int64(id), 10) + ":")
	return &result
}

// BuildingComplexStructureGet - Get function for Model ComplexStructure
func (service *Contractor) BuildingComplexStructureGet(ctx context.Context, id int) (*BuildingComplexStructure, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Building/ComplexStructure:"+strconv.FormatInt(int64(id), 10)+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingComplexStructure)
	result.cinp = service.cinp

	return result, nil
}

// BuildingComplexStructureGetURI - Get function for Model ComplexStructure vi URI
func (service *Contractor) BuildingComplexStructureGetURI(ctx context.Context, uri string) (*BuildingComplexStructure, error) {
	if !strings.HasPrefix(uri, "/api/v1/Building/ComplexStructure:") {
		return nil, fmt.Errorf("URI is not for a 'BuildingComplexStructure'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingComplexStructure)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model ComplexStructure
func (object *BuildingComplexStructure) Create(ctx context.Context) (*BuildingComplexStructure, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Building/ComplexStructure", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*BuildingComplexStructure), nil
}

// Update - Update function for Model ComplexStructure
func (object *BuildingComplexStructure) Update(ctx context.Context) (*BuildingComplexStructure, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*BuildingComplexStructure), nil
}

// Delete - Delete function for Model ComplexStructure
func (object *BuildingComplexStructure) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// BuildingComplexStructureListFilters - Return a slice of valid filter names ComplexStructure
func (service *Contractor) BuildingComplexStructureListFilters() [1]string {
	return [1]string{"complex"}
}

// BuildingComplexStructureList - List function for Model ComplexStructure
func (service *Contractor) BuildingComplexStructureList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *BuildingComplexStructure, error) {
	if filterName != "" {
		for _, item := range service.BuildingComplexStructureListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Building/ComplexStructure", reflect.TypeOf(BuildingComplexStructure{}), filterName, filterValues, 50)
	out := make(chan *BuildingComplexStructure)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*BuildingComplexStructure).cinp = service.cinp
			out <- (*v).(*BuildingComplexStructure)
		}
	}()
	return out, nil
}

// CallGetConfig calls getConfig
func (object *BuildingComplexStructure) CallGetConfig(ctx context.Context) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/ComplexStructure(getConfig)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// BuildingDependency - Model Dependency(/api/v1/Building/Dependency)
/*
Dependency(id, structure, dependency, foundation, script_structure, link, create_script_name, destroy_script_name, built_at, updated, created)
*/
type BuildingDependency struct {
	cinp.BaseObject
	cinp              cinp.CInPClient `json:"-"`
	Structure         *string         `json:"structure,omitempty"`
	Dependency        *string         `json:"dependency,omitempty"`
	Foundation        *string         `json:"foundation,omitempty"`
	ScriptStructure   *string         `json:"script_structure,omitempty"`
	Link              *string         `json:"link,omitempty"`
	CreateScriptName  *string         `json:"create_script_name,omitempty"`
	DestroyScriptName *string         `json:"destroy_script_name,omitempty"`
	BuiltAt           *time.Time      `json:"built_at,omitempty"`
	Updated           *time.Time      `json:"updated,omitempty"`
	Created           *time.Time      `json:"created,omitempty"`
	ID                *int            `json:"id,omitempty"`
	State             *string         `json:"state,omitempty"`
}

// BuildingDependencyNew - Make a new object of Model Dependency
func (service *Contractor) BuildingDependencyNew() *BuildingDependency {
	return &BuildingDependency{cinp: service.cinp}
}

// BuildingDependencyNewWithID - Make a new object of Model Dependency
func (service *Contractor) BuildingDependencyNewWithID(id int) *BuildingDependency {
	result := BuildingDependency{cinp: service.cinp}
	result.SetURI("/api/v1/Building/Dependency:" + strconv.FormatInt(int64(id), 10) + ":")
	return &result
}

// BuildingDependencyGet - Get function for Model Dependency
func (service *Contractor) BuildingDependencyGet(ctx context.Context, id int) (*BuildingDependency, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Building/Dependency:"+strconv.FormatInt(int64(id), 10)+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingDependency)
	result.cinp = service.cinp

	return result, nil
}

// BuildingDependencyGetURI - Get function for Model Dependency vi URI
func (service *Contractor) BuildingDependencyGetURI(ctx context.Context, uri string) (*BuildingDependency, error) {
	if !strings.HasPrefix(uri, "/api/v1/Building/Dependency:") {
		return nil, fmt.Errorf("URI is not for a 'BuildingDependency'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*BuildingDependency)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Dependency
func (object *BuildingDependency) Create(ctx context.Context) (*BuildingDependency, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Building/Dependency", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*BuildingDependency), nil
}

// Update - Update function for Model Dependency
func (object *BuildingDependency) Update(ctx context.Context) (*BuildingDependency, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*BuildingDependency), nil
}

// Delete - Delete function for Model Dependency
func (object *BuildingDependency) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// BuildingDependencyListFilters - Return a slice of valid filter names Dependency
func (service *Contractor) BuildingDependencyListFilters() [2]string {
	return [2]string{"foundation", "site"}
}

// BuildingDependencyList - List function for Model Dependency
func (service *Contractor) BuildingDependencyList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *BuildingDependency, error) {
	if filterName != "" {
		for _, item := range service.BuildingDependencyListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Building/Dependency", reflect.TypeOf(BuildingDependency{}), filterName, filterValues, 50)
	out := make(chan *BuildingDependency)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*BuildingDependency).cinp = service.cinp
			out <- (*v).(*BuildingDependency)
		}
	}()
	return out, nil
}

// CallGetJob calls getJob
func (object *BuildingDependency) CallGetJob(ctx context.Context) (string, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Building/Dependency(getJob)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}
func registerBuilding(cinp cinp.CInPClient) {
	cinp.RegisterType("/api/v1/Building/Foundation", reflect.TypeOf((*BuildingFoundation)(nil)).Elem())
	cinp.RegisterType("/api/v1/Building/Structure", reflect.TypeOf((*BuildingStructure)(nil)).Elem())
	cinp.RegisterType("/api/v1/Building/Complex", reflect.TypeOf((*BuildingComplex)(nil)).Elem())
	cinp.RegisterType("/api/v1/Building/ComplexStructure", reflect.TypeOf((*BuildingComplexStructure)(nil)).Elem())
	cinp.RegisterType("/api/v1/Building/Dependency", reflect.TypeOf((*BuildingDependency)(nil)).Elem())
}
