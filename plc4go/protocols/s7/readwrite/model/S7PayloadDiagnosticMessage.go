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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadDiagnosticMessage is the corresponding interface of S7PayloadDiagnosticMessage
type S7PayloadDiagnosticMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetEventId returns EventId (property field)
	GetEventId() uint16
	// GetPriorityClass returns PriorityClass (property field)
	GetPriorityClass() uint8
	// GetObNumber returns ObNumber (property field)
	GetObNumber() uint8
	// GetDatId returns DatId (property field)
	GetDatId() uint16
	// GetInfo1 returns Info1 (property field)
	GetInfo1() uint16
	// GetInfo2 returns Info2 (property field)
	GetInfo2() uint32
	// GetTimeStamp returns TimeStamp (property field)
	GetTimeStamp() DateAndTime
}

// S7PayloadDiagnosticMessageExactly can be used when we want exactly this type and not a type which fulfills S7PayloadDiagnosticMessage.
// This is useful for switch cases.
type S7PayloadDiagnosticMessageExactly interface {
	S7PayloadDiagnosticMessage
	isS7PayloadDiagnosticMessage() bool
}

// _S7PayloadDiagnosticMessage is the data-structure of this message
type _S7PayloadDiagnosticMessage struct {
	*_S7PayloadUserDataItem
	EventId       uint16
	PriorityClass uint8
	ObNumber      uint8
	DatId         uint16
	Info1         uint16
	Info2         uint32
	TimeStamp     DateAndTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadDiagnosticMessage) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadDiagnosticMessage) GetCpuFunctionType() uint8 {
	return 0x00
}

func (m *_S7PayloadDiagnosticMessage) GetCpuSubfunction() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadDiagnosticMessage) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
	m.DataLength = dataLength
}

func (m *_S7PayloadDiagnosticMessage) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadDiagnosticMessage) GetEventId() uint16 {
	return m.EventId
}

func (m *_S7PayloadDiagnosticMessage) GetPriorityClass() uint8 {
	return m.PriorityClass
}

func (m *_S7PayloadDiagnosticMessage) GetObNumber() uint8 {
	return m.ObNumber
}

func (m *_S7PayloadDiagnosticMessage) GetDatId() uint16 {
	return m.DatId
}

func (m *_S7PayloadDiagnosticMessage) GetInfo1() uint16 {
	return m.Info1
}

func (m *_S7PayloadDiagnosticMessage) GetInfo2() uint32 {
	return m.Info2
}

