/*
 * Copyright 2015-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package memory_test

import (
	"testing"

	"github.com/instana/java-buildpack-memory-calculator/v4/memory"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func TestMaxHeapYoungGeneration(t *testing.T) {
	spec.Run(t, "MaxHeapYoungGeneration", func(t *testing.T, when spec.G, it spec.S) {

		g := NewGomegaWithT(t)

		it("formats", func() {
			g.Expect(memory.MaxHeapYoungGeneration(memory.Kibi).String()).To(Equal("-Xmn1K"))
		})

		it("matches -Xmn", func() {
			g.Expect(memory.IsMaxHeapYoungGeneration("-Xmn1K")).To(BeTrue())
		})

		it("does not match non -Xmn", func() {
			g.Expect(memory.IsMaxHeapYoungGeneration("-Xss1K")).To(BeFalse())
		})

		it("parses", func() {
			g.Expect(memory.ParseMaxHeapYoungGeneration("-Xmn1K")).To(Equal(memory.MaxHeapYoungGeneration(memory.Kibi)))
		})

	})
}
