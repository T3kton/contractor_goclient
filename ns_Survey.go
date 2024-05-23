// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Survey/ at 2024-05-23T12:38:27.938159
package contractor

import (
	"context"
	"fmt"
	cinp "github.com/cinp/go"
	"reflect"
	"strings"
	"time"
)

// SurveyPlot - Model Plot(/api/v1/Survey/Plot)
/*
Plot(name, corners, parent, updated, created)
*/
type SurveyPlot struct {
	cinp.BaseObject
	cinp    *cinp.CInP `json:"-"`
	Name    *string    `json:"name,omitempty"`
	Corners *string    `json:"corners,omitempty"`
	Parent  *string    `json:"parent,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Created *time.Time `json:"created,omitempty"`
}

// SurveyPlotNew - Make a new object of Model Plot
func (service *Contractor) SurveyPlotNew() *SurveyPlot {
	return &SurveyPlot{cinp: service.cinp}
}

// SurveyPlotNewWithID - Make a new object of Model Plot
func (service *Contractor) SurveyPlotNewWithID(id string) *SurveyPlot {
	result := SurveyPlot{cinp: service.cinp}
	result.SetURI("/api/v1/Survey/Plot:" + id + ":")
	return &result
}

// SurveyPlotGet - Get function for Model Plot
func (service *Contractor) SurveyPlotGet(ctx context.Context, id string) (*SurveyPlot, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Survey/Plot:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*SurveyPlot)
	result.cinp = service.cinp

	return result, nil
}

// SurveyPlotGetURI - Get function for Model Plot vi URI
func (service *Contractor) SurveyPlotGetURI(ctx context.Context, uri string) (*SurveyPlot, error) {
	if !strings.HasPrefix(uri, "/api/v1/Survey/Plot:") {
		return nil, fmt.Errorf("URI is not for a 'SurveyPlot'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*SurveyPlot)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Plot
func (object *SurveyPlot) Create(ctx context.Context) (*SurveyPlot, error) {
	result, err := object.cinp.Create(ctx, "/api/v1/Survey/Plot", object)
	if err != nil {
		return nil, err
	}

	return (*result).(*SurveyPlot), nil
}

// Update - Update function for Model Plot
func (object *SurveyPlot) Update(ctx context.Context) (*SurveyPlot, error) {
	result, err := object.cinp.Update(ctx, object)
	if err != nil {
		return nil, err
	}

	return (*result).(*SurveyPlot), nil
}

// Delete - Delete function for Model Plot
func (object *SurveyPlot) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// SurveyPlotListFilters - Return a slice of valid filter names Plot
func (service *Contractor) SurveyPlotListFilters() [0]string {
	return [0]string{}
}

// SurveyPlotList - List function for Model Plot
func (service *Contractor) SurveyPlotList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *SurveyPlot, error) {
	if filterName != "" {
		for _, item := range service.SurveyPlotListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Survey/Plot", reflect.TypeOf(SurveyPlot{}), filterName, filterValues, 50)
	out := make(chan *SurveyPlot)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*SurveyPlot).cinp = service.cinp
			out <- (*v).(*SurveyPlot)
		}
	}()
	return out, nil
}

// SurveyCartographer - Model Cartographer(/api/v1/Survey/Cartographer)
/*
Cartographer(identifier, foundation, message, info_map, last_checkin, updated, created)
*/
type SurveyCartographer struct {
	cinp.BaseObject
	cinp        *cinp.CInP              `json:"-"`
	Identifier  *string                 `json:"identifier,omitempty"`
	Foundation  *string                 `json:"foundation,omitempty"`
	Message     *string                 `json:"message,omitempty"`
	InfoMap     *map[string]interface{} `json:"info_map,omitempty"`
	LastCheckin *time.Time              `json:"last_checkin,omitempty"`
	Updated     *time.Time              `json:"updated,omitempty"`
	Created     *time.Time              `json:"created,omitempty"`
}

// SurveyCartographerNew - Make a new object of Model Cartographer
func (service *Contractor) SurveyCartographerNew() *SurveyCartographer {
	return &SurveyCartographer{cinp: service.cinp}
}

// SurveyCartographerNewWithID - Make a new object of Model Cartographer
func (service *Contractor) SurveyCartographerNewWithID(id string) *SurveyCartographer {
	result := SurveyCartographer{cinp: service.cinp}
	result.SetURI("/api/v1/Survey/Cartographer:" + id + ":")
	return &result
}

// SurveyCartographerGet - Get function for Model Cartographer
func (service *Contractor) SurveyCartographerGet(ctx context.Context, id string) (*SurveyCartographer, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Survey/Cartographer:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*SurveyCartographer)
	result.cinp = service.cinp

	return result, nil
}

// SurveyCartographerGetURI - Get function for Model Cartographer vi URI
func (service *Contractor) SurveyCartographerGetURI(ctx context.Context, uri string) (*SurveyCartographer, error) {
	if !strings.HasPrefix(uri, "/api/v1/Survey/Cartographer:") {
		return nil, fmt.Errorf("URI is not for a 'SurveyCartographer'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*SurveyCartographer)
	result.cinp = service.cinp

	return result, nil
}

// Delete - Delete function for Model Cartographer
func (object *SurveyCartographer) Delete(ctx context.Context) error {
	if err := object.cinp.Delete(ctx, object); err != nil {
		return err
	}

	return nil
}

// SurveyCartographerListFilters - Return a slice of valid filter names Cartographer
func (service *Contractor) SurveyCartographerListFilters() [0]string {
	return [0]string{}
}

// SurveyCartographerList - List function for Model Cartographer
func (service *Contractor) SurveyCartographerList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *SurveyCartographer, error) {
	if filterName != "" {
		for _, item := range service.SurveyCartographerListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Survey/Cartographer", reflect.TypeOf(SurveyCartographer{}), filterName, filterValues, 50)
	out := make(chan *SurveyCartographer)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*SurveyCartographer).cinp = service.cinp
			out <- (*v).(*SurveyCartographer)
		}
	}()
	return out, nil
}

// CallAssign calls assign
func (object *SurveyCartographer) CallAssign(ctx context.Context, Foundation string) error {
	args := map[string]interface{}{
		"foundation": Foundation,
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Survey/Cartographer(assign)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// SurveyCartographerCallRegister calls register
func (service *Contractor) SurveyCartographerCallRegister(ctx context.Context, Identifier string) error {
	args := map[string]interface{}{
		"identifier": Identifier,
	}
	uri := "/api/v1/Survey/Cartographer(register)"

	result := ""

	if err := service.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallLookup calls lookup
func (object *SurveyCartographer) CallLookup(ctx context.Context, InfoMap map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{
		"info_map": InfoMap,
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Survey/Cartographer(lookup)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CallSetMessage calls setMessage
func (object *SurveyCartographer) CallSetMessage(ctx context.Context, Message string) error {
	args := map[string]interface{}{
		"message": Message,
	}
	_, _, _, ids, _, err := object.cinp.Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Survey/Cartographer(setMessage)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallDone calls done
func (object *SurveyCartographer) CallDone(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.UpdateIDs("/api/v1/Survey/Cartographer(done)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}
func registerSurvey(cinp *cinp.CInP) {
	cinp.RegisterType("/api/v1/Survey/Plot", reflect.TypeOf((*SurveyPlot)(nil)).Elem())
	cinp.RegisterType("/api/v1/Survey/Cartographer", reflect.TypeOf((*SurveyCartographer)(nil)).Elem())
}
