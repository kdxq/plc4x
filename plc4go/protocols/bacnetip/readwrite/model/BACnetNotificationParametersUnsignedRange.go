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

// BACnetNotificationParametersUnsignedRange is the corresponding interface of BACnetNotificationParametersUnsignedRange
type BACnetNotificationParametersUnsignedRange interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() BACnetContextTagUnsignedInteger
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetExceededLimit returns ExceededLimit (property field)
	GetExceededLimit() BACnetContextTagUnsignedInteger
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetNotificationParametersUnsignedRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersUnsignedRange.
// This is useful for switch cases.
type BACnetNotificationParametersUnsignedRangeExactly interface {
	BACnetNotificationParametersUnsignedRange
	isBACnetNotificationParametersUnsignedRange() bool
}

// _BACnetNotificationParametersUnsignedRange is the data-structure of this message
type _BACnetNotificationParametersUnsignedRange struct {
	*_BACnetNotificationParameters
	InnerOpeningTag BACnetOpeningTag
	SequenceNumber  BACnetContextTagUnsignedInteger
	StatusFlags     BACnetStatusFlagsTagged
	ExceededLimit   BACnetContextTagUnsignedInteger
	InnerClosingTag BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersUnsignedRange) InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersUnsignedRange) GetParent() BACnetNotificationParameters {
	return m._BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersUnsignedRange) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersUnsignedRange) GetSequenceNumber() BACnetContextTagUnsignedInteger {
	return m.SequenceNumber
}

func (m *_BACnetNotificationParametersUnsignedRange) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersUnsignedRange) GetExceededLimit() BACnetContextTagUnsignedInteger {
	return m.ExceededLimit
}

func (m *_BACnetNotificationParametersUnsignedRange) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersUnsignedRange factory function for _BACnetNotificationParametersUnsignedRange
func NewBACnetNotificationParametersUnsignedRange(innerOpeningTag BACnetOpeningTag, sequenceNumber BACnetContextTagUnsignedInteger, statusFlags BACnetStatusFlagsTagged, exceededLimit BACnetContextTagUnsignedInteger, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersUnsignedRange {
	_result := &_BACnetNotificationParametersUnsignedRange{
		InnerOpeningTag:               innerOpeningTag,
		SequenceNumber:                sequenceNumber,
		StatusFlags:                   statusFlags,
		ExceededLimit:                 exceededLimit,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersUnsignedRange(structType interface{}) BACnetNotificationParametersUnsignedRange {
	if casted, ok := structType.(BACnetNotificationParametersUnsignedRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersUnsignedRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersUnsignedRange) GetTypeName() string {
	return "BACnetNotificationParametersUnsignedRange"
}

func (m *_BACnetNotificationParametersUnsignedRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (sequenceNumber)
	lengthInBits += m.SequenceNumber.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (exceededLimit)
	lengthInBits += m.ExceededLimit.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersUnsignedRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersUnsignedRangeParse(theBytes []byte, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersUnsignedRange, error) {
	return BACnetNotificationParametersUnsignedRangeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber, tagNumber, objectTypeArgument)
}

func BACnetNotificationParametersUnsignedRangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersUnsignedRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersUnsignedRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersUnsignedRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field of BACnetNotificationParametersUnsignedRange")
	}
	innerOpeningTag := _innerOpeningTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Simple Field (sequenceNumber)
	if pullErr := readBuffer.PullContext("sequenceNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sequenceNumber")
	}
	_sequenceNumber, _sequenceNumberErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field of BACnetNotificationParametersUnsignedRange")
	}
	sequenceNumber := _sequenceNumber.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("sequenceNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sequenceNumber")
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusFlags")
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field of BACnetNotificationParametersUnsignedRange")
	}
	statusFlags := _statusFlags.(BACnetStatusFlagsTagged)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusFlags")
	}

	// Simple Field (exceededLimit)
	if pullErr := readBuffer.PullContext("exceededLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for exceededLimit")
	}
	_exceededLimit, _exceededLimitErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _exceededLimitErr != nil {
		return nil, errors.Wrap(_exceededLimitErr, "Error parsing 'exceededLimit' field of BACnetNotificationParametersUnsignedRange")
	}
	exceededLimit := _exceededLimit.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("exceededLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for exceededLimit")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field of BACnetNotificationParametersUnsignedRange")
	}
	innerClosingTag := _innerClosingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersUnsignedRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersUnsignedRange")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersUnsignedRange{
		_BACnetNotificationParameters: &_BACnetNotificationParameters{
			TagNumber:          tagNumber,
			ObjectTypeArgument: objectTypeArgument,
		},
		InnerOpeningTag: innerOpeningTag,
		SequenceNumber:  sequenceNumber,
		StatusFlags:     statusFlags,
		ExceededLimit:   exceededLimit,
		InnerClosingTag: innerClosingTag,
	}
	_child._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersUnsignedRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersUnsignedRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersUnsignedRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersUnsignedRange")
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
		}
		_innerOpeningTagErr := writeBuffer.WriteSerializable(ctx, m.GetInnerOpeningTag())
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerOpeningTag")
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (sequenceNumber)
		if pushErr := writeBuffer.PushContext("sequenceNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for sequenceNumber")
		}
		_sequenceNumberErr := writeBuffer.WriteSerializable(ctx, m.GetSequenceNumber())
		if popErr := writeBuffer.PopContext("sequenceNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for sequenceNumber")
		}
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusFlags")
		}
		_statusFlagsErr := writeBuffer.WriteSerializable(ctx, m.GetStatusFlags())
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusFlags")
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}

		// Simple Field (exceededLimit)
		if pushErr := writeBuffer.PushContext("exceededLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for exceededLimit")
		}
		_exceededLimitErr := writeBuffer.WriteSerializable(ctx, m.GetExceededLimit())
		if popErr := writeBuffer.PopContext("exceededLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for exceededLimit")
		}
		if _exceededLimitErr != nil {
			return errors.Wrap(_exceededLimitErr, "Error serializing 'exceededLimit' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
		}
		_innerClosingTagErr := writeBuffer.WriteSerializable(ctx, m.GetInnerClosingTag())
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerClosingTag")
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersUnsignedRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersUnsignedRange")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersUnsignedRange) isBACnetNotificationParametersUnsignedRange() bool {
	return true
}

func (m *_BACnetNotificationParametersUnsignedRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
