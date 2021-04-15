package qbuilder

import "sync"

type ConditionsCache struct {
	*sync.Mutex
}

type Condition struct{}

