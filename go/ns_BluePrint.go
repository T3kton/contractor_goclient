/*Package contractor(version: "0.1") - Automatically generated by cinp-codegen from /api/v1/BluePrint/ at 2020-07-29T04:55:03.557628
 */
package contractor

import (
	"reflect"
	"time"
	cinp "github.com/cinp/go"
)

//BlueprintBluePrint - Model BluePrint(/api/v1/BluePrint/BluePrint)
/*
BluePrint(name, description, config_values, updated, created)
 */
type BlueprintBluePrint struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Name string `json:"name"`
	Description string `json:"description"`
	ConfigValues map[string]interface{} `json:"config_values"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	ScriptMap map[string]interface{} `json:"script_map"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *BlueprintBluePrint) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"name": object.Name,
			"description": object.Description,
			"config_values": object.ConfigValues,
		}
	}
	return &map[string]interface{}{ 
		"description": object.Description,
		"config_values": object.ConfigValues,
	}
}

// BlueprintBluePrintNew - Make a new object of Model BluePrint
func (service *Contractor) BlueprintBluePrintNew() *BlueprintBluePrint {
	return &BlueprintBluePrint{cinp: service.cinp}
}

// BlueprintBluePrintNewWithID - Make a new object of Model BluePrint
func (service *Contractor) BlueprintBluePrintNewWithID(id string) *BlueprintBluePrint {
	result := BlueprintBluePrint{cinp: service.cinp}
	result.SetID("/api/v1/BluePrint/BluePrint:"+id+":")
	return &result
}


// CallGetConfig calls getConfig
func (object *BlueprintBluePrint) CallGetConfig() (map[string]interface{}, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/BluePrint/BluePrint(getConfig)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}


//BlueprintFoundationBluePrint - Model FoundationBluePrint(/api/v1/BluePrint/FoundationBluePrint)
/*
FoundationBluePrint(name, description, config_values, updated, created, blueprint_ptr, foundation_type_list, template, physical_interface_names)
 */
type BlueprintFoundationBluePrint struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Name string `json:"name"`
	Description string `json:"description"`
	ConfigValues map[string]interface{} `json:"config_values"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	FoundationTypeList []string `json:"foundation_type_list"`
	Template map[string]interface{} `json:"template"`
	PhysicalInterfaceNames []string `json:"physical_interface_names"`
	ParentList []string `json:"parent_list"`
	ScriptMap map[string]interface{} `json:"script_map"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *BlueprintFoundationBluePrint) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"name": object.Name,
			"description": object.Description,
			"config_values": object.ConfigValues,
			"foundation_type_list": object.FoundationTypeList,
			"template": object.Template,
			"physical_interface_names": object.PhysicalInterfaceNames,
			"parent_list": object.ParentList,
		}
	}
	return &map[string]interface{}{ 
		"description": object.Description,
		"config_values": object.ConfigValues,
		"foundation_type_list": object.FoundationTypeList,
		"template": object.Template,
		"physical_interface_names": object.PhysicalInterfaceNames,
		"parent_list": object.ParentList,
	}
}

// BlueprintFoundationBluePrintNew - Make a new object of Model FoundationBluePrint
func (service *Contractor) BlueprintFoundationBluePrintNew() *BlueprintFoundationBluePrint {
	return &BlueprintFoundationBluePrint{cinp: service.cinp}
}

// BlueprintFoundationBluePrintNewWithID - Make a new object of Model FoundationBluePrint
func (service *Contractor) BlueprintFoundationBluePrintNewWithID(id string) *BlueprintFoundationBluePrint {
	result := BlueprintFoundationBluePrint{cinp: service.cinp}
	result.SetID("/api/v1/BluePrint/FoundationBluePrint:"+id+":")
	return &result
}

