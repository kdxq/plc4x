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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataZoneTo is the corresponding interface of BACnetConstructedDataZoneTo
type BACnetConstructedDataZoneTo interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetZoneTo returns ZoneTo (property field)
	GetZoneTo() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
}

// BACnetConstructedDataZoneToExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataZoneTo.
// This is useful for switch cases.
type BACnetConstructedDataZoneToExactly interface {
	BACnetConstructedDataZoneTo
	isBACnetConstructedDataZoneTo() bool
}

// _BACnetConstructedDataZoneTo is the data-structure of this message
type _BACnetConstructedDataZoneTo struct {
	*_BACnetConstructedData
	ZoneTo BACnetDeviceObjectReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataZoneTo) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataZoneTo) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ZONE_TO
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataZoneTo) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataZoneTo) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataZoneTo) GetZoneTo() BACnetDeviceObjectReference {
	return m.ZoneTo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataZoneTo) GetActualValue() BACnetDeviceObjectReference {
	return CastBACnetDeviceObjectReference(m.GetZoneTo())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataZoneTo factory function for _BACnetConstructedDataZoneTo
func NewBACnetConstructedDataZoneTo(zoneTo BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataZoneTo {
	_result := &_BACnetConstructedDataZoneTo{
		ZoneTo:                 zoneTo,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataZoneTo(structType interface{}) BACnetConstructedDataZoneTo {
	if casted, ok := structType.(BACnetConstructedDataZoneTo); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataZoneTo); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataZoneTo) GetTypeName() string {
	return "BACnetConstructedDataZoneTo"
}

func (m *_BACnetConstructedDataZoneTo) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataZoneTo) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneTo)
	lengthInBits += m.ZoneTo.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataZoneTo) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataZoneToParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataZoneTo, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataZoneTo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataZoneTo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneTo)
	if pullErr := readBuffer.PullContext("zoneTo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneTo")
	}
	_zoneTo, _zoneToErr := BACnetDeviceObjectReferenceParse(readBuffer)
	if _zoneToErr != nil {
		return nil, errors.Wrap(_zoneToErr, "Error parsing 'zoneTo' field of BACnetConstructedDataZoneTo")
	}
	zoneTo := _zoneTo.(BACnetDeviceObjectReference)
	if closeErr := readBuffer.CloseContext("zoneTo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneTo")
	}

	// Virtual field
	_actualValue := zoneTo
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataZoneTo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataZoneTo")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataZoneTo{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ZoneTo: zoneTo,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataZoneTo) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataZoneTo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataZoneTo")
		}

		// Simple Field (zoneTo)
		if pushErr := writeBuffer.PushContext("zoneTo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneTo")
		}
		_zoneToErr := writeBuffer.WriteSerializable(m.GetZoneTo())
		if popErr := writeBuffer.PopContext("zoneTo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneTo")
		}
		if _zoneToErr != nil {
			return errors.Wrap(_zoneToErr, "Error serializing 'zoneTo' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataZoneTo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataZoneTo")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataZoneTo) isBACnetConstructedDataZoneTo() bool {
	return true
}

func (m *_BACnetConstructedDataZoneTo) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}