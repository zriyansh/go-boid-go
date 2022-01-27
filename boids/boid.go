package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	postion  Vector2d
	velocity Vector2d
	id       int
}

func (b *Boid) calcAcceleration() Vector2d {
	upper, lower := b.postion.AddV(viewRadius), b.postion.AddV(-viewRadius)
	avgVelocity := Vector2d{0, 0}
	avgPosition := Vector2d{0, 0}
	separation := Vector2d{0, 0}
	count := 0.0
	rWlock.RLock()
	// these 2 loops simply iterate to every single viewBox we have
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].postion.Distance(b.postion); dist < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].postion)                                    // sum of accumulation of all other boids
					separation = separation.Add(b.postion.Substract(boids[otherBoidId].postion).DivisionV(dist)) // gives resultant accel to move away from boids in a group
				}
			}
		}
	}
	rWlock.RUnlock()

	// accel := Vector2d{0, 0}
	accel := Vector2d{b.borderBounce(b.postion.x, screenWidth), b.borderBounce(b.postion.y, screenHeight)}
	if count > 0 {
		avgVelocity = avgVelocity.DivisionV(count)
		avgPosition = avgPosition.DivisionV(count)
		accelAlignment := avgVelocity.Substract(b.velocity).MultiplyV(adjRate)
		accelCohesion := avgPosition.Substract(b.postion).MultiplyV(adjRate)
		accelSeparation := separation.MultiplyV(adjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	}

	return accel
}

// pos->either on X or Y axis
// maxBorderPos -> how far it is from border axis
func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

// making boids move acc to velocity
// next determined where the boid is, inside frame or passed it
func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration() // if this call is inside lock, it wont be able to update itself
	rWlock.Lock()
	b.velocity = b.velocity.Add(acceleration).limit(-1, 1) // lower and upper inside (-1, 1)
	boidMap[int(b.postion.x)][int(b.postion.y)] = -1

	b.postion = b.postion.Add(b.velocity)
	boidMap[int(b.postion.x)][int(b.postion.y)] = b.id

	// next := b.postion.Add(b.velocity)
	// if next.x >= screenWidth || next.x < 0 {
	// 	b.velocity = Vector2d{-b.velocity.x, b.velocity.y} // changes the velocity, creates a bouncy effect upon striking the edge for x  axis
	// }
	// if next.y >= screenHeight || next.y < 0 {
	// 	b.velocity = Vector2d{b.velocity.x, -b.velocity.y} // changes the velocity, creates a bouncy effect upon striking the edge for y axis
	// } // this bouncy effect does not look cool, so we implement a more realsitic one
	rWlock.Unlock()
}

// Lock() is also Writers Lock(), no WLock() is made for it.

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int) {
	b := Boid{
		postion:  Vector2d{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2d{rand.Float64()*2 - 1.0, rand.Float64()*2 - 1.0},
		id:       bid,
	}
	boids[bid] = &b
	boidMap[int(b.postion.x)][int(b.postion.y)] = b.id
	go b.start()
}

// rand returns random number between 0 and 1
// value of velocity must be less than 1 pixel
// boids is the array
