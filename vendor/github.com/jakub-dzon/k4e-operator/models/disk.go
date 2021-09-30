// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Disk disk
//
// swagger:model disk
type Disk struct {

	// bootable
	Bootable bool `json:"bootable,omitempty"`

	// by-id is the World Wide Number of the device which guaranteed to be unique for every storage device
	ByID string `json:"by_id,omitempty"`

	// by-path is the shortest physical path to the device
	ByPath string `json:"by_path,omitempty"`

	// drive type
	DriveType string `json:"drive_type,omitempty"`

	// hctl
	Hctl string `json:"hctl,omitempty"`

	// Determine the disk's unique identifier which is the by-id field if it exists and fallback to the by-path field otherwise
	ID string `json:"id,omitempty"`

	// io perf
	IoPerf *IoPerf `json:"io_perf,omitempty"`

	// Whether the disk appears to be an installation media or not
	IsInstallationMedia bool `json:"is_installation_media,omitempty"`

	// model
	Model string `json:"model,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// serial
	Serial string `json:"serial,omitempty"`

	// size bytes
	SizeBytes int64 `json:"size_bytes,omitempty"`

	// smart
	Smart string `json:"smart,omitempty"`

	// vendor
	Vendor string `json:"vendor,omitempty"`

	// wwn
	Wwn string `json:"wwn,omitempty"`
}

// Validate validates this disk
func (m *Disk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIoPerf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Disk) validateIoPerf(formats strfmt.Registry) error {

	if swag.IsZero(m.IoPerf) { // not required
		return nil
	}

	if m.IoPerf != nil {
		if err := m.IoPerf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_perf")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Disk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Disk) UnmarshalBinary(b []byte) error {
	var res Disk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
