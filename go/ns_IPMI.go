/*Package contractor(version: "0.1") - Automatically generated by cinp-codegen from /api/v1/IPMI/ at 2020-04-16T20:18:01.673062
 */
package contractor

import (
	"reflect"
	"time"
	cinp "github.com/cinp/go"
)

//IpmiIPMIFoundation - Model IPMIFoundation(/api/v1/IPMI/IPMIFoundation)
/*
IPMIFoundation(locator, site, blueprint, id_map, located_at, built_at, updated, created, foundation_ptr, ipmi_username, ipmi_password, ipmi_ip_address, plot)
 */
type IpmiIPMIFoundation struct {
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
	IpmiUsername string `json:"ipmi_username"`
	IpmiPassword string `json:"ipmi_password"`
	IpmiIPAddress string `json:"ipmi_ip_address"`
	Plot string `json:"plot"`
	State string `json:"state"`
	Type string `json:"type"`
	ClassList string `json:"class_list"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *IpmiIPMIFoundation) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"locator": object.Locator,
			"site": object.Site,
			"blueprint": object.Blueprint,
			"id_map": object.IDMap,
			"ipmi_username": object.IpmiUsername,
			"ipmi_password": object.IpmiPassword,
			"ipmi_ip_address": object.IpmiIPAddress,
			"plot": object.Plot,
		}
	}
	return &map[string]interface{}{ 
		"site": object.Site,
		"blueprint": object.Blueprint,
		"id_map": object.IDMap,
		"ipmi_username": object.IpmiUsername,
		"ipmi_password": object.IpmiPassword,
		"ipmi_ip_address": object.IpmiIPAddress,
		"plot": object.Plot,
	}
}

// IpmiIPMIFoundationNew - Make a new object of Model IPMIFoundation
func (service *Contractor) IpmiIPMIFoundationNew() *IpmiIPMIFoundation {
	return &IpmiIPMIFoundation{cinp: service.cinp}
}

// IpmiIPMIFoundationNewWithID - Make a new object of Model IPMIFoundation
func (service *Contractor) IpmiIPMIFoundationNewWithID(id string) *IpmiIPMIFoundation {
	result := IpmiIPMIFoundation{cinp: service.cinp}
	result.SetID("/api/v1/IPMI/IPMIFoundation:"+id+":")
	return &result
}

// IpmiIPMIFoundationGet - Get function for Model IPMIFoundation
func (service *Contractor) IpmiIPMIFoundationGet(id string) (*IpmiIPMIFoundation, error) {
	object, err := service.cinp.Get("/api/v1/IPMI/IPMIFoundation:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*IpmiIPMIFoundation)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model IPMIFoundation
func (object *IpmiIPMIFoundation) Create() error {
	if err := object.cinp.Create("/api/v1/IPMI/IPMIFoundation", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model IPMIFoundation
func (object *IpmiIPMIFoundation) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model IPMIFoundation
func (object *IpmiIPMIFoundation) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// IpmiIPMIFoundationList - List function for Model IPMIFoundation
func (service *Contractor) IpmiIPMIFoundationList(filterName string, filterValues map[string]interface{}) <-chan *IpmiIPMIFoundation {
	in := service.cinp.ListObjects("/api/v1/IPMI/IPMIFoundation", reflect.TypeOf(IpmiIPMIFoundation{}), filterName, filterValues, 50)
	out := make(chan *IpmiIPMIFoundation)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*IpmiIPMIFoundation)
		}
	}()
	return out
}

func registerIPMI(cinp *cinp.CInP) { 
	cinp.RegisterType("/api/v1/IPMI/IPMIFoundation", reflect.TypeOf((*IpmiIPMIFoundation)(nil)).Elem())
}