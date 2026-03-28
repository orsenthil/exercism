package robot

import "fmt"

// See defs.go for other definitions

// Step 1
// Define N, E, S, W here.
const N = 0
const E = 1
const S = 2
const W = 3

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 4 - 1) % 4
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

func (d Dir) String() string {
	switch d {
	case N:
		return "N"
	case E:
		return "E"
	case S:
		return "S"
	case W:
		return "W"
	}
	return "N"
}

// // Step 2
// // Define Action type here.
type Action struct {
	cmd Command
}

func (robot *Step2Robot) Right() {
	robot.Dir = (robot.Dir + 1) % 4
}

func (robot *Step2Robot) Left() {
	robot.Dir = (robot.Dir + 4 - 1) % 4
}

func (robot *Step2Robot) Advance(extent Rect) {
	switch robot.Dir {
	case N:
		if robot.Pos.Northing < extent.Max.Northing {
			robot.Pos.Northing++
		}
	case E:
		if robot.Pos.Easting < extent.Max.Easting {
			robot.Pos.Easting++
		}
	case S:
		if robot.Pos.Northing > extent.Min.Northing {
			robot.Pos.Northing--
		}
	case W:
		if robot.Pos.Easting > extent.Min.Easting {
			robot.Pos.Easting--
		}
	}
}

func StartRobot(command chan Command, action chan Action) {
	for cmd := range command {
		action <- Action{cmd}
	}
	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	fmt.Println("Robot", robot.Pos, robot.Dir)
	for act := range action {
		fmt.Println("Performing action " + string(act.cmd))
		switch act.cmd {
		case 'R':
			robot.Right()
		case 'L':
			robot.Left()
		case 'A':
			robot.Advance(extent)
		}
		fmt.Println("Robot", robot.Pos, robot.Dir)
	}
	fmt.Println("-----------------")
	report <- robot
	close(report)
}

// // Step 3
// // Define Action3 type here.
type Action3 struct {
	cmd  rune
	name string
}

func (robot *Step3Robot) Right() {
	robot.Dir = (robot.Dir + 1) % 4
}

func (robot *Step3Robot) Left() {
	robot.Dir = (robot.Dir + 4 - 1) % 4
}

func (robot *Step3Robot) Advance(extent Rect, positions map[string]Pos) bool {
	nextPosition := robot.Pos

	switch robot.Dir {
	case N:
		if nextPosition.Northing < extent.Max.Northing {
			nextPosition.Northing++
		} else {
			return false
		}
	case E:
		if nextPosition.Easting < extent.Max.Easting {
			nextPosition.Easting++
		} else {
			return false
		}
	case S:
		if nextPosition.Northing > extent.Min.Northing {
			nextPosition.Northing--
		} else {
			return false
		}
	case W:
		if nextPosition.Easting > extent.Min.Easting {
			nextPosition.Easting--
		} else {
			return false
		}
	}

	for name, pos := range positions {
		if name == robot.Name {
			continue
		}
		if pos.Easting == nextPosition.Easting && pos.Northing == nextPosition.Northing {
			return false
		}
	}

	robot.Pos = nextPosition

	return true
}

func (robot *Step3Robot) IsInside(extent Rect) bool {
	return robot.Northing <= extent.Max.Northing &&
		robot.Easting <= extent.Max.Easting &&
		robot.Northing >= extent.Min.Northing &&
		robot.Easting >= extent.Min.Easting
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	for _, cmd := range script {
		action <- Action3{cmd, name}
	}
	action <- Action3{'0', name}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {

	positions := make(map[string]Pos)

	error := checkInitialState(robots, extent, log)
	if error {
		rep <- robots
		return
	}

	remaining := len(robots)

	for act := range action {

		robot := findRobot(robots, act.name)
		if robot == nil {
			log <- "Bad robot " + act.name
			break
		}

		_, ok := positions[act.name]
		if !ok {
			positions[act.name] = robot.Pos
		}

		if act.cmd == '0' {
			//This robot has finished
			remaining--
			if remaining == 0 {
				break
			} else {
				continue
			}
		}

		if act.cmd != 'R' && act.cmd != 'A' && act.cmd != 'L' {
			log <- "Bad command: " + string(act.cmd)
			break
		}

		fmt.Println("Performing action " + string(act.cmd))
		switch act.cmd {
		case 'R':
			robot.Right()
		case 'L':
			robot.Left()
		case 'A':
			if !robot.Advance(extent, positions) {
				log <- "Tried to advance into a wall or another robot"
			}
		}
		positions[act.name] = robot.Pos
		fmt.Println("Robot", robot.Pos, robot.Dir)
	}

	rep <- robots
}

func findRobot(robots []Step3Robot, name string) *Step3Robot {
	var robot *Step3Robot
	for r := range robots {
		if robots[r].Name == name {
			robot = &robots[r]
			break
		}
	}
	return robot
}

func checkInitialState(robots []Step3Robot, extent Rect, log chan string) bool {
	for i := range robots {
		if robots[i].Name == "" {
			log <- "A robot without a name"
			return true
		}

		if !robots[i].IsInside(extent) {
			log <- "Robot is outside bounds: " + robots[i].Name
			return true
		}

		for j := i + 1; j < len(robots); j++ {
			if robots[i].Name == robots[j].Name {
				log <- "Robots with the same name: " + robots[i].Name
				return true
			}
			if robots[i].Pos == robots[j].Pos {
				log <- "Robots with the same position: " + robots[i].Name
				return true
			}
		}
	}
	return false
}
