package broadpeakgo

type Broadpeak struct {
	apiKey string
}

func GetBroadPeak(apiKey string) Broadpeak {
	b := Broadpeak{apiKey}
	return b
}
