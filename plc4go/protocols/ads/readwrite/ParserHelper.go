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

	"github.com/apache/plc4x/plc4go/protocols/ads/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type AdsParserHelper struct {
}

func (m AdsParserHelper) Parse(typeName string, arguments []string, io utils.ReadBuffer) (interface{}, error) {
	switch typeName {
	case "AmsSerialFrame":
		return model.AmsSerialFrameParseWithBuffer(context.Background(), io)
	case "DataItem":
		plcValueType, _ := model.PlcValueTypeByName(arguments[0])
		stringLength, err := utils.StrToInt32(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.DataItemParseWithBuffer(context.Background(), io, plcValueType, stringLength)
	case "AdsTableSizes":
		return model.AdsTableSizesParseWithBuffer(context.Background(), io)
	case "AdsMultiRequestItem":
		indexGroup, err := utils.StrToUint32(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.AdsMultiRequestItemParseWithBuffer(context.Background(), io, indexGroup)
	case "AmsSerialAcknowledgeFrame":
		return model.AmsSerialAcknowledgeFrameParseWithBuffer(context.Background(), io)
	case "AdsDataTypeArrayInfo":
		return model.AdsDataTypeArrayInfoParseWithBuffer(context.Background(), io)
	case "AdsDataTypeTableEntry":
		return model.AdsDataTypeTableEntryParseWithBuffer(context.Background(), io)
	case "AmsNetId":
		return model.AmsNetIdParseWithBuffer(context.Background(), io)
	case "AdsStampHeader":
		return model.AdsStampHeaderParseWithBuffer(context.Background(), io)
	case "AmsSerialResetFrame":
		return model.AmsSerialResetFrameParseWithBuffer(context.Background(), io)
	case "AdsDataTypeTableChildEntry":
		return model.AdsDataTypeTableChildEntryParseWithBuffer(context.Background(), io)
	case "AdsConstants":
		return model.AdsConstantsParseWithBuffer(context.Background(), io)
	case "AdsNotificationSample":
		return model.AdsNotificationSampleParseWithBuffer(context.Background(), io)
	case "AdsSymbolTableEntry":
		return model.AdsSymbolTableEntryParseWithBuffer(context.Background(), io)
	case "AmsTCPPacket":
		return model.AmsTCPPacketParseWithBuffer(context.Background(), io)
	case "AmsPacket":
		return model.AmsPacketParseWithBuffer(context.Background(), io)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
