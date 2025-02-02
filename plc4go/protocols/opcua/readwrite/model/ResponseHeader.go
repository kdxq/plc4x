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

// ResponseHeader is the corresponding interface of ResponseHeader
type ResponseHeader interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() int64
	// GetRequestHandle returns RequestHandle (property field)
	GetRequestHandle() uint32
	// GetServiceResult returns ServiceResult (property field)
	GetServiceResult() StatusCode
	// GetServiceDiagnostics returns ServiceDiagnostics (property field)
	GetServiceDiagnostics() DiagnosticInfo
	// GetNoOfStringTable returns NoOfStringTable (property field)
	GetNoOfStringTable() int32
	// GetStringTable returns StringTable (property field)
	GetStringTable() []PascalString
	// GetAdditionalHeader returns AdditionalHeader (property field)
	GetAdditionalHeader() ExtensionObject
}

// ResponseHeaderExactly can be used when we want exactly this type and not a type which fulfills ResponseHeader.
// This is useful for switch cases.
type ResponseHeaderExactly interface {
	ResponseHeader
	isResponseHeader() bool
}

// _ResponseHeader is the data-structure of this message
type _ResponseHeader struct {
	*_ExtensionObjectDefinition
	Timestamp          int64
	RequestHandle      uint32
	ServiceResult      StatusCode
	ServiceDiagnostics DiagnosticInfo
	NoOfStringTable    int32
	StringTable        []PascalString
	AdditionalHeader   ExtensionObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ResponseHeader) GetIdentifier() string {
	return "394"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ResponseHeader) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ResponseHeader) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ResponseHeader) GetTimestamp() int64 {
	return m.Timestamp
}

func (m *_ResponseHeader) GetRequestHandle() uint32 {
	return m.RequestHandle
}

func (m *_ResponseHeader) GetServiceResult() StatusCode {
	return m.ServiceResult
}

func (m *_ResponseHeader) GetServiceDiagnostics() DiagnosticInfo {
	return m.ServiceDiagnostics
}

func (m *_ResponseHeader) GetNoOfStringTable() int32 {
	return m.NoOfStringTable
}

func (m *_ResponseHeader) GetStringTable() []PascalString {
	return m.StringTable
}

