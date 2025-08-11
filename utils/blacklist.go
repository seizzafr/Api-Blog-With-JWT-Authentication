package utils

import "sync"

var tokenBlacklist = make(map[string]bool)
var mutex sync.Mutex

func AddToBlacklist(token string) {
	mutex.Lock()
	defer mutex.Unlock()
	tokenBlacklist[token] = true
}

func IsTokenBlacklisted(token string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	return tokenBlacklist[token]
}
