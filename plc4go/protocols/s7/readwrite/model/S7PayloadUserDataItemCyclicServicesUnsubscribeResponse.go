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
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadUserDataItemCyclicServicesUnsubscribeResponse is the corresponding interface of S7PayloadUserDataItemCyclicServicesUnsubscribeResponse
type S7PayloadUserDataItemCyclicServicesUnsubscribeResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
}

// S7PayloadUserDataItemCyclicServicesUnsubscribeResponseExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCyclicServicesUnsubscribeResponse.
// This is useful for switch cases.
type S7PayloadUserDataItemCyclicServicesUnsubscribeResponseExactly interface {
	S7PayloadUserDataItemCyclicServicesUnsubscribeResponse
	isS7PayloadUserDataItemCyclicServicesUnsubscribeResponse() bool
}

// _S7PayloadUserDataItemCyclicServicesUnsubscribeResponse is the data-structure of this message
type _S7PayloadUserDataItemCyclicServicesUnsubscribeResponse struct {
	*_S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetCpuFunctionGroup() uint8 {
	return 0x02
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetCpuSubfunction() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
	m.DataLength = dataLength
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

// NewS7PayloadUserDataItemCyclicServicesUnsubscribeResponse factory function for _S7PayloadUserDataItemCyclicServicesUnsubscribeResponse
func NewS7PayloadUserDataItemCyclicServicesUnsubscribeResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse {
	_result := &_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse{
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCyclicServicesUnsubscribeResponse(structType any) S7PayloadUserDataItemCyclicServicesUnsubscribeResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCyclicServicesUnsubscribeResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCyclicServicesUnsubscribeResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCyclicServicesUnsubscribeResponse"
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataItemCyclicServicesUnsubscribeResponseParse(theBytes []byte, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCyclicServicesUnsubscribeResponse, error) {
	return S7PayloadUserDataItemCyclicServicesUnsubscribeResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemCyclicServicesUnsubscribeResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCyclicServicesUnsubscribeResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCyclicServicesUnsubscribeResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCyclicServicesUnsubscribeResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCyclicServicesUnsubscribeResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCyclicServicesUnsubscribeResponse")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCyclicServicesUnsubscribeResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCyclicServicesUnsubscribeResponse")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCyclicServicesUnsubscribeResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCyclicServicesUnsubscribeResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) isS7PayloadUserDataItemCyclicServicesUnsubscribeResponse() bool {
	return true
}

func (m *_S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
