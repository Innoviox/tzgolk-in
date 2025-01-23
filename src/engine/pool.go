package engine

import (
    "sync"
)

var (
    DeltaPool = sync.Pool{
        New: func() interface{} {
            return &Delta{}
        },
    }

    PlayerDeltaMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[Color]PlayerDelta, 4)
        },
    }
    
    WorkerDeltaMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]WorkerDelta, 32)
        },
    }
    
    WheelDeltaMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]WheelDelta, 5)
        },
    }
    
    IntMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]int, 4)
        },
    }

    PositionDeltaMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]PositionDelta, 10)
        },
    }

    TempleDeltaMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]TempleDelta, 3)
        },
    }

    PlayerLocationsMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[int]int, 4)
        },
    }

    LevelsMapPool = sync.Pool{
        New: func() interface{} {
            return make(map[Color]Levels, 4)
        },
    }
)

func GetDelta() *Delta {
    d := DeltaPool.Get().(*Delta)
    // d.reset()
    return d
}

// PutDelta returns a Delta to the pool
func (d *Delta) Put() {
    if d == nil {
        return
    }
    // d.cleanup()
    d.reset()
    DeltaPool.Put(d)
}

func (d *Delta) reset() {
    d.PlayerDeltas = nil
    d.WorkerDeltas = nil
    d.CalendarDelta = CalendarDelta{}
    d.TemplesDelta = TemplesDelta{}
    d.ResearchDelta = ResearchDelta{}
    d.Monuments = nil
    d.Buildings = nil
    d.CurrPlayer = 0
    d.FirstPlayer = 0
    d.AccumulatedCorn = 0
    d.Age = 0
    d.Day = 0
    d.Over = 0
    d.Description = ""
    d.BuildingNum = 0
}


func (d *Delta) cleanup() {
}

func DeltaPrewarmPools(count int) {
    for i := 0; i < count; i++ {
        DeltaPool.Put(GetDelta())
        PlayerDeltaMapPool.Put(make(map[Color]PlayerDelta, 4))
        WorkerDeltaMapPool.Put(make(map[int]WorkerDelta, 32))
        WheelDeltaMapPool.Put(make(map[int]WheelDelta, 5))
        IntMapPool.Put(make(map[int]int, 4))
        PositionDeltaMapPool.Put(make(map[int]PositionDelta, 10))
        TempleDeltaMapPool.Put(make(map[int]TempleDelta, 3))
        PlayerLocationsMapPool.Put(make(map[Color]int, 4))
        LevelsMapPool.Put(make(map[Color]Levels, 4))
    }
}