// BlueprintFoundationBluePrintGet - Get function for Model FoundationBluePrint
func (service *Contractor) BlueprintFoundationBluePrintGet(id string) (*BlueprintFoundationBluePrint, error) {
	object, err := service.cinp.Get("/api/v1/BluePrint/FoundationBluePrint:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BlueprintFoundationBluePrint)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model FoundationBluePrint
func (object *BlueprintFoundationBluePrint) Create() error {
	if err := object.cinp.Create("/api/v1/BluePrint/FoundationBluePrint", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model FoundationBluePrint
func (object *BlueprintFoundationBluePrint) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model FoundationBluePrint
func (object *BlueprintFoundationBluePrint) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// BlueprintFoundationBluePrintList - List function for Model FoundationBluePrint
func (service *Contractor) BlueprintFoundationBluePrintList(filterName string, filterValues map[string]interface{}) <-chan *BlueprintFoundationBluePrint {
	in := service.cinp.ListObjects("/api/v1/BluePrint/FoundationBluePrint", reflect.TypeOf(BlueprintFoundationBluePrint{}), filterName, filterValues, 50)
	out := make(chan *BlueprintFoundationBluePrint)
	go func() {
		defer close(out)
		for v := range in {
			v.(*BlueprintFoundationBluePrint).cinp = service.cinp
			out <- v.(*BlueprintFoundationBluePrint)
		}
	}()
	return out
}


// CallGetConfig calls getConfig
func (object *BlueprintFoundationBluePrint) CallGetConfig() (map[string]interface{}, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/BluePrint/FoundationBluePrint(getConfig)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}


//BlueprintStructureBluePrint - Model StructureBluePrint(/api/v1/BluePrint/StructureBluePrint)
/*
StructureBluePrint(name, description, config_values, updated, created, blueprint_ptr)
 */
type BlueprintStructureBluePrint struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Name string `json:"name"`
	Description string `json:"description"`
	ConfigValues map[string]interface{} `json:"config_values"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	ParentList []string `json:"parent_list"`
	FoundationBlueprintList []string `json:"foundation_blueprint_list"`
	ScriptMap map[string]interface{} `json:"script_map"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *BlueprintStructureBluePrint) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"name": object.Name,
			"description": object.Description,
			"config_values": object.ConfigValues,
			"parent_list": object.ParentList,
			"foundation_blueprint_list": object.FoundationBlueprintList,
		}
	}
	return &map[string]interface{}{ 
		"description": object.Description,
		"config_values": object.ConfigValues,
		"parent_list": object.ParentList,
		"foundation_blueprint_list": object.FoundationBlueprintList,
	}
}

// BlueprintStructureBluePrintNew - Make a new object of Model StructureBluePrint
func (service *Contractor) BlueprintStructureBluePrintNew() *BlueprintStructureBluePrint {
	return &BlueprintStructureBluePrint{cinp: service.cinp}
}

// BlueprintStructureBluePrintNewWithID - Make a new object of Model StructureBluePrint
func (service *Contractor) BlueprintStructureBluePrintNewWithID(id string) *BlueprintStructureBluePrint {
	result := BlueprintStructureBluePrint{cinp: service.cinp}
	result.SetID("/api/v1/BluePrint/StructureBluePrint:"+id+":")
	return &result
}

