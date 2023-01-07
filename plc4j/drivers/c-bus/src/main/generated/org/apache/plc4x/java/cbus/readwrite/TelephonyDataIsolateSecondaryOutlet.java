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
package org.apache.plc4x.java.cbus.readwrite;

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

public class TelephonyDataIsolateSecondaryOutlet extends TelephonyData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte isolateStatus;

  public TelephonyDataIsolateSecondaryOutlet(
      TelephonyCommandTypeContainer commandTypeContainer, byte argument, byte isolateStatus) {
    super(commandTypeContainer, argument);
    this.isolateStatus = isolateStatus;
  }

  public byte getIsolateStatus() {
    return isolateStatus;
  }

  public boolean getIsBehaveNormal() {
    return (boolean) ((getIsolateStatus()) == (0x00));
  }

  public boolean getIsToBeIsolated() {
    return (boolean) ((getIsolateStatus()) == (0x01));
  }

  @Override
  protected void serializeTelephonyDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("TelephonyDataIsolateSecondaryOutlet");

    // Simple Field (isolateStatus)
    writeSimpleField("isolateStatus", isolateStatus, writeByte(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isBehaveNormal = getIsBehaveNormal();
    writeBuffer.writeVirtual("isBehaveNormal", isBehaveNormal);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isToBeIsolated = getIsToBeIsolated();
    writeBuffer.writeVirtual("isToBeIsolated", isToBeIsolated);

    writeBuffer.popContext("TelephonyDataIsolateSecondaryOutlet");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TelephonyDataIsolateSecondaryOutlet _value = this;

    // Simple field (isolateStatus)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static TelephonyDataIsolateSecondaryOutletBuilder staticParseBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("TelephonyDataIsolateSecondaryOutlet");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    byte isolateStatus = readSimpleField("isolateStatus", readByte(readBuffer, 8));
    boolean isBehaveNormal =
        readVirtualField("isBehaveNormal", boolean.class, (isolateStatus) == (0x00));
    boolean isToBeIsolated =
        readVirtualField("isToBeIsolated", boolean.class, (isolateStatus) == (0x01));

    readBuffer.closeContext("TelephonyDataIsolateSecondaryOutlet");
    // Create the instance
    return new TelephonyDataIsolateSecondaryOutletBuilder(isolateStatus);
  }

  public static class TelephonyDataIsolateSecondaryOutletBuilder
      implements TelephonyData.TelephonyDataBuilder {
    private final byte isolateStatus;

    public TelephonyDataIsolateSecondaryOutletBuilder(byte isolateStatus) {

      this.isolateStatus = isolateStatus;
    }

    public TelephonyDataIsolateSecondaryOutlet build(
        TelephonyCommandTypeContainer commandTypeContainer, byte argument) {
      TelephonyDataIsolateSecondaryOutlet telephonyDataIsolateSecondaryOutlet =
          new TelephonyDataIsolateSecondaryOutlet(commandTypeContainer, argument, isolateStatus);
      return telephonyDataIsolateSecondaryOutlet;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TelephonyDataIsolateSecondaryOutlet)) {
      return false;
    }
    TelephonyDataIsolateSecondaryOutlet that = (TelephonyDataIsolateSecondaryOutlet) o;
    return (getIsolateStatus() == that.getIsolateStatus()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getIsolateStatus());
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