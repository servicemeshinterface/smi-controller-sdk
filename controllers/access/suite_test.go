/*


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

package access

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	//+kubebuilder:scaffold:imports
)

var cfg *rest.Config
var k8sClient client.Client
var testEnv *envtest.Environment

func TestAPIs(t *testing.T) {
	t.Cleanup(func() {
		err := testEnv.Stop()
		require.NoError(t, err)
	})

	setupSuite(t)

	// execute tests
	//t.Run("Create Traffic Split", testCreateTrafficSplit)
	//t.Run("Delete Traffic Split", testDeleteTrafficSplit)
}

func setupSuite(t *testing.T) {
	//logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))
	testEnv = &envtest.Environment{
		CRDDirectoryPaths:        []string{filepath.Join("..", "..", "config", "crd", "bases")},
		ErrorIfCRDPathMissing:    true,
		AttachControlPlaneOutput: false,
	}

	var err error
	cfg, err = testEnv.Start()
	require.NoError(t, err)
	require.NotNil(t, cfg)
}
