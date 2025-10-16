# Inverter Protocol Command and Data Format

This document provides a detailed breakdown of all commands available in the Axpert MAX communication protocol. It is intended as a technical blueprint for implementation.

## 1. Inquiry Commands

### 2.1. QPI: Device Protocol ID Inquiry
- **Description:** To request the device Protocol ID.
- **Request:** `QPI<CRC><cr>`
- **Response:** `(PI<NN><CRC><cr>`
- **Data Fields:**
    - `NN`: Protocol ID

### 2.2. QID: The device serial number inquiry
- **Description:** To request the device's 14-character serial number.
- **Request:** `QID<CRC><cr>`
- **Response:** `(XXXXXXXXXXXXXX<CRC><cr>`
- **Data Fields:**
    - `XXXXXXXXXXXXXX`: 14-character serial number

### 2.3. QSID: The device serial number inquiry (long)
- **Description:** To request a serial number longer than 14 characters.
- **Request:** `QSID<CRC><cr>`
- **Response:** `(NNXXXXXXXXXXXXXXXXXXXX<CRC><cr>`
- **Data Fields:**
    - `NN`: Serial number valid length
    - `XXXXXXXXXXXXXXXXXXXX`: 20-character serial number

### 2.4. QVFW: Main CPU Firmware version inquiry
- **Description:** To request the main CPU's firmware version.
- **Request:** `QVFW<CRC><cr>`
- **Response:** `(VERFW:<NNNNN.NN><CRC><cr>`
- **Data Fields:**
    - `NNNNN.NN`: Firmware version

### 2.5. QVFW3: Another CPU Firmware version inquiry
- **Description:** To request the firmware version of a secondary CPU (e.g., remote panel).
- **Request:** `QVFW3<CRC><cr>`
- **Response:** `(VERFW:<NNNNN.NN><CRC><cr>`
- **Data Fields:**
    - `NNNNN.NN`: Firmware version

### 2.7. QPIRI: Device Rating Information inquiry
- **Description:** To inquire about the device's rating information.
- **Request:** `QPIRI<CRC><cr>`
- **Response:** `(BBB.B CC.C DDD.D EE.E FF.F HHHH IIII JJ.J KK.K JJ.J KK.K LL.L O PP QQ0 OPQRSST U VV.V W X YYY Z CCC<CRC><cr>`
- **Data Fields:**
    - `BBB.B`: Grid Rating Voltage
    - `CC.C`: Grid Rating Current
    - `DDD.D`: AC Output Rating Voltage
    - `EE.E`: AC Output Rating Frequency
    - `FF.F`: AC Output Rating Current
    - `HHHH`: AC Output Rating Apparent Power
    - `IIII`: AC Output Rating Active Power
    - `JJ.J`: Battery Rating Voltage
    - `KK.K`: Battery Re-charge Voltage
    - `JJ.J`: Battery Under Voltage
    - `KK.K`: Battery Bulk Voltage
    - `LL.L`: Battery Float Voltage
    - `O`: Battery Type
    - `PP`: Max AC Charging Current
    - `QQ0`: Max Charging Current
    - `O`: Input Voltage Range
    - `P`: Output Source Priority
    - `Q`: Charger Source Priority
    - `R`: Parallel Max Number
    - `SS`: Machine Type
    - `T`: Topology
    - `U`: Output Mode
    - `VV.V`: Battery Re-discharge Voltage
    - `W`: PV OK Condition for Parallel
    - `X`: PV Power Balance
    - `YYY`: Max. Charging Time at C.V Stage
    - `Z`: Operation Logic
    - `CCC`: Max Discharging Current

### 2.9. QPIGS: Device general status parameters inquiry
- **Description:** To request the primary general status of the device.
- **Request:** `QPIGS<CRC><cr>`
- **Response:** `(BBB.BCC.CDDD.D EE.E FFFF GGGG HHH III JJ.JJ KKK OOO TTTT EE.E UUU.UWW.WW_PPPPP b7b6b5b4b3b2b1b0 QQ VV MMMMM b10b9b8 Y ZZ AAAA BB.B<CRC><cr>`
- **Data Fields:**
    - `BBB.B`: Grid Voltage
    - `CC.C`: Grid Frequency
    - `DDD.D`: AC Output Voltage
    - `EE.E`: AC Output Frequency
    - `FFFF`: AC Output Apparent Power
    - `GGGG`: AC Output Active Power
    - `HHH`: Output Load Percent
    - `III`: BUS Voltage
    - `JJ.JJ`: Battery Voltage
    - `KKK`: Battery Charging Current
    - `OOO`: Battery Capacity
    - `TTTT`: Inverter Heat Sink Temperature
    - `EE.E`: PV1 Input Current
    - `UUU.U`: PV1 Input Voltage
    - `WW.WW`: Battery Voltage from SCC
    - `PPPPP`: Battery Discharge Current
    - `b7b6b5b4b3b2b1b0`: Device Status 1 (bit-encoded)
    - `QQ`: Battery V offset for fans on
    - `VV`: EEPROM Version
    - `MMMMM`: PV1 Charging Power
    - `b10b9b8`: Device Status 2 (bit-encoded)
    - `Y`: Solar feed to grid status
    - `ZZ`: Set country regulation
    - `AAAA`: Solar feed to grid power
    - `BB.B`: Grid input current

