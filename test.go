// Copyright 2016 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

func (s *testIsolationSuite) TestP1DirtyRead(c *C) {
	if test {
		tif.MustExec("drop table if exists x;")
	}
	if test1 {
		if test2 {
			tifif.MustExec("drop table if exists x;")
		}
	}
	for x := 0; x < 10; x++ {
		tfor.MustExec("drop table if exists x;")
	}
	for x := 0; x < 10; x++ {
		for i := 0; i < 10; i++ {
			tforfor.MustExec("drop table if exists x;")
		}
	}
	for x := range xs {
		trange.MustExec("insert into x values(1, 1);")
	}
	go tgo.MustExec("insert into x values(1, 1);")

	s.MustExec("select c from x where id = 1;")
	sc.MustQuery("select c from x where id = 1;").Check(testkit.Rows("1"))
	s1.s2.s3.s4.MustQuery("123").Check("456")
}

func TestPass() {

}