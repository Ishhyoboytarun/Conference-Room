package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestConferenceRoomSystem(t *testing.T) {
	crs := NewConferenceRoomSystem()

	// Test AddBuilding
	crs.AddBuilding("b1")
	if _, ok := crs.Buildings["b1"]; !ok {
		t.Errorf("AddBuilding failed, building 'b1' not found in the system")
	}

	// Test AddFloor
	err := crs.AddFloor("b1", 1)
	if err != nil {
		t.Errorf("AddFloor failed: %v", err)
	}

	// Test AddConferenceRoom
	err = crs.AddConferenceRoom("b1", 1, "c1")
	if err != nil {
		t.Errorf("AddConferenceRoom failed: %v", err)
	}

	// Test BookConferenceRoom
	err = crs.BookConferenceRoom("1:3", "b1", 1, "c1")
	if err != nil {
		t.Errorf("BookConferenceRoom failed: %v", err)
	}

	// Test ListRooms
	expectedListRoomsOutput := "c1 1 b1 []\n"
	listRoomsOutput := captureOutput(func() {
		crs.ListRooms()
	})
	if listRoomsOutput != expectedListRoomsOutput {
		t.Errorf("ListRooms failed, expected output: '%s', got: '%s'", expectedListRoomsOutput, listRoomsOutput)
	}

	// Test ListBookings
	expectedListBookingsOutput := ""
	listBookingsOutput := captureOutput(func() {
		crs.ListBookings("b1", 1)
	})
	if listBookingsOutput != expectedListBookingsOutput {
		t.Errorf("ListBookings failed, expected output: '%s', got: '%s'", expectedListBookingsOutput,
			listBookingsOutput)
	}

	// Test CancelBooking
	err = crs.CancelBooking("1:3", "b1", 1, "c1")
	if err != nil {
		t.Errorf("CancelBooking failed: %v", err)
	}

	// Test ListRooms after cancellation
	expectedListRoomsOutputAfterCancellation := "c1 1 b1 []\n"
	listRoomsOutputAfterCancellation := captureOutput(func() {
		crs.ListRooms()
	})
	if listRoomsOutputAfterCancellation != expectedListRoomsOutputAfterCancellation {
		t.Errorf("ListRooms after cancellation failed, expected output: '%s', got: '%s'",
			expectedListRoomsOutputAfterCancellation, listRoomsOutputAfterCancellation)
	}
}

// Helper function to capture stdout output
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