func (m *_S7PayloadDiagnosticMessage) GetTimeStamp() DateAndTime {
	return m.TimeStamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadDiagnosticMessage factory function for _S7PayloadDiagnosticMessage
func NewS7PayloadDiagnosticMessage(eventId uint16, priorityClass uint8, obNumber uint8, datId uint16, info1 uint16, info2 uint32, timeStamp DateAndTime, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadDiagnosticMessage {
	_result := &_S7PayloadDiagnosticMessage{
		EventId:                eventId,
		PriorityClass:          priorityClass,
		ObNumber:               obNumber,
		DatId:                  datId,
		Info1:                  info1,
		Info2:                  info2,
		TimeStamp:              timeStamp,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadDiagnosticMessage(structType any) S7PayloadDiagnosticMessage {
	if casted, ok := structType.(S7PayloadDiagnosticMessage); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadDiagnosticMessage); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadDiagnosticMessage) GetTypeName() string {
	return "S7PayloadDiagnosticMessage"
}

func (m *_S7PayloadDiagnosticMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (eventId)
	lengthInBits += 16

	// Simple field (priorityClass)
	lengthInBits += 8

	// Simple field (obNumber)
	lengthInBits += 8

	// Simple field (datId)
	lengthInBits += 16

	// Simple field (info1)
	lengthInBits += 16

	// Simple field (info2)
	lengthInBits += 32

	// Simple field (timeStamp)
	lengthInBits += m.TimeStamp.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_S7PayloadDiagnosticMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadDiagnosticMessageParse(ctx context.Context, theBytes []byte, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadDiagnosticMessage, error) {
	return S7PayloadDiagnosticMessageParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadDiagnosticMessageParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadDiagnosticMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7PayloadDiagnosticMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadDiagnosticMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (eventId)
	_eventId, _eventIdErr := readBuffer.ReadUint16("eventId", 16)
	if _eventIdErr != nil {
		return nil, errors.Wrap(_eventIdErr, "Error parsing 'eventId' field of S7PayloadDiagnosticMessage")
	}
	eventId := _eventId

	// Simple Field (priorityClass)
	_priorityClass, _priorityClassErr := readBuffer.ReadUint8("priorityClass", 8)
	if _priorityClassErr != nil {
		return nil, errors.Wrap(_priorityClassErr, "Error parsing 'priorityClass' field of S7PayloadDiagnosticMessage")
	}
	priorityClass := _priorityClass

	// Simple Field (obNumber)
	_obNumber, _obNumberErr := readBuffer.ReadUint8("obNumber", 8)
	if _obNumberErr != nil {
		return nil, errors.Wrap(_obNumberErr, "Error parsing 'obNumber' field of S7PayloadDiagnosticMessage")
	}
	obNumber := _obNumber

	// Simple Field (datId)
	_datId, _datIdErr := readBuffer.ReadUint16("datId", 16)
	if _datIdErr != nil {
		return nil, errors.Wrap(_datIdErr, "Error parsing 'datId' field of S7PayloadDiagnosticMessage")
	}
	datId := _datId

	// Simple Field (info1)
	_info1, _info1Err := readBuffer.ReadUint16("info1", 16)
	if _info1Err != nil {
		return nil, errors.Wrap(_info1Err, "Error parsing 'info1' field of S7PayloadDiagnosticMessage")
	}
	info1 := _info1

	// Simple Field (info2)
	_info2, _info2Err := readBuffer.ReadUint32("info2", 32)
	if _info2Err != nil {
		return nil, errors.Wrap(_info2Err, "Error parsing 'info2' field of S7PayloadDiagnosticMessage")
	}
	info2 := _info2

	// Simple Field (timeStamp)
	if pullErr := readBuffer.PullContext("timeStamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeStamp")
	}
	_timeStamp, _timeStampErr := DateAndTimeParseWithBuffer(ctx, readBuffer)
	if _timeStampErr != nil {
		return nil, errors.Wrap(_timeStampErr, "Error parsing 'timeStamp' field of S7PayloadDiagnosticMessage")
	}
	timeStamp := _timeStamp.(DateAndTime)
	if closeErr := readBuffer.CloseContext("timeStamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeStamp")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadDiagnosticMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadDiagnosticMessage")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadDiagnosticMessage{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		EventId:                eventId,
		PriorityClass:          priorityClass,
		ObNumber:               obNumber,
		DatId:                  datId,
		Info1:                  info1,
		Info2:                  info2,
		TimeStamp:              timeStamp,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadDiagnosticMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadDiagnosticMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadDiagnosticMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadDiagnosticMessage")
		}

		// Simple Field (eventId)
		eventId := uint16(m.GetEventId())
		_eventIdErr := writeBuffer.WriteUint16("eventId", 16, uint16((eventId)))
		if _eventIdErr != nil {
			return errors.Wrap(_eventIdErr, "Error serializing 'eventId' field")
		}

		// Simple Field (priorityClass)
		priorityClass := uint8(m.GetPriorityClass())
		_priorityClassErr := writeBuffer.WriteUint8("priorityClass", 8, uint8((priorityClass)))
		if _priorityClassErr != nil {
			return errors.Wrap(_priorityClassErr, "Error serializing 'priorityClass' field")
		}

		// Simple Field (obNumber)
		obNumber := uint8(m.GetObNumber())
		_obNumberErr := writeBuffer.WriteUint8("obNumber", 8, uint8((obNumber)))
		if _obNumberErr != nil {
			return errors.Wrap(_obNumberErr, "Error serializing 'obNumber' field")
		}

		// Simple Field (datId)
		datId := uint16(m.GetDatId())
		_datIdErr := writeBuffer.WriteUint16("datId", 16, uint16((datId)))
		if _datIdErr != nil {
			return errors.Wrap(_datIdErr, "Error serializing 'datId' field")
		}

		// Simple Field (info1)
		info1 := uint16(m.GetInfo1())
		_info1Err := writeBuffer.WriteUint16("info1", 16, uint16((info1)))
		if _info1Err != nil {
			return errors.Wrap(_info1Err, "Error serializing 'info1' field")
		}

		// Simple Field (info2)
		info2 := uint32(m.GetInfo2())
		_info2Err := writeBuffer.WriteUint32("info2", 32, uint32((info2)))
		if _info2Err != nil {
			return errors.Wrap(_info2Err, "Error serializing 'info2' field")
		}

		// Simple Field (timeStamp)
		if pushErr := writeBuffer.PushContext("timeStamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeStamp")
		}
		_timeStampErr := writeBuffer.WriteSerializable(ctx, m.GetTimeStamp())
		if popErr := writeBuffer.PopContext("timeStamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeStamp")
		}
		if _timeStampErr != nil {
			return errors.Wrap(_timeStampErr, "Error serializing 'timeStamp' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadDiagnosticMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadDiagnosticMessage")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadDiagnosticMessage) isS7PayloadDiagnosticMessage() bool {
	return true
}

func (m *_S7PayloadDiagnosticMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
