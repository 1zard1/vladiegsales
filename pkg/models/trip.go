package models

import "time"

// User - Data Model To Store user Data in
type Trip struct {
	ID               string    //  Уникальный идентификатор поездки.
	FlightNumber     string    // Номер рейса.
	DepartureAirport string    // Аэропорт вылета.
	ArrivalAirport   string    // Аэропорт прибытия.
	DepartureTime    time.Time // Время вылета.
	ArrivalTime      time.Time // Время прибытия.
	SeatNumber       string    // Номер места.
	TicketPrice      float64   // Стоимость билета.
	LuggageWeight    float64   // вес багажа.
	Airline          string    //Авиакомпания.
	BookingReference string    // Номер бронирования.
}
