// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetTxListResultTxsTxValue get tx list result txs tx value
//
// swagger:model getTxListResult.txs.tx.value
type GetTxListResultTxsTxValue struct {

	// fee
	// Required: true
	Fee *GetTxListResultTxsTxValueFee `json:"fee"`

	// memo
	// Required: true
	Memo *string `json:"memo"`

	// msg
	// Required: true
	Msg []*GetTxListResultTxsTxValueMsg `json:"msg"`

	// signatures
	// Required: true
	Signatures []*GetTxListResultTxsTxValueSignatures `json:"signatures"`
}

// Validate validates this get tx list result txs tx value
func (m *GetTxListResultTxsTxValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTxListResultTxsTxValue) validateFee(formats strfmt.Registry) error {

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	if m.Fee != nil {
		if err := m.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fee")
			}
			return err
		}
	}

	return nil
}

func (m *GetTxListResultTxsTxValue) validateMemo(formats strfmt.Registry) error {

	if err := validate.Required("memo", "body", m.Memo); err != nil {
		return err
	}

	return nil
}

func (m *GetTxListResultTxsTxValue) validateMsg(formats strfmt.Registry) error {

	if err := validate.Required("msg", "body", m.Msg); err != nil {
		return err
	}

	for i := 0; i < len(m.Msg); i++ {
		if swag.IsZero(m.Msg[i]) { // not required
			continue
		}

		if m.Msg[i] != nil {
			if err := m.Msg[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("msg" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetTxListResultTxsTxValue) validateSignatures(formats strfmt.Registry) error {

	if err := validate.Required("signatures", "body", m.Signatures); err != nil {
		return err
	}

	for i := 0; i < len(m.Signatures); i++ {
		if swag.IsZero(m.Signatures[i]) { // not required
			continue
		}

		if m.Signatures[i] != nil {
			if err := m.Signatures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get tx list result txs tx value based on the context it is used
func (m *GetTxListResultTxsTxValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMsg(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTxListResultTxsTxValue) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

	if m.Fee != nil {
		if err := m.Fee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fee")
			}
			return err
		}
	}

	return nil
}

func (m *GetTxListResultTxsTxValue) contextValidateMsg(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Msg); i++ {

		if m.Msg[i] != nil {
			if err := m.Msg[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("msg" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetTxListResultTxsTxValue) contextValidateSignatures(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Signatures); i++ {

		if m.Signatures[i] != nil {
			if err := m.Signatures[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTxListResultTxsTxValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTxListResultTxsTxValue) UnmarshalBinary(b []byte) error {
	var res GetTxListResultTxsTxValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
