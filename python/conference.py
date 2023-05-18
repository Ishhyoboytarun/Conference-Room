class ConferenceRoom:
    def __init__(self, room_id, floor, building):
        self.ID = room_id
        self.Floor = floor
        self.Building = building
        self.Bookings = []


class Building:
    def __init__(self, name):
        self.Name = name
        self.Floors = {}
        self.ConferenceMap = {}


class ConferenceRoomSystem:
    def __init__(self):
        self.Buildings = {}

    def add_building(self, building_name):
        if building_name not in self.Buildings:
            self.Buildings[building_name] = Building(building_name)
            print(f"Added building {building_name} into the system.")
        else:
            print(f"Building {building_name} already exists in the system.")

    def add_floor(self, building_name, floor):
        if building_name in self.Buildings:
            building = self.Buildings[building_name]
            if floor not in building.Floors:
                building.Floors[floor] = []
                print(f"Added floor {floor} in building {building_name}.")
            else:
                print("Floor already exists.")
        else:
            print("Building not found.")

    def add_conference_room(self, building_name, floor, room_id):
        if building_name in self.Buildings:
            building = self.Buildings[building_name]
            if floor in building.Floors:
                rooms = building.Floors[floor]
                for room in rooms:
                    if room.ID == room_id:
                        print("Conference room already exists.")
                        return

                new_room = ConferenceRoom(room_id, floor, building_name)
                rooms.append(new_room)
                building.Floors[floor] = rooms
                building.ConferenceMap[room_id] = new_room
                print(f"Added conference room {room_id} in floor {floor} of building {building_name}.")
            else:
                print("Floor not found.")
        else:
            print("Building not found.")

    def list_rooms(self):
        for building in self.Buildings.values():
            for floors in building.Floors.values():
                for room in floors:
                    print(f"{room.ID} {room.Floor} {room.Building} {room.Bookings}")

    def book_conference_room(self, slot, building_name, floor, room_id):
        if building_name in self.Buildings:
            building = self.Buildings[building_name]
            if floor in building.Floors:
                room = building.ConferenceMap.get(room_id)
                if room:
                    for booking in room.Bookings:
                        if booking == slot:
                            print("Conference room already booked for this slot.")
                            return

                    room.Bookings.append(slot)
                    print(
                        f"Booked conference room {room_id} for slot {slot} on floor {floor} of building {building_name}.")
                else:
                    print("Conference room not found.")
            else:
                print("Floor not found.")
        else:
            print("Building not found.")

    def cancel_booking(self, slot, building_name, floor, room_id):
        if building_name in self.Buildings:
            building = self.Buildings[building_name]
            if floor in building.Floors:
                room = building.ConferenceMap.get(room_id)
                if room:
                    if slot in room.Bookings:
                        room.Bookings.remove(slot)
                        print(f"Cancelled booking for slot {slot} on floor {floor} of building {building_name}.")
                    else:
                        print("Booking not found for the specified slot.")
                else:
                    print("Conference room not found.")
            else:
                print("Floor not found.")
        else:
            print("Building not found.")

    def list_bookings(self, building_name, floor):
        building = self.Buildings.get(building_name)
        if building:
            rooms = building.Floors.get(floor)
            if rooms:
                for room in rooms:
                    for booking in room.Bookings:
                        print(f"{booking} {floor} {building_name} {room.ID}")
            else:
                print("Floor not found.")
        else:
            print("Building not found.")


def main():
    crs = ConferenceRoomSystem()

    crs.add_building("b1")
    crs.add_floor("b1", 1)
    crs.add_conference_room("b1", 1, "c1")
    crs.book_conference_room("1:3", "b1", 1, "c1")
    crs.list_rooms()
    crs.list_bookings("b1", 1)
    crs.cancel_booking("1:3", "b1", 1, "c1")
    crs.list_rooms()


if __name__ == "__main__":
    main()
