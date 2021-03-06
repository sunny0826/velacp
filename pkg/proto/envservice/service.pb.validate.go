// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/proto/envservice/service.proto

package envservice

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on PutEnvRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PutEnvRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetEnv()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PutEnvRequestValidationError{
				field:  "Env",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PutEnvRequestValidationError is the validation error returned by
// PutEnvRequest.Validate if the designated constraints aren't met.
type PutEnvRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutEnvRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutEnvRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutEnvRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutEnvRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutEnvRequestValidationError) ErrorName() string { return "PutEnvRequestValidationError" }

// Error satisfies the builtin error interface
func (e PutEnvRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutEnvRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutEnvRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutEnvRequestValidationError{}

// Validate checks the field values on PutEnvResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PutEnvResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// PutEnvResponseValidationError is the validation error returned by
// PutEnvResponse.Validate if the designated constraints aren't met.
type PutEnvResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutEnvResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutEnvResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutEnvResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutEnvResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutEnvResponseValidationError) ErrorName() string { return "PutEnvResponseValidationError" }

// Error satisfies the builtin error interface
func (e PutEnvResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutEnvResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutEnvResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutEnvResponseValidationError{}

// Validate checks the field values on GetEnvRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetEnvRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// GetEnvRequestValidationError is the validation error returned by
// GetEnvRequest.Validate if the designated constraints aren't met.
type GetEnvRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEnvRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEnvRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEnvRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEnvRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEnvRequestValidationError) ErrorName() string { return "GetEnvRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetEnvRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEnvRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEnvRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEnvRequestValidationError{}

// Validate checks the field values on GetEnvResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetEnvResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetEnv()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetEnvResponseValidationError{
				field:  "Env",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetEnvResponseValidationError is the validation error returned by
// GetEnvResponse.Validate if the designated constraints aren't met.
type GetEnvResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEnvResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEnvResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEnvResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEnvResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEnvResponseValidationError) ErrorName() string { return "GetEnvResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetEnvResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEnvResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEnvResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEnvResponseValidationError{}

// Validate checks the field values on ListEnvsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListEnvsRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListEnvsRequestValidationError is the validation error returned by
// ListEnvsRequest.Validate if the designated constraints aren't met.
type ListEnvsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEnvsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEnvsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEnvsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEnvsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEnvsRequestValidationError) ErrorName() string { return "ListEnvsRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListEnvsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEnvsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEnvsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEnvsRequestValidationError{}

// Validate checks the field values on ListEnvsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListEnvsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEnvs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListEnvsResponseValidationError{
					field:  fmt.Sprintf("Envs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListEnvsResponseValidationError is the validation error returned by
// ListEnvsResponse.Validate if the designated constraints aren't met.
type ListEnvsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEnvsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEnvsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEnvsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEnvsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEnvsResponseValidationError) ErrorName() string { return "ListEnvsResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListEnvsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEnvsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEnvsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEnvsResponseValidationError{}

// Validate checks the field values on DelEnvRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DelEnvRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// DelEnvRequestValidationError is the validation error returned by
// DelEnvRequest.Validate if the designated constraints aren't met.
type DelEnvRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DelEnvRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DelEnvRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DelEnvRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DelEnvRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DelEnvRequestValidationError) ErrorName() string { return "DelEnvRequestValidationError" }

// Error satisfies the builtin error interface
func (e DelEnvRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDelEnvRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DelEnvRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DelEnvRequestValidationError{}

// Validate checks the field values on DelEnvResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DelEnvResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DelEnvResponseValidationError is the validation error returned by
// DelEnvResponse.Validate if the designated constraints aren't met.
type DelEnvResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DelEnvResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DelEnvResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DelEnvResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DelEnvResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DelEnvResponseValidationError) ErrorName() string { return "DelEnvResponseValidationError" }

// Error satisfies the builtin error interface
func (e DelEnvResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDelEnvResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DelEnvResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DelEnvResponseValidationError{}

// Validate checks the field values on ListCapsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListCapsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for EnvName

	return nil
}

// ListCapsRequestValidationError is the validation error returned by
// ListCapsRequest.Validate if the designated constraints aren't met.
type ListCapsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCapsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCapsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCapsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCapsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCapsRequestValidationError) ErrorName() string { return "ListCapsRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListCapsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCapsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCapsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCapsRequestValidationError{}

// Validate checks the field values on ListCapsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListCapsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetCaps() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCapsResponseValidationError{
					field:  fmt.Sprintf("Caps[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListCapsResponseValidationError is the validation error returned by
// ListCapsResponse.Validate if the designated constraints aren't met.
type ListCapsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCapsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCapsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCapsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCapsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCapsResponseValidationError) ErrorName() string { return "ListCapsResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListCapsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCapsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCapsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCapsResponseValidationError{}
