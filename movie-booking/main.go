package main

import (
	"fmt"
	"sync"
)

type Movie struct {
	movieName string
	rating    int
}

type Screen struct {
	screenNo int
	seats    map[int]*Seat
}

type Seat struct {
	seatId int
}

type Show struct {
	movie       Movie
	startTime   int
	endTime     int
	screen      *Screen
	bookedSeats map[int]bool
	mu          sync.Mutex
}

func (s *Show) BookSeat(seatNumber int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.screen.seats[seatNumber]
	if !exists {
		return fmt.Errorf("seat %d does not exist", seatNumber)
	}

	if s.bookedSeats[seatNumber] {
		return fmt.Errorf("seat %d is already booked", seatNumber)
	}

	s.bookedSeats[seatNumber] = true

	fmt.Printf(
		"Seat %d booked for movie %s on Screen %d (%d-%d)\n",
		seatNumber,
		s.movie.movieName,
		s.screen.screenNo,
		s.startTime,
		s.endTime,
	)

	return nil
}

type Theatre struct {
	id       int
	name     string
	location string
	screen   map[int]*Screen
	shows    []*Show
	mu       sync.Mutex
}

func (th *Theatre) addShow(movie Movie, screenNo int, startTime int, endTime int) error {
	th.mu.Lock()
	defer th.mu.Unlock()
	// Check all shows
	for _, show := range th.shows {

		if show.screen.screenNo != screenNo {
			continue
		}

		if startTime < show.endTime &&
			show.startTime < endTime {

			return fmt.Errorf("screen already occupied during this time")
		}
	}

	screen, ok := th.screen[screenNo]
	if !ok {
		return fmt.Errorf("invalid screen")
	}

	newShow := &Show{
		movie:       movie,
		screen:      screen,
		startTime:   startTime,
		endTime:     endTime,
		bookedSeats: make(map[int]bool),
	}

	th.shows = append(th.shows, newShow)

	return nil
}
func (th *Theatre) getShows() []*Show {
	th.mu.Lock()
	defer th.mu.Unlock()
	return th.shows
}

type MovieBookingSystem struct {
	theatres      map[int]*Theatre
	nextTheatreId int
	mu            sync.Mutex
}

func (ms *MovieBookingSystem) addTheatre(name string, location string, screenCount int) *Theatre {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	screens := make(map[int]*Screen)

	for i := range screenCount {
		newScreen := &Screen{
			screenNo: i + 1,
			seats:    make(map[int]*Seat),
		}
		//50 seats in each screen
		for k := 1; k < 50+1; k++ {
			newScreen.seats[k] = &Seat{
				seatId: k,
			}
		}
		screens[i+1] = newScreen
	}

	th := &Theatre{
		id:       ms.nextTheatreId,
		name:     name,
		location: location,
		screen:   screens,
		shows:    make([]*Show, 0),
	}
	ms.theatres[ms.nextTheatreId] = th
	ms.nextTheatreId++
	return th
}

func main() {

	fmt.Println("adding a movie theatre")

	ms := MovieBookingSystem{
		theatres:      make(map[int]*Theatre),
		nextTheatreId: 1,
	}
	th := ms.addTheatre("inox", "gurugram", 4)
	movie := Movie{
		movieName: "avenegers",
		rating:    5,
	}
	movie2 := Movie{
		movieName: "jab we met",
		rating:    4,
	}
	th.addShow(movie, 2, 12, 14)
	th.addShow(movie2, 1, 9, 11)
	shows := th.getShows()
	fmt.Println("Available shows are : ")
	if len(shows) < 1 {
		fmt.Println("no shows available")
	} else {
		for _, show := range shows {
			fmt.Printf("\n Movie Name: %s ", show.movie.movieName)
			fmt.Printf("\n Screen Number: %d ", show.screen.screenNo)
			fmt.Printf("\n Show Timings : %d -- %d  ", show.startTime, show.endTime)
			fmt.Println("---------------------------------")
		}

		if err := shows[1].BookSeat(2); err != nil {
			fmt.Println(err)
		}

		if err := shows[1].BookSeat(2); err != nil {
			fmt.Println(err)
		}

		//booking parallely

		var wg sync.WaitGroup

		for i := 0; i < 2; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				if err := shows[1].BookSeat(3); err != nil {
					fmt.Println(err)
				}
			}()
		}

		wg.Wait()

	}

}
