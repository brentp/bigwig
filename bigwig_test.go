package bigwig_test

import (
	"fmt"
	"testing"

	"github.com/brentp/bigwig"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type BWSuite struct {
	bw *bigwig.BigWig
}

var _ = Suite(&BWSuite{})

func (s *BWSuite) SetUpTest(c *C) {

	var err error
	s.bw, err = bigwig.Open("test.bw")
	c.Assert(err, IsNil)
}

func (s *BWSuite) TestValues(c *C) {
	res := s.bw.Values("1", 0, 9)
	c.Assert(len(res), Equals, 3)
	c.Assert(fmt.Sprintf("%.1f", res[0]), Equals, "0.1")
	c.Assert(fmt.Sprintf("%.1f", res[1]), Equals, "0.2")
	c.Assert(fmt.Sprintf("%.1f", res[2]), Equals, "0.3")
}

func (s *BWSuite) TestMean(c *C) {
	res := s.bw.Mean("1", 0, 9)
	c.Assert(fmt.Sprintf("%.1f", res), Equals, "0.2")
}
