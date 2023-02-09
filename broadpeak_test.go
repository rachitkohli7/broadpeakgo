package broadpeakgo

import (
	"fmt"
	"testing"
)

const apiKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InRlc3RpbmciLCJ0ZW5hbnRJZCI6MTM2NCwiaWF0IjoxNjc0MjEzODk1fQ.DSYz972uYnEOjIUkYoWxdU47DLwVzo5uYsgMdbzy_qo"

func TestCreateAsse(t *testing.T) {
	broadpeak := GetBroadPeak(apiKey)
	fmt.Println(broadpeak)

	//broadpeak.GetAllSources()
	broadpeak.CreateAsset("rachit", "http://localhost:3000", "", "")

}
