package sdk

// API defines an object containing functions that the
// specific service meshes must implement to use this controller
type api struct {
	v1alpha *v1AlphaImpl
}

var apiInstance *api

// Returns a sinlgeton instance of the API
func API() *api {
	if apiInstance == nil {
		apiInstance = &api{
			v1alpha: &v1AlphaImpl{},
		}
	}

	return apiInstance
}

// register an interface which contains methods from the V1Alpha2 interface
// we register an interface to allow optional methods
func (a *api) RegisterV1Alpha(i interface{}) {
	a.v1alpha.RegisterV1Alpha(i)
}

func (a *api) V1Alpha() V1Alpha {
	return a.v1alpha
}
