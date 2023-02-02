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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// SALData is the corresponding interface of SALData
type SALData interface {
	utils.LengthAware
	utils.Serializable
	// GetApplicationId returns ApplicationId (discriminator field)
	GetApplicationId() ApplicationId
	// GetSalData returns SalData (property field)
	GetSalData() SALData
}

// SALDataExactly can be used when we want exactly this type and not a type which fulfills SALData.
// This is useful for switch cases.
type SALDataExactly interface {
	SALData
	isSALData() bool
}

// _SALData is the data-structure of this message
type _SALData struct {
	_SALDataChildRequirements
	SalData SALData
}

type _SALDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetApplicationId() ApplicationId
}

type SALDataParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child SALData, serializeChildFunction func() error) error
	GetTypeName() string
}

type SALDataChild interface {
	utils.Serializable
	InitializeParent(parent SALData, salData SALData)
	GetParent() *SALData

	GetTypeName() string
	SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALData) GetSalData() SALData {
	return m.SalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALData factory function for _SALData
func NewSALData(salData SALData) *_SALData {
	return &_SALData{SalData: salData}
}

// Deprecated: use the interface for direct cast
func CastSALData(structType interface{}) SALData {
	if casted, ok := structType.(SALData); ok {
		return casted
	}
	if casted, ok := structType.(*SALData); ok {
		return *casted
	}
	return nil
}

func (m *_SALData) GetTypeName() string {
	return "SALData"
}

func (m *_SALData) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (salData)
	if m.SalData != nil {
		lengthInBits += m.SalData.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_SALData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SALDataParse(theBytes []byte, applicationId ApplicationId) (SALData, error) {
	return SALDataParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type SALDataChildSerializeRequirement interface {
		SALData
		InitializeParent(SALData, SALData)
		GetParent() SALData
	}
	var _childTemp interface{}
	var _child SALDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case applicationId == ApplicationId_RESERVED: // SALDataReserved
		_childTemp, typeSwitchError = SALDataReservedParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_FREE_USAGE: // SALDataFreeUsage
		_childTemp, typeSwitchError = SALDataFreeUsageParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_TEMPERATURE_BROADCAST: // SALDataTemperatureBroadcast
		_childTemp, typeSwitchError = SALDataTemperatureBroadcastParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_ROOM_CONTROL_SYSTEM: // SALDataRoomControlSystem
		_childTemp, typeSwitchError = SALDataRoomControlSystemParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_LIGHTING: // SALDataLighting
		_childTemp, typeSwitchError = SALDataLightingParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_VENTILATION: // SALDataVentilation
		_childTemp, typeSwitchError = SALDataVentilationParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_IRRIGATION_CONTROL: // SALDataIrrigationControl
		_childTemp, typeSwitchError = SALDataIrrigationControlParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_POOLS_SPAS_PONDS_FOUNTAINS_CONTROL: // SALDataPoolsSpasPondsFountainsControl
		_childTemp, typeSwitchError = SALDataPoolsSpasPondsFountainsControlParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_HEATING: // SALDataHeating
		_childTemp, typeSwitchError = SALDataHeatingParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_AIR_CONDITIONING: // SALDataAirConditioning
		_childTemp, typeSwitchError = SALDataAirConditioningParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_TRIGGER_CONTROL: // SALDataTriggerControl
		_childTemp, typeSwitchError = SALDataTriggerControlParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_ENABLE_CONTROL: // SALDataEnableControl
		_childTemp, typeSwitchError = SALDataEnableControlParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_AUDIO_AND_VIDEO: // SALDataAudioAndVideo
		_childTemp, typeSwitchError = SALDataAudioAndVideoParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_SECURITY: // SALDataSecurity
		_childTemp, typeSwitchError = SALDataSecurityParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_METERING: // SALDataMetering
		_childTemp, typeSwitchError = SALDataMeteringParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_ACCESS_CONTROL: // SALDataAccessControl
		_childTemp, typeSwitchError = SALDataAccessControlParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_CLOCK_AND_TIMEKEEPING: // SALDataClockAndTimekeeping
		_childTemp, typeSwitchError = SALDataClockAndTimekeepingParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_TELEPHONY_STATUS_AND_CONTROL: // SALDataTelephonyStatusAndControl
		_childTemp, typeSwitchError = SALDataTelephonyStatusAndControlParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_MEASUREMENT: // SALDataMeasurement
		_childTemp, typeSwitchError = SALDataMeasurementParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_TESTING: // SALDataTesting
		_childTemp, typeSwitchError = SALDataTestingParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_MEDIA_TRANSPORT_CONTROL: // SALDataMediaTransport
		_childTemp, typeSwitchError = SALDataMediaTransportParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_ERROR_REPORTING: // SALDataErrorReporting
		_childTemp, typeSwitchError = SALDataErrorReportingParseWithBuffer(ctx, readBuffer, applicationId)
	case applicationId == ApplicationId_HVAC_ACTUATOR: // SALDataHvacActuator
		_childTemp, typeSwitchError = SALDataHvacActuatorParseWithBuffer(ctx, readBuffer, applicationId)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [applicationId=%v]", applicationId)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of SALData")
	}
	_child = _childTemp.(SALDataChildSerializeRequirement)

	// Optional Field (salData) (Can be skipped, if a given expression evaluates to false)
	var salData SALData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("salData"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for salData")
		}
		_val, _err := SALDataParseWithBuffer(ctx, readBuffer, applicationId)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'salData' field of SALData")
		default:
			salData = _val.(SALData)
			if closeErr := readBuffer.CloseContext("salData"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for salData")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("SALData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALData")
	}

	// Finish initializing
	_child.InitializeParent(_child, salData)
	return _child, nil
}

func (pm *_SALData) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child SALData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("SALData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SALData")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Optional Field (salData) (Can be skipped, if the value is null)
	var salData SALData = nil
	if m.GetSalData() != nil {
		if pushErr := writeBuffer.PushContext("salData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for salData")
		}
		salData = m.GetSalData()
		_salDataErr := writeBuffer.WriteSerializable(ctx, salData)
		if popErr := writeBuffer.PopContext("salData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for salData")
		}
		if _salDataErr != nil {
			return errors.Wrap(_salDataErr, "Error serializing 'salData' field")
		}
	}

	if popErr := writeBuffer.PopContext("SALData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SALData")
	}
	return nil
}

func (m *_SALData) isSALData() bool {
	return true
}

func (m *_SALData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
