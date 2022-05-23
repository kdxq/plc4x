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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// NLMRejectRouterToNetwork is the data-structure of this message
type NLMRejectRouterToNetwork struct {
	*NLM
	RejectReason              NLMRejectRouterToNetworkRejectReason
	DestinationNetworkAddress uint16

	// Arguments.
	ApduLength uint16
}

// INLMRejectRouterToNetwork is the corresponding interface of NLMRejectRouterToNetwork
type INLMRejectRouterToNetwork interface {
	INLM
	// GetRejectReason returns RejectReason (property field)
	GetRejectReason() NLMRejectRouterToNetworkRejectReason
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *NLMRejectRouterToNetwork) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *NLMRejectRouterToNetwork) InitializeParent(parent *NLM, vendorId *BACnetVendorId) {
	m.NLM.VendorId = vendorId
}

func (m *NLMRejectRouterToNetwork) GetParent() *NLM {
	return m.NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *NLMRejectRouterToNetwork) GetRejectReason() NLMRejectRouterToNetworkRejectReason {
	return m.RejectReason
}

func (m *NLMRejectRouterToNetwork) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMRejectRouterToNetwork factory function for NLMRejectRouterToNetwork
func NewNLMRejectRouterToNetwork(rejectReason NLMRejectRouterToNetworkRejectReason, destinationNetworkAddress uint16, vendorId *BACnetVendorId, apduLength uint16) *NLMRejectRouterToNetwork {
	_result := &NLMRejectRouterToNetwork{
		RejectReason:              rejectReason,
		DestinationNetworkAddress: destinationNetworkAddress,
		NLM:                       NewNLM(vendorId, apduLength),
	}
	_result.Child = _result
	return _result
}

func CastNLMRejectRouterToNetwork(structType interface{}) *NLMRejectRouterToNetwork {
	if casted, ok := structType.(NLMRejectRouterToNetwork); ok {
		return &casted
	}
	if casted, ok := structType.(*NLMRejectRouterToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(NLM); ok {
		return CastNLMRejectRouterToNetwork(casted.Child)
	}
	if casted, ok := structType.(*NLM); ok {
		return CastNLMRejectRouterToNetwork(casted.Child)
	}
	return nil
}

func (m *NLMRejectRouterToNetwork) GetTypeName() string {
	return "NLMRejectRouterToNetwork"
}

func (m *NLMRejectRouterToNetwork) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NLMRejectRouterToNetwork) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (rejectReason)
	lengthInBits += 8

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	return lengthInBits
}

func (m *NLMRejectRouterToNetwork) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMRejectRouterToNetworkParse(readBuffer utils.ReadBuffer, apduLength uint16, messageType uint8) (*NLMRejectRouterToNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMRejectRouterToNetwork"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (rejectReason)
	if pullErr := readBuffer.PullContext("rejectReason"); pullErr != nil {
		return nil, pullErr
	}
	_rejectReason, _rejectReasonErr := NLMRejectRouterToNetworkRejectReasonParse(readBuffer)
	if _rejectReasonErr != nil {
		return nil, errors.Wrap(_rejectReasonErr, "Error parsing 'rejectReason' field")
	}
	rejectReason := _rejectReason
	if closeErr := readBuffer.CloseContext("rejectReason"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (destinationNetworkAddress)
	_destinationNetworkAddress, _destinationNetworkAddressErr := readBuffer.ReadUint16("destinationNetworkAddress", 16)
	if _destinationNetworkAddressErr != nil {
		return nil, errors.Wrap(_destinationNetworkAddressErr, "Error parsing 'destinationNetworkAddress' field")
	}
	destinationNetworkAddress := _destinationNetworkAddress

	if closeErr := readBuffer.CloseContext("NLMRejectRouterToNetwork"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &NLMRejectRouterToNetwork{
		RejectReason:              rejectReason,
		DestinationNetworkAddress: destinationNetworkAddress,
		NLM:                       &NLM{},
	}
	_child.NLM.Child = _child
	return _child, nil
}

func (m *NLMRejectRouterToNetwork) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMRejectRouterToNetwork"); pushErr != nil {
			return pushErr
		}

		// Simple Field (rejectReason)
		if pushErr := writeBuffer.PushContext("rejectReason"); pushErr != nil {
			return pushErr
		}
		_rejectReasonErr := m.RejectReason.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("rejectReason"); popErr != nil {
			return popErr
		}
		if _rejectReasonErr != nil {
			return errors.Wrap(_rejectReasonErr, "Error serializing 'rejectReason' field")
		}

		// Simple Field (destinationNetworkAddress)
		destinationNetworkAddress := uint16(m.DestinationNetworkAddress)
		_destinationNetworkAddressErr := writeBuffer.WriteUint16("destinationNetworkAddress", 16, (destinationNetworkAddress))
		if _destinationNetworkAddressErr != nil {
			return errors.Wrap(_destinationNetworkAddressErr, "Error serializing 'destinationNetworkAddress' field")
		}

		if popErr := writeBuffer.PopContext("NLMRejectRouterToNetwork"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *NLMRejectRouterToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
