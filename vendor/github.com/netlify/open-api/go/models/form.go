// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Form form
// swagger:model form
type Form struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// fields
	Fields []interface{} `json:"fields"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// paths
	Paths []string `json:"paths"`

	// site id
	SiteID string `json:"site_id,omitempty"`

	// submission count
	SubmissionCount int32 `json:"submission_count,omitempty"`
}

// Validate validates this form
func (m *Form) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Form) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Form) UnmarshalBinary(b []byte) error {
	var res Form
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
