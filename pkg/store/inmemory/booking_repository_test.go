package inmemory

import (
	"testing"
	"time"

	"github.com/alewkinr/example-app-design-review/internal/booking"
)

func TestBookingRepository_isIntersected(t *testing.T) {
	const testRoomID, testHotelID = "lux", "reddison"
	type fields struct {
		store map[string][]booking.Booking
	}
	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"has intersections 1",
			fields{
				store: map[string][]booking.Booking{
					"1": {
						{
							Room: booking.Room{
								ID:      testRoomID,
								HotelID: testHotelID,
							},
							CheckInDateTime:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
							CheckOutDateTime: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			args{
				from: time.Date(2023, 12, 31, 23, 0, 0, 0, time.UTC),
				to:   time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
			},
			true,
		},

		{
			"has intersections 2",
			fields{
				store: map[string][]booking.Booking{
					"1": {
						{
							Room: booking.Room{
								ID:      testRoomID,
								HotelID: testHotelID,
							},
							CheckInDateTime:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
							CheckOutDateTime: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			args{
				from: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
				to:   time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			},
			true,
		},

		{
			"has intersections 3",
			fields{
				store: map[string][]booking.Booking{
					"1": {
						{
							Room: booking.Room{
								ID:      testRoomID,
								HotelID: testHotelID,
							},
							CheckInDateTime:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
							CheckOutDateTime: time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			args{
				from: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			},
			true,
		},

		{
			"has intersections 4",
			fields{
				store: map[string][]booking.Booking{
					"1": {
						{
							Room: booking.Room{
								ID:      testRoomID,
								HotelID: testHotelID,
							},
							CheckInDateTime:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
							CheckOutDateTime: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			args{
				from: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			true,
		},

		{
			"no intersections 1",
			fields{
				store: map[string][]booking.Booking{
					"1": {
						{
							Room: booking.Room{
								ID:      testRoomID,
								HotelID: testHotelID,
							},
							CheckInDateTime:  time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
							CheckOutDateTime: time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			args{
				from: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			false,
		},

		{
			"no intersections 2",
			fields{
				store: map[string][]booking.Booking{
					"1": {
						{
							Room: booking.Room{
								ID:      testRoomID,
								HotelID: testHotelID,
							},
							CheckInDateTime:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
							CheckOutDateTime: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			args{
				from: time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC),
			},
			false,
		},

		{
			"no intersections 3",
			fields{
				store: map[string][]booking.Booking{
					"1": {
						{
							Room: booking.Room{
								ID:      testRoomID,
								HotelID: testHotelID,
							},
							CheckInDateTime:  time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
							CheckOutDateTime: time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			args{
				from: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2024, 1, 2, 23, 59, 59, 59, time.UTC),
			},
			false,
		},

		{
			"no intersections 4",
			fields{
				store: map[string][]booking.Booking{
					"1": {
						{
							Room: booking.Room{
								ID:      testRoomID,
								HotelID: testHotelID,
							},
							CheckInDateTime:  time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
							CheckOutDateTime: time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			args{
				from: time.Date(2024, 1, 5, 0, 0, 0, 1, time.UTC),
				to:   time.Date(2024, 1, 6, 0, 0, 0, 0, time.UTC),
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &BookingRepository{
				store: tt.fields.store,
			}

			if got := r.isIntersected(tt.fields.store["1"][0], tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("isIntersected() = %v, want %v", got, tt.want)
			}
		})
	}
}
