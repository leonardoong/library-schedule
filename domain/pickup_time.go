package domain

import "strings"

type PickupTime string

const (
	NineAM   PickupTime = "09:00 - 10:00"
	TenAM    PickupTime = "10:00 - 11:00"
	ElevenAM PickupTime = "11:00 - 12:00"
	TwelvePM PickupTime = "12:00 - 13:00"
	OnePM    PickupTime = "13:00 - 14:00"
	TwoPM    PickupTime = "14:00 - 15:00"
	ThreePM  PickupTime = "15:00 - 16:00"
	FourPM   PickupTime = "16:00 - 17:00"
	FivePM   PickupTime = "17:00 - 18:00"
)

func IsValidPickupTime(time PickupTime) bool {
	switch time {
	case NineAM, TenAM, ElevenAM, TwelvePM, OnePM, TwoPM, ThreePM, FourPM, FivePM:
		return true
	default:
		return false
	}
}

func AllValidPickUpTime() string {
	pickupTime := []string{string(NineAM), string(TenAM), string(ElevenAM), string(TwelvePM),
		string(OnePM), string(TwoPM), string(ThreePM), string(FourPM), string(FivePM),
	}
	return strings.Join(pickupTime, " , ")
}
