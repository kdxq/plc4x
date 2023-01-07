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
package org.apache.plc4x.java.bacnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class BACnetEventTimestamps implements Message {

  // Properties.
  protected final BACnetTimeStamp toOffnormal;
  protected final BACnetTimeStamp toFault;
  protected final BACnetTimeStamp toNormal;

  public BACnetEventTimestamps(
      BACnetTimeStamp toOffnormal, BACnetTimeStamp toFault, BACnetTimeStamp toNormal) {
    super();
    this.toOffnormal = toOffnormal;
    this.toFault = toFault;
    this.toNormal = toNormal;
  }

  public BACnetTimeStamp getToOffnormal() {
    return toOffnormal;
  }

  public BACnetTimeStamp getToFault() {
    return toFault;
  }

  public BACnetTimeStamp getToNormal() {
    return toNormal;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetEventTimestamps");

    // Simple Field (toOffnormal)
    writeSimpleField("toOffnormal", toOffnormal, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (toFault)
    writeSimpleField("toFault", toFault, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (toNormal)
    writeSimpleField("toNormal", toNormal, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetEventTimestamps");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetEventTimestamps _value = this;

    // Simple field (toOffnormal)
    lengthInBits += toOffnormal.getLengthInBits();

    // Simple field (toFault)
    lengthInBits += toFault.getLengthInBits();

    // Simple field (toNormal)
    lengthInBits += toNormal.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetEventTimestamps staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static BACnetEventTimestamps staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetEventTimestamps");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetTimeStamp toOffnormal =
        readSimpleField(
            "toOffnormal",
            new DataReaderComplexDefault<>(
                () -> BACnetTimeStamp.staticParse(readBuffer), readBuffer));

    BACnetTimeStamp toFault =
        readSimpleField(
            "toFault",
            new DataReaderComplexDefault<>(
                () -> BACnetTimeStamp.staticParse(readBuffer), readBuffer));

    BACnetTimeStamp toNormal =
        readSimpleField(
            "toNormal",
            new DataReaderComplexDefault<>(
                () -> BACnetTimeStamp.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("BACnetEventTimestamps");
    // Create the instance
    BACnetEventTimestamps _bACnetEventTimestamps;
    _bACnetEventTimestamps = new BACnetEventTimestamps(toOffnormal, toFault, toNormal);
    return _bACnetEventTimestamps;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetEventTimestamps)) {
      return false;
    }
    BACnetEventTimestamps that = (BACnetEventTimestamps) o;
    return (getToOffnormal() == that.getToOffnormal())
        && (getToFault() == that.getToFault())
        && (getToNormal() == that.getToNormal())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getToOffnormal(), getToFault(), getToNormal());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}