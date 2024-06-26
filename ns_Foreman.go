// Package contractor - (version: "0.1") - Automatically generated by cinp-codegen from /api/v1/Foreman/ at 2024-05-24T15:42:51.144790
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

// ForemanBaseJob - Model BaseJob(/api/v1/Foreman/BaseJob)
/*
BaseJob(id, site, state, status, message, script_runner, script_name, updated, created)
*/
type ForemanBaseJob struct {
	cinp.BaseObject
	cinp       cinp.CInPClient `json:"-"`
	Site       *string         `json:"site,omitempty"`
	State      *string         `json:"state,omitempty"`
	Status     *string         `json:"status,omitempty"`
	Message    *string         `json:"message,omitempty"`
	ScriptName *string         `json:"script_name,omitempty"`
	Updated    *time.Time      `json:"updated,omitempty"`
	Created    *time.Time      `json:"created,omitempty"`
	ID         *int            `json:"id,omitempty"`
	CanStart   *string         `json:"can_start,omitempty"`
}

// ForemanBaseJobNew - Make a new object of Model BaseJob
func (service *Contractor) ForemanBaseJobNew() *ForemanBaseJob {
	return &ForemanBaseJob{cinp: service.cinp}
}

// ForemanBaseJobNewWithID - Make a new object of Model BaseJob
func (service *Contractor) ForemanBaseJobNewWithID(id int) *ForemanBaseJob {
	result := ForemanBaseJob{cinp: service.cinp}
	result.SetURI("/api/v1/Foreman/BaseJob:" + strconv.FormatInt(int64(id), 10) + ":")
	return &result
}