func (m *_ResponseHeader) GetAdditionalHeader() ExtensionObject {
	return m.AdditionalHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewResponseHeader factory function for _ResponseHeader
func NewResponseHeader(timestamp int64, requestHandle uint32, serviceResult StatusCode, serviceDiagnostics DiagnosticInfo, noOfStringTable int32, stringTable []PascalString, additionalHeader ExtensionObject) *_ResponseHeader {
	_result := &_ResponseHeader{
		Timestamp:                  timestamp,
		RequestHandle:              requestHandle,
		ServiceResult:              serviceResult,
		ServiceDiagnostics:         serviceDiagnostics,
		NoOfStringTable:            noOfStringTable,
		StringTable:                stringTable,
		AdditionalHeader:           additionalHeader,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastResponseHeader(structType any) ResponseHeader {
	if casted, ok := structType.(ResponseHeader); ok {
		return casted
	}
	if casted, ok := structType.(*ResponseHeader); ok {
		return *casted
	}
	return nil
}

func (m *_ResponseHeader) GetTypeName() string {
	return "ResponseHeader"
}

func (m *_ResponseHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (timestamp)
	lengthInBits += 64

	// Simple field (requestHandle)
	lengthInBits += 32

	// Simple field (serviceResult)
	lengthInBits += m.ServiceResult.GetLengthInBits(ctx)

	// Simple field (serviceDiagnostics)
	lengthInBits += m.ServiceDiagnostics.GetLengthInBits(ctx)

	// Simple field (noOfStringTable)
	lengthInBits += 32

	// Array field
	if len(m.StringTable) > 0 {
		for _curItem, element := range m.StringTable {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.StringTable), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (additionalHeader)
	lengthInBits += m.AdditionalHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ResponseHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ResponseHeaderParse(ctx context.Context, theBytes []byte, identifier string) (ResponseHeader, error) {
	return ResponseHeaderParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ResponseHeaderParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ResponseHeader, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ResponseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ResponseHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timestamp)
	_timestamp, _timestampErr := readBuffer.ReadInt64("timestamp", 64)
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field of ResponseHeader")
	}
	timestamp := _timestamp

	// Simple Field (requestHandle)
	_requestHandle, _requestHandleErr := readBuffer.ReadUint32("requestHandle", 32)
	if _requestHandleErr != nil {
		return nil, errors.Wrap(_requestHandleErr, "Error parsing 'requestHandle' field of ResponseHeader")
	}
	requestHandle := _requestHandle

	// Simple Field (serviceResult)
	if pullErr := readBuffer.PullContext("serviceResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serviceResult")
	}
	_serviceResult, _serviceResultErr := StatusCodeParseWithBuffer(ctx, readBuffer)
	if _serviceResultErr != nil {
		return nil, errors.Wrap(_serviceResultErr, "Error parsing 'serviceResult' field of ResponseHeader")
	}
	serviceResult := _serviceResult.(StatusCode)
	if closeErr := readBuffer.CloseContext("serviceResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serviceResult")
	}

	// Simple Field (serviceDiagnostics)
	if pullErr := readBuffer.PullContext("serviceDiagnostics"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serviceDiagnostics")
	}
	_serviceDiagnostics, _serviceDiagnosticsErr := DiagnosticInfoParseWithBuffer(ctx, readBuffer)
	if _serviceDiagnosticsErr != nil {
		return nil, errors.Wrap(_serviceDiagnosticsErr, "Error parsing 'serviceDiagnostics' field of ResponseHeader")
	}
	serviceDiagnostics := _serviceDiagnostics.(DiagnosticInfo)
	if closeErr := readBuffer.CloseContext("serviceDiagnostics"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serviceDiagnostics")
	}

	// Simple Field (noOfStringTable)
	_noOfStringTable, _noOfStringTableErr := readBuffer.ReadInt32("noOfStringTable", 32)
	if _noOfStringTableErr != nil {
		return nil, errors.Wrap(_noOfStringTableErr, "Error parsing 'noOfStringTable' field of ResponseHeader")
	}
	noOfStringTable := _noOfStringTable

	// Array field (stringTable)
	if pullErr := readBuffer.PullContext("stringTable", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for stringTable")
	}
	// Count array
	stringTable := make([]PascalString, utils.Max(noOfStringTable, 0))
	// This happens when the size is set conditional to 0
	if len(stringTable) == 0 {
		stringTable = nil
	}
	{
		_numItems := uint16(utils.Max(noOfStringTable, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'stringTable' field of ResponseHeader")
			}
			stringTable[_curItem] = _item.(PascalString)
		}
	}
	if closeErr := readBuffer.CloseContext("stringTable", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for stringTable")
	}

	// Simple Field (additionalHeader)
	if pullErr := readBuffer.PullContext("additionalHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for additionalHeader")
	}
	_additionalHeader, _additionalHeaderErr := ExtensionObjectParseWithBuffer(ctx, readBuffer, bool(bool(true)))
	if _additionalHeaderErr != nil {
		return nil, errors.Wrap(_additionalHeaderErr, "Error parsing 'additionalHeader' field of ResponseHeader")
	}
	additionalHeader := _additionalHeader.(ExtensionObject)
	if closeErr := readBuffer.CloseContext("additionalHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for additionalHeader")
	}

	if closeErr := readBuffer.CloseContext("ResponseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ResponseHeader")
	}

	// Create a partially initialized instance
	_child := &_ResponseHeader{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		Timestamp:                  timestamp,
		RequestHandle:              requestHandle,
		ServiceResult:              serviceResult,
		ServiceDiagnostics:         serviceDiagnostics,
		NoOfStringTable:            noOfStringTable,
		StringTable:                stringTable,
		AdditionalHeader:           additionalHeader,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ResponseHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ResponseHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ResponseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ResponseHeader")
		}

		// Simple Field (timestamp)
		timestamp := int64(m.GetTimestamp())
		_timestampErr := writeBuffer.WriteInt64("timestamp", 64, int64((timestamp)))
		if _timestampErr != nil {
			return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
		}

		// Simple Field (requestHandle)
		requestHandle := uint32(m.GetRequestHandle())
		_requestHandleErr := writeBuffer.WriteUint32("requestHandle", 32, uint32((requestHandle)))
		if _requestHandleErr != nil {
			return errors.Wrap(_requestHandleErr, "Error serializing 'requestHandle' field")
		}

		// Simple Field (serviceResult)
		if pushErr := writeBuffer.PushContext("serviceResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serviceResult")
		}
		_serviceResultErr := writeBuffer.WriteSerializable(ctx, m.GetServiceResult())
		if popErr := writeBuffer.PopContext("serviceResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serviceResult")
		}
		if _serviceResultErr != nil {
			return errors.Wrap(_serviceResultErr, "Error serializing 'serviceResult' field")
		}

		// Simple Field (serviceDiagnostics)
		if pushErr := writeBuffer.PushContext("serviceDiagnostics"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serviceDiagnostics")
		}
		_serviceDiagnosticsErr := writeBuffer.WriteSerializable(ctx, m.GetServiceDiagnostics())
		if popErr := writeBuffer.PopContext("serviceDiagnostics"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serviceDiagnostics")
		}
		if _serviceDiagnosticsErr != nil {
			return errors.Wrap(_serviceDiagnosticsErr, "Error serializing 'serviceDiagnostics' field")
		}

		// Simple Field (noOfStringTable)
		noOfStringTable := int32(m.GetNoOfStringTable())
		_noOfStringTableErr := writeBuffer.WriteInt32("noOfStringTable", 32, int32((noOfStringTable)))
		if _noOfStringTableErr != nil {
			return errors.Wrap(_noOfStringTableErr, "Error serializing 'noOfStringTable' field")
		}

		// Array Field (stringTable)
		if pushErr := writeBuffer.PushContext("stringTable", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for stringTable")
		}
		for _curItem, _element := range m.GetStringTable() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetStringTable()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'stringTable' field")
			}
		}
		if popErr := writeBuffer.PopContext("stringTable", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for stringTable")
		}

		// Simple Field (additionalHeader)
		if pushErr := writeBuffer.PushContext("additionalHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for additionalHeader")
		}
		_additionalHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetAdditionalHeader())
		if popErr := writeBuffer.PopContext("additionalHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for additionalHeader")
		}
		if _additionalHeaderErr != nil {
			return errors.Wrap(_additionalHeaderErr, "Error serializing 'additionalHeader' field")
		}

		if popErr := writeBuffer.PopContext("ResponseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ResponseHeader")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ResponseHeader) isResponseHeader() bool {
	return true
}

func (m *_ResponseHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
