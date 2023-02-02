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

// SysexCommandReportFirmwareRequest is the corresponding interface of SysexCommandReportFirmwareRequest
type SysexCommandReportFirmwareRequest interface {
	utils.LengthAware
	utils.Serializable
	SysexCommand
}

// SysexCommandReportFirmwareRequestExactly can be used when we want exactly this type and not a type which fulfills SysexCommandReportFirmwareRequest.
// This is useful for switch cases.
type SysexCommandReportFirmwareRequestExactly interface {
	SysexCommandReportFirmwareRequest
	isSysexCommandReportFirmwareRequest() bool
}

// _SysexCommandReportFirmwareRequest is the data-structure of this message
type _SysexCommandReportFirmwareRequest struct {
	*_SysexCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandReportFirmwareRequest) GetCommandType() uint8 {
	return 0x79
}

func (m *_SysexCommandReportFirmwareRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandReportFirmwareRequest) InitializeParent(parent SysexCommand) {}

func (m *_SysexCommandReportFirmwareRequest) GetParent() SysexCommand {
	return m._SysexCommand
}

// NewSysexCommandReportFirmwareRequest factory function for _SysexCommandReportFirmwareRequest
func NewSysexCommandReportFirmwareRequest() *_SysexCommandReportFirmwareRequest {
	_result := &_SysexCommandReportFirmwareRequest{
		_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandReportFirmwareRequest(structType interface{}) SysexCommandReportFirmwareRequest {
	if casted, ok := structType.(SysexCommandReportFirmwareRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandReportFirmwareRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandReportFirmwareRequest) GetTypeName() string {
	return "SysexCommandReportFirmwareRequest"
}

func (m *_SysexCommandReportFirmwareRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandReportFirmwareRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SysexCommandReportFirmwareRequestParse(theBytes []byte, response bool) (SysexCommandReportFirmwareRequest, error) {
	return SysexCommandReportFirmwareRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), response)
}

func SysexCommandReportFirmwareRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (SysexCommandReportFirmwareRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandReportFirmwareRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandReportFirmwareRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandReportFirmwareRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandReportFirmwareRequest")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandReportFirmwareRequest{
		_SysexCommand: &_SysexCommand{},
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandReportFirmwareRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandReportFirmwareRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandReportFirmwareRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandReportFirmwareRequest")
		}

		if popErr := writeBuffer.PopContext("SysexCommandReportFirmwareRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandReportFirmwareRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandReportFirmwareRequest) isSysexCommandReportFirmwareRequest() bool {
	return true
}

func (m *_SysexCommandReportFirmwareRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
