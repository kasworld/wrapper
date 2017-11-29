// Copyright 2017 SeukWon Kang (kasworld@gmail.com)
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

func TestWrapper_Wrap(t *testing.T) {
	wr := New(100)
	for i := -100; i < 200; i++ {
		t.Logf("%v %v", i, wr.Wrap(i))
	}
}

func BenchmarkWrap(b *testing.B) {
	wr := New(1000)
	for i := 0; i < b.N; i++ {
		for v := -wr.GetWidth(); v < wr.GetWidth()*2; v++ {
			wr.Wrap(v)
		}
	}
}

func BenchmarkWrapTable(b *testing.B) {
	wr := New(1000)
	for i := 0; i < b.N; i++ {
		for v := -wr.GetWidth(); v < wr.GetWidth()*2; v++ {
			wr.WrapTable(v)
		}
	}
}

func BenchmarkWrapSimple(b *testing.B) {
	wr := New(1000)
	for i := 0; i < b.N; i++ {
		for v := -wr.GetWidth(); v < wr.GetWidth()*2; v++ {
			wr.WrapSimple(v)
		}
	}
}

func BenchmarkWrapFull(b *testing.B) {
	wr := New(1000)
	for i := 0; i < b.N; i++ {
		for v := -wr.GetWidth(); v < wr.GetWidth()*2; v++ {
			wr.WrapFull(v)
		}
	}
}

func BenchmarkWrapIf(b *testing.B) {
	wr := New(1000)
	for i := 0; i < b.N; i++ {
		for v := -wr.GetWidth(); v < wr.GetWidth()*2; v++ {
			wr.WrapIf(v)
		}
	}
}

func BenchmarkWrapIfSimple(b *testing.B) {
	wr := New(1000)
	for i := 0; i < b.N; i++ {
		for v := -wr.GetWidth(); v < wr.GetWidth()*2; v++ {
			wr.WrapIfSimple(v)
		}
	}
}
