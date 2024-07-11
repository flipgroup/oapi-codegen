package util

import (
	"net/url"

	"github.com/getkin/kin-openapi/openapi3"
)

func LoadSwagger(filePath string) (swagger *openapi3.T, err error) {

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	u, err := url.Parse(filePath)
	if err == nil && u.Scheme != "" && u.Host != "" {
		return loader.LoadFromURI(u)
	} else {
		return loader.LoadFromFile(filePath)
	}
}

func LoadSwaggerWithCircularReferenceCount(filePath string, circularReferenceCount int) (swagger *openapi3.T, err error) {
	// kin-openapi v0.126.0
	//  * `openapi3.CircularReferenceError` and `openapi3.CircularReferenceCounter` are removed. `openapi3.Loader` now implements reference backtracking, so any kind of circular references should be properly resolved.
	//
	// Read more:
	//  https://github.com/getkin/kin-openapi/blob/1a819a1374ef609f37591483a559fd5f2cb81905/README.md?plain=1#L306
	//
	// // get a copy of the existing count
	// existingCircularReferenceCount := openapi3.CircularReferenceCounter
	// if circularReferenceCount > 0 {
	// 	openapi3.CircularReferenceCounter = circularReferenceCount
	// }

	swagger, err = LoadSwagger(filePath)

	// if circularReferenceCount > 0 {
	// 	// and make sure to reset it
	//	openapi3.CircularReferenceCounter = existingCircularReferenceCount
	// }

	return swagger, err
}
