// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: src/api/message/common/v1/common.proto

package v1

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
var _common_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on UUID with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *UUID) Validate() error {
	if m == nil {
		return nil
	}

	if !_UUID_Id_Pattern.MatchString(m.GetId()) {
		return UUIDValidationError{
			field:  "Id",
			reason: "value does not match regex pattern \"^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$\"",
		}
	}

	return nil
}

// UUIDValidationError is the validation error returned by UUID.Validate if the
// designated constraints aren't met.
type UUIDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UUIDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UUIDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UUIDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UUIDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UUIDValidationError) ErrorName() string { return "UUIDValidationError" }

// Error satisfies the builtin error interface
func (e UUIDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUUID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UUIDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UUIDValidationError{}

var _UUID_Id_Pattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Pagination with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Pagination) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLimit() < 0 {
		return PaginationValidationError{
			field:  "Limit",
			reason: "value must be greater than or equal to 0",
		}
	}

	if m.GetOffset() < 0 {
		return PaginationValidationError{
			field:  "Offset",
			reason: "value must be greater than or equal to 0",
		}
	}

	return nil
}

// PaginationValidationError is the validation error returned by
// Pagination.Validate if the designated constraints aren't met.
type PaginationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginationValidationError) ErrorName() string { return "PaginationValidationError" }

// Error satisfies the builtin error interface
func (e PaginationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPagination.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginationValidationError{}
