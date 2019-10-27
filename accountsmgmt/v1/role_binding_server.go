/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// RoleBindingServer represents the interface the manages the 'role_binding' resource.
type RoleBindingServer interface {

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the role binding.
	Delete(ctx context.Context, request *RoleBindingDeleteServerRequest, response *RoleBindingDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the role binding.
	Get(ctx context.Context, request *RoleBindingGetServerRequest, response *RoleBindingGetServerResponse) error
}

// RoleBindingDeleteServerRequest is the request for the 'delete' method.
type RoleBindingDeleteServerRequest struct {
}

// RoleBindingDeleteServerResponse is the response for the 'delete' method.
type RoleBindingDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *RoleBindingDeleteServerResponse) Status(value int) *RoleBindingDeleteServerResponse {
	r.status = value
	return r
}

// RoleBindingGetServerRequest is the request for the 'get' method.
type RoleBindingGetServerRequest struct {
}

// RoleBindingGetServerResponse is the response for the 'get' method.
type RoleBindingGetServerResponse struct {
	status int
	err    *errors.Error
	body   *RoleBinding
}

// Body sets the value of the 'body' parameter.
//
//
func (r *RoleBindingGetServerResponse) Body(value *RoleBinding) *RoleBindingGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *RoleBindingGetServerResponse) Status(value int) *RoleBindingGetServerResponse {
	r.status = value
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *RoleBindingGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// dispatchRoleBinding navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchRoleBinding(w http.ResponseWriter, r *http.Request, server RoleBindingServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case http.MethodDelete:
			adaptRoleBindingDeleteRequest(w, r, server)
		case http.MethodGet:
			adaptRoleBindingGetRequest(w, r, server)
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	} else {
		switch segments[0] {
		default:
			errors.SendNotFound(w, r)
			return
		}
	}
}

// readRoleBindingDeleteRequest reads the given HTTP requests and translates it
// into an object of type RoleBindingDeleteServerRequest.
func readRoleBindingDeleteRequest(r *http.Request) (*RoleBindingDeleteServerRequest, error) {
	var err error
	result := new(RoleBindingDeleteServerRequest)
	return result, err
}

// writeRoleBindingDeleteResponse translates the given request object into an
// HTTP response.
func writeRoleBindingDeleteResponse(w http.ResponseWriter, r *RoleBindingDeleteServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	return nil
}

// adaptRoleBindingDeleteRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptRoleBindingDeleteRequest(w http.ResponseWriter, r *http.Request, server RoleBindingServer) {
	request, err := readRoleBindingDeleteRequest(r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := new(RoleBindingDeleteServerResponse)
	response.status = http.StatusOK
	err = server.Delete(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeRoleBindingDeleteResponse(w, response)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// readRoleBindingGetRequest reads the given HTTP requests and translates it
// into an object of type RoleBindingGetServerRequest.
func readRoleBindingGetRequest(r *http.Request) (*RoleBindingGetServerRequest, error) {
	var err error
	result := new(RoleBindingGetServerRequest)
	return result, err
}

// writeRoleBindingGetResponse translates the given request object into an
// HTTP response.
func writeRoleBindingGetResponse(w http.ResponseWriter, r *RoleBindingGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}

// adaptRoleBindingGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptRoleBindingGetRequest(w http.ResponseWriter, r *http.Request, server RoleBindingServer) {
	request, err := readRoleBindingGetRequest(r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := new(RoleBindingGetServerResponse)
	response.status = http.StatusOK
	err = server.Get(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeRoleBindingGetResponse(w, response)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}