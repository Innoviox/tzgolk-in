package engine

import (
    "fmt"
)

// everything represents a delta
// booleans are ints; positive means true, negative means false

type Delta struct {
    PlayerDeltas map[Color]PlayerDelta
    WorkerDeltas map[int]WorkerDelta

    CalendarDelta CalendarDelta
    TemplesDelta TemplesDelta
    ResearchDelta ResearchDelta

    Monuments map[int]int
    Buildings map[int]int

    CurrPlayer int
    FirstPlayer int
    AccumulatedCorn int
    Age int
    Day int

    Over int

    Description string
    BuildingNum int // todo do we need this?
}

type PlayerDelta struct {
    Resources [4]int
    Corn int
    Points int
    CornTiles int
    WoodTiles int
    FreeWorkers int
    WorkerDeduction int

    LightSide int
    // 0 => unbuilt
    // 1 => built
    Buildings map[int]int
    Monuments map[int]int
}

type WorkerDelta struct {
    // all args required
    Available int
    Wheel_id int
    Position int
}

type CalendarDelta struct {
    WheelDeltas map[int]WheelDelta
    Rotation int

    FirstPlayer int
}

type WheelDelta struct {
    // todo how should this work?
    OldOccupied map[int]int
    NewOccupied map[int]int
    PositionDeltas map[int]PositionDelta
}

type PositionDelta struct {
    PData PalenqueDataDelta
    CData ChichenDataDelta
}

type PalenqueDataDelta struct {
    CornTiles int
    WoodTiles int
}

type ChichenDataDelta struct {
    Full int
}

// DO NOT INSTANTIATE TEMPLESDELTA, USE TEMPLES.STEP INSTEAD
type TemplesDelta struct {
    TempleDeltas map[int]TempleDelta
}

type TempleDelta struct {
    PlayerLocations map[Color]int
}

type ResearchDelta struct {
    Levels map[Color]Levels
}

func Bool(d int, m int, current bool) bool {
    if d * m > 0 {
        return true
    } else if d * m < 0 {
        return false
    }
    return current
}

