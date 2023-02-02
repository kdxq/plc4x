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

// ErrorClass is an enum
type ErrorClass uint16

type IErrorClass interface {
	utils.Serializable
}

const (
	ErrorClass_DEVICE                   ErrorClass = 0x0000
	ErrorClass_OBJECT                   ErrorClass = 0x0001
	ErrorClass_PROPERTY                 ErrorClass = 0x0002
	ErrorClass_RESOURCES                ErrorClass = 0x0003
	ErrorClass_SECURITY                 ErrorClass = 0x0004
	ErrorClass_SERVICES                 ErrorClass = 0x0005
	ErrorClass_VT                       ErrorClass = 0x0006
	ErrorClass_COMMUNICATION            ErrorClass = 0x0007
	ErrorClass_VENDOR_PROPRIETARY_VALUE ErrorClass = 0xFFFF
)

var ErrorClassValues []ErrorClass

func init() {
	_ = errors.New
	ErrorClassValues = []ErrorClass{
		ErrorClass_DEVICE,
		ErrorClass_OBJECT,
		ErrorClass_PROPERTY,
		ErrorClass_RESOURCES,
		ErrorClass_SECURITY,
		ErrorClass_SERVICES,
		ErrorClass_VT,
		ErrorClass_COMMUNICATION,
		ErrorClass_VENDOR_PROPRIETARY_VALUE,
	}
}

func ErrorClassByValue(value uint16) (enum ErrorClass, ok bool) {
	switch value {
	case 0xFFFF:
		return ErrorClass_VENDOR_PROPRIETARY_VALUE, true
	case 0x0000:
		return ErrorClass_DEVICE, true
	case 0x0001:
		return ErrorClass_OBJECT, true
	case 0x0002:
		return ErrorClass_PROPERTY, true
	case 0x0003:
		return ErrorClass_RESOURCES, true
	case 0x0004:
		return ErrorClass_SECURITY, true
	case 0x0005:
		return ErrorClass_SERVICES, true
	case 0x0006:
		return ErrorClass_VT, true
	case 0x0007:
		return ErrorClass_COMMUNICATION, true
	}
	return 0, false
}

func ErrorClassByName(value string) (enum ErrorClass, ok bool) {
	switch value {
	case "VENDOR_PROPRIETARY_VALUE":
		return ErrorClass_VENDOR_PROPRIETARY_VALUE, true
	case "DEVICE":
		return ErrorClass_DEVICE, true
	case "OBJECT":
		return ErrorClass_OBJECT, true
	case "PROPERTY":
		return ErrorClass_PROPERTY, true
	case "RESOURCES":
		return ErrorClass_RESOURCES, true
	case "SECURITY":
		return ErrorClass_SECURITY, true
	case "SERVICES":
		return ErrorClass_SERVICES, true
	case "VT":
		return ErrorClass_VT, true
	case "COMMUNICATION":
		return ErrorClass_COMMUNICATION, true
	}
	return 0, false
}

func ErrorClassKnows(value uint16) bool {
	for _, typeValue := range ErrorClassValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastErrorClass(structType interface{}) ErrorClass {
	castFunc := func(typ interface{}) ErrorClass {
		if sErrorClass, ok := typ.(ErrorClass); ok {
			return sErrorClass
		}
		return 0
	}
	return castFunc(structType)
}

func (m ErrorClass) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m ErrorClass) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorClassParse(ctx context.Context, theBytes []byte) (ErrorClass, error) {
	return ErrorClassParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorClassParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorClass, error) {
	val, err := readBuffer.ReadUint16("ErrorClass", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ErrorClass")
	}
	if enum, ok := ErrorClassByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return ErrorClass(val), nil
	} else {
		return enum, nil
	}
}

func (e ErrorClass) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ErrorClass) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("ErrorClass", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ErrorClass) PLC4XEnumName() string {
	switch e {
	case ErrorClass_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case ErrorClass_DEVICE:
		return "DEVICE"
	case ErrorClass_OBJECT:
		return "OBJECT"
	case ErrorClass_PROPERTY:
		return "PROPERTY"
	case ErrorClass_RESOURCES:
		return "RESOURCES"
	case ErrorClass_SECURITY:
		return "SECURITY"
	case ErrorClass_SERVICES:
		return "SERVICES"
	case ErrorClass_VT:
		return "VT"
	case ErrorClass_COMMUNICATION:
		return "COMMUNICATION"
	}
	return ""
}

func (e ErrorClass) String() string {
	return e.PLC4XEnumName()
}
