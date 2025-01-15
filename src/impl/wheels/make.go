package wheels

import (
	"fmt"
	"strings"
    . "tzgolkin/engine"
)

func MakeWheel(options []Options, Wheel_id int, wheel_name string) *Wheel {
	positions := make([]*Position, 0)

	for i := 0; i < len(options); i++ {
		positions = append(positions, &Position{
			Wheel_id: Wheel_id,
			Corn: i,
			GetOptions: options[i],
		})
	}

	for i := 6; i < 8; i++ {
		positions = append(positions, &Position{
			Wheel_id: Wheel_id,
			Corn: i,
			GetOptions: Flatten(options),
		})
	}

	return &Wheel{
		Id: Wheel_id,
		Size: len(positions),
		Occupied: make(map[int]int),
		Positions: positions, 
		Name: wheel_name,
		String: func (wheel *Wheel, workers []*Worker) string {
			var br strings.Builder
		
			fmt.Fprintf(&br, "| %-12s: ", wheel.Name)
		
			out := make([]string, wheel.Size)
		
			for k, v := range wheel.Occupied {
				out[k] = workers[v].Color.String()
			}
		
			for k, o := range out {
				if len(o) > 0 {
					fmt.Fprintf(&br, "  %s", o)
				} else {
					fmt.Fprintf(&br, "%3d", k)
				}
			}
			fmt.Fprintf(&br, "\n")
		
			return br.String()
		},
	}
}
