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

// BACnetAuthorizationMode is an enum
type BACnetAuthorizationMode uint16

type IBACnetAuthorizationMode interface {
	utils.Serializable
}

const (
	BACnetAuthorizationMode_AUTHORIZE                BACnetAuthorizationMode = 0
	BACnetAuthorizationMode_GRANT_ACTIVE             BACnetAuthorizationMode = 1
	BACnetAuthorizationMode_DENY_ALL                 BACnetAuthorizationMode = 2
	BACnetAuthorizationMode_VERIFICATION_REQUIRED    BACnetAuthorizationMode = 3
	BACnetAuthorizationMode_AUTHORIZATION_DELAYED    BACnetAuthorizationMode = 4
	BACnetAuthorizationMode_NONE                     BACnetAuthorizationMode = 5
	BACnetAuthorizationMode_VENDOR_PROPRIETARY_VALUE BACnetAuthorizationMode = 0xFFFF
)

var BACnetAuthorizationModeValues []BACnetAuthorizationMode

func init() {
	_ = errors.New
	BACnetAuthorizationModeValues = []BACnetAuthorizationMode{
		BACnetAuthorizationMode_AUTHORIZE,
		BACnetAuthorizationMode_GRANT_ACTIVE,
		BACnetAuthorizationMode_DENY_ALL,
		BACnetAuthorizationMode_VERIFICATION_REQUIRED,
		BACnetAuthorizationMode_AUTHORIZATION_DELAYED,
		BACnetAuthorizationMode_NONE,
		BACnetAuthorizationMode_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetAuthorizationModeByValue(value uint16) (enum BACnetAuthorizationMode, ok bool) {
	switch value {
	case 0:
		return BACnetAuthorizationMode_AUTHORIZE, true
	case 0xFFFF:
		return BACnetAuthorizationMode_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetAuthorizationMode_GRANT_ACTIVE, true
	case 2:
		return BACnetAuthorizationMode_DENY_ALL, true
	case 3:
		return BACnetAuthorizationMode_VERIFICATION_REQUIRED, true
	case 4:
		return BACnetAuthorizationMode_AUTHORIZATION_DELAYED, true
	case 5:
		return BACnetAuthorizationMode_NONE, true
	}
	return 0, false
}

func BACnetAuthorizationModeByName(value string) (enum BACnetAuthorizationMode, ok bool) {
	switch value {
	case "AUTHORIZE":
		return BACnetAuthorizationMode_AUTHORIZE, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetAuthorizationMode_VENDOR_PROPRIETARY_VALUE, true
	case "GRANT_ACTIVE":
		return BACnetAuthorizationMode_GRANT_ACTIVE, true
	case "DENY_ALL":
		return BACnetAuthorizationMode_DENY_ALL, true
	case "VERIFICATION_REQUIRED":
		return BACnetAuthorizationMode_VERIFICATION_REQUIRED, true
	case "AUTHORIZATION_DELAYED":
		return BACnetAuthorizationMode_AUTHORIZATION_DELAYED, true
	case "NONE":
		return BACnetAuthorizationMode_NONE, true
	}
	return 0, false
}

func BACnetAuthorizationModeKnows(value uint16) bool {
	for _, typeValue := range BACnetAuthorizationModeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAuthorizationMode(structType interface{}) BACnetAuthorizationMode {
	castFunc := func(typ interface{}) BACnetAuthorizationMode {
		if sBACnetAuthorizationMode, ok := typ.(BACnetAuthorizationMode); ok {
			return sBACnetAuthorizationMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAuthorizationMode) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetAuthorizationMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthorizationModeParse(ctx context.Context, theBytes []byte) (BACnetAuthorizationMode, error) {
	return BACnetAuthorizationModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAuthorizationModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthorizationMode, error) {
	val, err := readBuffer.ReadUint16("BACnetAuthorizationMode", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAuthorizationMode")
	}
	if enum, ok := BACnetAuthorizationModeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetAuthorizationMode(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAuthorizationMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAuthorizationMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetAuthorizationMode", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAuthorizationMode) PLC4XEnumName() string {
	switch e {
	case BACnetAuthorizationMode_AUTHORIZE:
		return "AUTHORIZE"
	case BACnetAuthorizationMode_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetAuthorizationMode_GRANT_ACTIVE:
		return "GRANT_ACTIVE"
	case BACnetAuthorizationMode_DENY_ALL:
		return "DENY_ALL"
	case BACnetAuthorizationMode_VERIFICATION_REQUIRED:
		return "VERIFICATION_REQUIRED"
	case BACnetAuthorizationMode_AUTHORIZATION_DELAYED:
		return "AUTHORIZATION_DELAYED"
	case BACnetAuthorizationMode_NONE:
		return "NONE"
	}
	return ""
}

func (e BACnetAuthorizationMode) String() string {
	return e.PLC4XEnumName()
}
