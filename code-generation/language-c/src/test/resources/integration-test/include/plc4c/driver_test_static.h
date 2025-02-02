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

#ifndef PLC4C_TEST_READ_WRITE_DRIVER_H_
#define PLC4C_TEST_READ_WRITE_DRIVER_H_

// Dummy
#include <stdint.h>
#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>

int8_t plc4c_test_read_write_crc_int8();
uint8_t plc4c_test_read_write_crc_uint8();
uint8_t plc4c_test_read_write_read_manual_field(plc4c_spi_read_buffer* readBuffer, uint8_t value);
plc4c_return_code plc4c_test_read_write_write_manual_field(plc4c_spi_write_buffer* writeBuffer, uint8_t value);

#ifdef __cplusplus
}
#endif

#endif  // PLC4C_TEST_READ_WRITE_DRIVER_H_
