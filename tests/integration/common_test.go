//go:build integration
// +build integration

package integration

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	testNamespace = "samba-operator-system"

	testFilesDir      = "../files"
	operatorConfigDir = "../../config"

	kustomizeCmd = "kustomize"

	testExpectedImage = "ghcr.io/zakuciael/samba-operator:latest"

	testClusteredShares = false

	testShuffleOrder = false

	testMinNodeCount = 0
)

func init() {
	ns := os.Getenv("SMBOP_TEST_NAMESPACE")
	if ns != "" {
		testNamespace = ns
	}

	fdir := os.Getenv("SMBOP_TEST_FILES_DIR")
	if fdir != "" {
		testFilesDir = fdir
	}

	cdir := os.Getenv("SMBOP_TEST_CONFIG_DIR")
	if cdir != "" {
		operatorConfigDir = cdir
	}

	km := os.Getenv("SMBOP_TEST_KUSTOMIZE")
	if km != "" {
		kustomizeCmd = km
	}
	km2 := os.Getenv("KUSTOMIZE")
	if km == "" && km2 != "" {
		kustomizeCmd = km2
	}

	timg := os.Getenv("SMBOP_TEST_EXPECT_MANAGER_IMG")
	if timg != "" {
		testExpectedImage = timg
	}

	testClustering := os.Getenv("SMBOP_TEST_CLUSTERED")
	if testClustering != "" {
		testClusteredShares = true
	}

	shuffleEnv := os.Getenv("SMBOP_TEST_SHUFFLE")
	if b, err := strconv.ParseBool(shuffleEnv); b && err == nil {
		testShuffleOrder = true
	}

	// specify the number of nodes available for the test run. Some tests
	// require >1 node or >N nodes to function, so this can be used to decide
	// what to run in those cases. This should only specify the number of
	// "worker" nodes that can run samba pods - it should not include control
	// plane nodes that are not schedulable by default.
	nc := os.Getenv("SMBOP_TEST_MIN_NODE_COUNT")
	if nc != "" {
		if v, err := strconv.ParseInt(nc, 10, 32); err == nil {
			testMinNodeCount = int(v)
		} else {
			fmt.Printf("Failed to convert %s to an int\n", nc)
		}
	}

	// ensure that tests run with a random seed. This can be removed once
	// we're certain we run only with Go 1.20+ OR we ought to make this
	// settable for test reproduction purposes.
	rand.Seed(time.Now().UnixNano())
}
