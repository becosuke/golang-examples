// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: node.proto

package service

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreateNodeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateNodeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNodeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateNodeRequestMultiError, or nil if none found.
func (m *CreateNodeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNodeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetNode() == nil {
		err := CreateNodeRequestValidationError{
			field:  "Node",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateNodeRequestValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateNodeRequestValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateNodeRequestValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateNodeRequestMultiError(errors)
	}
	return nil
}

// CreateNodeRequestMultiError is an error wrapping multiple validation errors
// returned by CreateNodeRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateNodeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNodeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNodeRequestMultiError) AllErrors() []error { return m }

// CreateNodeRequestValidationError is the validation error returned by
// CreateNodeRequest.Validate if the designated constraints aren't met.
type CreateNodeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNodeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNodeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNodeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNodeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNodeRequestValidationError) ErrorName() string {
	return "CreateNodeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNodeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNodeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNodeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNodeRequestValidationError{}

// Validate checks the field values on CreateNodeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateNodeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNodeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateNodeResponseMultiError, or nil if none found.
func (m *CreateNodeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNodeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateNodeResponseValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateNodeResponseValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateNodeResponseValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateNodeResponseMultiError(errors)
	}
	return nil
}

// CreateNodeResponseMultiError is an error wrapping multiple validation errors
// returned by CreateNodeResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateNodeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNodeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNodeResponseMultiError) AllErrors() []error { return m }

// CreateNodeResponseValidationError is the validation error returned by
// CreateNodeResponse.Validate if the designated constraints aren't met.
type CreateNodeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNodeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNodeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNodeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNodeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNodeResponseValidationError) ErrorName() string {
	return "CreateNodeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNodeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNodeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNodeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNodeResponseValidationError{}

// Validate checks the field values on GetNodeRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetNodeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNodeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetNodeRequestMultiError,
// or nil if none found.
func (m *GetNodeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNodeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetNodeKey() == nil {
		err := GetNodeRequestValidationError{
			field:  "NodeKey",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetNodeKey()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetNodeRequestValidationError{
					field:  "NodeKey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetNodeRequestValidationError{
					field:  "NodeKey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNodeKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetNodeRequestValidationError{
				field:  "NodeKey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetNodeRequestMultiError(errors)
	}
	return nil
}

// GetNodeRequestMultiError is an error wrapping multiple validation errors
// returned by GetNodeRequest.ValidateAll() if the designated constraints
// aren't met.
type GetNodeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNodeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNodeRequestMultiError) AllErrors() []error { return m }

// GetNodeRequestValidationError is the validation error returned by
// GetNodeRequest.Validate if the designated constraints aren't met.
type GetNodeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNodeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNodeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNodeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNodeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNodeRequestValidationError) ErrorName() string { return "GetNodeRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetNodeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNodeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNodeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNodeRequestValidationError{}

// Validate checks the field values on GetNodeResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetNodeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNodeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNodeResponseMultiError, or nil if none found.
func (m *GetNodeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNodeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetNodeResponseValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetNodeResponseValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetNodeResponseValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetNodeResponseMultiError(errors)
	}
	return nil
}

// GetNodeResponseMultiError is an error wrapping multiple validation errors
// returned by GetNodeResponse.ValidateAll() if the designated constraints
// aren't met.
type GetNodeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNodeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNodeResponseMultiError) AllErrors() []error { return m }

