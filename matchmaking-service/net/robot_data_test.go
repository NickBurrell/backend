package net

import "bytes"
import "testing"

var sampleData = []byte{
	0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 63, 0, 0, 0, 63, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 63, 128, 0, 0, 63, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 63, 0, 0, 0, 63, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 63, 128, 0, 0, 63, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 63, 192, 0, 0, 63, 192, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 63, 0, 0, 0, 63, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 63, 128, 0, 0, 63, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 63, 0, 0, 0, 63, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 63, 128, 0, 0, 63, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 63, 192, 0, 0, 63, 192, 0, 0, 0, 128, 0, 0,
}

var sampleRobot = RobotData{
	DIO: [2]DIOData{
		{
			RelayForward:  1,
			RelayReverse:  0,
			DigitalOutput: 255,
			PWMValues: [10]float32{
				0, 0, .5, .5, 0, 0, 1, 1, 0, 0,
			},
			CANValues: [10]float32{
				.5, .5, 0, 0, 1, 1, 0, 0, 1.5, 1.5,
			},
		},
		{
			RelayForward:  0,
			RelayReverse:  1,
			DigitalOutput: 1024,
			PWMValues: [10]float32{
				0, 0, .5, .5, 0, 0, 1, 1, 0, 0,
			},
			CANValues: [10]float32{
				.5, .5, 0, 0, 1, 1, 0, 0, 1.5, 1.5,
			},
		},
	},
	Solenoid: SolenoidData{
		State: byte(128),
	},
}

func TestMarshal(t *testing.T) {
	data, _ := Marshal(sampleRobot)
	if !bytes.Equal(data, sampleData) {
		t.Errorf("Marshal failed: \nExpected: \n%+v,\nGot: \n%+v", sampleData, data)
	}

}

func BenchmarkMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Marshal(sampleRobot)
	}

}

func TestUnmarshal(t *testing.T) {
	var r RobotData
	Unmarshal(sampleData, &r)
	if r != sampleRobot {
		t.Errorf("Marshal failed: \nExpected: \n%#v,\nGot: \n%#v", sampleRobot, r)
	}

}

func BenchmarkUnmarshal(b *testing.B) {
	var r RobotData
	for i := 0; i < b.N; i++ {
		Unmarshal(sampleData, &r)
	}
}
