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

// BACnetNotificationParametersChangeOfDiscreteValueNewValue is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValue
type BACnetNotificationParametersChangeOfDiscreteValueNewValue interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetPeekedIsContextTag returns PeekedIsContextTag (virtual field)
	GetPeekedIsContextTag() bool
}

// BACnetNotificationParametersChangeOfDiscreteValueNewValueExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfDiscreteValueNewValue.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfDiscreteValueNewValueExactly interface {
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	isBACnetNotificationParametersChangeOfDiscreteValueNewValue() bool
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValue is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValue struct {
	_BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

type _BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
}

type BACnetNotificationParametersChangeOfDiscreteValueNewValueParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetNotificationParametersChangeOfDiscreteValueNewValue, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetNotificationParametersChangeOfDiscreteValueNewValueChild interface {
	utils.Serializable
	InitializeParent(parent BACnetNotificationParametersChangeOfDiscreteValueNewValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag)
	GetParent() *BACnetNotificationParametersChangeOfDiscreteValueNewValue

	GetTypeName() string
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetPeekedIsContextTag() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValue factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValue
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return &_BACnetNotificationParametersChangeOfDiscreteValueNewValue{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValue(structType interface{}) BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValue"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueParse(theBytes []byte, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValue, error) {
	return BACnetNotificationParametersChangeOfDiscreteValueNewValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetNotificationParametersChangeOfDiscreteValueNewValue")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Virtual field
	_peekedIsContextTag := bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))
	peekedIsContextTag := bool(_peekedIsContextTag)
	_ = peekedIsContextTag

	// Validation
	if !(bool((!(peekedIsContextTag))) || bool((bool(bool(peekedIsContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x6)))) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"unexpected opening or closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetNotificationParametersChangeOfDiscreteValueNewValueChildSerializeRequirement interface {
		BACnetNotificationParametersChangeOfDiscreteValueNewValue
		InitializeParent(BACnetNotificationParametersChangeOfDiscreteValueNewValue, BACnetOpeningTag, BACnetTagHeader, BACnetClosingTag)
		GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValue
	}
	var _childTemp interface{}
	var _child BACnetNotificationParametersChangeOfDiscreteValueNewValueChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == 0x1 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanParseWithBuffer(ctx, readBuffer, tagNumber)
	case peekedTagNumber == 0x2 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsignedParseWithBuffer(ctx, readBuffer, tagNumber)
	case peekedTagNumber == 0x3 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueIntegerParseWithBuffer(ctx, readBuffer, tagNumber)
	case peekedTagNumber == 0x9 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumeratedParseWithBuffer(ctx, readBuffer, tagNumber)
	case peekedTagNumber == 0x7 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringParseWithBuffer(ctx, readBuffer, tagNumber)
	case peekedTagNumber == 0x6 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetStringParseWithBuffer(ctx, readBuffer, tagNumber)
	case peekedTagNumber == 0xA && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateParseWithBuffer(ctx, readBuffer, tagNumber)
	case peekedTagNumber == 0xB && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeParseWithBuffer(ctx, readBuffer, tagNumber)
	case peekedTagNumber == 0xC && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifierParseWithBuffer(ctx, readBuffer, tagNumber)
	case peekedTagNumber == uint8(0) && peekedIsContextTag == bool(true): // BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetimeParseWithBuffer(ctx, readBuffer, tagNumber)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v, peekedIsContextTag=%v]", peekedTagNumber, peekedIsContextTag)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetNotificationParametersChangeOfDiscreteValueNewValue")
	}
	_child = _childTemp.(BACnetNotificationParametersChangeOfDiscreteValueNewValueChildSerializeRequirement)

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetNotificationParametersChangeOfDiscreteValueNewValue")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValue")
	}

	// Finish initializing
	_child.InitializeParent(_child, openingTag, peekedTagHeader, closingTag)
	return _child, nil
}

func (pm *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetNotificationParametersChangeOfDiscreteValueNewValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValue")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	if _peekedIsContextTagErr := writeBuffer.WriteVirtual(ctx, "peekedIsContextTag", m.GetPeekedIsContextTag()); _peekedIsContextTagErr != nil {
		return errors.Wrap(_peekedIsContextTagErr, "Error serializing 'peekedIsContextTag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) isBACnetNotificationParametersChangeOfDiscreteValueNewValue() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
