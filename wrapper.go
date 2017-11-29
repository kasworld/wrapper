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
	"fmt"

	"github.com/kasworld/wrapper/wrapperi"
)

var _ wrapperi.WrapperI = &Wrapper{}

type Wrapper struct {
	wrapTable []int
	width     int
}

func (wr Wrapper) String() string {
	return fmt.Sprintf("Wrapper[%v]", wr.width)
}

func New(w int) *Wrapper {
	wr := &Wrapper{
		wrapTable: make([]int, w*3),
		width:     w,
	}
	for i, _ := range wr.wrapTable {
		wr.wrapTable[i] = i % w
	}
	return wr
}

func (wr *Wrapper) GetWidth() int {
	return wr.width
}

func (wr *Wrapper) Wrap(i int) int {
	return wr.wrapTable[i+wr.width]
}

func (wr *Wrapper) WrapSafe(i int) int {
	return (i%wr.width + wr.width) % wr.width
}

func (wr *Wrapper) WrapTable(i int) int {
	return wr.wrapTable[i+wr.width]
}

func (wr *Wrapper) WrapSimple(i int) int {
	return (i + wr.width) % wr.width
}

func (wr *Wrapper) WrapFull(i int) int {
	return (i%wr.width + wr.width) % wr.width
}

func (wr *Wrapper) WrapIf(i int) int {
	if i < 0 {
		return i%wr.width + wr.width
	} else if i >= wr.width {
		return i % wr.width
	} else {
		return i
	}
}

func (wr *Wrapper) WrapIfSimple(i int) int {
	if i < 0 {
		return i + wr.width
	} else if i >= wr.width {
		return i - wr.width
	} else {
		return i
	}
}
