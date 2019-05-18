// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "identity provider": Application User Types
//
// Command:
// $ goagen
// --design=github.com/Microkubes/identity-provider/design
// --out=$(GOPATH)/src/github.com/Microkubes/identity-provider
// --version=v1.3.1

package app

import (
	"github.com/keitaroinc/goa"
)

// DeleteSPPayload
type deleteSPPayload struct {
	// ID of service provider
	ServiceID *string `form:"serviceId,omitempty" json:"serviceId,omitempty" yaml:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

// Validate validates the deleteSPPayload type instance.
func (ut *deleteSPPayload) Validate() (err error) {
	if ut.ServiceID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "serviceId"))
	}
	return
}

// Publicize creates DeleteSPPayload from deleteSPPayload
func (ut *deleteSPPayload) Publicize() *DeleteSPPayload {
	var pub DeleteSPPayload
	if ut.ServiceID != nil {
		pub.ServiceID = *ut.ServiceID
	}
	return &pub
}

// DeleteSPPayload
type DeleteSPPayload struct {
	// ID of service provider
	ServiceID string `form:"serviceId" json:"serviceId" yaml:"serviceId" xml:"serviceId"`
}

// Validate validates the DeleteSPPayload type instance.
func (ut *DeleteSPPayload) Validate() (err error) {
	if ut.ServiceID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "serviceId"))
	}
	return
}

// DeleteSessionPayload
type deleteSessionPayload struct {
	// ID of the session
	SessionID *string `form:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

// Validate validates the deleteSessionPayload type instance.
func (ut *deleteSessionPayload) Validate() (err error) {
	if ut.SessionID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "sessionId"))
	}
	return
}

// Publicize creates DeleteSessionPayload from deleteSessionPayload
func (ut *deleteSessionPayload) Publicize() *DeleteSessionPayload {
	var pub DeleteSessionPayload
	if ut.SessionID != nil {
		pub.SessionID = *ut.SessionID
	}
	return &pub
}

// DeleteSessionPayload
type DeleteSessionPayload struct {
	// ID of the session
	SessionID string `form:"sessionId" json:"sessionId" yaml:"sessionId" xml:"sessionId"`
}

// Validate validates the DeleteSessionPayload type instance.
func (ut *DeleteSessionPayload) Validate() (err error) {
	if ut.SessionID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "sessionId"))
	}
	return
}
