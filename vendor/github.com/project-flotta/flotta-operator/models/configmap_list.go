// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// ConfigmapList List of configmaps used by the workload
//
// swagger:model configmap-list
type ConfigmapList []string

// Validate validates this configmap list
func (m ConfigmapList) Validate(formats strfmt.Registry) error {
	return nil
}
