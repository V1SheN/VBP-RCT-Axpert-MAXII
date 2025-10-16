# Implemented and Missing Protocol Commands

The current implementation focuses on querying real-time status and ratings with these four commands:
*   `QPIGS`
*   `QPIGS2`
*   `QPIRI`
*   `QPIWS`

### Missing Inquiry Commands
The following inquiry commands are defined in the protocol but are **not** implemented in the Go application:

*   **Device Identification:**
    *   `QPI`: Device Protocol ID Inquiry
    *   `QID`: Device serial number inquiry
    *   `QSID`: Device serial number inquiry (for long serials)
    *   `QMN`: Query model name
    *   `QGMN`: Query general model name
*   **Firmware Version:**
    *   `QVFW`: Main CPU Firmware version inquiry
    *   `QVFW3`: Another CPU Firmware version inquiry
    *   `VERFW`: Bluetooth version inquiry
*   **Status & Settings Inquiry:**
    *   `QFLAG`: Device flag status inquiry
    *   `QPGSn`: Parallel Information inquiry
    *   `QMOD`: Device Mode inquiry
    *   `QDI`: Default setting value information
    *   `QMCHGCR`: Query selectable max charging currents
    *   `QMUCHGCR`: Query selectable max utility charging currents
*   **Energy & Time Queries:**
    *   `QT`: Time inquiry
    *   `QET`: Total PV generated energy
    *   `QEYyyyy`: Yearly PV generated energy
    *   `QEMyyyymm`: Monthly PV generated energy
    *   `QEDyyyymmdd`: Daily PV generated energy
    *   `QLT`: Total output load energy
    *   `QLYyyyy`: Yearly output load energy
    *   `QLMyyyymm`: Monthly output load energy
    *   `QLDyyyymmdd`: Daily output load energy
*   **Battery Equalization:**
    *   `QBEQI`: Battery equalization status inquiry
*   **Other:**
    *   `QBMS` / `PBMS`: BMS message inquiry
    *   `QLED`: LED status inquiry
    *   `QWFS`: Wi-Fi module status query

### Missing Setting Commands
**None of the setting commands** appear to be implemented. The application is currently read-only; it queries data but does not send any commands to change the inverter's configuration. This includes commands for:
*   Changing device settings (`PF`, `POP`, `PCP`, `PBT`, etc.)
*   Setting voltages or currents (`PSDV`, `PCVV`, `PBFT`, `MNCHGC`, etc.)
*   Managing battery equalization (`PBEQE`, `PBEQA`, etc.)
*   Resetting data (`RTEY`, `RTDL`)
*   Setting the date and time (`DAT`)