// BlueprintStructureBluePrintGet - Get function for Model StructureBluePrint
func (service *Contractor) BlueprintStructureBluePrintGet(id string) (*BlueprintStructureBluePrint, error) {
	object, err := service.cinp.Get("/api/v1/BluePrint/StructureBluePrint:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BlueprintStructureBluePrint)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model StructureBluePrint
func (object *BlueprintStructureBluePrint) Create() error {
	if err := object.cinp.Create("/api/v1/BluePrint/StructureBluePrint", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model StructureBluePrint
func (object *BlueprintStructureBluePrint) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model StructureBluePrint
func (object *BlueprintStructureBluePrint) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// BlueprintStructureBluePrintList - List function for Model StructureBluePrint
func (service *Contractor) BlueprintStructureBluePrintList(filterName string, filterValues map[string]interface{}) <-chan *BlueprintStructureBluePrint {
	in := service.cinp.ListObjects("/api/v1/BluePrint/StructureBluePrint", reflect.TypeOf(BlueprintStructureBluePrint{}), filterName, filterValues, 50)
	out := make(chan *BlueprintStructureBluePrint)
	go func() {
		defer close(out)
		for v := range in {
			v.(*BlueprintStructureBluePrint).cinp = service.cinp
			out <- v.(*BlueprintStructureBluePrint)
		}
	}()
	return out
}


// CallGetConfig calls getConfig
func (object *BlueprintStructureBluePrint) CallGetConfig() (map[string]interface{}, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/BluePrint/StructureBluePrint(getConfig)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}


//BlueprintScript - Model Script(/api/v1/BluePrint/Script)
/*
Script(name, description, script, updated, created)
 */
type BlueprintScript struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Name string `json:"name"`
	Description string `json:"description"`
	Script string `json:"script"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *BlueprintScript) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"name": object.Name,
			"description": object.Description,
			"script": object.Script,
		}
	}
	return &map[string]interface{}{ 
		"description": object.Description,
		"script": object.Script,
	}
}

// BlueprintScriptNew - Make a new object of Model Script
func (service *Contractor) BlueprintScriptNew() *BlueprintScript {
	return &BlueprintScript{cinp: service.cinp}
}

// BlueprintScriptNewWithID - Make a new object of Model Script
func (service *Contractor) BlueprintScriptNewWithID(id string) *BlueprintScript {
	result := BlueprintScript{cinp: service.cinp}
	result.SetID("/api/v1/BluePrint/Script:"+id+":")
	return &result
}

// BlueprintScriptGet - Get function for Model Script
func (service *Contractor) BlueprintScriptGet(id string) (*BlueprintScript, error) {
	object, err := service.cinp.Get("/api/v1/BluePrint/Script:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BlueprintScript)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Script
func (object *BlueprintScript) Create() error {
	if err := object.cinp.Create("/api/v1/BluePrint/Script", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model Script
func (object *BlueprintScript) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model Script
func (object *BlueprintScript) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// BlueprintScriptList - List function for Model Script
func (service *Contractor) BlueprintScriptList(filterName string, filterValues map[string]interface{}) <-chan *BlueprintScript {
	in := service.cinp.ListObjects("/api/v1/BluePrint/Script", reflect.TypeOf(BlueprintScript{}), filterName, filterValues, 50)
	out := make(chan *BlueprintScript)
	go func() {
		defer close(out)
		for v := range in {
			v.(*BlueprintScript).cinp = service.cinp
			out <- v.(*BlueprintScript)
		}
	}()
	return out
}



//BlueprintBluePrintScript - Model BluePrintScript(/api/v1/BluePrint/BluePrintScript)
/*
BluePrintScript(id, blueprint, script, name, updated, created)
 */
type BlueprintBluePrintScript struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Blueprint string `json:"blueprint"`
	Script string `json:"script"`
	Name string `json:"name"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *BlueprintBluePrintScript) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"blueprint": object.Blueprint,
			"script": object.Script,
			"name": object.Name,
		}
	}
	return &map[string]interface{}{ 
		"blueprint": object.Blueprint,
		"script": object.Script,
		"name": object.Name,
	}
}

// BlueprintBluePrintScriptNew - Make a new object of Model BluePrintScript
func (service *Contractor) BlueprintBluePrintScriptNew() *BlueprintBluePrintScript {
	return &BlueprintBluePrintScript{cinp: service.cinp}
}

// BlueprintBluePrintScriptNewWithID - Make a new object of Model BluePrintScript
func (service *Contractor) BlueprintBluePrintScriptNewWithID(id string) *BlueprintBluePrintScript {
	result := BlueprintBluePrintScript{cinp: service.cinp}
	result.SetID("/api/v1/BluePrint/BluePrintScript:"+id+":")
	return &result
}