### 2.10. QPIGS2: Device general status parameters inquiry (for MAXII)
- **Description:** To request the secondary general status (PV2) of the device.
- **Request:** `QPIGS2<CRC><cr>`
- **Response:** `(BB.B CCC.C DDDDD<CRC><cr>`
- **Data Fields:**
    - `BB.B`: PV2 Input Current
    - `CCC.C`: PV2 Input Voltage
    - `DDDDD`: PV2 Charging Power

### 2.12. QMOD: Device Mode inquiry
- **Description:** To request the current working mode of the device.
- **Request:** `QMOD<CRC><cr>`
- **Response:** `(M<CRC><cr>`
- **Data Fields:**
    - `M`: Device Mode

### 2.13. QPIWS: Device Warning Status inquiry
- **Description:** To request the device's warning status as a bitmask.
- **Request:** `QPIWS<CRC><cr>`
- **Response:** `(a0a1.....a30a31<CRC><cr>`
- **Data Fields:**
    - `a0...a31`: 32-bit warning status. Each bit corresponds to a specific warning.

### 2.14. QDI: The default setting value information
- **Description:** To request the device's default factory settings.
- **Request:** `QDI<CRC><cr>`
- **Response:** `(BBB.B CC.C 00DD EE.E FF.F GG.G HH.H II J K L M N O P Q R S T U V W YY.Y X Z aaa bbb<CRC><cr>`
- **Data Fields:**
    - `BBB.B`: AC Output Voltage
    - `CC.C`: AC Output Frequency
    - `DD`: Max AC Charging Current
    - `EE.E`: Battery Under Voltage
    - `FF.F`: Battery Float Charging Voltage
    - `GG.G`: Battery Bulk Charging (C.V.) Voltage
    - `HH.H`: Battery Re-charge Voltage
    - `II`: Max Charging Current
    - `J`: Input Voltage Range
    - `K`: Output Source Priority
    - `L`: Charger Source Priority
    - `M`: Battery Type
    - `N..W`: (undocumented)
    - `YY.Y`: Battery Re-discharge voltage
    - `X`: (undocumented)
    - `Z`: Machine Type
    - `aaa`: (undocumented)
    - `bbb`: (undocumented)

### 2.23. QET: Query total PV generated energy
- **Description:** To request the total PV generated energy since reset.
- **Request:** `QET<CRC><cr>`
- **Response:** `(NNNNNNNN<CRC><cr>`
- **Data Fields:**
    - `NNNNNNNN`: Total generated energy in KWH

### 2.26. QEDyyyymmdd: Query PV generated energy of day
- **Description:** To request the PV generated energy for a specific day.
- **Request:** `QEDyyyymmdd<CRC><cr>`
- **Response:** `(NNNNNNNN<CRC><cr>`
- **Data Fields:**
    - `NNNNNNNN`: Generated energy for the day in KWH

---
## 2. Setting Commands

Setting commands typically respond with `(ACK<CRC><cr>` on success or `(NAK<CRC><cr>` on failure.

### 3.6. PF: Setting control parameter to default value
- **Description:** Resets device settings to factory defaults.
- **Request:** `PF<CRC><cr>`

### 3.12. POP: Setting device output source priority
- **Description:** Sets the priority for the output source (e.g., Solar > Utility > Battery).
- **Request:** `POP<NN><CRC><cr>`
- **Parameters:**
    - `NN`: Priority (e.g., `00` for Utility-Solar-Battery)

### 3.17. PBT: Setting battery type
- **Description:** Sets the connected battery type.
- **Request:** `PBT<NN><CRC><cr>`
- **Parameters:**
    - `NN`: Battery Type (e.g., `00` for AGM, `01` for Flooded, `02` for User)

### 3.20. PSDV: Setting battery cut-off voltage
- **Description:** Sets the battery cut-off voltage (battery under voltage).
- **Request:** `PSDV<nn.n><CRC><cr>`
- **Parameters:**
    - `nn.n`: Voltage

### 3.35. DAT: Date and time
- **Description:** Sets the device's internal date and time.
- **Request:** `DAT<YYMMDDHHMMSS><CRC><cr>`
- **Parameters:**
    - `YYMMDDHHMMSS`: Year, Month, Day, Hour, Minute, Second

*(Note: This is a partial list highlighting key commands. A complete implementation would require creating entries for all commands listed in Protocol.md)*
