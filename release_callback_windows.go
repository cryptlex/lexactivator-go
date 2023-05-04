// +build windows
package lexactivator

import "C"
import (
	"encoding/json"
	
)
//export newReleaseUpdateCallbackWrapper
func newReleaseUpdateCallbackWrapper(status int, releaseJson *C.ushort) {
	releaseJsonStr := ctoGoString(releaseJson)
	if releaseCallbackFunction != nil {
	   if releaseJsonStr != "" {
		  release := &Release{}
		  json.Unmarshal([]byte(releaseJsonStr), release)
		  releaseCallbackFunction(status, release, releaseCallbackFunctionUserData)
	   } else {
		  releaseCallbackFunction(status, nil, releaseCallbackFunctionUserData)
	   }
	}
}