// Copyright 2016 Huitse Tai. All rights reserved.
// Use of this source code is governed by BSD 3-clause
// license that can be found in the LICENSE file.

package jormungand

import (
	"github.com/setekhid/jormungand/tungo"
	_ "golang.org/x/net/ipv4"
	"io"
)

type Router struct {
	tun SysTun
}

type RouterConf struct {
	IfInfo tungo.IfInfo
}

func NewRouter(conf RouterConf) *Router {

	return &Router{
		tun: SysTun{
			IfInfo: conf.IfInfo,
		},
	}
}

func (this *Router) Start() error {

	var err error

	err = this.tun.Open()
	if err != nil {
		return err
	}

	return nil
}

func (this *Router) Stop() {

	this.tun.Close()
}

// Override comet.Auth2ReadWriteCloser.Auth
func (this *Router) Auth(token string, writable int64) (io.ReadWriteCloser, int64) {

	// TODO
	return nil, 0
}
