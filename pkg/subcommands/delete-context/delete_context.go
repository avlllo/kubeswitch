// Copyright 2021 The Kubeswitch authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package setcontext

import (
	"fmt"
	kubeconfigutil "github.com/danielfoehrkn/kubeswitch/pkg/util/kubectx_copied"
	"os"
)

func DeleteContext(desiredContext string) error {
	kcPath := os.Getenv("KUBECONFIG")
	kubeconfig, err := kubeconfigutil.NewKubeconfigForPath(kcPath)
	if err != nil {
		return err
	}
	if err := kubeconfig.RemoveContext(desiredContext); err != nil {
		return err
	}

	if _, err := kubeconfig.WriteKubeconfigFile(); err != nil {
		return fmt.Errorf("failed to write kubeconfig file: %v", err)
	}

	return nil
}