// GetNodeResponseValidationError is the validation error returned by
// GetNodeResponse.Validate if the designated constraints aren't met.
type GetNodeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNodeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNodeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNodeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNodeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNodeResponseValidationError) ErrorName() string { return "GetNodeResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetNodeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNodeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNodeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNodeResponseValidationError{}

// Validate checks the field values on UpdateNodeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateNodeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateNodeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateNodeRequestMultiError, or nil if none found.
func (m *UpdateNodeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNodeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetNode() == nil {
		err := UpdateNodeRequestValidationError{
			field:  "Node",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNodeRequestValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNodeRequestValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNodeRequestValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateNodeRequestMultiError(errors)
	}
	return nil
}

// UpdateNodeRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateNodeRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateNodeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNodeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNodeRequestMultiError) AllErrors() []error { return m }

// UpdateNodeRequestValidationError is the validation error returned by
// UpdateNodeRequest.Validate if the designated constraints aren't met.
type UpdateNodeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNodeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNodeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNodeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNodeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNodeRequestValidationError) ErrorName() string {
	return "UpdateNodeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateNodeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNodeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNodeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNodeRequestValidationError{}

// Validate checks the field values on UpdateNodeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateNodeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateNodeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateNodeResponseMultiError, or nil if none found.
func (m *UpdateNodeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNodeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNodeResponseValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNodeResponseValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNodeResponseValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateNodeResponseMultiError(errors)
	}
	return nil
}

// UpdateNodeResponseMultiError is an error wrapping multiple validation errors
// returned by UpdateNodeResponse.ValidateAll() if the designated constraints
// aren't met.
type UpdateNodeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNodeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNodeResponseMultiError) AllErrors() []error { return m }

// UpdateNodeResponseValidationError is the validation error returned by
// UpdateNodeResponse.Validate if the designated constraints aren't met.
type UpdateNodeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNodeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNodeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNodeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNodeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNodeResponseValidationError) ErrorName() string {
	return "UpdateNodeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateNodeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNodeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNodeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNodeResponseValidationError{}

// Validate checks the field values on DeleteNodeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteNodeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteNodeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteNodeRequestMultiError, or nil if none found.
func (m *DeleteNodeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteNodeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetNodeKey() == nil {
		err := DeleteNodeRequestValidationError{
			field:  "NodeKey",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetNodeKey()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeleteNodeRequestValidationError{
					field:  "NodeKey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeleteNodeRequestValidationError{
					field:  "NodeKey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNodeKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteNodeRequestValidationError{
				field:  "NodeKey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DeleteNodeRequestMultiError(errors)
	}
	return nil
}

// DeleteNodeRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteNodeRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteNodeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteNodeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteNodeRequestMultiError) AllErrors() []error { return m }

// DeleteNodeRequestValidationError is the validation error returned by
// DeleteNodeRequest.Validate if the designated constraints aren't met.
type DeleteNodeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteNodeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteNodeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteNodeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteNodeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteNodeRequestValidationError) ErrorName() string {
	return "DeleteNodeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteNodeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteNodeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteNodeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteNodeRequestValidationError{}

// Validate checks the field values on Node with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Node) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Node with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NodeMultiError, or nil if none found.
func (m *Node) ValidateAll() error {
	return m.validate(true)
}

func (m *Node) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetKey()); l < 1 || l > 128 {
		err := NodeValidationError{
			field:  "Key",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetValue()); l < 1 || l > 128 {
		err := NodeValidationError{
			field:  "Value",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return NodeMultiError(errors)
	}
	return nil
}

// NodeMultiError is an error wrapping multiple validation errors returned by
// Node.ValidateAll() if the designated constraints aren't met.
type NodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NodeMultiError) AllErrors() []error { return m }

// NodeValidationError is the validation error returned by Node.Validate if the
// designated constraints aren't met.
type NodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeValidationError) ErrorName() string { return "NodeValidationError" }

// Error satisfies the builtin error interface
func (e NodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeValidationError{}

// Validate checks the field values on NodeKey with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NodeKey) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NodeKey with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NodeKeyMultiError, or nil if none found.
func (m *NodeKey) ValidateAll() error {
	return m.validate(true)
}

func (m *NodeKey) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetKey()); l < 1 || l > 128 {
		err := NodeKeyValidationError{
			field:  "Key",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return NodeKeyMultiError(errors)
	}
	return nil
}

// NodeKeyMultiError is an error wrapping multiple validation errors returned
// by NodeKey.ValidateAll() if the designated constraints aren't met.
type NodeKeyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NodeKeyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NodeKeyMultiError) AllErrors() []error { return m }

// NodeKeyValidationError is the validation error returned by NodeKey.Validate
// if the designated constraints aren't met.
type NodeKeyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeKeyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeKeyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeKeyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeKeyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeKeyValidationError) ErrorName() string { return "NodeKeyValidationError" }

// Error satisfies the builtin error interface
func (e NodeKeyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeKey.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeKeyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeKeyValidationError{}