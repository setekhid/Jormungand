// Copyright 2016 Huitse Tai. All rights reserved.
// Use of this source code is governed by BSD 3-clause
// license that can be found in the LICENSE file.

package jargs

import (
	"errors"
	"sync"
)

var (
	entries = map[string]ModuleEntry{}
)

type ModuleEntry func()

func RunInMain(modules []string) {

	waiter := sync.WaitGroup{}

	for _, m := range modules {

		entry := entries[m]

		waiter.Add(1)
		go func() {
			defer waiter.Done()
			entry()
		}()
	}

	waiter.Wait()
}

func RegistEntry(module string, entry ModuleEntry) {

	if _, ok := entries[module]; ok {
		panic(errors.New("duplicated module entry " + module))
	}

	entries[module] = entry
}
