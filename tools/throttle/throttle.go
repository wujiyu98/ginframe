package throttle

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// IPRateLimiter .
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// NewIPRateLimiter .
//限制每r分钟多少b限制
func NewIPRateLimiter(r time.Duration, b int) *IPRateLimiter {

	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   rate.Every(time.Second*60*r) * rate.Limit(b),
		b:   b,
	}

	return i
}

// AddIP 创建了一个新的速率限制器，并将其添加到 ips 映射中,
// 使用 IP地址作为密钥
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()
	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter

	return limiter
}

// GetLimiter 返回所提供的IP地址的速率限制器(如果存在的话).
// 否则调用 AddIP 将 IP 地址添加到映射中
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	i.clearIp()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}
	i.mu.Unlock()
	return limiter
}

//为防止保存IP过多，定量删除
func (i *IPRateLimiter) clearIp() {
	if len(i.ips) > 100000 {
		fmt.Println("clear")
		for k := range i.ips {
			delete(i.ips, k)
		}
	}
}
