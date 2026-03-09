package xiter

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestBuffer(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		now := time.Now()

		producerReadAt := map[int]time.Time{}
		fastProducer := func(yield func(int) bool) {
			for i := 1; i <= 5; i++ {
				c := yield(i)
				producerReadAt[i] = time.Now()
				if !c {
					return
				}
			}
		}

		var result []int
		slowConsumer := func(v int) bool {
			time.Sleep(1 * time.Second) // Simulate slow consumer
			result = append(result, v)
			return true
		}

		Buffer(fastProducer, 2)(slowConsumer)

		// Verify that the values are collected in order
		expected := []struct {
			value       int
			collectedAt time.Time
		}{
			{value: 1, collectedAt: now},
			{value: 2, collectedAt: now},
			{value: 3, collectedAt: now},
			{value: 4, collectedAt: now.Add(1 * time.Second)},
			{value: 5, collectedAt: now.Add(2 * time.Second)},
		}

		if len(result) != len(expected) {
			t.Fatalf("expected %d results, got %d", len(expected), len(result))
		}

		for i := range expected {
			if result[i] != expected[i].value {
				t.Errorf("[%d] expected value %d, got %d", i, expected[i].value, result[i])
			}

			if !expected[i].collectedAt.Equal(producerReadAt[expected[i].value]) {
				t.Errorf("[%d] expected collectedAt %s, got %s", i, expected[i].collectedAt, producerReadAt[expected[i].value])
			}
		}

		synctest.Wait()
	})
}

func TestBufferWithBreak(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		producerStoppedAt := -1
		producer := func(yield func(int) bool) {
			for i := 1; i <= 20; i++ {
				if !yield(i) {
					producerStoppedAt = i
					return
				}
			}
		}

		var result []int
		consumer := func(v int) bool {
			time.Sleep(1 * time.Second) // Simulate slow consumer
			result = append(result, v)
			return v < 3 // Break after consuming 3
		}

		Buffer(producer, 5)(consumer)

		expected := []int{1, 2, 3}
		if len(result) != len(expected) {
			t.Fatalf("expected %d results, got %d", len(expected), len(result))
		}

		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("[%d] expected value %d, got %d", i, expected[i], result[i])
			}
		}

		// wait for producer goroutine to finish
		synctest.Wait()

		// Buffer size is 5,
		// 1. Producer sends 5 values (1-5) to the buffer
		// 2. Consumer consumes 3 values (1-3)
		// 3. Producer sends another 5 values (4-8), to keep buffer full
		// 4. Producer sends another 1 value  (9  ), but then finds out consumer has broken and exits.
		// Setting 9 as the last value
		if producerStoppedAt != 9 {
			t.Errorf("expected producer to stop at 9, got %d", producerStoppedAt)
		}
	})
}
