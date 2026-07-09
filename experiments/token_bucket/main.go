package main

import (
	"fmt"
	"math"
	"time"
)

type TokenBucket struct {
	Tokens         float64
	Capacity       float64
	RefillRate     float64
	LastRefillTime time.Time
}

func NewTokenBucket(capacity, refillRate float64) *TokenBucket {
	return &TokenBucket{
		Tokens:         capacity,
		Capacity:       capacity,
		RefillRate:     refillRate,
		LastRefillTime: time.Now(),
	}
}

func (tb *TokenBucket) Allow() bool {

	tb.refill()

	if tb.Tokens < 1 {
		return false
	}

	tb.Tokens--

	return true
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.LastRefillTime).Seconds()
	generated := elapsed * tb.RefillRate

	tb.Tokens += generated

	tb.Tokens = math.Min(tb.Capacity, tb.Tokens)

	tb.LastRefillTime = now
}

func main() {
	bucket := NewTokenBucket(5, 1)

	fmt.Println("Initial Bucket:")
	fmt.Printf("%+v\n\n", bucket)

	for i := 1; i <= 6; i++ {
		fmt.Printf(
			"Request %d -> Allowed: %v | Tokens Left: %.2f\n",
			i,
			bucket.Allow(),
			bucket.Tokens,
		)
	}

	fmt.Println("\nWaiting 3 seconds...")

	time.Sleep(3 * time.Second)

	for i := 1; i <= 4; i++ {
		fmt.Printf(
			"Request %d -> Allowed: %v | Tokens Left: %.2f\n",
			i,
			bucket.Allow(),
			bucket.Tokens,
		)
	}
}