// CallPause calls pause
func (object *ForemanBaseJob) CallPause(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/BaseJob(pause)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallResume calls resume
func (object *ForemanBaseJob) CallResume(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/BaseJob(resume)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallReset calls reset
func (object *ForemanBaseJob) CallReset(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/BaseJob(reset)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallRollback calls rollback
func (object *ForemanBaseJob) CallRollback(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/BaseJob(rollback)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallClearDispatched calls clearDispatched
func (object *ForemanBaseJob) CallClearDispatched(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/BaseJob(clearDispatched)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// ForemanBaseJobCallJobStats calls jobStats
func (service *Contractor) ForemanBaseJobCallJobStats(ctx context.Context, Site string) (map[string]interface{}, error) {
	args := map[string]interface{}{
		"site": Site,
	}
	uri := "/api/v1/Foreman/BaseJob(jobStats)"

	result := map[string]interface{}{}

	if err := service.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CallJobRunnerVariables calls jobRunnerVariables
func (object *ForemanBaseJob) CallJobRunnerVariables(ctx context.Context) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/BaseJob(jobRunnerVariables)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CallJobRunnerState calls jobRunnerState
func (object *ForemanBaseJob) CallJobRunnerState(ctx context.Context) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/BaseJob(jobRunnerState)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CallSignalComplete calls signalComplete
func (object *ForemanBaseJob) CallSignalComplete(ctx context.Context, Cookie string) (string, error) {
	args := map[string]interface{}{
		"cookie": Cookie,
	}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/BaseJob(signalComplete)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// CallSignalAlert calls signalAlert
func (object *ForemanBaseJob) CallSignalAlert(ctx context.Context, Msg string) (string, error) {
	args := map[string]interface{}{
		"msg": Msg,
	}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/BaseJob(signalAlert)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// CallPostMessage calls postMessage
func (object *ForemanBaseJob) CallPostMessage(ctx context.Context, Msg string) (string, error) {
	args := map[string]interface{}{
		"msg": Msg,
	}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return "", err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/BaseJob(postMessage)", ids)
	if err != nil {
		return "", err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// ForemanFoundationJob - Model FoundationJob(/api/v1/Foreman/FoundationJob)
/*
FoundationJob(id, site, state, status, message, script_runner, script_name, updated, created, basejob_ptr, foundation)
*/
type ForemanFoundationJob struct {
	cinp.BaseObject
	cinp       cinp.CInPClient `json:"-"`
	Site       *string         `json:"site,omitempty"`
	State      *string         `json:"state,omitempty"`
	Status     *string         `json:"status,omitempty"`
	Message    *string         `json:"message,omitempty"`
	ScriptName *string         `json:"script_name,omitempty"`
	Updated    *time.Time      `json:"updated,omitempty"`
	Created    *time.Time      `json:"created,omitempty"`
	Foundation *string         `json:"foundation,omitempty"`
	ID         *int            `json:"id,omitempty"`
	CanStart   *string         `json:"can_start,omitempty"`
}

// ForemanFoundationJobNew - Make a new object of Model FoundationJob
func (service *Contractor) ForemanFoundationJobNew() *ForemanFoundationJob {
	return &ForemanFoundationJob{cinp: service.cinp}
}

// ForemanFoundationJobNewWithID - Make a new object of Model FoundationJob
func (service *Contractor) ForemanFoundationJobNewWithID(id int) *ForemanFoundationJob {
	result := ForemanFoundationJob{cinp: service.cinp}
	result.SetURI("/api/v1/Foreman/FoundationJob:" + strconv.FormatInt(int64(id), 10) + ":")
	return &result
}

// ForemanFoundationJobGet - Get function for Model FoundationJob
func (service *Contractor) ForemanFoundationJobGet(ctx context.Context, id int) (*ForemanFoundationJob, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Foreman/FoundationJob:"+strconv.FormatInt(int64(id), 10)+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*ForemanFoundationJob)
	result.cinp = service.cinp

	return result, nil
}

// ForemanFoundationJobGetURI - Get function for Model FoundationJob vi URI
func (service *Contractor) ForemanFoundationJobGetURI(ctx context.Context, uri string) (*ForemanFoundationJob, error) {
	if !strings.HasPrefix(uri, "/api/v1/Foreman/FoundationJob:") {
		return nil, fmt.Errorf("URI is not for a 'ForemanFoundationJob'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*ForemanFoundationJob)
	result.cinp = service.cinp

	return result, nil
}

// ForemanFoundationJobListFilters - Return a slice of valid filter names FoundationJob
func (service *Contractor) ForemanFoundationJobListFilters() [1]string {
	return [1]string{"site"}
}

// ForemanFoundationJobList - List function for Model FoundationJob
func (service *Contractor) ForemanFoundationJobList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *ForemanFoundationJob, error) {
	if filterName != "" {
		for _, item := range service.ForemanFoundationJobListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Foreman/FoundationJob", reflect.TypeOf(ForemanFoundationJob{}), filterName, filterValues, 50)
	out := make(chan *ForemanFoundationJob)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*ForemanFoundationJob).cinp = service.cinp
			out <- (*v).(*ForemanFoundationJob)
		}
	}()
	return out, nil
}

// CallPause calls pause
func (object *ForemanFoundationJob) CallPause(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/FoundationJob(pause)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallResume calls resume
func (object *ForemanFoundationJob) CallResume(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/FoundationJob(resume)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallReset calls reset
func (object *ForemanFoundationJob) CallReset(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/FoundationJob(reset)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallRollback calls rollback
func (object *ForemanFoundationJob) CallRollback(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/FoundationJob(rollback)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallJobRunnerVariables calls jobRunnerVariables
func (object *ForemanFoundationJob) CallJobRunnerVariables(ctx context.Context) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/FoundationJob(jobRunnerVariables)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CallJobRunnerState calls jobRunnerState
func (object *ForemanFoundationJob) CallJobRunnerState(ctx context.Context) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/FoundationJob(jobRunnerState)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ForemanStructureJob - Model StructureJob(/api/v1/Foreman/StructureJob)
/*
StructureJob(id, site, state, status, message, script_runner, script_name, updated, created, basejob_ptr, structure)
*/
type ForemanStructureJob struct {
	cinp.BaseObject
	cinp       cinp.CInPClient `json:"-"`
	Site       *string         `json:"site,omitempty"`
	State      *string         `json:"state,omitempty"`
	Status     *string         `json:"status,omitempty"`
	Message    *string         `json:"message,omitempty"`
	ScriptName *string         `json:"script_name,omitempty"`
	Updated    *time.Time      `json:"updated,omitempty"`
	Created    *time.Time      `json:"created,omitempty"`
	Structure  *string         `json:"structure,omitempty"`
	ID         *int            `json:"id,omitempty"`
	CanStart   *string         `json:"can_start,omitempty"`
}

// ForemanStructureJobNew - Make a new object of Model StructureJob
func (service *Contractor) ForemanStructureJobNew() *ForemanStructureJob {
	return &ForemanStructureJob{cinp: service.cinp}
}

// ForemanStructureJobNewWithID - Make a new object of Model StructureJob
func (service *Contractor) ForemanStructureJobNewWithID(id int) *ForemanStructureJob {
	result := ForemanStructureJob{cinp: service.cinp}
	result.SetURI("/api/v1/Foreman/StructureJob:" + strconv.FormatInt(int64(id), 10) + ":")
	return &result
}

// ForemanStructureJobGet - Get function for Model StructureJob
func (service *Contractor) ForemanStructureJobGet(ctx context.Context, id int) (*ForemanStructureJob, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Foreman/StructureJob:"+strconv.FormatInt(int64(id), 10)+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*ForemanStructureJob)
	result.cinp = service.cinp

	return result, nil
}

// ForemanStructureJobGetURI - Get function for Model StructureJob vi URI
func (service *Contractor) ForemanStructureJobGetURI(ctx context.Context, uri string) (*ForemanStructureJob, error) {
	if !strings.HasPrefix(uri, "/api/v1/Foreman/StructureJob:") {
		return nil, fmt.Errorf("URI is not for a 'ForemanStructureJob'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*ForemanStructureJob)
	result.cinp = service.cinp

	return result, nil
}

// ForemanStructureJobListFilters - Return a slice of valid filter names StructureJob
func (service *Contractor) ForemanStructureJobListFilters() [1]string {
	return [1]string{"site"}
}

// ForemanStructureJobList - List function for Model StructureJob
func (service *Contractor) ForemanStructureJobList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *ForemanStructureJob, error) {
	if filterName != "" {
		for _, item := range service.ForemanStructureJobListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Foreman/StructureJob", reflect.TypeOf(ForemanStructureJob{}), filterName, filterValues, 50)
	out := make(chan *ForemanStructureJob)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*ForemanStructureJob).cinp = service.cinp
			out <- (*v).(*ForemanStructureJob)
		}
	}()
	return out, nil
}

// CallPause calls pause
func (object *ForemanStructureJob) CallPause(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/StructureJob(pause)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallResume calls resume
func (object *ForemanStructureJob) CallResume(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/StructureJob(resume)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallReset calls reset
func (object *ForemanStructureJob) CallReset(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/StructureJob(reset)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallRollback calls rollback
func (object *ForemanStructureJob) CallRollback(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/StructureJob(rollback)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallJobRunnerVariables calls jobRunnerVariables
func (object *ForemanStructureJob) CallJobRunnerVariables(ctx context.Context) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/StructureJob(jobRunnerVariables)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CallJobRunnerState calls jobRunnerState
func (object *ForemanStructureJob) CallJobRunnerState(ctx context.Context) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/StructureJob(jobRunnerState)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ForemanDependencyJob - Model DependencyJob(/api/v1/Foreman/DependencyJob)
/*
DependencyJob(id, site, state, status, message, script_runner, script_name, updated, created, basejob_ptr, dependency)
*/
type ForemanDependencyJob struct {
	cinp.BaseObject
	cinp       cinp.CInPClient `json:"-"`
	Site       *string         `json:"site,omitempty"`
	State      *string         `json:"state,omitempty"`
	Status     *string         `json:"status,omitempty"`
	Message    *string         `json:"message,omitempty"`
	ScriptName *string         `json:"script_name,omitempty"`
	Updated    *time.Time      `json:"updated,omitempty"`
	Created    *time.Time      `json:"created,omitempty"`
	Dependency *string         `json:"dependency,omitempty"`
	ID         *int            `json:"id,omitempty"`
	C          *string         `json:"c,omitempty"`
	A          *string         `json:"a,omitempty"`
	N          *string         `json:"n,omitempty"`
	*string    `json:"_,omitempty"`
	S          *string `json:"s,omitempty"`
	T          *string `json:"t,omitempty"`
	R          *string `json:"r,omitempty"`
}

// ForemanDependencyJobNew - Make a new object of Model DependencyJob
func (service *Contractor) ForemanDependencyJobNew() *ForemanDependencyJob {
	return &ForemanDependencyJob{cinp: service.cinp}
}

// ForemanDependencyJobNewWithID - Make a new object of Model DependencyJob
func (service *Contractor) ForemanDependencyJobNewWithID(id int) *ForemanDependencyJob {
	result := ForemanDependencyJob{cinp: service.cinp}
	result.SetURI("/api/v1/Foreman/DependencyJob:" + strconv.FormatInt(int64(id), 10) + ":")
	return &result
}

// ForemanDependencyJobGet - Get function for Model DependencyJob
func (service *Contractor) ForemanDependencyJobGet(ctx context.Context, id int) (*ForemanDependencyJob, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Foreman/DependencyJob:"+strconv.FormatInt(int64(id), 10)+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*ForemanDependencyJob)
	result.cinp = service.cinp

	return result, nil
}

// ForemanDependencyJobGetURI - Get function for Model DependencyJob vi URI
func (service *Contractor) ForemanDependencyJobGetURI(ctx context.Context, uri string) (*ForemanDependencyJob, error) {
	if !strings.HasPrefix(uri, "/api/v1/Foreman/DependencyJob:") {
		return nil, fmt.Errorf("URI is not for a 'ForemanDependencyJob'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*ForemanDependencyJob)
	result.cinp = service.cinp

	return result, nil
}

// ForemanDependencyJobListFilters - Return a slice of valid filter names DependencyJob
func (service *Contractor) ForemanDependencyJobListFilters() [1]string {
	return [1]string{"site"}
}

// ForemanDependencyJobList - List function for Model DependencyJob
func (service *Contractor) ForemanDependencyJobList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *ForemanDependencyJob, error) {
	if filterName != "" {
		for _, item := range service.ForemanDependencyJobListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Foreman/DependencyJob", reflect.TypeOf(ForemanDependencyJob{}), filterName, filterValues, 50)
	out := make(chan *ForemanDependencyJob)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*ForemanDependencyJob).cinp = service.cinp
			out <- (*v).(*ForemanDependencyJob)
		}
	}()
	return out, nil
}

// CallPause calls pause
func (object *ForemanDependencyJob) CallPause(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/DependencyJob(pause)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallResume calls resume
func (object *ForemanDependencyJob) CallResume(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/DependencyJob(resume)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallReset calls reset
func (object *ForemanDependencyJob) CallReset(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/DependencyJob(reset)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallRollback calls rollback
func (object *ForemanDependencyJob) CallRollback(ctx context.Context) error {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/DependencyJob(rollback)", ids)
	if err != nil {
		return err
	}

	result := ""

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// CallJobRunnerVariables calls jobRunnerVariables
func (object *ForemanDependencyJob) CallJobRunnerVariables(ctx context.Context) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/DependencyJob(jobRunnerVariables)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CallJobRunnerState calls jobRunnerState
func (object *ForemanDependencyJob) CallJobRunnerState(ctx context.Context) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	_, _, _, ids, _, err := object.cinp.GetURI().Split(object.GetURI())
	if err != nil {
		return nil, err
	}
	uri, err := object.cinp.GetURI().UpdateIDs("/api/v1/Foreman/DependencyJob(jobRunnerState)", ids)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}

	if err := object.cinp.Call(ctx, uri, &args, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ForemanJobLog - Model JobLog(/api/v1/Foreman/JobLog)
/*
JobLog(id, site, job_id, target_id, target_class, target_description, script_name, creator, started_at, finished_at, canceled_by, canceled_at, updated, created)
*/
type ForemanJobLog struct {
	cinp.BaseObject
	cinp              cinp.CInPClient `json:"-"`
	Site              *string         `json:"site,omitempty"`
	JobID             *int            `json:"job_id,omitempty"`
	TargetID          *string         `json:"target_id,omitempty"`
	TargetClass       *string         `json:"target_class,omitempty"`
	TargetDescription *string         `json:"target_description,omitempty"`
	ScriptName        *string         `json:"script_name,omitempty"`
	Creator           *string         `json:"creator,omitempty"`
	StartedAt         *time.Time      `json:"started_at,omitempty"`
	FinishedAt        *time.Time      `json:"finished_at,omitempty"`
	CanceledBy        *string         `json:"canceled_by,omitempty"`
	CanceledAt        *time.Time      `json:"canceled_at,omitempty"`
	Updated           *time.Time      `json:"updated,omitempty"`
	Created           *time.Time      `json:"created,omitempty"`
	ID                *int            `json:"id,omitempty"`
}

// ForemanJobLogNew - Make a new object of Model JobLog
func (service *Contractor) ForemanJobLogNew() *ForemanJobLog {
	return &ForemanJobLog{cinp: service.cinp}
}

// ForemanJobLogNewWithID - Make a new object of Model JobLog
func (service *Contractor) ForemanJobLogNewWithID(id int) *ForemanJobLog {
	result := ForemanJobLog{cinp: service.cinp}
	result.SetURI("/api/v1/Foreman/JobLog:" + strconv.FormatInt(int64(id), 10) + ":")
	return &result
}

// ForemanJobLogGet - Get function for Model JobLog
func (service *Contractor) ForemanJobLogGet(ctx context.Context, id int) (*ForemanJobLog, error) {
	object, err := service.cinp.Get(ctx, "/api/v1/Foreman/JobLog:"+strconv.FormatInt(int64(id), 10)+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*ForemanJobLog)
	result.cinp = service.cinp

	return result, nil
}

// ForemanJobLogGetURI - Get function for Model JobLog vi URI
func (service *Contractor) ForemanJobLogGetURI(ctx context.Context, uri string) (*ForemanJobLog, error) {
	if !strings.HasPrefix(uri, "/api/v1/Foreman/JobLog:") {
		return nil, fmt.Errorf("URI is not for a 'ForemanJobLog'")
	}

	object, err := service.cinp.Get(ctx, uri)
	if err != nil {
		return nil, err
	}
	result := (*object).(*ForemanJobLog)
	result.cinp = service.cinp

	return result, nil
}

// ForemanJobLogListFilters - Return a slice of valid filter names JobLog
func (service *Contractor) ForemanJobLogListFilters() [5]string {
	return [5]string{"site", "job", "structure", "foundation", "dependency"}
}

// ForemanJobLogList - List function for Model JobLog
func (service *Contractor) ForemanJobLogList(ctx context.Context, filterName string, filterValues map[string]interface{}) (<-chan *ForemanJobLog, error) {
	if filterName != "" {
		for _, item := range service.ForemanJobLogListFilters() {
			if item == filterName {
				goto good
			}
		}
		return nil, fmt.Errorf("filter '%s' is invalid", filterName)
	}
good:

	in := service.cinp.ListObjects(ctx, "/api/v1/Foreman/JobLog", reflect.TypeOf(ForemanJobLog{}), filterName, filterValues, 50)
	out := make(chan *ForemanJobLog)
	go func() {
		defer close(out)
		for v := range in {
			(*v).(*ForemanJobLog).cinp = service.cinp
			out <- (*v).(*ForemanJobLog)
		}
	}()
	return out, nil
}
func registerForeman(cinp cinp.CInPClient) {
	cinp.RegisterType("/api/v1/Foreman/BaseJob", reflect.TypeOf((*ForemanBaseJob)(nil)).Elem())
	cinp.RegisterType("/api/v1/Foreman/FoundationJob", reflect.TypeOf((*ForemanFoundationJob)(nil)).Elem())
	cinp.RegisterType("/api/v1/Foreman/StructureJob", reflect.TypeOf((*ForemanStructureJob)(nil)).Elem())
	cinp.RegisterType("/api/v1/Foreman/DependencyJob", reflect.TypeOf((*ForemanDependencyJob)(nil)).Elem())
	cinp.RegisterType("/api/v1/Foreman/JobLog", reflect.TypeOf((*ForemanJobLog)(nil)).Elem())
}
