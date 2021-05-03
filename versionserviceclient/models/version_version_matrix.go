// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VersionVersionMatrix VersionMatrix represents set of possible product versions.
//
// swagger:model versionVersionMatrix
type VersionVersionMatrix struct {

	// backup
	Backup map[string]VersionVersion `json:"backup,omitempty"`

	// haproxy
	Haproxy map[string]VersionVersion `json:"haproxy,omitempty"`

	// log collector
	LogCollector map[string]VersionVersion `json:"logCollector,omitempty"`

	// mongod
	Mongod map[string]VersionVersion `json:"mongod,omitempty"`

	// operator
	Operator map[string]VersionVersion `json:"operator,omitempty"`

	// pmm
	Pmm map[string]VersionVersion `json:"pmm,omitempty"`

	// proxysql
	Proxysql map[string]VersionVersion `json:"proxysql,omitempty"`

	// pxc
	Pxc map[string]VersionVersion `json:"pxc,omitempty"`
}

// Validate validates this version version matrix
func (m *VersionVersionMatrix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHaproxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogCollector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMongod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePxc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionVersionMatrix) validateBackup(formats strfmt.Registry) error {

	if swag.IsZero(m.Backup) { // not required
		return nil
	}

	for k := range m.Backup {

		if err := validate.Required("backup"+"."+k, "body", m.Backup[k]); err != nil {
			return err
		}
		if val, ok := m.Backup[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validateHaproxy(formats strfmt.Registry) error {

	if swag.IsZero(m.Haproxy) { // not required
		return nil
	}

	for k := range m.Haproxy {

		if err := validate.Required("haproxy"+"."+k, "body", m.Haproxy[k]); err != nil {
			return err
		}
		if val, ok := m.Haproxy[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validateLogCollector(formats strfmt.Registry) error {

	if swag.IsZero(m.LogCollector) { // not required
		return nil
	}

	for k := range m.LogCollector {

		if err := validate.Required("logCollector"+"."+k, "body", m.LogCollector[k]); err != nil {
			return err
		}
		if val, ok := m.LogCollector[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validateMongod(formats strfmt.Registry) error {

	if swag.IsZero(m.Mongod) { // not required
		return nil
	}

	for k := range m.Mongod {

		if err := validate.Required("mongod"+"."+k, "body", m.Mongod[k]); err != nil {
			return err
		}
		if val, ok := m.Mongod[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	for k := range m.Operator {

		if err := validate.Required("operator"+"."+k, "body", m.Operator[k]); err != nil {
			return err
		}
		if val, ok := m.Operator[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validatePmm(formats strfmt.Registry) error {

	if swag.IsZero(m.Pmm) { // not required
		return nil
	}

	for k := range m.Pmm {

		if err := validate.Required("pmm"+"."+k, "body", m.Pmm[k]); err != nil {
			return err
		}
		if val, ok := m.Pmm[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validateProxysql(formats strfmt.Registry) error {

	if swag.IsZero(m.Proxysql) { // not required
		return nil
	}

	for k := range m.Proxysql {

		if err := validate.Required("proxysql"+"."+k, "body", m.Proxysql[k]); err != nil {
			return err
		}
		if val, ok := m.Proxysql[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validatePxc(formats strfmt.Registry) error {

	if swag.IsZero(m.Pxc) { // not required
		return nil
	}

	for k := range m.Pxc {

		if err := validate.Required("pxc"+"."+k, "body", m.Pxc[k]); err != nil {
			return err
		}
		if val, ok := m.Pxc[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionVersionMatrix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionVersionMatrix) UnmarshalBinary(b []byte) error {
	var res VersionVersionMatrix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}