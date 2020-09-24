// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package always

import (
	"github.com/goharbor/harbor/src/lib/selector"
	"github.com/goharbor/harbor/src/pkg/retention/policy/action"
	"github.com/goharbor/harbor/src/pkg/retention/policy/rule"
)

const (
	// TemplateID of the always retain rule
	TemplateID = "always"
)

type evaluator struct{}

// Process for the "always" Evaluator simply returns the input with no error
func (e *evaluator) Process(artifacts []*selector.Candidate) ([]*selector.Candidate, error) {
	return artifacts, nil
}

func (e *evaluator) Action() string {
	return action.Retain
}

// New returns an "always" Evaluator. It requires no parameters.
func New(_ rule.Parameters) rule.Evaluator {
	return &evaluator{}
}
