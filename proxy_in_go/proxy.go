package main

import "errors"

//Proxy server include pakages and limiter
type ProxyServer struct{ 
	packages *SoftwarePackeges
	maxAllowedRequest int 
	limiter map[string]int
}

// add ProxyServer to new values
func newProxyServer() *ProxyServer {
    return &ProxyServer{
        packages:       &SoftwarePackeges{},
        maxAllowedRequest: 2,
        limiter:       make(map[string]int),
    }
}

func (n *ProxyServer) checkRateLimiting(packageName string) bool {
    if n.limiter[packageName] == 0 {
        n.limiter[packageName] = 1
    }
	// check limit 
    if n.limiter[packageName] > n.maxAllowedRequest {
        return false
    }
	// every request add user requested package until exceed the limiter
    n.limiter[packageName] = n.limiter[packageName] + 1
    return true
}

// check r
func (n *ProxyServer) clientRequest(packageName string) (string, error) {
    allowed := n.checkRateLimiting(packageName)
	// checking limit going to checkRateLimiting function.
    if !allowed {
        return "",errors.New("all of the licenses have all been taken up and retry later")
    }
    return n.packages.clientRequest(packageName),nil
}