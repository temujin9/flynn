package main

import (
	"net/http"
// 	"time"

	c "github.com/flynn/go-check"
)

type RouterSuite struct {
	Helper
}

var _ = c.ConcurrentSuite(&RouterSuite{})

func (s *RouterSuite) TestAdditionalHttpPorts(t *c.C) {
	// boot 1 node cluster
	x := s.bootCluster(t, 1)
	defer x.Destroy()

	// Test that setting added HTTP and HTTPS ports succeeds
	t.Assert(x.flynn("/", "-a", "router", "env", "set", "ADDITIONAL_HTTP_PORTS=8080"), Succeeds)
// 	time.Sleep(1 * time.Second)
	t.Assert(x.flynn("/", "-a", "router", "env", "set", "ADDITIONAL_HTTP_PORTS=8080"), Succeeds)
// 	time.Sleep(1 * time.Second)

	// check a non-routed HTTP request to an additional port fails
	req, err := http.NewRequest("GET", "http://dashboard."+x.Domain+":8080", nil)
	t.Assert(err, c.IsNil)
	req.SetBasicAuth("", x.Key)
	res, err := http.DefaultClient.Do(req)
	t.Assert(err, c.IsNil)
	t.Assert(res.StatusCode, c.Equals, http.StatusNotFound)

	// add a controller route on the new port
	t.Assert(x.flynn("/", "-a", "dashboard", "route", "add", "http", "-s", "dashboard-web", "-p", "8080", "dashboard."+x.Domain), Succeeds)
// 	time.Sleep(1 * time.Second)

	// check a routed HTTP request succeeds
	req, err = http.NewRequest("GET", "http://dashboard."+x.Domain+":8080", nil)
	t.Assert(err, c.IsNil)
	req.SetBasicAuth("", x.Key)
	res, err = http.DefaultClient.Do(req)
	t.Assert(err, c.IsNil)
	t.Assert(res.StatusCode, c.Equals, http.StatusOK)

	// check that a HTTP request to the default port succeeds
	req, err = http.NewRequest("GET", "http://dashboard."+x.Domain, nil)
	t.Assert(err, c.IsNil)
	req.SetBasicAuth("", x.Key)
	res, err = http.DefaultClient.Do(req)
	t.Assert(err, c.IsNil)
	t.Assert(res.StatusCode, c.Equals, http.StatusOK)
}
