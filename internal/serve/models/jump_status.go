// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JumpStatus Status of the jump connection for an experiment
//
// swagger:model JumpStatus
type JumpStatus struct {

	// is the experiment currently actively sending?
	// Required: true
	Active *bool `json:"active"`

	// number of clients connected (0 if just the experiment)
	// Required: true
	Clients *int64 `json:"clients"`

	// is the experiment currently connected to jump?
	// Required: true
	Connected *bool `json:"connected"`

	// duration since last send by experiment
	// Required: true
	Last *string `json:"last"`

	// name of the experiment
	// Example: spin30
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this jump status
func (m *JumpStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnected(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JumpStatus) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

func (m *JumpStatus) validateClients(formats strfmt.Registry) error {

	if err := validate.Required("clients", "body", m.Clients); err != nil {
		return err
	}

	return nil
}

func (m *JumpStatus) validateConnected(formats strfmt.Registry) error {

	if err := validate.Required("connected", "body", m.Connected); err != nil {
		return err
	}

	return nil
}

func (m *JumpStatus) validateLast(formats strfmt.Registry) error {

	if err := validate.Required("last", "body", m.Last); err != nil {
		return err
	}

	return nil
}

func (m *JumpStatus) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this jump status based on context it is used
func (m *JumpStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JumpStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JumpStatus) UnmarshalBinary(b []byte) error {
	var res JumpStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
