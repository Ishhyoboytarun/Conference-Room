package main

import (
	"errors"
	"fmt"
)

// ConferenceRoom represents a conference room with its attributes
type ConferenceRoom struct {
	ID       string
	Floor    int
	Building string
	Bookings []string
}

// Building represents a building with its conference rooms
type Building struct {
	Name          string
	Floors        map[int][]ConferenceRoom
	ConferenceMap map[string]ConferenceRoom
}

// ConferenceRoomSystem represents the conference room management system
type ConferenceRoomSystem struct {
	Buildings map[string]Building
}

// NewConferenceRoomSystem creates a new instance of ConferenceRoomSystem
func NewConferenceRoomSystem() *ConferenceRoomSystem {
	return &ConferenceRoomSystem{
		Buildings: make(map[string]Building),
	}
}

// AddBuilding adds a new building to the system
func (crs *ConferenceRoomSystem) AddBuilding(buildingName string) {
	if _, ok := crs.Buildings[buildingName]; !ok {
		crs.Buildings[buildingName] = Building{
			Name:          buildingName,
			Floors:        make(map[int][]ConferenceRoom),
			ConferenceMap: make(map[string]ConferenceRoom),
		}
		fmt.Printf("Added building %s into the system.\n", buildingName)
	} else {
		fmt.Printf("Building %s already exists in the system.\n", buildingName)
	}
}

// AddFloor adds a new floor to the specified building
func (crs *ConferenceRoomSystem) AddFloor(buildingName string, floor int) error {
	building, ok := crs.Buildings[buildingName]
	if !ok {
		return errors.New("building not found")
	}

	if _, ok := building.Floors[floor]; !ok {
		building.Floors[floor] = make([]ConferenceRoom, 0)
		crs.Buildings[buildingName] = building
		fmt.Printf("Added floor %d in building %s.\n", floor, buildingName)
		return nil
	}

	return errors.New("floor already exists")
}

// AddConferenceRoom adds a new conference room to the specified floor of a building
func (crs *ConferenceRoomSystem) AddConferenceRoom(buildingName string, floor int, roomID string) error {
	building, ok := crs.Buildings[buildingName]
	if !ok {
		return errors.New("building not found")
	}

	rooms, ok := building.Floors[floor]
	if !ok {
		return errors.New("floor not found")
	}

	for _, room := range rooms {
		if room.ID == roomID {
			return errors.New("conference room already exists")
		}
	}

	newRoom := ConferenceRoom{
		ID:       roomID,
		Floor:    floor,
		Building: buildingName,
	}

	rooms = append(rooms, newRoom)
	building.Floors[floor] = rooms
	building.ConferenceMap[roomID] = newRoom
	crs.Buildings[buildingName] = building

	fmt.Printf("Added conference room %s in floor %d of building %s.\n", roomID, floor, buildingName)
	return nil
}

// ListRooms lists all the available conference rooms
func (crs *ConferenceRoomSystem) ListRooms() {
	for _, building := range crs.Buildings {
		for _, floors := range building.Floors {
			for _, room := range floors {
				fmt.Printf("%s %d %s %v\n", room.ID, room.Floor, room.Building, room.Bookings)
			}
		}
	}
}

// BookConferenceRoom books a conference room for the given slot
func (crs *ConferenceRoomSystem) BookConferenceRoom(slot string, buildingName string, floor int, roomID string) error {
	building, ok := crs.Buildings[buildingName]
	if !ok {
		return errors.New("building not found")
	}

	room, ok := building.ConferenceMap[roomID]
	if !ok {
		return errors.New("conference room not found")
	}

	for _, booking := range room.Bookings {
		if booking == slot {
			return errors.New("conference room already booked for this slot")
		}
	}

	room.Bookings = append(room.Bookings, slot)
	building.ConferenceMap[roomID] = room
	crs.Buildings[buildingName] = building

	fmt.Printf("Booked conference room %s for slot %s on floor %d of building %s.\n", roomID, slot, floor, buildingName)
	return nil
}

// CancelBooking cancels the booking for the given slot
func (crs *ConferenceRoomSystem) CancelBooking(slot string, buildingName string, floor int, roomID string) error {
	building, ok := crs.Buildings[buildingName]
	if !ok {
		return errors.New("building not found")
	}

	room, ok := building.ConferenceMap[roomID]
	if !ok {
		return errors.New("conference room not found")
	}

	var index int
	found := false
	for i, booking := range room.Bookings {
		if booking == slot {
			index = i
			found = true
			break
		}
	}

	if found {
		room.Bookings = append(room.Bookings[:index], room.Bookings[index+1:]...)
		building.ConferenceMap[roomID] = room
		crs.Buildings[buildingName] = building

		fmt.Printf("Cancelled booking for slot %s on floor %d of building %s.\n", slot, floor, buildingName)
		return nil
	}

	return errors.New("booking not found for the specified slot")
}

// ListBookings lists all the bookings
func (crs *ConferenceRoomSystem) ListBookings(buildingName string, floor int) {
	building, ok := crs.Buildings[buildingName]
	if !ok {
		fmt.Println("Building not found.")
		return
	}

	rooms, ok := building.Floors[floor]
	if !ok {
		fmt.Println("Floor not found.")
		return
	}

	for _, room := range rooms {
		for _, booking := range room.Bookings {
			fmt.Printf("%s %d %s %s\n", booking, floor, buildingName, room.ID)
		}
	}
}

func main() {
	crs := NewConferenceRoomSystem()

	crs.AddBuilding("b1")
	crs.AddFloor("b1", 1)
	crs.AddConferenceRoom("b1", 1, "c1")
	crs.BookConferenceRoom("1:3", "b1", 1, "c1")
	crs.ListRooms()
	crs.ListBookings("b1", 1)
	crs.CancelBooking("1:3", "b1", 1, "c1")
	crs.ListRooms()
}
