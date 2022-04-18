package complexity

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCounter_Limit_Value(t *testing.T) {
	t.Parallel()
	counter := NewCounter(10)
	require.EqualValues(t, 10, counter.Limit())
}

func TestCounter_Limit_Exceeded(t *testing.T) {
	t.Parallel()
	counter := NewCounter(10)
	require.EqualValues(t, 10, counter.Limit())

	// no error on below limit.
	err := counter.Add(5)
	require.NoError(t, err)
	// no error on exact limit
	err = counter.Add(5)
	require.NoError(t, err)
	// error above limit
	err = counter.Add(1)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrLimitExceeded)
}

func TestCounter_Value_Empty(t *testing.T) {
	t.Parallel()
	counter := NewCounter(10)
	counter.MakeReady()
	require.EqualValues(t, 0, counter.Value())
}

func TestCounter_Value_Incremented(t *testing.T) {
	t.Parallel()
	counter := NewCounter(10)
	err := counter.Add(5)
	require.NoError(t, err)
	counter.MakeReady()
	require.EqualValues(t, 5, counter.Value())
}

func TestCounter_Value_Concurrent(t *testing.T) {
	t.Parallel()
	counter := NewCounter(10)

	// firing up 10 goroutines
	// each will try to get the counter value.
	// they should all get the same value and,
	// wait until the counter is calculated.
	valWg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		valWg.Add(1)
		go func() {
			defer valWg.Done()
			val := counter.Value()
			require.EqualValues(t, 10, val)
		}()
	}

	// firing up 10 goroutines
	// each will try to increment the counter.
	incWg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		incWg.Add(1)
		go func() {
			defer incWg.Done()
			err := counter.Add(1)
			require.NoError(t, err)
		}()
	}

	// waiting for the increment goroutines to finish
	incWg.Wait()
	// calculating the counter
	counter.MakeReady()
	// waiting for the value goroutines to finish the assertions.
	valWg.Wait()
}
