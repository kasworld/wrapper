// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wrapper

import (
	"testing"
)

// var wr *Wrapper

// func TestMain(m *testing.M) {
// 	wr = New(1000)
// }

func TestWrapper_WrapPowerOfTwo(t *testing.T) {
	wr := New(1024)
	for i := -wr.GetWidth() * 1; i < wr.GetWidth()*2; i++ {
		if wr.Wrap(i) != wr.wrapByCalc(i) {
			t.Fail()
		}
	}
}

func TestWrapper_Wrap(t *testing.T) {
	wr := New(1023)
	for i := -wr.GetWidth() * 1; i < wr.GetWidth()*2; i++ {
		if wr.Wrap(i) != wr.wrapByCalc(i) {
			t.Fail()
		}
	}
}

func wrapBench(b *testing.B, wr *Wrapper, wrfn func(i int) int) int {
	sum := 0
	for i := 0; i < b.N; i++ {
		for v := -wr.GetWidth(); v < wr.GetWidth()*2; v++ {
			sum += wrfn(v)
		}
	}
	return sum
}

func BenchmarkWrap1023(b *testing.B) {
	wr := New(1023)
	wrapBench(b, wr, wr.Wrap)
}
func BenchmarkWrap1024(b *testing.B) {
	wr := New(1024)
	wrapBench(b, wr, wr.Wrap)
}
func BenchmarkGetWrapFn1023(b *testing.B) {
	wr := New(1023)
	wrapBench(b, wr, wr.GetWrapFn())
}
func BenchmarkGetWrapFn1024(b *testing.B) {
	wr := New(1024)
	wrapBench(b, wr, wr.GetWrapFn())
}

func BenchmarkWrapSafe(b *testing.B) {
	wr := New(1023)
	wrapBench(b, wr, wr.WrapSafe)
}
func BenchmarkWrapSafe1024(b *testing.B) {
	wr := New(1024)
	wrapBench(b, wr, wr.WrapSafe)
}

func BenchmarkWrapByTable(b *testing.B) {
	wr := New(1023)
	wrapBench(b, wr, wr.wrapByTable)
}

func BenchmarkWrapSimple(b *testing.B) {
	wr := New(1023)
	wrapBench(b, wr, wr.wrapSimple)
}

func BenchmarkWrapByCalc(b *testing.B) {
	wr := New(1023)
	wrapBench(b, wr, wr.wrapByCalc)
}

func BenchmarkWrapIf(b *testing.B) {
	wr := New(1023)
	wrapBench(b, wr, wr.wrapIf)
}

func BenchmarkWrapIfSimple(b *testing.B) {
	wr := New(1023)
	wrapBench(b, wr, wr.wrapIfSimple)
}
