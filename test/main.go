package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/cucumber/messages-go/v10"
	"github.com/go-logr/logr"
	"github.com/hashicorp/go-hclog"
	accessv1alpha1 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha1"
	accessv1alpha2 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha2"
	accessv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha3"
	"github.com/servicemeshinterface/smi-controller-sdk/controllers/helpers"
	"github.com/servicemeshinterface/smi-controller-sdk/sdk"
	"github.com/servicemeshinterface/smi-controller-sdk/sdk/controller"
	"github.com/shipyard-run/shipyard/pkg/clients"
	"github.com/stretchr/testify/mock"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubectl/pkg/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var opts = &godog.Options{
	Format: "pretty",
	Output: colors.Colored(os.Stdout),
}

var mockAPI *helpers.MockAPI
var logger logr.Logger
var k8sClient clients.Kubernetes
var config controller.Config

// store a reference to any objects submitted to the controller for later cleanup
var crds []interface{}

func main() {
	godog.BindFlags("godog.", flag.CommandLine, opts)
	flag.Parse()

	status := godog.TestSuite{
		Name:                 "SDK Functional Tests",
		ScenarioInitializer:  initializeScenario,
		TestSuiteInitializer: initializeSuite,
		Options:              opts,
	}.Run()

	os.Exit(status)
}

func setupClient() error {

	err := setupSplits()
	if err != nil {
		return err
	}

	err = setupAccess()
	if err != nil {
		return err
	}

	err = setupSpecs()
	if err != nil {
		return err
	}

	k8sClient = clients.NewKubernetes(10*time.Millisecond, hclog.NewNullLogger())
	k8sClient, err = k8sClient.SetConfig(os.Getenv("KUBECONFIG"))

	return err
}

func initializeSuite(ctx *godog.TestSuiteContext) {
	logger = Log()

	err := setupClient()
	if err != nil {
		panic(err)
	}

	// create and start the controller
	setupMockAPI()

	sdk.API().RegisterV1Alpha(mockAPI)

	config = controller.DefaultConfig()
	config.WebhooksEnabled = true
	config.Logger = logger

	go controller.Start(config)
}

func initializeScenario(ctx *godog.ScenarioContext) {
	crds = []interface{}{}

	ctx.Step(`^the server is running$`, theServerIsRunning)
	ctx.Step(`^I create the following resource$`, iCreateTheFollowingResource)
	ctx.Step(`^I expect "([^"]*)" to be called (\d+) time$`, iExpectToBeCalled)

	ctx.AfterScenario(func(s *messages.Pickle, err error) {
		cleanupResources()

		if err != nil {
			fmt.Println("Error occurred running the tests", err)
			fmt.Println(logger.(StringLogger).String())
		}

		// wait for server to have cleaned up objects and exit
		// as deleting an object is not immediate.
		// We should probably handle this eventual consistency in the cleanup
		// function however this 5 delay should be fine.
		// If you are raising a PR to fix this "should be fine" be sure to shame me in the comments
		time.Sleep(5 * time.Second)
	})
}

func cleanupResources() {
	c := getK8sConfig()
	kc, err := client.New(c, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	kc.DeleteAllOf(
		ctx,
		&accessv1alpha3.TrafficTarget{}, client.InNamespace("default"))

	if err != nil {
		fmt.Println("Error removing v3 TrafficTargets", err)
	}

	kc.DeleteAllOf(
		ctx,
		&accessv1alpha2.TrafficTarget{}, client.InNamespace("default"))

	if err != nil {
		fmt.Println("Error removing v2 TrafficTargets", err)
	}

	kc.DeleteAllOf(
		ctx,
		&accessv1alpha1.TrafficTarget{}, client.InNamespace("default"))

	if err != nil {
		fmt.Println("Error removing v1 TrafficTargets", err)
	}
}

func setupMockAPI() {
	mockAPI = &helpers.MockAPI{}
	mockAPI.On("UpsertTrafficTarget", mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(ctrl.Result{}, nil)

	mockAPI.On("DeleteTrafficTarget", mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(ctrl.Result{}, nil)

	mockAPI.On("UpsertTrafficSplit", mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(ctrl.Result{}, nil)

	mockAPI.On("DeleteTrafficSplit", mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(ctrl.Result{}, nil)

	mockAPI.On("UpsertHTTPRouteGroup", mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(ctrl.Result{}, nil)

	mockAPI.On("DeleteHTTPRouteGroup", mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(ctrl.Result{}, nil)

	mockAPI.On("UpsertTCPRoute", mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(ctrl.Result{}, nil)

	mockAPI.On("DeleteTCPRoute", mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(ctrl.Result{}, nil)

	mockAPI.On("UpsertUDPRoute", mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(ctrl.Result{}, nil)

	mockAPI.On("DeleteUDPRoute", mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(ctrl.Result{}, nil)
}

func theServerIsRunning() error {
	return waitForComplete(
		30*time.Second,
		func() error {
			resp, err := http.Get(fmt.Sprintf("http://%s/readyz", config.HealthProbeBindAddress))
			if err == nil {
				if resp != nil && resp.StatusCode == http.StatusOK {
					return nil
				}
			}

			return fmt.Errorf("Timeout waiting for service to become ready")
		},
	)
}

func iCreateTheFollowingResource(arg1 *messages.PickleStepArgument_PickleDocString) error {
	// save the document to a temporary file
	f, err := ioutil.TempFile("", "*.yaml")
	if err != nil {
		return err
	}

	// cleanup
	defer os.Remove(f.Name())

	// write the document to the file
	_, err = f.WriteString(arg1.GetContent())
	if err != nil {
		return err
	}

	// import the file to the kubernetes cluster
	err = k8sClient.Apply([]string{f.Name()}, true)
	return err
}

// The controller is eventually consistent so we need to check this in a loop
func iExpectToBeCalled(method string, n int) error {
	return waitForComplete(
		30*time.Second,
		func() error {

			count := 0
			for _, call := range mockAPI.Calls {
				if call.Method == method {
					count++
				}
			}

			if count != n {
				return fmt.Errorf("Expected %s to be called %d time(s), it was called %d time(s)", method, n, count)
			}

			return nil
		},
	)
}

func getK8sConfig() *rest.Config {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		panic(err.Error())
	}

	return config
}

// helper function to loop until a condition is met
func waitForComplete(duration time.Duration, f func() error) error {
	// wait for the server to mark it is ready
	done := make(chan struct{})
	timeout := time.After(duration)

	var err error

	go func() {
		for {
			err = f()
			if err == nil {
				done <- struct{}{}
				break
			}

			// retry after 1s
			time.Sleep(1 * time.Second)
		}
	}()

	select {
	case <-timeout:
		return err
	case <-done:
		return nil
	}
}
