# goutils

***You can install every package with*** ```go get github.com/zrzf/goutils```
- [semaphore](#semaphore)   ```go get github.com/zrzf/goutils/semaphore```
- [slices](#slices)         ```go get github.com/zrzf/goutils/slices```
- [color](#color)           ```go get github.com/zrzf/goutils/color```
- [cursor](#cursor)         ```go get github.com/zrzf/goutils/cursor```
- [file](#file)             ```go get github.com/zrzf/goutils/file```
- [looper](#looper)         ```go get github.com/zrzf/goutils/looper```
- [strings](#strings)       ```go get github.com/zrzf/goutils/strings```

# semaphore
```go
s := semaphore.New(6)

for i:=0;i<20;i++{
    s.Add()
    go func(){
        fmt.Println("Working")
        time.Sleep(time.Second*5)
        s.Done()
    }()
}

s.Wait()
fmt.Println("Done")
```

# slices
```IndexInt```, ```IndexStr```
```go
slices.IndexInt([1,2,3,10], 10)
>>> 3
```
---
```ContainInt```, ```ContainStr```
```go
slices.ContainInt([1,2,3,10], 10)
>>> true
```
---
```SumInt```, ```SumUint```, ```SumFloat32```, ```SumFloat64```
```go
slices.SumInt([1,2,3,10])
>>> 16
```
---
```UniqueInt```, ```UniqueStr```, ```UniqueUint```
```go
slices.UniqueInt(&[1,2,3,3,1])
>>> [1 2 3]
```
---
```CountInt```, ```CountStr```
```go
slices.CountInt([1,2,3,10,2,2], 2)
>>> 3
```
---
```FilterInt```, ```FilterStr```
```go
slices.FilterInt(&[1,2,3,10,2], 2)
>>> [1 3 10]
```
---
```RangeInt```, ```RangeUint```
```go
slices.RangeInt(0, 5)
>>> [0 1 2 3 4 5]
```
---
```ShuffleInt```, ```ShuffleUint```, ```ShuffleStr```
```go
slices.ShuffleInt([1,2,3,4], seed)
>>> [3 1 4 2]
```
---
```ChunkInt```, ```ChunkStr```
```go
slices.ChunkInt([1,2,3,4,5], 2)
>>> [[1 2] [3 4] [5]]
```
---
```ReverseInt```, ```ReverseStr```
```go
slices.ReverseInt([1,2,3,4])
>>> [4 3 2 1]
```

# color
```go
fmt.Println(color.Red + color.Bold + "Bold red" + color.End)
fmt.Println(color.SlowBlink + color.RGB(0,0,255) + "Blinking Blue" + color.End)
fmt.Println(color.BMagenta + "Magenta Background" + color.End)
```
