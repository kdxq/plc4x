/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetDoorSecuredStatusTagged is the corresponding interface of BACnetDoorSecuredStatusTagged
type BACnetDoorSecuredStatusTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetDoorSecuredStatus
}

// BACnetDoorSecuredStatusTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetDoorSecuredStatusTagged.
// This is useful for switch cases.
type BACnetDoorSecuredStatusTaggedExactly interface {
	BACnetDoorSecuredStatusTagged
	isBACnetDoorSecuredStatusTagged() bool
}

// _BACnetDoorSecuredStatusTagged is the data-structure of this message
type _BACnetDoorSecuredStatusTagged struct {
	Header BACnetTagHeader
	Value  BACnetDoorSecuredStatus

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDoorSecuredStatusTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetDoorSecuredStatusTagged) GetValue() BACnetDoorSecuredStatus {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDoorSecuredStatusTagged factory function for _BACnetDoorSecuredStatusTagged
func NewBACnetDoorSecuredStatusTagged(header BACnetTagHeader, value BACnetDoorSecuredStatus, tagNumber uint8, tagClass TagClass) *_BACnetDoorSecuredStatusTagged {
	return &_BACnetDoorSecuredStatusTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetDoorSecuredStatusTagged(structType interface{}) BACnetDoorSecuredStatusTagged {
	if casted, ok := structType.(BACnetDoorSecuredStatusTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDoorSecuredStatusTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDoorSecuredStatusTagged) GetTypeName() string {
	return "BACnetDoorSecuredStatusTagged"
}

func (m *_BACnetDoorSecuredStatusTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetDoorSecuredStatusTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDoorSecuredStatusTaggedParse(theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetDoorSecuredStatusTagged, error) {
	return BACnetDoorSecuredStatusTaggedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetDoorSecuredStatusTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetDoorSecuredStatusTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDoorSecuredStatusTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDoorSecuredStatusTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetDoorSecuredStatusTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetDoorSecuredStatus_SECURED)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetDoorSecuredStatusTagged")
	}
	var value BACnetDoorSecuredStatus
	if _value != nil {
		value = _value.(BACnetDoorSecuredStatus)
	}

	if closeErr := readBuffer.CloseContext("BACnetDoorSecuredStatusTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDoorSecuredStatusTagged")
	}

	// Create the instance
	return &_BACnetDoorSecuredStatusTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetDoorSecuredStatusTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDoorSecuredStatusTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetDoorSecuredStatusTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDoorSecuredStatusTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDoorSecuredStatusTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDoorSecuredStatusTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetDoorSecuredStatusTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetDoorSecuredStatusTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetDoorSecuredStatusTagged) isBACnetDoorSecuredStatusTagged() bool {
	return true
}

func (m *_BACnetDoorSecuredStatusTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
