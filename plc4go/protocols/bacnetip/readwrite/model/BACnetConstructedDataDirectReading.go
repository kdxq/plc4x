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

// BACnetConstructedDataDirectReading is the corresponding interface of BACnetConstructedDataDirectReading
type BACnetConstructedDataDirectReading interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDirectReading returns DirectReading (property field)
	GetDirectReading() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataDirectReadingExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDirectReading.
// This is useful for switch cases.
type BACnetConstructedDataDirectReadingExactly interface {
	BACnetConstructedDataDirectReading
	isBACnetConstructedDataDirectReading() bool
}

// _BACnetConstructedDataDirectReading is the data-structure of this message
type _BACnetConstructedDataDirectReading struct {
	*_BACnetConstructedData
	DirectReading BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDirectReading) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDirectReading) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DIRECT_READING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDirectReading) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDirectReading) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDirectReading) GetDirectReading() BACnetApplicationTagReal {
	return m.DirectReading
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDirectReading) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetDirectReading())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDirectReading factory function for _BACnetConstructedDataDirectReading
func NewBACnetConstructedDataDirectReading(directReading BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDirectReading {
	_result := &_BACnetConstructedDataDirectReading{
		DirectReading:          directReading,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDirectReading(structType interface{}) BACnetConstructedDataDirectReading {
	if casted, ok := structType.(BACnetConstructedDataDirectReading); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDirectReading); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDirectReading) GetTypeName() string {
	return "BACnetConstructedDataDirectReading"
}

func (m *_BACnetConstructedDataDirectReading) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDirectReading) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (directReading)
	lengthInBits += m.DirectReading.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDirectReading) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDirectReadingParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDirectReading, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDirectReading"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDirectReading")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (directReading)
	if pullErr := readBuffer.PullContext("directReading"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for directReading")
	}
	_directReading, _directReadingErr := BACnetApplicationTagParse(readBuffer)
	if _directReadingErr != nil {
		return nil, errors.Wrap(_directReadingErr, "Error parsing 'directReading' field of BACnetConstructedDataDirectReading")
	}
	directReading := _directReading.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("directReading"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for directReading")
	}

	// Virtual field
	_actualValue := directReading
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDirectReading"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDirectReading")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDirectReading{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DirectReading: directReading,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDirectReading) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDirectReading"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDirectReading")
		}

		// Simple Field (directReading)
		if pushErr := writeBuffer.PushContext("directReading"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for directReading")
		}
		_directReadingErr := writeBuffer.WriteSerializable(m.GetDirectReading())
		if popErr := writeBuffer.PopContext("directReading"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for directReading")
		}
		if _directReadingErr != nil {
			return errors.Wrap(_directReadingErr, "Error serializing 'directReading' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDirectReading"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDirectReading")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDirectReading) isBACnetConstructedDataDirectReading() bool {
	return true
}

func (m *_BACnetConstructedDataDirectReading) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}