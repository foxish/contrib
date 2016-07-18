/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

// CleanStringSlice will take a string slice with a single empty value and return an empty slice.
// This is extremely useful for StringSlice flags, so the user can do --flag="" and instead
// of getting []string{""} they will get []string{}
func CleanStringSlice(in []string) []string {
	if len(in) == 1 && len(in[0]) == 0 {
		return []string{}
	}
	return in
}
