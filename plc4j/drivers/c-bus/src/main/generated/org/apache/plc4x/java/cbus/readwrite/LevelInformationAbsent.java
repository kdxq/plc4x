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

public class LevelInformationAbsent extends LevelInformation implements Message {

  // Accessors for discriminator values.
  // Reserved Fields
  private Integer reservedField0;

  public LevelInformationAbsent(int raw) {
    super(raw);
  }

  @Override
  protected void serializeLevelInformationChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("LevelInformationAbsent");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (int) 0x0000,
        writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("LevelInformationAbsent");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    LevelInformationAbsent _value = this;

    // Reserved Field (reserved)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static LevelInformationAbsentBuilder staticParseBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("LevelInformationAbsent");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    Integer reservedField0 =
        readReservedField("reserved", readUnsignedInt(readBuffer, 16), (int) 0x0000);

    readBuffer.closeContext("LevelInformationAbsent");
    // Create the instance
    return new LevelInformationAbsentBuilder(reservedField0);
  }

  public static class LevelInformationAbsentBuilder
      implements LevelInformation.LevelInformationBuilder {
    private final Integer reservedField0;

    public LevelInformationAbsentBuilder(Integer reservedField0) {
      this.reservedField0 = reservedField0;
    }

    public LevelInformationAbsent build(int raw) {
      LevelInformationAbsent levelInformationAbsent = new LevelInformationAbsent(raw);
      levelInformationAbsent.reservedField0 = reservedField0;
      return levelInformationAbsent;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof LevelInformationAbsent)) {
      return false;
    }
    LevelInformationAbsent that = (LevelInformationAbsent) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
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