# Axpert MAX Communication Protocol

This document describes the communication protocol for the Axpert MAXII, VMIV, and MKSIV inverters.

## 1. Communication Format

### 1.1. RS232

*   **Baud rate:** 2400
*   **Start bit:** 1
*   **Data bit:** 8
*   **Parity bit:** N
*   **Stop bit:** 1

## 2. Inquiry Commands

### 2.1. QPI: Device Protocol ID Inquiry

*   **Computer:** `QPI<CRC><cr>`
*   **Device:** `(PI<NN><CRC><cr>`
*   **Description:** To request the device Protocol ID.
*   **Protocol ID distribution:** 30 for Axpert KS series

### 2.2. QID: The device serial number inquiry

*   **Computer:** `QID<CRC><cr>`
*   **Device:** `(XXXXXXXXXXXXXX<CRC><cr>`

### 2.3. QSID: The device serial number inquiry (the length is more than 14)

*   **Computer:** `QSID<CRC><cr>`
*   **Device:** `(NNXXXXXXXXXXXXXXXXXXXX<CRC><cr>`
*   **Description:** NN: Serial number valid length, X: Serial number, invalid part is filled as ‘0', total X is 20.

### 2.4. QVFW: Main CPU Firmware version inquiry

*   **Computer:** `QVFW<CRC><cr>`
*   **Device:** `(VERFW:<NNNNN.NN><CRC><cr>`

### 2.5. QVFW3: Another CPU (remote panel) Firmware version inquiry

*   **Computer:** `QVFW3<CRC><cr>`
*   **Device:** `(VERFW:<NNNNN.NN><CRC><cr>`

### 2.6. VERFW: Bluetooth version inquiry

*   **Computer:** `VERFW:<CRC><cr>`
*   **Device:** `(VERFW:<NNNNN.NN><CRC><cr>`

### 2.7. QPIRI: Device Rating Information inquiry

*   **Computer:** `QPIRI<CRC><cr>`
*   **Device:** `(BBB.B CC.C DDD.D EE.E FF.F HHHH IIII JJ.J KK.K JJ.J KK.K LL.L O PP QQ0 OPQRSST U VV.V W X YYY Z CCC<CRC><cr>`

### 2.8. QFLAG: Device flag status inquiry

*   **Computer:** `QFLAG<CRC><cr>`
*   **Device:** `(ExxxDxxx<CRC><cr>`
*   **Description:** ExxxDxxx is the flag status. E means enable, D means disable

### 2.9. QPIGS: Device general status parameters inquiry

*   **Computer:** `QPIGS<CRC><cr>`
*   **Device:** `(BBB.BCC.CDDD.D EE.E FFFF GGGG HHH III JJ.JJ KKK OOO TTTT EE.E UUU.UWW.WW_PPPPP b7b6b5b4b3b2b1b0 QQ VV MMMMM b10b9b8 Y ZZ AAAA BB.B<CRC><cr>`

### 2.10. QPIGS2: Device general status parameters inquiry (for MAXII)

*   **Computer:** `QPIGS2<CRC><cr>`
*   **Device:** `(BB.B CCC.C DDDDD<CRC><cr>`

### 2.11. QPGSn: Parallel Information inquiry (for MAXII & MKSIV)

*   **Computer:** `QPGSn<CRC><cr>`
*   **Device:** `(A BBBBBBBBBBBBBB C DD EEE.E FF.FF GGG.G HH.HH_IIII JJJJ KKK LL.L MMM NNN OOO.O PPP QQQQQ RRRRR SSS b7b6b5b4b3b2b1b0 T U VVV WWW ZZ XX YYY 000.O XX<CRC><cr>`

### 2.12. QMOD: Device Mode inquiry

*   **Computer:** `QMOD<CRC><cr>`
*   **Device:** `(M<CRC><cr>`

### 2.13. QPIWS: Device Warning Status inquiry

*   **Computer:** `QPIWS<CRC><cr>`
*   **Device:** `(a0a1.....a30a31<CRC><cr>`

### 2.14. QDI: The default setting value information

*   **Computer:** `QDI<CRC><cr>`
*   **Device:** `(BBB.B CC.C 00DD EE.E FF.F GG.G HH.H II J K L M N O P Q R S T U V W YY.Y X Z aaa bbb<CRC><cr>`

