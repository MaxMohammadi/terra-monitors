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

// SigningInfo signing info
//
// swagger:model SigningInfo
type SigningInfo struct {

	// address
	// Required: true
	Address *string `json:"address"`

	// index offset
	// Required: true
	IndexOffset *string `json:"index_offset"`

	// missed blocks counter
	// Required: true
	MissedBlocksCounter *string `json:"missed_blocks_counter"`

	// start height
	// Required: true
	StartHeight *string `json:"start_height"`

	// tombstoned
	// Required: true
	Tombstoned *bool `json:"tombstoned"`
}

// Validate validates this signing info
func (m *SigningInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexOffset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMissedBlocksCounter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTombstoned(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SigningInfo) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *SigningInfo) validateIndexOffset(formats strfmt.Registry) error {

	if err := validate.Required("index_offset", "body", m.IndexOffset); err != nil {
		return err
	}

	return nil
}

func (m *SigningInfo) validateMissedBlocksCounter(formats strfmt.Registry) error {

	if err := validate.Required("missed_blocks_counter", "body", m.MissedBlocksCounter); err != nil {
		return err
	}

	return nil
}

func (m *SigningInfo) validateStartHeight(formats strfmt.Registry) error {

	if err := validate.Required("start_height", "body", m.StartHeight); err != nil {
		return err
	}

	return nil
}

func (m *SigningInfo) validateTombstoned(formats strfmt.Registry) error {

	if err := validate.Required("tombstoned", "body", m.Tombstoned); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this signing info based on context it is used
func (m *SigningInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SigningInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SigningInfo) UnmarshalBinary(b []byte) error {
	var res SigningInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
