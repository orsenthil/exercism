package tournament

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Tally(reader io.Reader, writer io.Writer) error {
	var line string
	var records []string

	wins := make(map[string]int)
	losses := make(map[string]int)
	draws := make(map[string]int)

	teams := make(map[string]int)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line = scanner.Text()
		records = strings.Split(line, ";")
		first, second, result := records[0], records[1], records[2]
		if result == "draw" {
			draws[first] += 1
			draws[second] += 1
		}
		if result == "win" {
			wins[first] += 1
			losses[second] += 1
		}
		if result == "loss" {
			losses[first] += 1
			wins[second] += 1
		}

		teams[first] = 0
		teams[second] = 0
	}

	bufw := bufio.NewWriter(writer)

	_, err := fmt.Fprint(bufw, fmt.Sprintf("Team                           | MP |  W |  D |  L |  P"))

	if err != nil {
		return err
	}

	for team, _ := range teams {
		_, err = fmt.Fprint(bufw, fmt.Sprintf("%30s|%d|%d|%d|%d", team, wins[team], draws[team], losses[team], wins[team]+draws[team]+losses[team]))
		if err != nil {
			return err
		}
	}

	err = bufw.Flush()

	if err != nil {
		return err
	}

	return nil
}