// BlueprintBluePrintScriptGet - Get function for Model BluePrintScript
func (service *Contractor) BlueprintBluePrintScriptGet(id string) (*BlueprintBluePrintScript, error) {
	object, err := service.cinp.Get("/api/v1/BluePrint/BluePrintScript:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BlueprintBluePrintScript)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model BluePrintScript
func (object *BlueprintBluePrintScript) Create() error {
	if err := object.cinp.Create("/api/v1/BluePrint/BluePrintScript", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model BluePrintScript
func (object *BlueprintBluePrintScript) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model BluePrintScript
func (object *BlueprintBluePrintScript) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// BlueprintBluePrintScriptList - List function for Model BluePrintScript
func (service *Contractor) BlueprintBluePrintScriptList(filterName string, filterValues map[string]interface{}) <-chan *BlueprintBluePrintScript {
	in := service.cinp.ListObjects("/api/v1/BluePrint/BluePrintScript", reflect.TypeOf(BlueprintBluePrintScript{}), filterName, filterValues, 50)
	out := make(chan *BlueprintBluePrintScript)
	go func() {
		defer close(out)
		for v := range in {
			v.(*BlueprintBluePrintScript).cinp = service.cinp
			out <- v.(*BlueprintBluePrintScript)
		}
	}()
	return out
}



//BlueprintPXE - Model PXE(/api/v1/BluePrint/PXE)
/*
PXE(name, boot_script, template, updated, created)
 */
type BlueprintPXE struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Name string `json:"name"`
	BootScript string `json:"boot_script"`
	Template string `json:"template"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *BlueprintPXE) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"name": object.Name,
			"boot_script": object.BootScript,
			"template": object.Template,
		}
	}
	return &map[string]interface{}{ 
		"boot_script": object.BootScript,
		"template": object.Template,
	}
}

// BlueprintPXENew - Make a new object of Model PXE
func (service *Contractor) BlueprintPXENew() *BlueprintPXE {
	return &BlueprintPXE{cinp: service.cinp}
}

// BlueprintPXENewWithID - Make a new object of Model PXE
func (service *Contractor) BlueprintPXENewWithID(id string) *BlueprintPXE {
	result := BlueprintPXE{cinp: service.cinp}
	result.SetID("/api/v1/BluePrint/PXE:"+id+":")
	return &result
}

// BlueprintPXEGet - Get function for Model PXE
func (service *Contractor) BlueprintPXEGet(id string) (*BlueprintPXE, error) {
	object, err := service.cinp.Get("/api/v1/BluePrint/PXE:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*BlueprintPXE)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model PXE
func (object *BlueprintPXE) Create() error {
	if err := object.cinp.Create("/api/v1/BluePrint/PXE", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model PXE
func (object *BlueprintPXE) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model PXE
func (object *BlueprintPXE) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// BlueprintPXEList - List function for Model PXE
func (service *Contractor) BlueprintPXEList(filterName string, filterValues map[string]interface{}) <-chan *BlueprintPXE {
	in := service.cinp.ListObjects("/api/v1/BluePrint/PXE", reflect.TypeOf(BlueprintPXE{}), filterName, filterValues, 50)
	out := make(chan *BlueprintPXE)
	go func() {
		defer close(out)
		for v := range in {
			v.(*BlueprintPXE).cinp = service.cinp
			out <- v.(*BlueprintPXE)
		}
	}()
	return out
}


func registerBluePrint(cinp *cinp.CInP) { 
	cinp.RegisterType("/api/v1/BluePrint/BluePrint", reflect.TypeOf((*BlueprintBluePrint)(nil)).Elem())
	cinp.RegisterType("/api/v1/BluePrint/FoundationBluePrint", reflect.TypeOf((*BlueprintFoundationBluePrint)(nil)).Elem())
	cinp.RegisterType("/api/v1/BluePrint/StructureBluePrint", reflect.TypeOf((*BlueprintStructureBluePrint)(nil)).Elem())
	cinp.RegisterType("/api/v1/BluePrint/Script", reflect.TypeOf((*BlueprintScript)(nil)).Elem())
	cinp.RegisterType("/api/v1/BluePrint/BluePrintScript", reflect.TypeOf((*BlueprintBluePrintScript)(nil)).Elem())
	cinp.RegisterType("/api/v1/BluePrint/PXE", reflect.TypeOf((*BlueprintPXE)(nil)).Elem())
}