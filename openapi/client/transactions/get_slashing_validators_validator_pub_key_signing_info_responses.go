// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/lidofinance/terra-monitors/openapi/models"
)

// GetSlashingValidatorsValidatorPubKeySigningInfoReader is a Reader for the GetSlashingValidatorsValidatorPubKeySigningInfo structure.
type GetSlashingValidatorsValidatorPubKeySigningInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSlashingValidatorsValidatorPubKeySigningInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSlashingValidatorsValidatorPubKeySigningInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSlashingValidatorsValidatorPubKeySigningInfoOK creates a GetSlashingValidatorsValidatorPubKeySigningInfoOK with default headers values
func NewGetSlashingValidatorsValidatorPubKeySigningInfoOK() *GetSlashingValidatorsValidatorPubKeySigningInfoOK {
	return &GetSlashingValidatorsValidatorPubKeySigningInfoOK{}
}

/* GetSlashingValidatorsValidatorPubKeySigningInfoOK describes a response with status code 200, with default header values.

Success
*/
type GetSlashingValidatorsValidatorPubKeySigningInfoOK struct {
	Payload *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOK) Error() string {
	return fmt.Sprintf("[GET /slashing/validators/{validatorPubKey}/signing_info][%d] getSlashingValidatorsValidatorPubKeySigningInfoOK  %+v", 200, o.Payload)
}
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOK) GetPayload() *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody {
	return o.Payload
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSlashingValidatorsValidatorPubKeySigningInfoOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSlashingValidatorsValidatorPubKeySigningInfoBadRequest creates a GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest with default headers values
func NewGetSlashingValidatorsValidatorPubKeySigningInfoBadRequest() *GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest {
	return &GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest{}
}

/* GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest describes a response with status code 400, with default header values.

Error
*/
type GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest struct {
	Payload *GetSlashingValidatorsValidatorPubKeySigningInfoBadRequestBody
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /slashing/validators/{validatorPubKey}/signing_info][%d] getSlashingValidatorsValidatorPubKeySigningInfoBadRequest  %+v", 400, o.Payload)
}
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest) GetPayload() *GetSlashingValidatorsValidatorPubKeySigningInfoBadRequestBody {
	return o.Payload
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSlashingValidatorsValidatorPubKeySigningInfoBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSlashingValidatorsValidatorPubKeySigningInfoBadRequestBody get slashing validators validator pub key signing info bad request body
swagger:model GetSlashingValidatorsValidatorPubKeySigningInfoBadRequestBody
*/
type GetSlashingValidatorsValidatorPubKeySigningInfoBadRequestBody struct {

	// code
	Code float64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this get slashing validators validator pub key signing info bad request body
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get slashing validators validator pub key signing info bad request body based on context it is used
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetSlashingValidatorsValidatorPubKeySigningInfoBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSlashingValidatorsValidatorPubKeySigningInfoOKBody get slashing validators validator pub key signing info o k body
swagger:model GetSlashingValidatorsValidatorPubKeySigningInfoOKBody
*/
type GetSlashingValidatorsValidatorPubKeySigningInfoOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result *models.SigningInfo `json:"result,omitempty"`
}

// Validate validates this get slashing validators validator pub key signing info o k body
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSlashingValidatorsValidatorPubKeySigningInfoOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get slashing validators validator pub key signing info o k body based on the context it is used
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if o.Result != nil {
		if err := o.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSlashingValidatorsValidatorPubKeySigningInfoOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody) UnmarshalBinary(b []byte) error {
	var res GetSlashingValidatorsValidatorPubKeySigningInfoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
