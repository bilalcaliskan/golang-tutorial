package main

import (
	"net/http"
	"golang.org/x/time/rate"
)


var limiter = rate.NewLimiter(1, 3)

func runRateLimit()  {

}