import unittest
from python import conference


class ConferenceRoomSystemTests(unittest.TestCase):

    def setUp(self):
        self.crs = conference.ConferenceRoomSystem()

    def test_add_building(self):
        self.crs.add_building("b1")
        self.assertEqual(len(self.crs.Buildings), 1)

        self.crs.add_building("b2")
        self.assertEqual(len(self.crs.Buildings), 2)

        self.crs.add_building("b1")
        self.assertEqual(len(self.crs.Buildings), 2)

    def test_add_floor(self):
        self.crs.add_building("b1")
        self.crs.add_floor("b1", 1)
        self.assertEqual(len(self.crs.Buildings["b1"].Floors), 1)

        self.crs.add_floor("b1", 2)
        self.assertEqual(len(self.crs.Buildings["b1"].Floors), 2)

        self.crs.add_floor("b2", 1)

    def test_add_conference_room(self):
        self.crs.add_building("b1")
        self.crs.add_floor("b1", 1)
        self.crs.add_conference_room("b1", 1, "c1")
        self.assertEqual(len(self.crs.Buildings["b1"].Floors[1]), 1)

        self.crs.add_conference_room("b1", 1, "c1")

        self.crs.add_conference_room("b2", 1, "c2")

        self.crs.add_conference_room("b1", 2, "c2")

    def test_book_conference_room(self):
        self.crs.add_building("b1")
        self.crs.add_floor("b1", 1)
        self.crs.add_conference_room("b1", 1, "c1")

        self.crs.book_conference_room("1:3", "b1", 1, "c1")
        self.assertEqual(len(self.crs.Buildings["b1"].ConferenceMap["c1"].Bookings), 1)

        self.crs.book_conference_room("1:3", "b1", 1, "c2")

        self.crs.book_conference_room("1:3", "b2", 1, "c1")

    def test_cancel_booking(self):
        self.crs.add_building("b1")
        self.crs.add_floor("b1", 1)
        self.crs.add_conference_room("b1", 1, "c1")
        self.crs.book_conference_room("1:3", "b1", 1, "c1")

        self.crs.cancel_booking("1:3", "b1", 1, "c1")
        self.assertEqual(len(self.crs.Buildings["b1"].ConferenceMap["c1"].Bookings), 0)

        self.crs.cancel_booking("1:3", "b1", 1, "c2")

        self.crs.cancel_booking("1:3", "b2", 1, "c1")

    def test_list_rooms(self):
        self.crs.add_building("b1")
        self.crs.add_floor("b1", 1)
        self.crs.add_conference_room("b1", 1, "c1")
        self.crs.add_building("b2")
        self.crs.add_floor("b2", 2)
        self.crs.add_conference_room("b2", 2, "c2")

        expected_output = None
        self.assertEqual(self.crs.list_rooms(), expected_output)

    def test_list_bookings(self):
        self.crs.add_building("b1")
        self.crs.add_floor("b1", 1)
        self.crs.add_conference_room("b1", 1, "c1")
        self.crs.book_conference_room("1:3", "b1", 1, "c1")

        expected_output = None
        self.assertEqual(self.crs.list_bookings("b1", 1), expected_output)


if __name__ == '__main__':
    unittest.main()
