package data

import (
	"time"

	"github.com/hunter32292/go-server-example/pkg/models"
)

func GetHardcodedTrips() []models.Trip {
	Hardtrip1 := models.Trip{
		ID:               "123123",                                                  // Уникальный идентификатор поездки.
		FlightNumber:     "1010",                                                    // Номер рейса.
		DepartureAirport: "Minsk",                                                   // Аэропорт вылета.
		ArrivalAirport:   "Moscow",                                                  // Аэропорт прибытия.
		DepartureTime:    time.Date(2024, time.November, 24, 14, 0, 0, 0, time.UTC), // Время вылета.
		ArrivalTime:      time.Date(2024, time.November, 25, 14, 0, 0, 0, time.UTC), // Время прибытия.
		SeatNumber:       "20",                                                      // Номер места.
		TicketPrice:      500.01,                                                    // Стоимость билета.
		LuggageWeight:    10.5,                                                      // Вес багажа.
		Airline:          "Belavia",                                                 // Авиакомпания.
		BookingReference: "12312320",                                                // Номер бронирования.
	}

	Hardtrip2 := models.Trip{
		ID:               "123124",                                                  // Уникальный идентификатор поездки.
		FlightNumber:     "1011",                                                    // Номер рейса.
		DepartureAirport: "Minsk",                                                   // Аэропорт вылета.
		ArrivalAirport:   "Moscow",                                                  // Аэропорт прибытия.
		DepartureTime:    time.Date(2024, time.November, 25, 14, 0, 0, 0, time.UTC), // Время вылета.
		ArrivalTime:      time.Date(2024, time.November, 26, 14, 0, 0, 0, time.UTC), // Время прибытия.
		SeatNumber:       "25",                                                      // Номер места.
		TicketPrice:      500.01,                                                    // Стоимость билета.
		LuggageWeight:    10.5,                                                      // Вес багажа.
		Airline:          "Belavia",                                                 // Авиакомпания.
		BookingReference: "12312324",                                                // Номер бронирования.
	}

	Hardtrip3 := models.Trip{
		ID:               "123125",                                                  // Уникальный идентификатор поездки.
		FlightNumber:     "1012",                                                    // Номер рейса.
		DepartureAirport: "Minsk",                                                   // Аэропорт вылета.
		ArrivalAirport:   "Moscow",                                                  // Аэропорт прибытия.
		DepartureTime:    time.Date(2024, time.November, 26, 14, 0, 0, 0, time.UTC), // Время вылета.
		ArrivalTime:      time.Date(2024, time.November, 27, 14, 0, 0, 0, time.UTC), // Время прибытия.
		SeatNumber:       "20",                                                      // Номер места.
		TicketPrice:      500.01,                                                    // Стоимость билета.
		LuggageWeight:    10.5,                                                      // Вес багажа.
		Airline:          "Belavia",                                                 // Авиакомпания.
		BookingReference: "12312325",                                                // Номер бронирования.
	}

	Hardtrip4 := models.Trip{
		ID:               "123126",                                                  // Уникальный идентификатор поездки.
		FlightNumber:     "1013",                                                    // Номер рейса.
		DepartureAirport: "Minsk",                                                   // Аэропорт вылета.
		ArrivalAirport:   "Moscow",                                                  // Аэропорт прибытия.
		DepartureTime:    time.Date(2024, time.November, 27, 14, 0, 0, 0, time.UTC), // Время вылета.
		ArrivalTime:      time.Date(2024, time.November, 28, 14, 0, 0, 0, time.UTC), // Время прибытия.
		SeatNumber:       "20",                                                      // Номер места.
		TicketPrice:      500.01,                                                    // Стоимость билета.
		LuggageWeight:    10.5,                                                      // Вес багажа.
		Airline:          "Belavia",                                                 // Авиакомпания.
		BookingReference: "12312326",                                                // Номер бронирования.
	}

	Hardtrip5 := models.Trip{
		ID:               "123127",                                                  // Уникальный идентификатор поездки.
		FlightNumber:     "1014",                                                    // Номер рейса.
		DepartureAirport: "Minsk",                                                   // Аэропорт вылета.
		ArrivalAirport:   "Moscow",                                                  // Аэропорт прибытия.
		DepartureTime:    time.Date(2024, time.November, 28, 14, 0, 0, 0, time.UTC), // Время вылета.
		ArrivalTime:      time.Date(2024, time.November, 29, 14, 0, 0, 0, time.UTC), // Время прибытия.
		SeatNumber:       "20",                                                      // Номер места.
		TicketPrice:      500.01,                                                    // Стоимость билета.
		LuggageWeight:    10.5,                                                      // Вес багажа.
		Airline:          "Belavia",                                                 // Авиакомпания.
		BookingReference: "12312327",                                                // Номер бронирования.
	}

	return []models.Trip{Hardtrip1, Hardtrip2, Hardtrip3, Hardtrip4, Hardtrip5}
}
