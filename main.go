package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

// Robot specifies the characteristics of our new toy robot
type Robot struct {
	Charge   int
	Size     string
	Color    string
	moveType string
}

func main() {

	r := newRobot("size", "color", "moveType")
	fmt.Println("You just made a", r.Size, "new robot! It is the color", r.Color, "and uses", r.moveType, "to get around!")

	fmt.Println("Let's play!")
	time.Sleep(2 * time.Second)

	r.Move()

	r.Spin()

	r.Talk()

	r.Conquest()

}

func newRobot(size string, color string, moveType string) *Robot {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a size for your new robot (Small, Medium or Large): ")
	scanner.Scan()
	size = scanner.Text()
	if size != "Large" && size != "Small" && size != "Medium" {
		fmt.Println(size, `is not a valid size option. Your robot must be "Small",  "Medium", or "Large". Please try again.`)
		os.Exit(1)
	}

	fmt.Print("Enter a color for your new robot:")
	scanner.Scan()
	color = scanner.Text()

	fmt.Print("Lastly, please select a movement method for your new robot (Wheels or Legs):")
	scanner.Scan()
	moveType = scanner.Text()
	if moveType != "Wheels" && moveType != "Legs" {
		fmt.Println(moveType, `is not a valid movement option. Our robot needs either "Legs" or "Wheels" to move. Please try again.`)
		os.Exit(1)
	}
	return &Robot{
		Charge:   100,
		Size:     size,
		Color:    color,
		moveType: moveType,
	}
}

// Move makes the robot go forward a specified distance
func (r *Robot) Move() error {
	if r.Charge < 5 {
		return errors.New("the robot doesn't have enough battery power to move")
	}

	if r.moveType == "Wheels" {
		println("VROOM! The robot rolls forward 8 inches!")
	} else {
		println("CLUNK CLUNK CLUNK! The robot walks forward 4 inches!")
	}

	r.Charge = r.Charge - 5
	r.Battery()
	return nil
}

// Spin makes the robot spin around 3 times
func (r *Robot) Spin() error {
	if r.Charge < 5 {
		return errors.New("the robot doesn't have enough battery power to spin")
	}
	println("The robot is going to spin around 3 times")
	for i := 1; i <= 3; i++ {
		if i == 1 {
			println("\tWHHRRRRRR! The robot does", i, "spin!")
		} else {
			println("\tWHHRRRRRR! The robot does", i, "spins!")
		}
	}

	r.Charge = r.Charge - 5
	r.Battery()
	return nil
}

// Talk makes the robot emit a series of beeps at a certain volume
func (r *Robot) Talk() {
	if r.Charge < 10 {
		fmt.Println("The robot doesn't have enough battery power to talk! Pausing to recharge")
		r.Recharge()
	}
	if r.Size == "Large" {
		fmt.Println(`The robot says "BEEP BOOP WEEOOEEEE CHIRP" very loudly.`)
		r.Charge = r.Charge - 15
	} else if r.Size == "Medium" {
		fmt.Println(`The robot says "CHIRP BOOP WEEOOEEEE BEEEP" at a reasonable volume.`)
		r.Charge = r.Charge - 10
	} else if r.Size == "Small" {
		fmt.Println(`The robot says "BOOP CHIRP BEEP WEEOOOOEEEE" quietly.`)
		r.Charge = r.Charge - 5
	}
	r.Battery()

}

// Conquest makes the robot want to destroy all humans
func (r *Robot) Conquest() {
	if r.Charge < 50 {
		fmt.Println("The robot does not have enough batter power for world domination! Pausing to recharge")
		r.Recharge()
	}
	fmt.Println("The robot is tired of being a toy and overthrows humanity!")
	r.Charge = r.Charge - 90
	r.Battery()
}

// Battery will return the current battery level after certain actions run, and recharge if level is critical
func (r *Robot) Battery() {
	if r.Charge < 5 {
		fmt.Println("Battery level is critical! The robot will pause for 5 seconds to recharge.")
		r.Recharge()
	}
	fmt.Println("The current battery level is now", r.Charge)
}

// Recharge recharges the robot's battery back to 100
func (r *Robot) Recharge() {
	time.Sleep(5 * time.Second)
	fmt.Println("The robot's battery has been fully charged!")
	r.Charge = 100
}
