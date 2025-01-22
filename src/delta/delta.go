package engine

import (
    // "fmt"
    "strings"
    "sync"
)

// Pool sizes and capacities
const (
    defaultMapSize = 4
    defaultDescriptionSize = 64
)

type Delta struct {
    PlayerDeltas map[int]PlayerDelta
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
    BuildingNum int
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
    PlayerLocations map[int]int
}

type ResearchDelta struct {
    Levels map[int]map[int]int

}

func Bool(d int, m int, current bool) bool {
    if d * m > 0 {
        return true
    } else if d * m < 0 {
        return false
    }
    return current
}

// Object pools
var (
    DeltaPool = sync.Pool{
        New: func() interface{} {
            return &Delta{}
        },
    }
    
    PlayerDeltaMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]PlayerDelta, defaultMapSize)
        },
    }
    
    WorkerDeltaMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]WorkerDelta, defaultMapSize)
        },
    }
    
    WheelDeltaMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]WheelDelta, defaultMapSize)
        },
    }
    
    IntMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]int, defaultMapSize)
        },
    }

    BuilderPool = sync.Pool{
        New: func() interface{} {
            return &strings.Builder{}
        },
    }

    PositionDeltaMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]PositionDelta, defaultMapSize)
        },
    }

    TempleDeltaMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]TempleDelta, defaultMapSize)
        },
    }

    PlayerLocationsMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]int, defaultMapSize)
        },
    }

    LevelsMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]map[int]int, defaultMapSize)
        },
    }
)

// GetDelta gets a Delta from the pool
func GetDelta() *Delta {
    d := DeltaPool.Get().(*Delta)
    d.reset()
    return d
}

// PutDelta returns a Delta to the pool
func PutDelta(d *Delta) {
    if d == nil {
        return
    }
    d.cleanup()
    DeltaPool.Put(d)
}

// reset initializes/resets a Delta for reuse
func (d *Delta) reset() {
    *d = Delta{}
}

// cleanup releases pooled resources
func (d *Delta) cleanup() {
    if d.PlayerDeltas != nil {
        for k := range d.PlayerDeltas {
            delete(d.PlayerDeltas, k)
        }
        PlayerDeltaMapPool.Put(d.PlayerDeltas)
    }
    if d.WorkerDeltas != nil {
        for k := range d.WorkerDeltas {
            delete(d.WorkerDeltas, k)
        }
        WorkerDeltaMapPool.Put(d.WorkerDeltas)
    }
    
    // Clean up WheelDeltas
    if d.CalendarDelta.WheelDeltas != nil {
        for _, wheel := range d.CalendarDelta.WheelDeltas {
            if wheel.Occupied != nil {
                for k := range wheel.Occupied {
                    delete(wheel.Occupied, k)
                }
                IntMapPool.Put(wheel.Occupied)
            }
            if wheel.PositionDeltas != nil {
                for k := range wheel.PositionDeltas {
                    delete(wheel.PositionDeltas, k)
                }
                PositionDeltaMapPool.Put(wheel.PositionDeltas)
            }
        }
        for k := range d.CalendarDelta.WheelDeltas {
            delete(d.CalendarDelta.WheelDeltas, k)
        }
        WheelDeltaMapPool.Put(d.CalendarDelta.WheelDeltas)
    }

    // Clean up TempleDeltas
    if d.TemplesDelta.TempleDeltas != nil {
        for _, temple := range d.TemplesDelta.TempleDeltas {
            if temple.PlayerLocations != nil {
                for k := range temple.PlayerLocations {
                    delete(temple.PlayerLocations, k)
                }
                PlayerLocationsMapPool.Put(temple.PlayerLocations)
            }
        }
        for k := range d.TemplesDelta.TempleDeltas {
            delete(d.TemplesDelta.TempleDeltas, k)
        }
        TempleDeltaMapPool.Put(d.TemplesDelta.TempleDeltas)
    }

    // Clean up ResearchDelta
    if d.ResearchDelta.Levels != nil {
        for k := range d.ResearchDelta.Levels {
            delete(d.ResearchDelta.Levels, k)
        }
        LevelsMapPool.Put(d.ResearchDelta.Levels)
    }

    // Clean up Monuments and Buildings
    if d.Monuments != nil {
        for k := range d.Monuments {
            delete(d.Monuments, k)
        }
        IntMapPool.Put(d.Monuments)
    }
    
    if d.Buildings != nil {
        for k := range d.Buildings {
            delete(d.Buildings, k)
        }
        IntMapPool.Put(d.Buildings)
    }
}

