// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/nested.proto

/*
Package validator_examples is a generated protocol buffer package.

It is generated from these files:
	examples/nested.proto

It has these top-level messages:
	InnerMessage
	OuterMessage
*/
package validator_examples

import regexp "regexp"
import fmt "fmt"
import go_proto_validators "github.com/simplesurance/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/simplesurance/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *InnerMessage) Validate() error {
	if !(this.SomeInteger > 0) {
		return go_proto_validators.FieldError("SomeInteger", fmt.Errorf(`value '%v' must be greater than '0'`, this.SomeInteger))
	}
	if !(this.SomeInteger < 100) {
		return go_proto_validators.FieldError("SomeInteger", fmt.Errorf(`value '%v' must be less than '100'`, this.SomeInteger))
	}
	return nil
}

var _regex_OuterMessage_ImportantString = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
var _regex_OuterMessage_UserId = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
var _regex_OuterMessage_UserIdd = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)

func (this *OuterMessage) Validate() error {
	if !_regex_OuterMessage_ImportantString.MatchString(this.ImportantString) {
		return go_proto_validators.FieldError("ImportantString", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"`, this.ImportantString))
	}
	if nil == this.Inner {
		return go_proto_validators.FieldError("Inner", fmt.Errorf("message must exist"))
	}
	if this.Inner != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Inner); err != nil {
			return go_proto_validators.FieldError("Inner", err)
		}
	}
	if !_regex_OuterMessage_UserId.MatchString(this.UserId) {
		return go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"`, this.UserId))
	}
	if !_regex_OuterMessage_UserIdd.MatchString(this.UserIdd) {
		return go_proto_validators.FieldError("UserIdd", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"`, this.UserIdd))
	}
	return nil
}