// todo MarkDelta function or something
func (d *Delta) Add(o *Delta) {
    // fmt.Printf("%v + %v\n", d, o)
    // fmt.Print("a")
    if o.PlayerDeltas != nil {
        if d.PlayerDeltas == nil {
            d.PlayerDeltas = o.PlayerDeltas
        } else {
            for k, v := range o.PlayerDeltas {
                p, ok := d.PlayerDeltas[k]
                if !ok {
                    d.PlayerDeltas[k] = v
                    continue
                }
                p.Resources[0] += v.Resources[0]
                p.Resources[1] += v.Resources[1]
                p.Resources[2] += v.Resources[2]
                p.Resources[3] += v.Resources[3]
                p.Corn += v.Corn
                p.Points += v.Points
                p.CornTiles += v.CornTiles
                p.WoodTiles += v.WoodTiles
                p.FreeWorkers += v.FreeWorkers
                p.WorkerDeduction += v.WorkerDeduction
                p.LightSide += v.LightSide

                if p.Buildings == nil {
                    p.Buildings = v.Buildings
                } else {
                    for k2, v2 := range v.Buildings {
                        p.Buildings[k2] += v2
                    }
                }

                if p.Monuments == nil {
                    p.Monuments = v.Monuments
                } else {
                    for k2, v2 := range v.Monuments {
                        p.Monuments[k2] += v2
                    }
                }
            }
        }
    }
    // fmt.Print("b")

    if o.WorkerDeltas != nil {
        if d.WorkerDeltas == nil {
            d.WorkerDeltas = o.WorkerDeltas
        } else {
            for k, v := range o.WorkerDeltas {
                w, ok := d.WorkerDeltas[k]
                if !ok {
                    d.WorkerDeltas[k] = v
                    continue
                }
                w.Available += v.Available
                w.Wheel_id += v.Wheel_id
                w.Position += v.Position
            }
        }
    }
    // fmt.Print("c")


    if o.CalendarDelta.WheelDeltas != nil {
        if d.CalendarDelta.WheelDeltas == nil {
            d.CalendarDelta.WheelDeltas = o.CalendarDelta.WheelDeltas
        } else {
            for k, v := range o.CalendarDelta.WheelDeltas {
                w, ok := d.CalendarDelta.WheelDeltas[k]
                if !ok {
                    d.CalendarDelta.WheelDeltas[k] = v
                    continue
                }

                if len(v.OldOccupied) > 0 {
                    w.OldOccupied = v.OldOccupied
                }

                if len(v.NewOccupied) > 0 {
                    w.NewOccupied = v.NewOccupied
                }

                if v.PositionDeltas != nil {
                    for k2, v2 := range v.PositionDeltas {
                        p := w.PositionDeltas[k2]

                        p.PData.CornTiles += v2.PData.CornTiles
                        p.PData.WoodTiles += v2.PData.WoodTiles
                        p.CData.Full += v2.CData.Full
                    }
                }
            }
        }
    }
    // fmt.Print("d")
    d.CalendarDelta.Rotation += o.CalendarDelta.Rotation
    d.CalendarDelta.FirstPlayer += o.CalendarDelta.FirstPlayer

    if o.TemplesDelta.TempleDeltas != nil {
        if d.TemplesDelta.TempleDeltas == nil {
            d.TemplesDelta.TempleDeltas = o.TemplesDelta.TempleDeltas
        } else {
            for k, v := range o.TemplesDelta.TempleDeltas {
                pl := d.TemplesDelta.TempleDeltas[k]
                if pl.PlayerLocations == nil {
                    pl.PlayerLocations = v.PlayerLocations
                } else {
                    for k2, v2 := range v.PlayerLocations {
                        pl.PlayerLocations[k2] += v2
                    }
                }
            }
        }
    }
    // fmt.Print("e")

    // if o.ResearchDelta.Levels != nil {
    //     if d.ResearchDelta.Levels == nil {
    //         d.ResearchDelta.Levels = o.ResearchDelta.Levels
    //     } else {
    //         for k, v := range o.ResearchDelta.Levels {
    //             for k2, v2 := range v {
    //                 d.ResearchDelta.Levels[k][k2] += v2
    //             }
    //         }
    //     }
    //     // fmt.Printf("%v %v\n", o.ResearchDelta, d.ResearchDelta)
    // }
    // fmt.Print("f")

    if o.Monuments != nil {
        if d.Monuments == nil {
            d.Monuments = o.Monuments
        } else {
            for k, v := range o.Monuments {
                d.Monuments[k] += v
            }
        }
    }
    // fmt.Print("g")

    if o.Buildings != nil {
        if d.Buildings == nil {
            d.Buildings = o.Buildings
        } else {
            for k, v := range o.Buildings {
                d.Buildings[k] += v
            }
        }
    }
    // fmt.Print("h")

    d.CurrPlayer += o.CurrPlayer
    d.FirstPlayer += o.FirstPlayer
    d.AccumulatedCorn += o.AccumulatedCorn
    d.Age += o.Age
    d.Day += o.Day
    d.Over += d.Over

    d.Description = fmt.Sprintf("%s; %s", d.Description, o.Description)

    if o.BuildingNum != 0 {
        d.BuildingNum = o.BuildingNum
    }
    // fmt.Print("i!\n")
    // fmt.Printf("%v\n", d)
}

func Combine(d1 *Delta, d2 *Delta) *Delta {
    d := &Delta{}
    d.Add(d1)
    d.Add(d2)
    return d
}


func ResourcesDelta(color Color, old [4]int, new [4]int) *Delta {
    return &Delta{PlayerDeltas: map[Color]PlayerDelta{color: PlayerDelta{Resources: 
        [4]int{new[0] - old[0], new[1] - old[1], new[2] - old[2], new[3] - old[3]},
    }}}
}

func PlayerDeltaWrapper(color Color, pd PlayerDelta) *Delta {
    return &Delta{PlayerDeltas: map[Color]PlayerDelta{color: pd}}
}