// Copyright 2021 rex lv
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

package goutil

type ObjectSlice struct {
	data  []interface{}
	idSet map[string]int
}

func NewObjectSlice() *ObjectSlice {
	return &ObjectSlice{
		data:  make([]interface{}, 0),
		idSet: make(map[string]int),
	}
}

func (set *ObjectSlice) Reset() {
	set.data = make([]interface{}, 0)
	set.idSet = make(map[string]int)
}

func (set *ObjectSlice) Append(obj Object) bool {
	var id = obj.ID()
	if _, ok := set.IndexID(id); ok {
		return false
	}

	set.data = append(set.data, obj)
	set.idSet[id] = len(set.data)
	return true
}

func (set *ObjectSlice) IndexID(id string) (int, bool) {
	ind, ok := set.idSet[id]
	return ind, ok
}

func (set *ObjectSlice) IndexObject(obj Object) (int, bool) {
	return set.IndexID(obj.ID())
}
