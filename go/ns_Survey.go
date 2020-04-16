/*Package contractor(version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Survey/ at 2020-04-16T20:18:01.673062
 */
package contractor

import (
	"reflect"
	"time"
	cinp "github.com/cinp/go"
)

//SurveyPlot - Model Plot(/api/v1/Survey/Plot)
/*
Plot(name, corners, parent, updated, created)
 */
type SurveyPlot struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Name string `json:"name"`
	Corners string `json:"corners"`
	Parent string `json:"parent"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *SurveyPlot) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"name": object.Name,
			"corners": object.Corners,
			"parent": object.Parent,
		}
	}
	return &map[string]interface{}{ 
		"corners": object.Corners,
		"parent": object.Parent,
	}
}

// SurveyPlotNew - Make a new object of Model Plot
func (service *Contractor) SurveyPlotNew() *SurveyPlot {
	return &SurveyPlot{cinp: service.cinp}
}

// SurveyPlotNewWithID - Make a new object of Model Plot
func (service *Contractor) SurveyPlotNewWithID(id string) *SurveyPlot {
	result := SurveyPlot{cinp: service.cinp}
	result.SetID("/api/v1/Survey/Plot:"+id+":")
	return &result
}

// SurveyPlotGet - Get function for Model Plot
func (service *Contractor) SurveyPlotGet(id string) (*SurveyPlot, error) {
	object, err := service.cinp.Get("/api/v1/Survey/Plot:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*SurveyPlot)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Plot
func (object *SurveyPlot) Create() error {
	if err := object.cinp.Create("/api/v1/Survey/Plot", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model Plot
func (object *SurveyPlot) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model Plot
func (object *SurveyPlot) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// SurveyPlotList - List function for Model Plot
func (service *Contractor) SurveyPlotList(filterName string, filterValues map[string]interface{}) <-chan *SurveyPlot {
	in := service.cinp.ListObjects("/api/v1/Survey/Plot", reflect.TypeOf(SurveyPlot{}), filterName, filterValues, 50)
	out := make(chan *SurveyPlot)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*SurveyPlot)
		}
	}()
	return out
}


//SurveyCartographer - Model Cartographer(/api/v1/Survey/Cartographer)
/*
Cartographer(identifier, foundation, message, updated, created)
 */
type SurveyCartographer struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Identifier string `json:"identifier"`
	Foundation string `json:"foundation"`
	Message string `json:"message"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *SurveyCartographer) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"identifier": object.Identifier,
			"foundation": object.Foundation,
			"message": object.Message,
		}
	}
	return &map[string]interface{}{ 
		"foundation": object.Foundation,
		"message": object.Message,
	}
}

// SurveyCartographerNew - Make a new object of Model Cartographer
func (service *Contractor) SurveyCartographerNew() *SurveyCartographer {
	return &SurveyCartographer{cinp: service.cinp}
}

// SurveyCartographerNewWithID - Make a new object of Model Cartographer
func (service *Contractor) SurveyCartographerNewWithID(id string) *SurveyCartographer {
	result := SurveyCartographer{cinp: service.cinp}
	result.SetID("/api/v1/Survey/Cartographer:"+id+":")
	return &result
}

// SurveyCartographerGet - Get function for Model Cartographer
func (service *Contractor) SurveyCartographerGet(id string) (*SurveyCartographer, error) {
	object, err := service.cinp.Get("/api/v1/Survey/Cartographer:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*SurveyCartographer)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Cartographer
func (object *SurveyCartographer) Create() error {
	if err := object.cinp.Create("/api/v1/Survey/Cartographer", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model Cartographer
func (object *SurveyCartographer) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model Cartographer
func (object *SurveyCartographer) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// SurveyCartographerList - List function for Model Cartographer
func (service *Contractor) SurveyCartographerList(filterName string, filterValues map[string]interface{}) <-chan *SurveyCartographer {
	in := service.cinp.ListObjects("/api/v1/Survey/Cartographer", reflect.TypeOf(SurveyCartographer{}), filterName, filterValues, 50)
	out := make(chan *SurveyCartographer)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*SurveyCartographer)
		}
	}()
	return out
}

// SurveyCartographerCallRegister calls registerNone
func (service *Contractor) SurveyCartographerCallRegister(identifier string) (string, error) {
	args := map[string]interface{}{
		"identifier": identifier,
	}
	uri := "/api/v1/Survey/Cartographer(register)"

	result := ""

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// CallLookup calls lookupNone
func (object *SurveyCartographer) CallLookup(info_map map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{
		"info_map": info_map,
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Survey/Cartographer(lookup)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CallSetMessage calls setMessageNone
func (object *SurveyCartographer) CallSetMessage(message string) (string, error) {
	args := map[string]interface{}{
		"message": message,
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Survey/Cartographer(setMessage)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// CallDone calls done
func (object *SurveyCartographer) CallDone() (string, error) {
	args := map[string]interface{}{
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetID())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Survey/Cartographer(done)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

func registerSurvey(cinp *cinp.CInP) { 
	cinp.RegisterType("/api/v1/Survey/Plot", reflect.TypeOf((*SurveyPlot)(nil)).Elem())
	cinp.RegisterType("/api/v1/Survey/Cartographer", reflect.TypeOf((*SurveyCartographer)(nil)).Elem())
}