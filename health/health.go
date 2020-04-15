package health 

import (
    "net/http"
    "sync"
)

var (
    healthzStatus = http.StatusOK
    readinessStatus = http.StatusOK
    mu   sync.RWMutex
)

func HealthzStatus() int {
    mu.RLock()
    defer mu.RUnlock()
    return healthzStatus
}

func ReadinessStatus() int {
    mu.RLock()
    defer mu.RUnlock()
    return readinessStatus
}

func SetHealthzStatus(status int) {
    mu.Lock()
    healthzStatus = status
    mu.Unlock()
}

func SetReadinessStatus(status int) {
    mu.Lock()
    readinessStatus = status
    mu.Unlock()
}

// HealthzHandler responds to health check requests
func HealthzHandler(response http.ResponseWriter, request *http.Request) {
    response.WriteHeader(HealthzStatus())
}

// ReadinessHandler responds to readiness check requests
func ReadinessHandler(response http.ResponseWriter, request *http.Request) {
    response.WriteHeader(ReadinessStatus())
}

func ReadinessStatusHandler(response http.ResponseWriter, request *http.Request) {

    switch ReadinessStatus() {
        case http.StatusOK:
               SetReadinessStatus(http.StatusServiceUnavailable)
        case http.StatusServiceUnavailable:
               SetReadinessStatus(http.StatusOK)
    }

    response.WriteHeader(http.StatusOK)
}


func HealthzStatusHandler(response http.ResponseWriter, request *http.Request) {
    switch HealthzStatus() {
        case http.StatusOK:
            SetHealthzStatus(http.StatusServiceUnavailable)
        case http.StatusServiceUnavailable:
            SetHealthzStatus(http.StatusOK)
    }
    response.WriteHeader(http.StatusOK)
}