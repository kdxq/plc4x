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
package org.apache.plc4x.java.s7.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.math.BigInteger;
import java.time.*;
import java.time.temporal.ChronoField;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;
import org.apache.plc4x.java.spi.values.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

// Code generated by code-generation. DO NOT EDIT.

public class DataItem {

  private static final Logger LOGGER = LoggerFactory.getLogger(DataItem.class);

  public static PlcValue staticParse(
      ReadBuffer readBuffer,
      String dataProtocolId,
      ControllerType controllerType,
      Integer stringLength)
      throws ParseException {
    if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_BOOL")) { // BOOL
      Byte reservedField0 =
          readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

      boolean value = readSimpleField("value", readBoolean(readBuffer));
      return new PlcBOOL(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_BYTE")) { // BYTE
      short value = readSimpleField("value", readUnsignedShort(readBuffer, 8));
      return new PlcBYTE(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_WORD")) { // WORD
      int value = readSimpleField("value", readUnsignedInt(readBuffer, 16));
      return new PlcWORD(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_DWORD")) { // DWORD
      long value = readSimpleField("value", readUnsignedLong(readBuffer, 32));
      return new PlcDWORD(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_LWORD")) { // LWORD
      BigInteger value = readSimpleField("value", readUnsignedBigInteger(readBuffer, 64));
      return new PlcLWORD(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_SINT")) { // SINT
      byte value = readSimpleField("value", readSignedByte(readBuffer, 8));
      return new PlcSINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_USINT")) { // USINT
      short value = readSimpleField("value", readUnsignedShort(readBuffer, 8));
      return new PlcUSINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_INT")) { // INT
      short value = readSimpleField("value", readSignedShort(readBuffer, 16));
      return new PlcINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_UINT")) { // UINT
      int value = readSimpleField("value", readUnsignedInt(readBuffer, 16));
      return new PlcUINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_DINT")) { // DINT
      int value = readSimpleField("value", readSignedInt(readBuffer, 32));
      return new PlcDINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_UDINT")) { // UDINT
      long value = readSimpleField("value", readUnsignedLong(readBuffer, 32));
      return new PlcUDINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_LINT")) { // LINT
      long value = readSimpleField("value", readSignedLong(readBuffer, 64));
      return new PlcLINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_ULINT")) { // ULINT
      BigInteger value = readSimpleField("value", readUnsignedBigInteger(readBuffer, 64));
      return new PlcULINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_REAL")) { // REAL
      float value = readSimpleField("value", readFloat(readBuffer, 32));
      return new PlcREAL(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_LREAL")) { // LREAL
      double value = readSimpleField("value", readDouble(readBuffer, 64));
      return new PlcLREAL(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_CHAR")) { // CHAR
      String value =
          readSimpleField("value", readString(readBuffer, 8), WithOption.WithEncoding("UTF-8"));
      return new PlcCHAR(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_WCHAR")) { // CHAR
      String value =
          readSimpleField("value", readString(readBuffer, 16), WithOption.WithEncoding("UTF-16"));
      return new PlcCHAR(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_STRING")) { // STRING
      String value =
          readManualField(
              "value",
              readBuffer,
              () ->
                  (String)
                      (org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.parseS7String(
                          readBuffer, stringLength, "UTF-8")),
              WithOption.WithEncoding("UTF-8"));
      return new PlcSTRING(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_WSTRING")) { // STRING
      String value =
          readManualField(
              "value",
              readBuffer,
              () ->
                  (String)
                      (org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.parseS7String(
                          readBuffer, stringLength, "UTF-16")),
              WithOption.WithEncoding("UTF-16"));
      return new PlcSTRING(value);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_TIME")) { // TIME
      long milliseconds = readSimpleField("milliseconds", readUnsignedLong(readBuffer, 32));
      return PlcTIME.ofMilliseconds(milliseconds);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "S7_S5TIME")) { // TIME
      long milliseconds =
          readManualField(
              "milliseconds",
              readBuffer,
              () ->
                  (long)
                      (org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.parseS5Time(
                          readBuffer)));
      return PlcTIME.ofMilliseconds(milliseconds);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_LTIME")) { // LTIME
      BigInteger nanoseconds =
          readSimpleField("nanoseconds", readUnsignedBigInteger(readBuffer, 64));
      return PlcLTIME.ofNanoseconds(nanoseconds);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_DATE")) { // DATE
      int daysSinceEpoch =
          readManualField(
              "daysSinceEpoch",
              readBuffer,
              () ->
                  (int)
                      (org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.parseTiaDate(
                          readBuffer)));
      return PlcDATE.ofDaysSinceEpoch(daysSinceEpoch);
    } else if (EvaluationHelper.equals(
        dataProtocolId, (String) "IEC61131_TIME_OF_DAY")) { // TIME_OF_DAY
      long millisecondsSinceMidnight =
          readSimpleField("millisecondsSinceMidnight", readUnsignedLong(readBuffer, 32));
      return PlcTIME_OF_DAY.ofMillisecondsSinceMidnight(millisecondsSinceMidnight);
    } else if (EvaluationHelper.equals(
        dataProtocolId, (String) "IEC61131_LTIME_OF_DAY")) { // LTIME_OF_DAY
      BigInteger nanosecondsSinceMidnight =
          readSimpleField("nanosecondsSinceMidnight", readUnsignedBigInteger(readBuffer, 64));
      return PlcLTIME_OF_DAY.ofNanosecondsSinceMidnight(nanosecondsSinceMidnight);
    } else if (EvaluationHelper.equals(
        dataProtocolId, (String) "IEC61131_DATE_AND_TIME")) { // DATE_AND_TIME
      short year =
          readManualField(
              "year",
              readBuffer,
              () ->
                  (short)
                      (org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.parseSiemensYear(
                          readBuffer)),
              WithOption.WithEncoding("BCD"));

      short month =
          readSimpleField(
              "month", readUnsignedShort(readBuffer, 8), WithOption.WithEncoding("BCD"));

      short day =
          readSimpleField("day", readUnsignedShort(readBuffer, 8), WithOption.WithEncoding("BCD"));

      short hour =
          readSimpleField("hour", readUnsignedShort(readBuffer, 8), WithOption.WithEncoding("BCD"));

      short minutes =
          readSimpleField(
              "minutes", readUnsignedShort(readBuffer, 8), WithOption.WithEncoding("BCD"));

      short seconds =
          readSimpleField(
              "seconds", readUnsignedShort(readBuffer, 8), WithOption.WithEncoding("BCD"));

      short millisecondsOfSecond =
          readSimpleField(
              "millisecondsOfSecond",
              readUnsignedShort(readBuffer, 12),
              WithOption.WithEncoding("BCD"));

      byte dayOfWeek =
          readSimpleField(
              "dayOfWeek", readUnsignedByte(readBuffer, 4), WithOption.WithEncoding("BCD"));
      return PlcDATE_AND_TIME.ofSegments(
          year,
          (month == 0) ? 1 : month,
          (day == 0) ? 1 : day,
          hour,
          minutes,
          seconds,
          millisecondsOfSecond * 1000000);
    } else if (EvaluationHelper.equals(
        dataProtocolId, (String) "IEC61131_DATE_AND_LTIME")) { // DATE_AND_LTIME
      BigInteger nanosecondsSinceEpoch =
          readSimpleField("nanosecondsSinceEpoch", readUnsignedBigInteger(readBuffer, 64));
      return PlcDATE_AND_LTIME.ofNanosecondsSinceEpoch(nanosecondsSinceEpoch);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_DTL")) { // DATE_AND_LTIME
      int year = readSimpleField("year", readUnsignedInt(readBuffer, 16));

      short month = readSimpleField("month", readUnsignedShort(readBuffer, 8));

      short day = readSimpleField("day", readUnsignedShort(readBuffer, 8));

      short dayOfWeek = readSimpleField("dayOfWeek", readUnsignedShort(readBuffer, 8));

      short hour = readSimpleField("hour", readUnsignedShort(readBuffer, 8));

      short minutes = readSimpleField("minutes", readUnsignedShort(readBuffer, 8));

      short seconds = readSimpleField("seconds", readUnsignedShort(readBuffer, 8));

      long nannosecondsOfSecond =
          readSimpleField("nannosecondsOfSecond", readUnsignedLong(readBuffer, 32));
      return PlcDATE_AND_LTIME.ofSegments(
          year,
          (month == 0) ? 1 : month,
          (day == 0) ? 1 : day,
          hour,
          minutes,
          seconds,
          nannosecondsOfSecond);
    }
    return null;
  }

  public static int getLengthInBytes(
      PlcValue _value, String dataProtocolId, ControllerType controllerType, Integer stringLength) {
    return (int)
        Math.ceil(
            (float) getLengthInBits(_value, dataProtocolId, controllerType, stringLength) / 8.0);
  }

  public static int getLengthInBits(
      PlcValue _value, String dataProtocolId, ControllerType controllerType, Integer stringLength) {
    int lengthInBits = 0;
    if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_BOOL")) { // BOOL
      // Reserved Field (reserved)
      lengthInBits += 7;

      // Simple field (value)
      lengthInBits += 1;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_BYTE")) { // BYTE
      // Simple field (value)
      lengthInBits += 8;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_WORD")) { // WORD
      // Simple field (value)
      lengthInBits += 16;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_DWORD")) { // DWORD
      // Simple field (value)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_LWORD")) { // LWORD
      // Simple field (value)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_SINT")) { // SINT
      // Simple field (value)
      lengthInBits += 8;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_USINT")) { // USINT
      // Simple field (value)
      lengthInBits += 8;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_INT")) { // INT
      // Simple field (value)
      lengthInBits += 16;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_UINT")) { // UINT
      // Simple field (value)
      lengthInBits += 16;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_DINT")) { // DINT
      // Simple field (value)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_UDINT")) { // UDINT
      // Simple field (value)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_LINT")) { // LINT
      // Simple field (value)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_ULINT")) { // ULINT
      // Simple field (value)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_REAL")) { // REAL
      // Simple field (value)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_LREAL")) { // LREAL
      // Simple field (value)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_CHAR")) { // CHAR
      // Simple field (value)
      lengthInBits += 8;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_WCHAR")) { // CHAR
      // Simple field (value)
      lengthInBits += 16;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_STRING")) { // STRING
      // Manual Field (value)
      lengthInBits += (((stringLength) * (8))) + (16);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_WSTRING")) { // STRING
      // Manual Field (value)
      lengthInBits += (((stringLength) * (16))) + (32);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_TIME")) { // TIME
      // Simple field (milliseconds)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "S7_S5TIME")) { // TIME
      // Manual Field (milliseconds)
      lengthInBits += 2;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_LTIME")) { // LTIME
      // Simple field (nanoseconds)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_DATE")) { // DATE
      // Manual Field (daysSinceEpoch)
      lengthInBits += 16;
    } else if (EvaluationHelper.equals(
        dataProtocolId, (String) "IEC61131_TIME_OF_DAY")) { // TIME_OF_DAY
      // Simple field (millisecondsSinceMidnight)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(
        dataProtocolId, (String) "IEC61131_LTIME_OF_DAY")) { // LTIME_OF_DAY
      // Simple field (nanosecondsSinceMidnight)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(
        dataProtocolId, (String) "IEC61131_DATE_AND_TIME")) { // DATE_AND_TIME
      // Manual Field (year)
      lengthInBits += 8;

      // Simple field (month)
      lengthInBits += 8;

      // Simple field (day)
      lengthInBits += 8;

      // Simple field (hour)
      lengthInBits += 8;

      // Simple field (minutes)
      lengthInBits += 8;

      // Simple field (seconds)
      lengthInBits += 8;

      // Simple field (millisecondsOfSecond)
      lengthInBits += 12;

      // Simple field (dayOfWeek)
      lengthInBits += 4;
    } else if (EvaluationHelper.equals(
        dataProtocolId, (String) "IEC61131_DATE_AND_LTIME")) { // DATE_AND_LTIME
      // Simple field (nanosecondsSinceEpoch)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_DTL")) { // DATE_AND_LTIME
      // Simple field (year)
      lengthInBits += 16;

      // Simple field (month)
      lengthInBits += 8;

      // Simple field (day)
      lengthInBits += 8;

      // Simple field (dayOfWeek)
      lengthInBits += 8;

      // Simple field (hour)
      lengthInBits += 8;

      // Simple field (minutes)
      lengthInBits += 8;

      // Simple field (seconds)
      lengthInBits += 8;

      // Simple field (nannosecondsOfSecond)
      lengthInBits += 32;
    }

    return lengthInBits;
  }

  public static void staticSerialize(
      WriteBuffer writeBuffer,
      PlcValue _value,
      String dataProtocolId,
      ControllerType controllerType,
      Integer stringLength)
      throws SerializationException {
    staticSerialize(
        writeBuffer, _value, dataProtocolId, controllerType, stringLength, ByteOrder.BIG_ENDIAN);
  }

  public static void staticSerialize(
      WriteBuffer writeBuffer,
      PlcValue _value,
      String dataProtocolId,
      ControllerType controllerType,
      Integer stringLength,
      ByteOrder byteOrder)
      throws SerializationException {
    if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_BOOL")) { // BOOL
      // Reserved Field (reserved)
      writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

      // Simple Field (value)
      writeSimpleField("value", (boolean) _value.getBoolean(), writeBoolean(writeBuffer));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_BYTE")) { // BYTE
      // Simple Field (value)
      writeSimpleField("value", (short) _value.getShort(), writeUnsignedShort(writeBuffer, 8));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_WORD")) { // WORD
      // Simple Field (value)
      writeSimpleField("value", (int) _value.getInteger(), writeUnsignedInt(writeBuffer, 16));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_DWORD")) { // DWORD
      // Simple Field (value)
      writeSimpleField("value", (long) _value.getLong(), writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_LWORD")) { // LWORD
      // Simple Field (value)
      writeSimpleField(
          "value", (BigInteger) _value.getBigInteger(), writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_SINT")) { // SINT
      // Simple Field (value)
      writeSimpleField("value", (byte) _value.getByte(), writeSignedByte(writeBuffer, 8));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_USINT")) { // USINT
      // Simple Field (value)
      writeSimpleField("value", (short) _value.getShort(), writeUnsignedShort(writeBuffer, 8));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_INT")) { // INT
      // Simple Field (value)
      writeSimpleField("value", (short) _value.getShort(), writeSignedShort(writeBuffer, 16));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_UINT")) { // UINT
      // Simple Field (value)
      writeSimpleField("value", (int) _value.getInteger(), writeUnsignedInt(writeBuffer, 16));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_DINT")) { // DINT
      // Simple Field (value)
      writeSimpleField("value", (int) _value.getInteger(), writeSignedInt(writeBuffer, 32));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_UDINT")) { // UDINT
      // Simple Field (value)
      writeSimpleField("value", (long) _value.getLong(), writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_LINT")) { // LINT
      // Simple Field (value)
      writeSimpleField("value", (long) _value.getLong(), writeSignedLong(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_ULINT")) { // ULINT
      // Simple Field (value)
      writeSimpleField(
          "value", (BigInteger) _value.getBigInteger(), writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_REAL")) { // REAL
      // Simple Field (value)
      writeSimpleField("value", (float) _value.getFloat(), writeFloat(writeBuffer, 32));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_LREAL")) { // LREAL
      // Simple Field (value)
      writeSimpleField("value", (double) _value.getDouble(), writeDouble(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_CHAR")) { // CHAR
      // Simple Field (value)
      writeSimpleField(
          "value",
          (String) _value.getString(),
          writeString(writeBuffer, 8),
          WithOption.WithEncoding("UTF-8"));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_WCHAR")) { // CHAR
      // Simple Field (value)
      writeSimpleField(
          "value",
          (String) _value.getString(),
          writeString(writeBuffer, 16),
          WithOption.WithEncoding("UTF-16"));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_STRING")) { // STRING
      // Manual Field (value)
      writeManualField(
          "value",
          () ->
              org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.serializeS7String(
                  writeBuffer, _value, stringLength, "UTF-8"),
          writeBuffer,
          WithOption.WithEncoding("UTF-8"));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_WSTRING")) { // STRING
      // Manual Field (value)
      writeManualField(
          "value",
          () ->
              org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.serializeS7String(
                  writeBuffer, _value, stringLength, "UTF-16"),
          writeBuffer,
          WithOption.WithEncoding("UTF-16"));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_TIME")) { // TIME
      // Simple Field (milliseconds)
      writeSimpleField(
          "milliseconds",
          (long) _value.getDuration().toMillis(),
          writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "S7_S5TIME")) { // TIME
      // Manual Field (milliseconds)
      writeManualField(
          "milliseconds",
          () ->
              org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.serializeS5Time(
                  writeBuffer, _value),
          writeBuffer);
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_LTIME")) { // LTIME
      // Simple Field (nanoseconds)
      writeSimpleField(
          "nanoseconds",
          (BigInteger) BigInteger.valueOf(_value.getDuration().toNanos()),
          writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_DATE")) { // DATE
      // Manual Field (daysSinceEpoch)
      writeManualField(
          "daysSinceEpoch",
          () ->
              org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.serializeTiaDate(
                  writeBuffer, _value),
          writeBuffer);
    } else if (EvaluationHelper.equals(
        dataProtocolId, (String) "IEC61131_TIME_OF_DAY")) { // TIME_OF_DAY
      // Simple Field (millisecondsSinceMidnight)
      writeSimpleField(
          "millisecondsSinceMidnight",
          (long) _value.getTime().getLong(ChronoField.MILLI_OF_DAY),
          writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(
        dataProtocolId, (String) "IEC61131_LTIME_OF_DAY")) { // LTIME_OF_DAY
      // Simple Field (nanosecondsSinceMidnight)
      writeSimpleField(
          "nanosecondsSinceMidnight",
          (BigInteger) BigInteger.valueOf(_value.getTime().getLong(ChronoField.NANO_OF_DAY)),
          writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(
        dataProtocolId, (String) "IEC61131_DATE_AND_TIME")) { // DATE_AND_TIME
      // Manual Field (year)
      writeManualField(
          "year",
          () ->
              org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.serializeSiemensYear(
                  writeBuffer, _value),
          writeBuffer,
          WithOption.WithEncoding("BCD"));

      // Simple Field (month)
      writeSimpleField(
          "month",
          (short) _value.getDate().getMonthValue(),
          writeUnsignedShort(writeBuffer, 8),
          WithOption.WithEncoding("BCD"));

      // Simple Field (day)
      writeSimpleField(
          "day",
          (short) _value.getDate().getDayOfMonth(),
          writeUnsignedShort(writeBuffer, 8),
          WithOption.WithEncoding("BCD"));

      // Simple Field (hour)
      writeSimpleField(
          "hour",
          (short) _value.getTime().getHour(),
          writeUnsignedShort(writeBuffer, 8),
          WithOption.WithEncoding("BCD"));

      // Simple Field (minutes)
      writeSimpleField(
          "minutes",
          (short) _value.getTime().getMinute(),
          writeUnsignedShort(writeBuffer, 8),
          WithOption.WithEncoding("BCD"));

      // Simple Field (seconds)
      writeSimpleField(
          "seconds",
          (short) _value.getTime().getSecond(),
          writeUnsignedShort(writeBuffer, 8),
          WithOption.WithEncoding("BCD"));

      // Simple Field (millisecondsOfSecond)
      writeSimpleField(
          "millisecondsOfSecond",
          (short) _value.getTime().get(ChronoField.MILLI_OF_SECOND),
          writeUnsignedShort(writeBuffer, 12),
          WithOption.WithEncoding("BCD"));

      // Simple Field (dayOfWeek)
      writeSimpleField(
          "dayOfWeek",
          (byte) _value.getDate().getDayOfWeek().getValue(),
          writeUnsignedByte(writeBuffer, 4),
          WithOption.WithEncoding("BCD"));
    } else if (EvaluationHelper.equals(
        dataProtocolId, (String) "IEC61131_DATE_AND_LTIME")) { // DATE_AND_LTIME
      // Simple Field (nanosecondsSinceEpoch)
      writeSimpleField(
          "nanosecondsSinceEpoch",
          (BigInteger)
              BigInteger.valueOf(_value.getDateTime().toEpochSecond(ZoneOffset.UTC))
                  .multiply(BigInteger.valueOf(1000000000))
                  .add(BigInteger.valueOf(_value.getDateTime().getNano())),
          writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataProtocolId, (String) "IEC61131_DTL")) { // DATE_AND_LTIME
      // Simple Field (year)
      writeSimpleField("year", (int) _value.getDate().getYear(), writeUnsignedInt(writeBuffer, 16));

      // Simple Field (month)
      writeSimpleField(
          "month", (short) _value.getDate().getMonthValue(), writeUnsignedShort(writeBuffer, 8));

      // Simple Field (day)
      writeSimpleField(
          "day", (short) _value.getDate().getDayOfMonth(), writeUnsignedShort(writeBuffer, 8));

      // Simple Field (dayOfWeek)
      writeSimpleField(
          "dayOfWeek",
          (short) _value.getDate().getDayOfWeek().getValue(),
          writeUnsignedShort(writeBuffer, 8));

      // Simple Field (hour)
      writeSimpleField(
          "hour", (short) _value.getTime().getHour(), writeUnsignedShort(writeBuffer, 8));

      // Simple Field (minutes)
      writeSimpleField(
          "minutes", (short) _value.getTime().getMinute(), writeUnsignedShort(writeBuffer, 8));

      // Simple Field (seconds)
      writeSimpleField(
          "seconds", (short) _value.getTime().getSecond(), writeUnsignedShort(writeBuffer, 8));

      // Simple Field (nannosecondsOfSecond)
      writeSimpleField(
          "nannosecondsOfSecond",
          (long) _value.getTime().getLong(ChronoField.NANO_OF_SECOND),
          writeUnsignedLong(writeBuffer, 32));
    }
  }
}
