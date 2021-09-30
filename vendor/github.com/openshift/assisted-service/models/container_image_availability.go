// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainerImageAvailability container image availability
//
// swagger:model container_image_availability
type ContainerImageAvailability struct {

	// The rate of size/time in seconds MBps.
	DownloadRate float64 `json:"download_rate,omitempty"`

	// A fully qualified image name (FQIN).
	Name string `json:"name,omitempty"`

	// result
	Result ContainerImageAvailabilityResult `json:"result,omitempty"`

	// Size of the image in bytes.
	SizeBytes float64 `json:"size_bytes,omitempty"`

	// Seconds it took to pull the image.
	Time float64 `json:"time,omitempty"`
}

// Validate validates this container image availability
func (m *ContainerImageAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerImageAvailability) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if err := m.Result.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("result")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerImageAvailability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerImageAvailability) UnmarshalBinary(b []byte) error {
	var res ContainerImageAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
