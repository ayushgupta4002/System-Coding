package main

import (
	"errors"
	"fmt"
)

type vehicleType string

const (
	BikeType  vehicleType = "BIKE"
	CarType   vehicleType = "CAR"
	TruckType vehicleType = "TRUCK"
)

type VehicleInterface interface {
	GetVehicleType() vehicleType
	GetVehicleNumber() string
}

type Vehicle struct {
	vehicleType   vehicleType
	vehicleNumber string
}

func (v *Vehicle) GetVehicleType() vehicleType {
	return v.vehicleType
}

func (v *Vehicle) GetVehicleNumber() string {
	return v.vehicleNumber
}

// type of vehicles

type Bike struct {
	Vehicle
}

type Car struct {
	Vehicle
}

type Truck struct {
	Vehicle
}

// parking Lot

type ParkingSpot struct {
	occupied bool
	vehicle  *Vehicle
	spotType vehicleType
}

type Ticket struct {
	ticketId int
	spot     *ParkingSpot
	vehicle  *Vehicle
}

type ParkingLot struct {
	spots        []*ParkingSpot
	tickets      map[int]*Ticket
	nextTicketId int
}

func NewParkingLot(pSpots []*ParkingSpot) ParkingLot {
	return ParkingLot{
		spots:        pSpots,
		tickets:      make(map[int]*Ticket),
		nextTicketId: 1,
	}
}

func (p *ParkingLot) park(v *Vehicle) (int, error) {
	for _, s := range p.spots {
		if !s.occupied && s.spotType == v.GetVehicleType() {
			s.occupied = true
			s.vehicle = v

			newTicket := &Ticket{
				ticketId: p.nextTicketId,
				spot:     s,
				vehicle:  v,
			}
			p.tickets[p.nextTicketId] = newTicket
			p.nextTicketId++
			return newTicket.ticketId, nil
		}
	}

	return 0, errors.New("spot not found")
}

func (p *ParkingLot) unpark(ticketId int) error {

	v, exists := p.tickets[ticketId]

	if !exists {
		return errors.New("Ticket not found")
	}
	v.spot.occupied = false
	v.spot.vehicle = nil
	delete(p.tickets, ticketId)

	return nil

}

func main() {
	pSpots := []*ParkingSpot{
		{
			occupied: false,
			vehicle:  nil,
			spotType: BikeType,
		},
		{
			occupied: false,
			vehicle:  nil,
			spotType: BikeType,
		},
		{
			occupied: false,
			vehicle:  nil,
			spotType: CarType,
		},
		{
			occupied: false,
			vehicle:  nil,
			spotType: TruckType,
		},
	}
	p := NewParkingLot(pSpots)
	v := Vehicle{
		vehicleNumber: "3241",
		vehicleType:   CarType,
	}

	v2 := Vehicle{
		vehicleNumber: "3243",
		vehicleType:   CarType,
	}

	ticketId, err := p.park(&v)
	if err != nil {
		fmt.Printf("err is %s", err)
	}
	fmt.Printf("vehicle with ticket : %d ", ticketId)
	p.unpark(ticketId)
	ticketId2, err2 := p.park(&v2)

	if err2 != nil {
		fmt.Printf("err is %s", err2)
	}
	fmt.Printf("\n vehicle with ticket : %d ", ticketId2)

}
