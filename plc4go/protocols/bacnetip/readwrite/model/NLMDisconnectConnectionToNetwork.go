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

// NLMDisconnectConnectionToNetwork is the data-structure of this message
type NLMDisconnectConnectionToNetwork struct {
	*NLM
	DestinationNetworkAddress uint16

	// Arguments.
	ApduLength uint16
}

// INLMDisconnectConnectionToNetwork is the corresponding interface of NLMDisconnectConnectionToNetwork
type INLMDisconnectConnectionToNetwork interface {
	INLM
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

func (m *NLMDisconnectConnectionToNetwork) GetMessageType() uint8 {
	return 0x09
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *NLMDisconnectConnectionToNetwork) InitializeParent(parent *NLM, vendorId *BACnetVendorId) {
	m.NLM.VendorId = vendorId
}

func (m *NLMDisconnectConnectionToNetwork) GetParent() *NLM {
	return m.NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *NLMDisconnectConnectionToNetwork) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMDisconnectConnectionToNetwork factory function for NLMDisconnectConnectionToNetwork
func NewNLMDisconnectConnectionToNetwork(destinationNetworkAddress uint16, vendorId *BACnetVendorId, apduLength uint16) *NLMDisconnectConnectionToNetwork {
	_result := &NLMDisconnectConnectionToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		NLM:                       NewNLM(vendorId, apduLength),
	}
	_result.Child = _result
	return _result
}

func CastNLMDisconnectConnectionToNetwork(structType interface{}) *NLMDisconnectConnectionToNetwork {
	if casted, ok := structType.(NLMDisconnectConnectionToNetwork); ok {
		return &casted
	}
	if casted, ok := structType.(*NLMDisconnectConnectionToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(NLM); ok {
		return CastNLMDisconnectConnectionToNetwork(casted.Child)
	}
	if casted, ok := structType.(*NLM); ok {
		return CastNLMDisconnectConnectionToNetwork(casted.Child)
	}
	return nil
}

func (m *NLMDisconnectConnectionToNetwork) GetTypeName() string {
	return "NLMDisconnectConnectionToNetwork"
}

func (m *NLMDisconnectConnectionToNetwork) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NLMDisconnectConnectionToNetwork) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	return lengthInBits
}

func (m *NLMDisconnectConnectionToNetwork) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMDisconnectConnectionToNetworkParse(readBuffer utils.ReadBuffer, apduLength uint16, messageType uint8) (*NLMDisconnectConnectionToNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMDisconnectConnectionToNetwork"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (destinationNetworkAddress)
	_destinationNetworkAddress, _destinationNetworkAddressErr := readBuffer.ReadUint16("destinationNetworkAddress", 16)
	if _destinationNetworkAddressErr != nil {
		return nil, errors.Wrap(_destinationNetworkAddressErr, "Error parsing 'destinationNetworkAddress' field")
	}
	destinationNetworkAddress := _destinationNetworkAddress

	if closeErr := readBuffer.CloseContext("NLMDisconnectConnectionToNetwork"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &NLMDisconnectConnectionToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		NLM:                       &NLM{},
	}
	_child.NLM.Child = _child
	return _child, nil
}

func (m *NLMDisconnectConnectionToNetwork) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMDisconnectConnectionToNetwork"); pushErr != nil {
			return pushErr
		}

		// Simple Field (destinationNetworkAddress)
		destinationNetworkAddress := uint16(m.DestinationNetworkAddress)
		_destinationNetworkAddressErr := writeBuffer.WriteUint16("destinationNetworkAddress", 16, (destinationNetworkAddress))
		if _destinationNetworkAddressErr != nil {
			return errors.Wrap(_destinationNetworkAddressErr, "Error serializing 'destinationNetworkAddress' field")
		}

		if popErr := writeBuffer.PopContext("NLMDisconnectConnectionToNetwork"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *NLMDisconnectConnectionToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
