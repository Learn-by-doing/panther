// Code generated by go-swagger; DO NOT EDIT.

package operations

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/panther-labs/panther/api/gateway/analysis/models"
)

// GetRuleReader is a Reader for the GetRule structure.
type GetRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRuleOK creates a GetRuleOK with default headers values
func NewGetRuleOK() *GetRuleOK {
	return &GetRuleOK{}
}

/*GetRuleOK handles this case with default header values.

OK
*/
type GetRuleOK struct {
	Payload *models.Rule
}

func (o *GetRuleOK) Error() string {
	return fmt.Sprintf("[GET /rule][%d] getRuleOK  %+v", 200, o.Payload)
}

func (o *GetRuleOK) GetPayload() *models.Rule {
	return o.Payload
}

func (o *GetRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuleBadRequest creates a GetRuleBadRequest with default headers values
func NewGetRuleBadRequest() *GetRuleBadRequest {
	return &GetRuleBadRequest{}
}

/*GetRuleBadRequest handles this case with default header values.

Bad request
*/
type GetRuleBadRequest struct {
	Payload *models.Error
}

func (o *GetRuleBadRequest) Error() string {
	return fmt.Sprintf("[GET /rule][%d] getRuleBadRequest  %+v", 400, o.Payload)
}

func (o *GetRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuleNotFound creates a GetRuleNotFound with default headers values
func NewGetRuleNotFound() *GetRuleNotFound {
	return &GetRuleNotFound{}
}

/*GetRuleNotFound handles this case with default header values.

Rule does not exist
*/
type GetRuleNotFound struct {
}

func (o *GetRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /rule][%d] getRuleNotFound ", 404)
}

func (o *GetRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRuleInternalServerError creates a GetRuleInternalServerError with default headers values
func NewGetRuleInternalServerError() *GetRuleInternalServerError {
	return &GetRuleInternalServerError{}
}

/*GetRuleInternalServerError handles this case with default header values.

Internal server error
*/
type GetRuleInternalServerError struct {
}

func (o *GetRuleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /rule][%d] getRuleInternalServerError ", 500)
}

func (o *GetRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
