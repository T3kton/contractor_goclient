/*Package contractor(version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Records/ at 2020-07-29T04:55:03.557628
 */
package contractor

import (
	"reflect"
	cinp "github.com/cinp/go"
)

//RecordsRecorder - Model Recorder(/api/v1/Records/Recorder)
/*

 */
type RecordsRecorder struct {
	cinp.BaseObject
	cinp *cinp.CInP
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *RecordsRecorder) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
		}
	}
	return &map[string]interface{}{ 
	}
}

// RecordsRecorderNew - Make a new object of Model Recorder
func (service *Contractor) RecordsRecorderNew() *RecordsRecorder {
	return &RecordsRecorder{cinp: service.cinp}
}

// RecordsRecorderNewWithID - Make a new object of Model Recorder
func (service *Contractor) RecordsRecorderNewWithID(id string) *RecordsRecorder {
	result := RecordsRecorder{cinp: service.cinp}
	result.SetID("/api/v1/Records/Recorder:"+id+":")
	return &result
}


// RecordsRecorderCallQuery calls queryNoneNoneNoneNone
func (service *Contractor) RecordsRecorderCallQuery(group string, query string, fields string, max_results int) ([]string, error) {
	args := map[string]interface{}{
		"group": group,
		"query": query,
		"fields": fields,
		"max_results": max_results,
	}
	uri := "/api/v1/Records/Recorder(query)"

	result := []string{}

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return []string{}, err
	}

	return result, nil
}

// RecordsRecorderCallQueryObjects calls queryObjectsNoneNoneNone
func (service *Contractor) RecordsRecorderCallQueryObjects(group string, query string, max_results int) (string, error) {
	args := map[string]interface{}{
		"group": group,
		"query": query,
		"max_results": max_results,
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