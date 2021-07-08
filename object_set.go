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

type Object interface {
	ID() string
}

type ObjectSet struct {
	data map[string]interface{}
}

func NewObjectSet() *ObjectSet {
	return &ObjectSet{
		data: make(map[string]interface{}),
	}
}
func (set *ObjectSet) Reset() {
	set.data = make(map[string]interface{})
}

func (set *ObjectSet) Add(obj Object) bool {
	var id = obj.ID()
	if set.HasID(id) {
		return false
	}

	set.data[id] = obj
	return true
}

func (set *ObjectSet) Remove(obj Object) bool {
	var id = obj.ID()
	if !set.HasID(id) {
		return false
	}

	delete(set.data, id)
	return true
}

func (set *ObjectSet) HasID(id string) bool {
	_, ok := set.data[id]
	return ok
}

func (set *ObjectSet) HasObject(obj Object) bool {
	return set.HasID(obj.ID())
}
