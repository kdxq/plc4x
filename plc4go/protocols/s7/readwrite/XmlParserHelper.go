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

package readwrite

import (
	"context"
	"strconv"
	"strings"

	"github.com/apache/plc4x/plc4go/protocols/s7/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type S7XmlParserHelper struct {
}

// Temporary imports to silent compiler warnings (TODO: migrate from static to emission based imports)
func init() {
	_ = strconv.ParseUint
	_ = strconv.ParseInt
	_ = strings.Join
	_ = utils.Dump
}

func (m S7XmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
	switch typeName {
	case "DataItem":
		// TODO: find a way to parse the sub types
		var dataProtocolId string
		parsedInt1, err := strconv.ParseInt(parserArguments[1], 10, 32)
		if err != nil {
			return nil, err
		}
		stringLength := int32(parsedInt1)
		return model.DataItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), dataProtocolId, stringLength)
	case "SzlId":
		return model.SzlIdParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AlarmMessageObjectAckType":
		return model.AlarmMessageObjectAckTypeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AlarmMessageAckPushType":
		return model.AlarmMessageAckPushTypeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "S7Message":
		return model.S7MessageParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "S7VarPayloadStatusItem":
		return model.S7VarPayloadStatusItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "S7Parameter":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		messageType := uint8(parsedUint0)
		return model.S7ParameterParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), messageType)
	case "S7DataAlarmMessage":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 4)
		if err != nil {
			return nil, err
		}
		cpuFunctionType := uint8(parsedUint0)
		return model.S7DataAlarmMessageParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cpuFunctionType)
	case "SzlDataTreeItem":
		return model.SzlDataTreeItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "COTPPacket":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		cotpLen := uint16(parsedUint0)
		return model.COTPPacketParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cotpLen)
	case "S7PayloadUserDataItem":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 4)
		if err != nil {
			return nil, err
		}
		cpuFunctionType := uint8(parsedUint0)
		parsedUint1, err := strconv.ParseUint(parserArguments[1], 10, 8)
		if err != nil {
			return nil, err
		}
		cpuSubfunction := uint8(parsedUint1)
		return model.S7PayloadUserDataItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cpuFunctionType, cpuSubfunction)
	case "DateAndTime":
		return model.DateAndTimeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "COTPParameter":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		rest := uint8(parsedUint0)
		return model.COTPParameterParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), rest)
	case "AlarmMessageObjectPushType":
		return model.AlarmMessageObjectPushTypeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "State":
		return model.StateParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AlarmMessagePushType":
		return model.AlarmMessagePushTypeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "TPKTPacket":
		return model.TPKTPacketParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AlarmMessageAckType":
		return model.AlarmMessageAckTypeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AssociatedValueType":
		return model.AssociatedValueTypeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AlarmMessageAckObjectPushType":
		return model.AlarmMessageAckObjectPushTypeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "S7Payload":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		messageType := uint8(parsedUint0)
		// TODO: find a way to parse the sub types
		var parameter model.S7Parameter
		return model.S7PayloadParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), messageType, parameter)
	case "S7VarRequestParameterItem":
		return model.S7VarRequestParameterItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "S7VarPayloadDataItem":
		return model.S7VarPayloadDataItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AlarmMessageQueryType":
		return model.AlarmMessageQueryTypeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AlarmMessageAckResponseType":
		return model.AlarmMessageAckResponseTypeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AlarmMessageObjectQueryType":
		return model.AlarmMessageObjectQueryTypeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "S7Address":
		return model.S7AddressParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "S7ParameterUserDataItem":
		return model.S7ParameterUserDataItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