### 2.15. QMCHGCR: Enquiry selectable value about max charging current

*   **Computer:** `QMCHGCR<CRC><cr>`
*   **Device:** `(AAA BBB CCC DDD......<CRC><cr>`

### 2.16. QMUCHGCR: Enquiry selectable value about max utility charging current

*   **Computer:** `QMUCHGCR<CRC><cr>`
*   **Device:** `(AAA BBB CCC DDD......<CRC><cr>`

### 2.17. QOPPT: The device output source priority time order inquiry

*   **Computer:** `QOPPT<CRC><cr>`
*   **Device:** `(M M M M M M M M M M M M M M M M M M M M M M M N O O<CRC><cr>`

### 2.18. QCHPT: The device charger source priority time order inquiry

*   **Computer:** `QCHPT<CRC><cr>`
*   **Device:** `(M M M M M M M M M M M M M M M M M M M M M M M N O O<CRC><cr>`

### 2.19. QT: Time inquiry

*   **Computer:** `QT<cr>`
*   **Device:** `(YYYYMMDDHHMMSS<cr>`

### 2.20. QBEQI: Battery equalization status parameters inquiry

*   **Computer:** `QBEQI<CRC><cr>`
*   **Device:** `(B CCC DDD EEE FFF GG.GG HHH III J KKKK<CRC><cr>`

### 2.21. QMN: Query model name

*   **Computer:** `QMN<CRC><cr>`
*   **Device:** `(MMMMM-NNNN<CRC><cr>`

### 2.22. QGMN: Query general model name

*   **Computer:** `QGMN<CRC><cr>`
*   **Device:** `(NNN<CRC><cr>`

### 2.23. QET: Query total PV generated energy

*   **Computer:** `QET<CRC><cr>`
*   **Device:** `(NNNNNNNN<CRC><cr>`

### 2.24. QEYyyyy: Query PV generated energy of year

*   **Computer:** `QEYyyyy<CRC><cr>`
*   **Device:** `(NNNNNNNN<CRC><cr>`

### 2.25. QEMyyyymm: Query PV generated energy of month

*   **Computer:** `QEMyyyymm<CRC><cr>`
*   **Device:** `(NNNNNNNN<CRC><cr>`

### 2.26. QEDyyyymmdd: Query PV generated energy of day

*   **Computer:** `QEDyyyymmdd<CRC><cr>`
*   **Device:** `(NNNNNNNN<CRC><cr>`

### 2.27. QLT: Query total output load energy

*   **Computer:** `QLT<CRC><cr>`
*   **Device:** `(NNNNNNNN<CRC><cr>`

### 2.28. QLYyyyy: Query output load energy of year

*   **Computer:** `QLYyyyy<CRC><cr>`
*   **Device:** `(NNNNNNNN<CRC><cr>`

### 2.29. QLMyyyymm: Query output load energy of month

*   **Computer:** `QLMyyyymm<CRC><cr>`
*   **Device:** `(NNNNNNNN<CRC><cr>`

### 2.30. QLDyyyymmdd: Query output load energy of day

*   **Computer:** `QLDyyyymmdd<CRC><cr>`
*   **Device:** `(NNNNNNNN<CRC><cr>`

### 2.31. QBMS: BMS message

*   **Computer:** `QBMS<CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 2.32. PBMS: BMS message

*   **Remote box:** `PBMSa bbb c d e fff ggg hhh iiii jjjj<CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 2.33. QLED: LED status parameters inquiry

*   **Computer:** `QLED<cr>`
*   **UPS:** `(A B C D E aaa1bbb1ccc1 aaa2bbb2ccc2 aaa3bbb3ccc3<cr>`

### 2.34. QWFS: Wi-Fi module RS232 communication status query

*   **Computer:** `QWFS<CRC><cr>`
*   **Device:** `(N<CRC><cr>`

## 3. Setting parameters Command

### 3.1. ATE1: Start ATE test, remote panel stop polling

*   **Computer:** `ATE1<CRC><cr>`

### 3.2. ATE0: End ATE test, remote panel polling

*   **Computer:** `ATE0<CRC><cr>`

### 3.3. LOGO: Setting logo LED enable/disable

*   **Computer:** `LOGO<n><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.4. WEL: Setting Welcome show

*   **Computer:** `WEL<n><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.5. PE<X> / PD<X>: Setting some status enable/disable

