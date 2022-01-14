/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetConstructedData struct {
	OpeningTag *BACnetOpeningTag
	Data       []*BACnetConstructedDataElement
	ClosingTag *BACnetClosingTag
	HasData    bool
	Child      IBACnetConstructedDataChild
}

// The corresponding interface
type IBACnetConstructedData interface {
	ObjectType() BACnetObjectType
	PropertyIdentifierArgument() IBACnetContextTagPropertyIdentifier
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetConstructedDataParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetConstructedData, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetConstructedDataChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, data []*BACnetConstructedDataElement, closingTag *BACnetClosingTag, hasData bool)
	GetTypeName() string
	IBACnetConstructedData
}

func NewBACnetConstructedData(openingTag *BACnetOpeningTag, data []*BACnetConstructedDataElement, closingTag *BACnetClosingTag, hasData bool) *BACnetConstructedData {
	return &BACnetConstructedData{OpeningTag: openingTag, Data: data, ClosingTag: closingTag, HasData: hasData}
}

func CastBACnetConstructedData(structType interface{}) *BACnetConstructedData {
	castFunc := func(typ interface{}) *BACnetConstructedData {
		if casted, ok := typ.(BACnetConstructedData); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConstructedData); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConstructedData) GetTypeName() string {
	return "BACnetConstructedData"
}

func (m *BACnetConstructedData) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetConstructedData) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetConstructedData) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Optional Field (openingTag)
	if m.OpeningTag != nil {
		lengthInBits += (*m.OpeningTag).LengthInBits()
	}

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.LengthInBits()
		}
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (closingTag)
	if m.ClosingTag != nil {
		lengthInBits += (*m.ClosingTag).LengthInBits()
	}

	return lengthInBits
}

func (m *BACnetConstructedData) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConstructedDataParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, propertyIdentifierArgument *BACnetContextTagPropertyIdentifier) (*BACnetConstructedData, error) {
	if pullErr := readBuffer.PullContext("BACnetConstructedData"); pullErr != nil {
		return nil, pullErr
	}

	// Optional Field (openingTag) (Can be skipped, if a given expression evaluates to false)
	var openingTag *BACnetOpeningTag = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, tagNumber, BACnetDataType_OPENING_TAG)
		switch {
		case _err != nil && _err != utils.ParseAssertError:
			return nil, errors.Wrap(_err, "Error parsing 'openingTag' field")
		case _err == utils.ParseAssertError:
			readBuffer.Reset(currentPos)
		default:
			openingTag = CastBACnetOpeningTag(_val)
			if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	data := make([]*BACnetConstructedDataElement, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, bool((objectType) == (BACnetObjectType_LIFE_SAFETY_ZONE)), tagNumber)) {
			_item, _err := BACnetConstructedDataElementParse(readBuffer, objectType, propertyIdentifierArgument)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field")
			}
			data = append(data, CastBACnetConstructedDataElement(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_hasData := bool((len(data)) == (0))
	hasData := bool(_hasData)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetConstructedData
	var typeSwitchError error
	switch {
	case objectType == BACnetObjectType_LIFE_SAFETY_ZONE: // BACnetConstructedDataLifeSafetyZone
		_parent, typeSwitchError = BACnetConstructedDataLifeSafetyZoneParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument)
	case true: // BACnetConstructedDataUnspecified
		_parent, typeSwitchError = BACnetConstructedDataUnspecifiedParse(readBuffer, tagNumber, objectType, propertyIdentifierArgument, hasData)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Optional Field (closingTag) (Can be skipped, if a given expression evaluates to false)
	var closingTag *BACnetClosingTag = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, tagNumber, BACnetDataType_CLOSING_TAG)
		switch {
		case _err != nil && _err != utils.ParseAssertError:
			return nil, errors.Wrap(_err, "Error parsing 'closingTag' field")
		case _err == utils.ParseAssertError:
			readBuffer.Reset(currentPos)
		default:
			closingTag = CastBACnetClosingTag(_val)
			if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedData"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, openingTag, data, closingTag, hasData)
	return _parent, nil
}

func (m *BACnetConstructedData) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetConstructedData) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetConstructedData, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetConstructedData"); pushErr != nil {
		return pushErr
	}

	// Optional Field (openingTag) (Can be skipped, if the value is null)
	var openingTag *BACnetOpeningTag = nil
	if m.OpeningTag != nil {
		if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
			return pushErr
		}
		openingTag = m.OpeningTag
		_openingTagErr := openingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
			return popErr
		}
		if _openingTagErr != nil {
			return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
		}
	}

	// Array Field (data)
	if m.Data != nil {
		if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.Data {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'data' field")
			}
		}
		if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}
	// Virtual field
	if _hasDataErr := writeBuffer.WriteVirtual("hasData", m.HasData); _hasDataErr != nil {
		return errors.Wrap(_hasDataErr, "Error serializing 'hasData' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Optional Field (closingTag) (Can be skipped, if the value is null)
	var closingTag *BACnetClosingTag = nil
	if m.ClosingTag != nil {
		if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
			return pushErr
		}
		closingTag = m.ClosingTag
		_closingTagErr := closingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
			return popErr
		}
		if _closingTagErr != nil {
			return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetConstructedData"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetConstructedData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}