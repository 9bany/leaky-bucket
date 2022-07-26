package leaky

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func TestNewCollector(t *testing.T) {
// 	c := NewCollector(true)

// 	if c.buckets == nil {
// 		t.Fatal("Didn't initialize priority?!")
// 	}
// 	if c.heap == nil {
// 		t.Fatal("Didn't initialize priority?!")
// 	}

// 	c.Free()
// }

// var collectorSimple = testSet{
// 	capacity: int64(5),
// 	rate:     1.0,
// 	set: []actionSet{
// 		{},
// 		{1, "add", 1},
// 		{1, "time-set", time.Nanosecond},
// 		{1, "till", time.Second - time.Nanosecond},
// 		{1, "time-set", time.Second - time.Nanosecond},
// 		{1, "till", time.Nanosecond},
// 		{0, "time-set", time.Second},
// 		{0, "till", time.Duration(0)},
// 		{1, "add", 1},
// 		{1, "time-add", time.Second / 2},
// 		{1, "till", time.Second / 2},
// 		{2, "add", 1},
// 		{2, "time-add", time.Second/2 - time.Nanosecond},
// 		{0, "time-add", time.Second * time.Duration(5)},
// 		{1, "add", 1},
// 		{2, "add", 1},
// 		{3, "add", 1},
// 		{4, "add", 1},
// 		{5, "add", 1},
// 		{5, "add", 1},
// 		{5, "till", time.Second * 5},
// 	},
// }

// var collectorVaried = testSet{
// 	capacity: 1000,
// 	rate:     60.0,
// 	set: []actionSet{
// 		{},
// 		{100, "add", 100},
// 		{100, "time-set", time.Nanosecond},
// 		{1000, "add", 1000},
// 		{1000, "add", 1},
// 		{940, "time-set", time.Second},
// 	},
// }

// func runKey(t *testing.T, c *Collector, key string, test *testSet) {
// 	setElapsed(0)
// 	// capacity := c.Capacity()

// 	for i, v := range test.set {
// 		switch v.action {
// 		case "add":
// 			count := c.Count(key)
// 			remaining := test.capacity - count
// 			amount := int64(v.value.(int))
// 			n := c.Add(key, amount)
// 			if n < amount {
// 				// The bucket should be full now.
// 				if n < remaining {
// 					t.Fatalf("Test %d: Bucket should have been filled by this Add()?!", i)
// 				}
// 			}
// 		case "time-set":
// 			setElapsed(v.value.(time.Duration))
// 		case "time-add":
// 			addToElapsed(v.value.(time.Duration))
// 		case "till":
// 			dt := c.TillEmpty(key)
// 			if dt != v.value.(time.Duration) {
// 				t.Fatalf("%s -> Test %d: Bad TillEmpty(). Expected %v, got %v", key, i, v.value, dt)
// 			}
// 		}
// 		count := c.Count(key)
// 		if count != v.count {
// 			t.Fatalf("%s -> Test %d: Bad count. Expected %d, got %d", key, i, v.count, count)
// 		}
// 		if count > capacity {
// 			t.Fatalf("%s -> Test %d: Went over capacity?!", key, i)
// 		}
// 		if c.Remaining(key) != test.capacity-v.count {
// 			t.Fatalf("Test %d: Expected remaining value %d, got %d",
// 				i, test.capacity-v.count, c.Remaining(key))
// 		}
// 	}
// }

// func TestCollector(t *testing.T) {
// 	setElapsed(0)

// 	tests := []testSet{
// 		collectorSimple,
// 		collectorSimple,
// 		collectorVaried,
// 	}

// 	for i, test := range tests {
// 		fmt.Println("Running testSet:", i)

// 		key := "127.0.0.1"

// 		c := NewCollector(test.rate, test.capacity, false)

// 		// Run and test Remove()
// 		runKey(t, c, key, &test)
// 		c.Remove(key)
// 		if c.Count(key) > 0 {
// 			t.Fatal("Key still has a count after removal?!")
// 		}

// 		// Run again and test Prune()
// 		runKey(t, c, "127.0.0.1", &test)
// 		c.Prune()
// 		setElapsed(time.Hour)
// 		c.Prune()

// 		// Run again and test Reset().
// 		runKey(t, c, "127.0.0.1", &test)
// 		c.Reset()
// 		if c.Count(key) != 0 {
// 			t.Fatal("Key still has a count after removal?!")
// 		}
// 		if c.TillEmpty(key) != 0 {
// 			t.Fatal("Key still has time till empty?!")
// 		}

// 		// Try to remove a non-exist bucket.
// 		c.Remove("fake")
// 		if c.Count("fake") != 0 {
// 			t.Fatal("Key still has a count after removal?!")
// 		}
// 	}
// }

// func TestPeriodicPrune(t *testing.T) {
// 	setElapsed(0)
// 	key := "localhost"
// 	c := NewCollector(1e7, 8, false)
// 	c.PeriodicPrune(time.Microsecond)
// 	n := c.Add(key, 100)
// 	if n != 8 {
// 		t.Fatal("Didn't fill bucket?!")
// 	}

// 	fmt.Printf("TillEmpty(): %v\n", c.TillEmpty(key))

// 	// Wait for the periodic prune.
// 	wait := time.Millisecond
// 	time.Sleep(wait)
// 	setElapsed(wait)

// 	count := c.Count(key)
// 	if count != 0 {
// 		t.Fatalf("Key's bucket is not empty: %d?!", count)
// 	}

// 	c.Free()
// }