*   **Computer:** `PE<X> / PD<X><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.6. PF: Setting control parameter to default value

*   **Computer:** `PF<CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.7. MNCHGC: Setting max charging current

*   **Computer:** `MNCHGC<mnnn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.8. MUCHGC: Setting utility max charging current

*   **Computer:** `MUCHGC<mnn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.9. F: Setting Inverter output rating frequency

*   **Computer:** `F<nn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.10. V: Setting device output rating voltage

*   **Computer:** `V<nnn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.11. POPV: Setting device output rating voltage (for MKSIV)

*   **Computer:** `POPV<nnnn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.12. POP: Setting device output source priority

*   **Computer:** `POP<NN><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.13. PBCV: Set battery re-charge voltage

*   **Computer:** `PBCV<nn.n><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.14. PBDV: Set battery re-discharge voltage

*   **Computer:** `PBDV<nn.n><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.15. PCP: Setting device charger priority

*   **Computer:** `PCP<NN><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.16. PGR: Setting device grid working range

*   **Computer:** `PGR<NN><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.17. PBT: Setting battery type

*   **Computer:** `PBT<NN><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.18. POPM: Set output mode

*   **Computer:** `POPM<nn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.19. PPCP: Setting parallel device charger priority

*   **Computer:** `PPCP<MNN><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.20. PSDV: Setting battery cut-off voltage (Battery under voltage)

*   **Computer:** `PSDV<nn.n><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.21. PCVV: Setting battery C.V. (constant voltage) charging voltage

*   **Computer:** `PCVV<nn.n><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.22. PBFT: Setting battery float charging voltage

*   **Computer:** `PBFT<nn.n><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.23. BTA1: Battery voltage adjust point one

*   **Computer:** `BTA1<nnn.nn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.24. BTA2: Battery voltage adjust point two

*   **Computer:** `BTA2<nnn.nn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.25. BTA0: Set battery voltage adjust parameters to be default value

*   **Computer:** `BTA0<CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.26. RTEY: Reset all stored data for PV/load energy

*   **Computer:** `RTEY<CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.27. RTDL: Erase all data log

*   **Computer:** `RTDL<CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.28. PBEQE: Enable or disable battery equalization

*   **Computer:** `PBEQE<n><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.29. PBEQT: Set battery equalization time

*   **Computer:** `PBEQT<nnn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.30. PBEQP: Set battery equalization period

*   **Computer:** `PBEQP<nnn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.31. PBEQV: Set battery equalization voltage

*   **Computer:** `PBEQV<nn.nn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.32. PBEQOT: Set battery equalization over time

*   **Computer:** `PBEQOT<nnn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.33. PBEQA: Active or inactive battery equalization now

*   **Computer:** `PBEQA<n><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.34. PCVT: Setting max charging time at C.V stage

*   **Computer:** `PCVT<nnn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.35. DAT: Date and time

*   **Computer:** `DAT<YYMMDDHHMMSS><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.36. PBATCD: Battery charge/discharge controlling command

*   **Computer:** `PBATCD<abc><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.37. PBATMAXDISC: Setting max discharging current

*   **Computer:** `PBATMAXDISC<nnn><CRC><cr>`
*   **Device:** `(ACK<CRC><cr>`

### 3.38. PLEDE: Enable/disable LED function

*   **Computer:** `PLEDE<n><cr>`
*   **UPS:** `(ACK<cr>`

### 3.39. PLEDS: set LED speed

*   **Computer:** `PLEDS<n><cr>`
*   **UPS:** `(ACK<cr>`

### 3.40. PLEDM: set LED effect

*   **Computer:** `PLEDM<n><cr>`
*   **UPS:** `(ACK<cr>`

### 3.41. PLEDB: set LED brightness

*   **Computer:** `PLEDB<n><cr>`
*   **UPS:** `(ACK<cr>`

### 3.42. PLEDD: set LED color presentation

*   **Computer:** `PLEDD<n><cr>`
*   **UPS:** `(ACK<cr>`

### 3.43. PLEDC: set LED color

*   **Computer:** `PLEDC<n><aaabbbccc><cr>`
*   **UPS:** `(ACK<cr>`

## 4. Appendix

### 4.1. CRC calibration method

The document mentions a `CRC.c` file, but the content is not provided.
