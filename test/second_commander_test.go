package test

import (
	"random-commander/services"
	"strings"
	"testing"
)

func TestDoctorsCompanionReturnsNewDoctor(t *testing.T) {
	//oracleText := "Doctor's Companion"
	want := "Time Lord Doctor"
	secondCmdr := services.FetchSecondCommander("Time Lord Doctor")

	if !strings.Contains(secondCmdr.TypeLine, want) {
		t.Fatalf("Did not get a doctor")
	}
}

func TestDoctorReturnsDoctorsCompanion(t *testing.T) {
	want := "Doctor's companion"
	secondCmdr := services.FetchSecondCommander("Doctor's Companion")

	if !strings.Contains(secondCmdr.OracleText, want) {
		t.Fatalf("Did not get a doctor's companion")
	}
}

func TestChooseABackgroundReturnsBackground(t *testing.T) {
	want := "Choose a Background"
	secondCmdr := services.FetchSecondCommander("Choose a Background")

	if !strings.Contains(secondCmdr.OracleText, want) {
		t.Fatalf("Failed at getting a choose a background commander")
	}
}

func TestBackground(t *testing.T) {
	want := "Background"
	secondCmdr := services.FetchSecondCommander("Background")
	if !strings.Contains(secondCmdr.TypeLine, want) {
		t.Fatalf("Did not get a background commander")
	}
}

func TestPartner(t *testing.T) {
	want := "Partner"
	secondCmdr := services.FetchSecondCommander("Partner")
	if !strings.Contains(secondCmdr.OracleText, want) {
		t.Fatalf("Did not get a partner commander")
	}
}

func TestFriendsForever(t *testing.T) {
	want := "Friends Forever"
	secondCmdr := services.FetchSecondCommander("Friends Forever")
	if !strings.Contains(secondCmdr.OracleText, want) {
		t.Fatalf("Did not get a friend forever commander")
	}
}

func TestPartnerWith(t *testing.T) {
	t.Fatalf("Unimplemented")
}