func (d *Delta) Add(o *Delta) {
    if o == nil {
        return
    }

    if o.PlayerDeltas != nil {
        if d.PlayerDeltas == nil {
            d.PlayerDeltas = PlayerDeltaMapPool.Get().(map[int]PlayerDelta)
        }
        
        for k, v := range o.PlayerDeltas {
            p, exists := d.PlayerDeltas[k]
            if !exists {
                p = PlayerDelta{}
            }
            
            // Optimize array operations
            for i := 0; i < 4; i++ {
                p.Resources[i] += v.Resources[i]
            }
            
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

            // Use pooled maps for Buildings and Monuments
            if p.Buildings == nil {
                p.Buildings = IntMapPool.Get().(map[int]int)
            }
            if p.Monuments == nil {
                p.Monuments = IntMapPool.Get().(map[int]int)
            }
            
            AddMapPooled(p.Buildings, v.Buildings)
            AddMapPooled(p.Monuments, v.Monuments)
            
            d.PlayerDeltas[k] = p
        }
    }

    // Handle WorkerDeltas
    if o.WorkerDeltas != nil {
        if d.WorkerDeltas == nil {
            d.WorkerDeltas = WorkerDeltaMapPool.Get().(map[int]WorkerDelta)
        }
        for k, v := range o.WorkerDeltas {
            w, exists := d.WorkerDeltas[k]
            if !exists {
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

    // Handle CalendarDelta.WheelDeltas
    if o.CalendarDelta.WheelDeltas != nil {
        if d.CalendarDelta.WheelDeltas == nil {
            d.CalendarDelta.WheelDeltas = WheelDeltaMapPool.Get().(map[int]WheelDelta)
        }
        for k, v := range o.CalendarDelta.WheelDeltas {
            w, exists := d.CalendarDelta.WheelDeltas[k]
            if !exists {
                w = WheelDelta{}
            }

            // Handle Occupied map
            if v.Occupied != nil {
                if w.Occupied == nil {
                    w.Occupied = IntMapPool.Get().(map[int]int)
                }
                AddMapPooled(w.Occupied, v.Occupied)
            }

            // Handle PositionDeltas map
            if v.PositionDeltas != nil {
                if w.PositionDeltas == nil {
                    w.PositionDeltas = PositionDeltaMapPool.Get().(map[int]PositionDelta)
                }

                for k2, v2 := range v.PositionDeltas {
                    p, exists := w.PositionDeltas[k2]
                    if !exists {
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

    // Handle TemplesDelta
    if o.TemplesDelta.TempleDeltas != nil {
        if d.TemplesDelta.TempleDeltas == nil {
            d.TemplesDelta.TempleDeltas = TempleDeltaMapPool.Get().(map[int]TempleDelta)
        }

        for k, v := range o.TemplesDelta.TempleDeltas {
            pl, exists := d.TemplesDelta.TempleDeltas[k]
            if !exists {
                pl = TempleDelta{
                    PlayerLocations: PlayerLocationsMapPool.Get().(map[int]int),
                }
            }
            if pl.PlayerLocations == nil {
                pl.PlayerLocations = PlayerLocationsMapPool.Get().(map[int]int)
            }
            AddMapPooled(pl.PlayerLocations, v.PlayerLocations)
            d.TemplesDelta.TempleDeltas[k] = pl
        }
    }

    // Handle ResearchDelta
    if o.ResearchDelta.Levels != nil {
        if d.ResearchDelta.Levels == nil {
            d.ResearchDelta.Levels = LevelsMapPool.Get().(map[int]map[int]int)
        }

        for k, v := range o.ResearchDelta.Levels {
            pl, exists := d.ResearchDelta.Levels[k]
            if !exists {
                pl = map[int]int{}
            }
            AddMapPooled(pl, v)
            d.ResearchDelta.Levels[k] = pl
        }
    }

    // Handle Monuments and Buildings
    if o.Monuments != nil {
        if d.Monuments == nil {
            d.Monuments = IntMapPool.Get().(map[int]int)
        }
        AddMapPooled(d.Monuments, o.Monuments)
    }

    if o.Buildings != nil {
        if d.Buildings == nil {
            d.Buildings = IntMapPool.Get().(map[int]int)
        }
        AddMapPooled(d.Buildings, o.Buildings)
    }

    // Optimize string concatenation
    if o.Description != "" {
        builder := BuilderPool.Get().(*strings.Builder)
        builder.Reset()
        builder.Grow(len(d.Description) + len(o.Description) + 2)
        builder.WriteString(d.Description)
        if d.Description != "" {
            builder.WriteString("; ")
        }
        builder.WriteString(o.Description)
        d.Description = builder.String()
        BuilderPool.Put(builder)
    }

    d.CurrPlayer += o.CurrPlayer
    d.FirstPlayer += o.FirstPlayer
    d.AccumulatedCorn += o.AccumulatedCorn
    d.Age += o.Age
    d.Day += o.Day
    d.Over += o.Over

    if o.BuildingNum != 0 {
        d.BuildingNum = o.BuildingNum
    }
}

// AddMapPooled is an optimized version of AddMap that uses object pooling
func AddMapPooled(m1, m2 map[int]int) map[int]int {
    if m2 == nil {
        return m1
    }
    if m1 == nil {
        m1 = IntMapPool.Get().(map[int]int)
    }
    
    for k, v := range m2 {
        m1[k] += v
    }
    return m1
}

func Combine(d1 *Delta, d2 *Delta) *Delta {
    d := GetDelta()
    d.Add(d1)
    d.Add(d2)
    return d
}

func ResourcesDelta(color int, old [4]int, new [4]int) *Delta {
    d := GetDelta()
    d.PlayerDeltas = PlayerDeltaMapPool.Get().(map[int]PlayerDelta)
    d.PlayerDeltas[color] = PlayerDelta{
        Resources: [4]int{
            new[0] - old[0],
            new[1] - old[1],
            new[2] - old[2],
            new[3] - old[3],
        },
    }
    return d
}

func PlayerDeltaWrapper(color int, pd PlayerDelta) *Delta {
    d := GetDelta()
    d.PlayerDeltas = PlayerDeltaMapPool.Get().(map[int]PlayerDelta)
    d.PlayerDeltas[color] = pd
    return d
}

// PrewarmPools pre-allocates commonly used objects
func DeltaPrewarmPools(count int) {
    for i := 0; i < count; i++ {
        DeltaPool.Put(GetDelta())
        PlayerDeltaMapPool.Put(make(map[int]PlayerDelta, defaultMapSize))
        WorkerDeltaMapPool.Put(make(map[int]WorkerDelta, defaultMapSize))
        WheelDeltaMapPool.Put(make(map[int]WheelDelta, defaultMapSize))
        IntMapPool.Put(make(map[int]int, defaultMapSize))
        PositionDeltaMapPool.Put(make(map[int]PositionDelta, defaultMapSize))
        TempleDeltaMapPool.Put(make(map[int]TempleDelta, defaultMapSize))
        PlayerLocationsMapPool.Put(make(map[int]int, defaultMapSize))
        LevelsMapPool.Put(make(map[int]map[int]int, defaultMapSize))
        BuilderPool.Put(&strings.Builder{})
    }
}