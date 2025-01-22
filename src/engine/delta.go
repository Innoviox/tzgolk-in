package engine

import (
    "fmt"
    // "runtime/debug"
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
    Unlocked int
    // Wheel_id int
    // Position int
}

type CalendarDelta struct {
    WheelDeltas map[int]WheelDelta
    Rotation int

    FirstPlayer int
}

type WheelDelta struct {
    // todo how should this work?
    // OldOccupied map[int]int
    // NewOccupied map[int]int
    // Sign int
    Occupied map[int]int
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
    if o.PlayerDeltas != nil {
        if d.PlayerDeltas == nil {
            d.PlayerDeltas = map[Color]PlayerDelta{}
        }

        for k, v := range o.PlayerDeltas {
            p, ok := d.PlayerDeltas[k]
            if !ok {
                p = PlayerDelta{}
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
            if p.LightSide == 0 {
                p.LightSide = v.LightSide
            } else if v.LightSide != 0 {
                p.LightSide *= v.LightSide
            }

            p.Buildings = AddMap(p.Buildings, v.Buildings)
            p.Monuments = AddMap(p.Monuments, v.Monuments)

            d.PlayerDeltas[k] = p
        }
    }

    if o.WorkerDeltas != nil {
        if d.WorkerDeltas == nil {
            d.WorkerDeltas = map[int]WorkerDelta{}
        }
        for k, v := range o.WorkerDeltas {
            w, ok := d.WorkerDeltas[k]
            if !ok {
                w = WorkerDelta{}
            }
            if w.Unlocked == 0 {
                w.Unlocked = v.Unlocked
            } else if v.Unlocked != 0 {
                w.Unlocked *= v.Unlocked
            }
            
            d.WorkerDeltas[k] = w
        }
    }

    if o.CalendarDelta.WheelDeltas != nil {
        if d.CalendarDelta.WheelDeltas == nil {
            d.CalendarDelta.WheelDeltas = map[int]WheelDelta{}
        } 
        for k, v := range o.CalendarDelta.WheelDeltas {
            w, ok := d.CalendarDelta.WheelDeltas[k]
            if !ok {
                w = WheelDelta{}
            }
            w.Occupied = AddMap(w.Occupied, v.Occupied)
            if v.PositionDeltas != nil {
                if w.PositionDeltas == nil {
                    w.PositionDeltas = map[int]PositionDelta{}
                }

                for k2, v2 := range v.PositionDeltas {
                    p, ok := w.PositionDeltas[k2]
                    if !ok {
                        p = PositionDelta{
                            PData: PalenqueDataDelta{},
                            CData: ChichenDataDelta{},
                        }
                    }

                    p.PData.CornTiles += v2.PData.CornTiles
                    p.PData.WoodTiles += v2.PData.WoodTiles
                    p.CData.Full += v2.CData.Full
                    w.PositionDeltas[k2] = p
                }
            }
            d.CalendarDelta.WheelDeltas[k] = w
        }
    }
    d.CalendarDelta.Rotation += o.CalendarDelta.Rotation
    d.CalendarDelta.FirstPlayer += o.CalendarDelta.FirstPlayer

    if o.TemplesDelta.TempleDeltas != nil {
        if d.TemplesDelta.TempleDeltas == nil {
            d.TemplesDelta.TempleDeltas = map[int]TempleDelta{}
        }

        for k, v := range o.TemplesDelta.TempleDeltas {
            pl, ok := d.TemplesDelta.TempleDeltas[k]
            if !ok {
                pl = TempleDelta{PlayerLocations: map[Color]int{}}
            }
            pl.PlayerLocations = AddMap(pl.PlayerLocations, v.PlayerLocations)
            d.TemplesDelta.TempleDeltas[k] = pl
        }
    }

    if o.ResearchDelta.Levels != nil {
        if d.ResearchDelta.Levels == nil {
            d.ResearchDelta.Levels = map[Color]Levels{}
        }

        for k, v := range o.ResearchDelta.Levels {
            pl, ok := d.ResearchDelta.Levels[k]
            if !ok {
                pl = Levels{}
            }
            pl = AddMap(pl, v)
            d.ResearchDelta.Levels[k] = pl
        }
    }
    d.Monuments = AddMap(d.Monuments, o.Monuments)
    d.Buildings = AddMap(d.Buildings, o.Buildings)

    d.CurrPlayer += o.CurrPlayer
    d.FirstPlayer += o.FirstPlayer
    d.AccumulatedCorn += o.AccumulatedCorn
    d.Age += o.Age
    d.Day += o.Day
    d.Over += o.Over

    d.Description = fmt.Sprintf("%s; %s", d.Description, o.Description)

    if o.BuildingNum != 0 {
        d.BuildingNum = o.BuildingNum
    }
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