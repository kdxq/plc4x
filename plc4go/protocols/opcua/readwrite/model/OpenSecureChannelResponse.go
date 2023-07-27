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

// OpenSecureChannelResponse is the corresponding interface of OpenSecureChannelResponse
type OpenSecureChannelResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetServerProtocolVersion returns ServerProtocolVersion (property field)
	GetServerProtocolVersion() uint32
	// GetSecurityToken returns SecurityToken (property field)
	GetSecurityToken() ExtensionObjectDefinition
	// GetServerNonce returns ServerNonce (property field)
	GetServerNonce() PascalByteString
}

// OpenSecureChannelResponseExactly can be used when we want exactly this type and not a type which fulfills OpenSecureChannelResponse.
// This is useful for switch cases.
type OpenSecureChannelResponseExactly interface {
	OpenSecureChannelResponse
	isOpenSecureChannelResponse() bool
}

// _OpenSecureChannelResponse is the data-structure of this message
type _OpenSecureChannelResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader        ExtensionObjectDefinition
	ServerProtocolVersion uint32
	SecurityToken         ExtensionObjectDefinition
	ServerNonce           PascalByteString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpenSecureChannelResponse) GetIdentifier() string {
	return "449"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpenSecureChannelResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_OpenSecureChannelResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpenSecureChannelResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_OpenSecureChannelResponse) GetServerProtocolVersion() uint32 {
	return m.ServerProtocolVersion
}

func (m *_OpenSecureChannelResponse) GetSecurityToken() ExtensionObjectDefinition {
	return m.SecurityToken
}

func (m *_OpenSecureChannelResponse) GetServerNonce() PascalByteString {
	return m.ServerNonce
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOpenSecureChannelResponse factory function for _OpenSecureChannelResponse
func NewOpenSecureChannelResponse(responseHeader ExtensionObjectDefinition, serverProtocolVersion uint32, securityToken ExtensionObjectDefinition, serverNonce PascalByteString) *_OpenSecureChannelResponse {
	_result := &_OpenSecureChannelResponse{
		ResponseHeader:             responseHeader,
		ServerProtocolVersion:      serverProtocolVersion,
		SecurityToken:              securityToken,
		ServerNonce:                serverNonce,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastOpenSecureChannelResponse(structType any) OpenSecureChannelResponse {
	if casted, ok := structType.(OpenSecureChannelResponse); ok {
		return casted
	}
	if casted, ok := structType.(*OpenSecureChannelResponse); ok {
		return *casted
	}
	return nil
}

func (m *_OpenSecureChannelResponse) GetTypeName() string {
	return "OpenSecureChannelResponse"
}

func (m *_OpenSecureChannelResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (serverProtocolVersion)
	lengthInBits += 32

	// Simple field (securityToken)
	lengthInBits += m.SecurityToken.GetLengthInBits(ctx)

	// Simple field (serverNonce)
	lengthInBits += m.ServerNonce.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpenSecureChannelResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpenSecureChannelResponseParse(ctx context.Context, theBytes []byte, identifier string) (OpenSecureChannelResponse, error) {
	return OpenSecureChannelResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func OpenSecureChannelResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (OpenSecureChannelResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("OpenSecureChannelResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpenSecureChannelResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of OpenSecureChannelResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	// Simple Field (serverProtocolVersion)
	_serverProtocolVersion, _serverProtocolVersionErr := readBuffer.ReadUint32("serverProtocolVersion", 32)
	if _serverProtocolVersionErr != nil {
		return nil, errors.Wrap(_serverProtocolVersionErr, "Error parsing 'serverProtocolVersion' field of OpenSecureChannelResponse")
	}
	serverProtocolVersion := _serverProtocolVersion

	// Simple Field (securityToken)
	if pullErr := readBuffer.PullContext("securityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityToken")
	}
	_securityToken, _securityTokenErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("443"))
	if _securityTokenErr != nil {
		return nil, errors.Wrap(_securityTokenErr, "Error parsing 'securityToken' field of OpenSecureChannelResponse")
	}
	securityToken := _securityToken.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("securityToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityToken")
	}

	// Simple Field (serverNonce)
	if pullErr := readBuffer.PullContext("serverNonce"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serverNonce")
	}
	_serverNonce, _serverNonceErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _serverNonceErr != nil {
		return nil, errors.Wrap(_serverNonceErr, "Error parsing 'serverNonce' field of OpenSecureChannelResponse")
	}
	serverNonce := _serverNonce.(PascalByteString)
	if closeErr := readBuffer.CloseContext("serverNonce"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serverNonce")
	}

	if closeErr := readBuffer.CloseContext("OpenSecureChannelResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpenSecureChannelResponse")
	}

	// Create a partially initialized instance
	_child := &_OpenSecureChannelResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
		ServerProtocolVersion:      serverProtocolVersion,
		SecurityToken:              securityToken,
		ServerNonce:                serverNonce,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_OpenSecureChannelResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpenSecureChannelResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpenSecureChannelResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpenSecureChannelResponse")
		}

		// Simple Field (responseHeader)
		if pushErr := writeBuffer.PushContext("responseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for responseHeader")
		}
		_responseHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetResponseHeader())
		if popErr := writeBuffer.PopContext("responseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for responseHeader")
		}
		if _responseHeaderErr != nil {
			return errors.Wrap(_responseHeaderErr, "Error serializing 'responseHeader' field")
		}

		// Simple Field (serverProtocolVersion)
		serverProtocolVersion := uint32(m.GetServerProtocolVersion())
		_serverProtocolVersionErr := writeBuffer.WriteUint32("serverProtocolVersion", 32, (serverProtocolVersion))
		if _serverProtocolVersionErr != nil {
			return errors.Wrap(_serverProtocolVersionErr, "Error serializing 'serverProtocolVersion' field")
		}

		// Simple Field (securityToken)
		if pushErr := writeBuffer.PushContext("securityToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityToken")
		}
		_securityTokenErr := writeBuffer.WriteSerializable(ctx, m.GetSecurityToken())
		if popErr := writeBuffer.PopContext("securityToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityToken")
		}
		if _securityTokenErr != nil {
			return errors.Wrap(_securityTokenErr, "Error serializing 'securityToken' field")
		}

		// Simple Field (serverNonce)
		if pushErr := writeBuffer.PushContext("serverNonce"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverNonce")
		}
		_serverNonceErr := writeBuffer.WriteSerializable(ctx, m.GetServerNonce())
		if popErr := writeBuffer.PopContext("serverNonce"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverNonce")
		}
		if _serverNonceErr != nil {
			return errors.Wrap(_serverNonceErr, "Error serializing 'serverNonce' field")
		}

		if popErr := writeBuffer.PopContext("OpenSecureChannelResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpenSecureChannelResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpenSecureChannelResponse) isOpenSecureChannelResponse() bool {
	return true
}

func (m *_OpenSecureChannelResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}