package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// IPRateLimiter .
type IPRateLimiter struct {
    ips map[string]*rate.Limiter
    mu  *sync.RWMutex
    r   rate.Limit
    b   int
}

func newIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
    i := &IPRateLimiter{
        ips: make(map[string]*rate.Limiter),
        mu:  &sync.RWMutex{},
        r:   r,
        b:   b,
    }

    return i
}

func (i *IPRateLimiter) addIP(ip string) *rate.Limiter {
    i.mu.Lock()
    defer i.mu.Unlock()

    limiter := rate.NewLimiter(i.r, i.b)

    i.ips[ip] = limiter

    return limiter
}

func (i *IPRateLimiter) getLimiter(ip string) *rate.Limiter {
    i.mu.Lock()
    limiter, exists := i.ips[ip]

    if !exists {
        i.mu.Unlock()
        return i.addIP(ip)
    }

    i.mu.Unlock()

    return limiter
}

var limiter = newIPRateLimiter(1, 5)


func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limiter := limiter.getLimiter(ctx.ClientIP())
        if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{})
            return
        }
	}
}
