// +----------------+         +----------------+         +-----------------+
// |    User        |<--------|   Booking       |-------->|    Payment      |
// |----------------|         |----------------|         |-----------------|
// | Attributes:    |         | Attributes:    |         | Attributes:     |
// | userId         |         | bookingId       |        | paymentId       |
// | name           |         | bookingTime     |        | amount          |
// | email          |         | totalPrice      |        | paymentStatus   |
// | phoneNumber    |         | status          |        | paymentTime     |
// |                |         |                 |        +-----------------+
// | Functions:     |         | Functions:      |               |
// | searchMovie()  |         | confirmBooking()|               |
// | bookTicket()   |         | cancelBooking() |               |
// | cancelTicket() |         +----------------+               |
// +----------------+                |                          |
//                                    v                          v
// +----------------+         +----------------+         +-----------------+
// |    Movie       |<--------|     Show        |-------->|     Seat        |
// |----------------|         |----------------|         |-----------------|
// | Attributes:    |         | Attributes:    |         | Attributes:     |
// | movieId        |         | showId          |        | seatId          |
// | title          |         | startTime       |        | seatNumber      |
// | duration       |         | endTime         |        | isBooked        |
// | genre          |         | date            |        | price           |
// | language       |         |                 |        |                 |
// | rating         |         | Functions:      |        | Functions:      |
// |                |         | getAvailableSeats()|     | bookSeat()      |
// | Functions:     |         | reserveSeats()  |        | cancelBooking() |
// | getDetails()   |         +----------------+         +-----------------+
// | getShows()     |                |                          |
// +----------------+                v                          v
//                         +----------------+         +-----------------+
//                         |    Screen      |<--------|    Theater       |
//                         |----------------|         |-----------------|
//                         | Attributes:    |         | Attributes:     |
//                         | screenId       |         | theaterId       |
//                         | screenName     |         | name            |
//                         | totalSeats     |         | location        |
//                         |                |         | totalScreens    |
//                         | Functions:     |         |                 |
//                         | getSeatAvailability()|    | Functions:      |
//                         |                 |        | getShowsByMovie()|
//                         +----------------+         | getScreens()     |
//                                                     +-----------------+
//                                                             ^
//                                                             |
//                                                       +----------------+
//                                                       |     City        |
//                                                       |-----------------|
//                                                       | Attributes:     |
//                                                       | cityId          |
//                                                       | cityName        |
//                                                       | state           |
//                                                       | country         |
//                                                       |                 |
//                                                       | Functions:      |
//                                                       | getTheaters()   |
//                                                       +-----------------+

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// User instance
	user := User{userId: 1, name: "Ashish", email: "ashish@example.com", phoneNumber: "1234567890"}
	user.SearchMovie()
	user.BookTicket()

	// Movie instance
	movie := Movie{movieId: 101, title: "Inception", duration: 148, genre: "Sci-Fi", language: "English", rating: 4.8}
	movie.GetDetails()
	movie.GetShows()

	// Show instance
	show := Show{showId: 201, startTime: time.Now(), endTime: time.Now().Add(2 * time.Hour), date: time.Now()}
	show.GetAvailableSeats()

	// Seat instance
	seat := Seat{seatId: 301, seatNumber: "A1", isBooked: false, price: 200.0}
	seat.BookSeat()
	seat.CancelBooking()

	// Booking instance
	booking := Booking{bookingId: 401, bookingTime: time.Now(), totalPrice: 400.0, status: "pending"}
	booking.ConfirmBooking()

	// Payment instance
	payment := Payment{paymentId: 501, amount: 400.0, paymentStatus: "pending", paymentTime: time.Now()}
	payment.ProcessPayment()

	// Screen instance
	screen := Screen{screenId: 601, screenName: "Screen 1", totalSeats: 150}
	screen.GetSeatAvailability()

	// Theater instance
	theater := Theater{theaterId: 701, name: "PVR Cinemas", location: "Mumbai", totalScreens: 5}
	theater.GetShowsByMovie()

	// City instance
	city := City{cityId: 801, cityName: "Mumbai", state: "Maharashtra", country: "India"}
	city.GetTheaters()

	// Create a seat
	seat = Seat{seatId: 301, seatNumber: "A1", isBooked: false, price: 200.0}

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Simulate multiple users trying to book the same seat concurrently
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(userId int) {
			defer wg.Done()
			fmt.Printf("User %d is trying to book seat %s...\n", userId, seat.seatNumber)
			seat.BookSeat()
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Simulate a cancellation
	fmt.Println("Cancelling the seat booking...")
	seat.CancelBooking()

	// Simulate another user trying to book the seat again after cancellation
	fmt.Println("User 6 is trying to book the seat after cancellation...")
	seat.BookSeat()

	// Simulate another user trying to book an already booked seat
	fmt.Println("User 7 is trying to book the seat...")
	seat.BookSeat()

	time.Sleep(1 * time.Second) // Give some time for the goroutines to finish
}
