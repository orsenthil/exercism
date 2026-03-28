package elon

import "fmt"

// TODO: define the 'Drive()' method

func (car *Car) Drive() {
    if (car.battery == 0 || (car.batteryDrain > car.battery)) {
        // return car
    }
	car.battery = car.battery - car.batteryDrain
    car.distance = car.distance + car.speed    
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
    for car.distance < trackDistance {
        if (car.battery == 0 || car.batteryDrain > car.battery) {
            return false
        }   
    	car.Drive()        
    }
	return true
}

// TODO: define the 'DisplayDistance() string' method

func (car Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", car.distance)
}

// TODO: define the 'DisplayBattery() string' method

func (car Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", car.battery)
}
