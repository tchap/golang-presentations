// map[T]T - maps // HL
var m map[string]uint
m = make(map[string]uint)

m := map[string]uint{"foo": 1, "bar": 2} // literal + type inference

// [N]T - arrays (not used very often) // HL
var a [2]int
a = [...]int{1, 2} // literal

// []T - slices // HL
s := make([]int, 0, 10) // make([]int, length, capacity)
s := make([]int, 10)    // ~ make([]int, 10, 10)

// chan T - channels // HL
ch := make(chan *Person)  // unbuffered channel
ch := make(chan bool, 10) // buffered channel
