// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Records/ at 2024-03-26T14:17:37.955685
package contractor

import (
	cinp "github.com/cinp/go"
	"reflect"
)

// RecordsRecorder - Model Recorder(/api/v1/Records/Recorder)
/*

 */
type RecordsRecorder struct {
	cinp.BaseObject
	cinp *cinp.CInP
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *RecordsRecorder) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{}
	}
	return &map[string]interface{}{}
}

// RecordsRecorderNew - Make a new object of Model Recorder
func (service *Contractor) RecordsRecorderNew() *RecordsRecorder {
	return &RecordsRecorder{cinp: service.cinp}
}

// RecordsRecorderCallQuery calls query
func (service *Contractor) RecordsRecorderCallQuery(Group string, Query string, Fields string, MaxResults int) ([]string, error) {
	args := map[string]interface{}{
		"group":       Group,
		"query":       Query,
		"fields":      Fields,
		"max_results": MaxResults,
	}
	uri := "/api/v1/Records/Recorder(query)"

	result := []string{}

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return []string{}, err
	}

	return result, nil
}

// RecordsRecorderCallQueryObjects calls queryObjects
func (service *Contractor) RecordsRecorderCallQueryObjects(Group string, Query string, MaxResults int) (string, error) {
	args := map[string]interface{}{
		"group":       Group,
		"query":       Query,
		"max_results": MaxResults,
	}
	uri := "/api/v1/Records/Recorder(queryObjects)"

	result := ""

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}
func registerRecords(cinp *cinp.CInP) {
	cinp.RegisterType("/api/v1/Records/Recorder", reflect.TypeOf((*RecordsRecorder)(nil)).Elem())
}
