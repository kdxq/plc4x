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

// CALDataIdentify is the corresponding interface of CALDataIdentify
type CALDataIdentify interface {
	utils.LengthAware
	utils.Serializable
	CALData
	// GetAttribute returns Attribute (property field)
	GetAttribute() Attribute
}

// CALDataIdentifyExactly can be used when we want exactly this type and not a type which fulfills CALDataIdentify.
// This is useful for switch cases.
type CALDataIdentifyExactly interface {
	CALDataIdentify
	isCALDataIdentify() bool
}

// _CALDataIdentify is the data-structure of this message
type _CALDataIdentify struct {
	*_CALData
	Attribute Attribute
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataIdentify) InitializeParent(parent CALData, commandTypeContainer CALCommandTypeContainer, additionalData CALData) {
	m.CommandTypeContainer = commandTypeContainer
	m.AdditionalData = additionalData
}

func (m *_CALDataIdentify) GetParent() CALData {
	return m._CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataIdentify) GetAttribute() Attribute {
	return m.Attribute
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataIdentify factory function for _CALDataIdentify
func NewCALDataIdentify(attribute Attribute, commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALDataIdentify {
	_result := &_CALDataIdentify{
		Attribute: attribute,
		_CALData:  NewCALData(commandTypeContainer, additionalData, requestContext),
	}
	_result._CALData._CALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataIdentify(structType interface{}) CALDataIdentify {
	if casted, ok := structType.(CALDataIdentify); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataIdentify); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataIdentify) GetTypeName() string {
	return "CALDataIdentify"
}

func (m *_CALDataIdentify) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (attribute)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CALDataIdentify) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CALDataIdentifyParse(theBytes []byte, requestContext RequestContext) (CALDataIdentify, error) {
	return CALDataIdentifyParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), requestContext)
}

func CALDataIdentifyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, requestContext RequestContext) (CALDataIdentify, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataIdentify"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataIdentify")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (attribute)
	if pullErr := readBuffer.PullContext("attribute"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for attribute")
	}
	_attribute, _attributeErr := AttributeParseWithBuffer(ctx, readBuffer)
	if _attributeErr != nil {
		return nil, errors.Wrap(_attributeErr, "Error parsing 'attribute' field of CALDataIdentify")
	}
	attribute := _attribute
	if closeErr := readBuffer.CloseContext("attribute"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for attribute")
	}

	if closeErr := readBuffer.CloseContext("CALDataIdentify"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataIdentify")
	}

	// Create a partially initialized instance
	_child := &_CALDataIdentify{
		_CALData: &_CALData{
			RequestContext: requestContext,
		},
		Attribute: attribute,
	}
	_child._CALData._CALDataChildRequirements = _child
	return _child, nil
}

func (m *_CALDataIdentify) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataIdentify) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataIdentify"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataIdentify")
		}

		// Simple Field (attribute)
		if pushErr := writeBuffer.PushContext("attribute"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for attribute")
		}
		_attributeErr := writeBuffer.WriteSerializable(ctx, m.GetAttribute())
		if popErr := writeBuffer.PopContext("attribute"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for attribute")
		}
		if _attributeErr != nil {
			return errors.Wrap(_attributeErr, "Error serializing 'attribute' field")
		}

		if popErr := writeBuffer.PopContext("CALDataIdentify"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataIdentify")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataIdentify) isCALDataIdentify() bool {
	return true
}

func (m *_CALDataIdentify) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
