package net

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"unsafe"
)

type DIOData struct {
	RelayForward, RelayReverse, DigitalOutput uint32
	PWMValues, CANValues                      [10]float32
}

const DIOLength int = int(unsafe.Sizeof(DIOData{}))

type SolenoidData struct {
	State byte
}

const SolenoidLength int = int(unsafe.Sizeof(SolenoidData{}))

type RobotData struct {
	DIO      [2]DIOData
	Solenoid SolenoidData
}

func Unmarshal(data []byte, r *RobotData) error {
	if len(data) != int(unsafe.Sizeof(RobotData{})) {
		return errors.New(fmt.Sprintf("Error: Incorrect number of bytes in packet.\n"))
	}
	for i, dio := range r.DIO {
		offset := i * DIOLength
		r.DIO[i].RelayForward = binary.BigEndian.Uint32(data[offset : offset+4])
		r.DIO[i].RelayReverse = binary.BigEndian.Uint32(data[offset+4 : offset+8])
		r.DIO[i].DigitalOutput = binary.BigEndian.Uint32(data[offset+8 : offset+12])

		for j := range dio.PWMValues {
			pwmOffset := offset + 12 + (4 * j)
			r.DIO[i].PWMValues[j] = math.Float32frombits(
				binary.BigEndian.Uint32(data[pwmOffset : pwmOffset+4]))
		}

		for k := range dio.CANValues {
			canOffset := offset + 52 + (4 * k)
			r.DIO[i].CANValues[k] = math.Float32frombits(
				binary.BigEndian.Uint32(data[canOffset : canOffset+4]))
		}

	}
	r.Solenoid.State = data[2*DIOLength+SolenoidLength]
	return nil
}

func Marshal(r RobotData) ([]byte, error) {
	b := make([]byte, int(unsafe.Sizeof(r)))
	for i, dio := range r.DIO {
		offset := i * DIOLength
		binary.BigEndian.PutUint32(b[offset:], dio.RelayForward)
		binary.BigEndian.PutUint32(b[offset+4:], dio.RelayReverse)
		binary.BigEndian.PutUint32(b[offset+8:], dio.DigitalOutput)

		for j, pwm := range dio.PWMValues {
			pwmOffset := offset + 12 + (4 * j)
			binary.BigEndian.PutUint32(b[pwmOffset:], math.Float32bits(pwm))
		}

		for k, can := range dio.CANValues {
			canOffset := offset + 52 + (4 * k)
			binary.BigEndian.PutUint32(b[canOffset:], math.Float32bits(can))
		}
	}

	b[2*DIOLength+SolenoidLength] = r.Solenoid.State
	return b, nil
}
