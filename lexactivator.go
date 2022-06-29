// Copyright 2020 Cryptlex, LLC. All rights reserved.

package lexactivator

/*
#cgo linux,!arm64 LDFLAGS: -Wl,-Bstatic -L${SRCDIR} -lLexActivator_amd64 -Wl,-Bdynamic -lm -lstdc++ -lpthread
#cgo linux,arm64 LDFLAGS: -Wl,-Bstatic -L${SRCDIR} -lLexActivator_arm64 -Wl,-Bdynamic -lm -lstdc++ -lpthread
#cgo darwin LDFLAGS: -L${SRCDIR}/libs/darwin_amd64 -lLexActivator -lc++ -framework CoreFoundation -framework SystemConfiguration -framework Security
#cgo windows LDFLAGS: -L${SRCDIR}/libs/windows_amd64 -lLexActivator
#include "LexActivator.h"
#include <stdlib.h>
void licenseCallbackCgoGateway(int status);
void releaseUpdateCallbackCgoGateway(int status);
*/
import "C"
import (
	"unsafe"
)

type callbackType func(int)

const (
	LA_USER      uint = 1
	LA_SYSTEM    uint = 2
	LA_IN_MEMORY uint = 4
)

var licenseCallbackFuncion callbackType

var releaseUpdateCallbackFuncion callbackType

//export licenseCallbackWrapper
func licenseCallbackWrapper(status int) {
	if licenseCallbackFuncion != nil {
		licenseCallbackFuncion(status)
	}
}

//export releaseUpdateCallbackWrapper
func releaseUpdateCallbackWrapper(status int) {
	if releaseUpdateCallbackFuncion != nil {
		releaseUpdateCallbackFuncion(status)
	}
}

/*
   FUNCTION: SetProductFile()

   PURPOSE: Sets the absolute path of the Product.dat file.

   This function must be called on every start of your program
   before any other functions are called.

   PARAMETERS:
   * filePath - absolute path of the product file (Product.dat)

   RETURN CODES: LA_OK, LA_E_FILE_PATH, LA_E_PRODUCT_FILE

   NOTE: If this function fails to set the path of product file, none of the
   other functions will work.
*/
func SetProductFile(filePath string) int {
	cFilePath := goToCString(filePath)
	status := C.SetProductFile(cFilePath)
	freeCString(cFilePath)
	return int(status)
}

/*
   FUNCTION: SetProductData()

   PURPOSE: Embeds the Product.dat file in the application.

   It can be used instead of SetProductFile() in case you want
   to embed the Product.dat file in your application.

   This function must be called on every start of your program
   before any other functions are called.

   PARAMETERS:
   * productData - content of the Product.dat file

   RETURN CODES: LA_OK, LA_E_PRODUCT_DATA

   NOTE: If this function fails to set the product data, none of the
   other functions will work.
*/
func SetProductData(productData string) int {
	cProductData := goToCString(productData)
	status := C.SetProductData(cProductData)
	freeCString(cProductData)
	return int(status)
}

/*
   FUNCTION: SetProductId()

   PURPOSE: Sets the product id of your application.

   This function must be called on every start of your program before
   any other functions are called, with the exception of SetProductFile()
   or SetProductData() function.

   PARAMETERS:
   * productId - the unique product id of your application as mentioned
     on the product page in the dashboard.

   * flags - depending upon whether your application requires admin/root
     permissions to run or not, this parameter can have one of the following
     values: LA_SYSTEM, LA_USER, LA_IN_MEMORY

   RETURN CODES: LA_OK, LA_E_WMIC, LA_E_PRODUCT_FILE, LA_E_PRODUCT_DATA, LA_E_PRODUCT_ID,
   LA_E_SYSTEM_PERMISSION

   NOTE: If this function fails to set the product id, none of the other
   functions will work.
*/
func SetProductId(productId string, flags uint) int {
	cProductId := goToCString(productId)
	cFlags := (C.uint)(flags)
	status := C.SetProductId(cProductId, cFlags)
	freeCString(cProductId)
	return int(status)
}

/*
   FUNCTION: SetDataDirectory()

   PURPOSE: In case you want to change the default directory used by LexActivator to
   store the activation data on Linux and macOS, this function can be used to
   set a different directory.

   If you decide to use this function, then it must be called on every start of
   your program before calling SetProductFile() or SetProductData() function.

   Please ensure that the directory exists and your app has read and write
   permissions in the directory.

   PARAMETERS:
   * directoryPath - absolute path of the directory.

   RETURN CODES: LA_OK, LA_E_FILE_PERMISSION

*/
func SetDataDirectory(directoryPath string) int {
	cDirectoryPath := goToCString(directoryPath)
	status := C.SetDataDirectory(cDirectoryPath)
	freeCString(cDirectoryPath)
	return int(status)
}

/*
   FUNCTION: SetCustomDeviceFingerprint()

   PURPOSE: In case you don't want to use the LexActivator's advanced
   device fingerprinting algorithm, this function can be used to set a custom
   device fingerprint.

   If you decide to use your own custom device fingerprint then this function must be
   called on every start of your program immediately after calling SetProductFile()
   or SetProductData() function.

   The license fingerprint matching strategy is ignored if this function is used.

   PARAMETERS:
   * fingerprint - string of minimum length 64 characters and maximum length 256 characters.

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_CUSTOM_FINGERPRINT_LENGTH
*/
func SetCustomDeviceFingerprint(fingerprint string) int {
	cFingerprint := goToCString(fingerprint)
	status := C.SetCustomDeviceFingerprint(cFingerprint)
	freeCString(cFingerprint)
	return int(status)
}

/*
   FUNCTION: SetLicenseKey()

   PURPOSE: Sets the license key required to activate the license.

   PARAMETERS:
   * licenseKey - a valid license key.

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY
*/
func SetLicenseKey(licenseKey string) int {
	cLicenseKey := goToCString(licenseKey)
	status := C.SetLicenseKey(cLicenseKey)
	freeCString(cLicenseKey)
	return int(status)
}

/*
   FUNCTION: SetLicenseUserCredential()

   PURPOSE: Sets the license user email and password for authentication.

   This function must be called before ActivateLicense() or IsLicenseGenuine()
   function if 'requireAuthentication' property of the license is set to true.

   PARAMETERS:
   * email - user email address.
   * password - user password.

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY
*/
func SetLicenseUserCredential(email string, password string) int {
	cEmail := goToCString(email)
	cPassword := goToCString(password)
	status := C.SetLicenseUserCredential(cEmail, cPassword)
	freeCString(cEmail)
	freeCString(cPassword)
	return int(status)
}

/*
   FUNCTION: SetLicenseCallback()

   PURPOSE: Sets server sync callback function.

   Whenever the server sync occurs in a separate thread, and server returns the response,
   license callback function gets invoked with the following status codes:
   LA_OK, LA_EXPIRED, LA_SUSPENDED,
   LA_E_REVOKED, LA_E_ACTIVATION_NOT_FOUND, LA_E_MACHINE_FINGERPRINT
   LA_E_AUTHENTICATION_FAILED, LA_E_COUNTRY, LA_E_INET, LA_E_SERVER,
   LA_E_RATE_LIMIT, LA_E_IP

   PARAMETERS:
   * callback - name of the callback function

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY
*/
func SetLicenseCallback(callbackFunction func(int)) int {
	status := C.SetLicenseCallback((C.CallbackType)(unsafe.Pointer(C.licenseCallbackCgoGateway)))
	licenseCallbackFuncion = callbackFunction
	return int(status)
}

/*
    FUNCTION: SetActivationLeaseDuration()

    PURPOSE: Sets the lease duration for the activation.

    The activation lease duration is honoured when the allow client
    lease duration property is enabled.

    PARAMETERS:
    * leaseDuration - value of the lease duration.

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY
*/
func SetActivationLeaseDuration(leaseDuration uint) int {
   cLeaseDuration := (C.uint)(leaseDuration)
   status := C.SetActivationLeaseDuration(cLeaseDuration)
   return int(status)
}

/*
   FUNCTION: SetActivationMetadata()

   PURPOSE: Sets the activation metadata.

   The  metadata appears along with the activation details of the license
   in dashboard.

   PARAMETERS:
   * key - string of maximum length 256 characters with utf-8 encoding.
   * value - string of maximum length 256 characters with utf-8 encoding.

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_METADATA_KEY_LENGTH,
   LA_E_METADATA_VALUE_LENGTH, LA_E_ACTIVATION_METADATA_LIMIT
*/
func SetActivationMetadata(key string, value string) int {
	cKey := goToCString(key)
	cValue := goToCString(value)
	status := C.SetActivationMetadata(cKey, cValue)
	freeCString(cKey)
	freeCString(cValue)
	return int(status)
}

/*
   FUNCTION: SetTrialActivationMetadata()

   PURPOSE: Sets the trial activation metadata.

   The  metadata appears along with the trial activation details of the product
   in dashboard.

   PARAMETERS:
   * key - string of maximum length 256 characters with utf-8 encoding.
   * value - string of maximum length 256 characters with utf-8 encoding.

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_METADATA_KEY_LENGTH,
   LA_E_METADATA_VALUE_LENGTH, LA_E_TRIAL_ACTIVATION_METADATA_LIMIT
*/
func SetTrialActivationMetadata(key string, value string) int {
	cKey := goToCString(key)
	cValue := goToCString(value)
	status := C.SetTrialActivationMetadata(cKey, cValue)
	freeCString(cKey)
	freeCString(cValue)
	return int(status)
}

/*
   FUNCTION: SetAppVersion()

   PURPOSE: Sets the current app version of your application.

   The app version appears along with the activation details in dashboard. It
   is also used to generate app analytics.

   PARAMETERS:
   * appVersion - string of maximum length 256 characters with utf-8 encoding.

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_APP_VERSION_LENGTH
*/
func SetAppVersion(appVersion string) int {
	cAppVersion := goToCString(appVersion)
	status := C.SetAppVersion(cAppVersion)
	freeCString(cAppVersion)
	return int(status)
}

/*
   FUNCTION: SetReleaseVersion()

   PURPOSE: Sets the current release version of your application.

   The release version appears along with the activation details in dashboard.

   PARAMETERS:
   * releaseVersion - string in following allowed formats: x.x, x.x.x, x.x.x.x

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_RELEASE_VERSION_FORMAT
*/
func SetReleaseVersion(releaseVersion string) int {
	cReleaseVersion := goToCString(releaseVersion)
	status := C.SetReleaseVersion(cReleaseVersion)
	freeCString(cReleaseVersion)
	return int(status)
}


/*
   FUNCTION: SetOfflineActivationRequestMeterAttributeUses()

   PURPOSE: Sets the meter attribute uses for the offline activation request.

   This function should only be called before GenerateOfflineActivationRequest()
   function to set the meter attributes in case of offline activation.

   PARAMETERS:
   * name - name of the meter attribute
   * uses - the uses value

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY

*/
func SetOfflineActivationRequestMeterAttributeUses(name string, uses uint) int {
	cName := goToCString(name)
	cUses := (C.uint)(uses)
	status := C.SetOfflineActivationRequestMeterAttributeUses(cName, cUses)
	freeCString(cName)
	return int(status)
}

/*
   FUNCTION: SetNetworkProxy()

   PURPOSE: Sets the network proxy to be used when contacting Cryptlex servers.

   The proxy format should be: [protocol://][username:password@]machine[:port]

   Following are some examples of the valid proxy strings:
       - http://127.0.0.1:8000/
       - http://user:pass@127.0.0.1:8000/
       - socks5://127.0.0.1:8000/

   PARAMETERS:
   * proxy - proxy string having correct proxy format

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_NET_PROXY

   NOTE: Proxy settings of the computer are automatically detected. So, in most of the
   cases you don't need to care whether your user is behind a proxy server or not.
*/
func SetNetworkProxy(proxy string) int {
	cProxy := goToCString(proxy)
	status := C.SetNetworkProxy(cProxy)
	freeCString(cProxy)
	return int(status)
}

/*
   FUNCTION: SetCryptlexHost()

   PURPOSE: In case you are running Cryptlex on-premise, you can set the
   host for your on-premise server.

   PARAMETERS:
   * host - the address of the Cryptlex on-premise server

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_HOST_URL
*/
func SetCryptlexHost(host string) int {
	cHost := goToCString(host)
	status := C.SetCryptlexHost(cHost)
	freeCString(cHost)
	return int(status)
}

/*
   FUNCTION: GetProductMetadata()

   PURPOSE: Gets the product metadata as set in the dashboard.

   This is available for trial as well as license activations.

   PARAMETERS:
   * key - metadata key to retrieve the value
   * value - pointer to a string that receives the value

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_METADATA_KEY_NOT_FOUND, LA_E_BUFFER_SIZE
*/
func GetProductMetadata(key string, value *string) int {
	cKey := goToCString(key)
	var cValue = getCArray()
	status := C.GetProductMetadata(cKey, &cValue[0], maxCArrayLength)
	*value = ctoGoString(&cValue[0])
	freeCString(cKey)
	return int(status)
}

/*
   FUNCTION: GetProductVersionName()

   PURPOSE: Gets the product version name.

   PARAMETERS:
   * name - pointer to a buffer that receives the value of the string
   * length - size of the buffer pointed to by the name parameter

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED, LA_E_PRODUCT_VERSION_NOT_LINKED,
   LA_E_BUFFER_SIZE
*/

func GetProductVersionName(name *string) int {
	var cName = getCArray()
	status := C.GetProductVersionName(&cName[0], maxCArrayLength)
	*name = ctoGoString(&cName[0])
	return int(status)
}

/*
   FUNCTION: GetProductVersionDisplayName()

   PURPOSE: Gets the product version display name.

   PARAMETERS:
   * displayName - pointer to a buffer that receives the value of the string
   * length - size of the buffer pointed to by the displayName parameter

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED, LA_E_PRODUCT_VERSION_NOT_LINKED,
   LA_E_BUFFER_SIZE
*/
func GetProductVersionDisplayName(displayName *string) int {
	var cDisplayName = getCArray()
	status := C.GetProductVersionDisplayName(&cDisplayName[0], maxCArrayLength)
	*displayName = ctoGoString(&cDisplayName[0])
	return int(status)
}

/*
   FUNCTION: GetProductVersionFeatureFlag()

   PURPOSE: Gets the product version feature flag.

   PARAMETERS:
   * name - name of the feature flag
   * enabled - pointer to the integer that receives the value - 0 or 1
   * data - pointer to a buffer that receives the value of the string
   * length - size of the buffer pointed to by the data parameter

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED, LA_E_PRODUCT_VERSION_NOT_LINKED,
   LA_E_FEATURE_FLAG_NOT_FOUND, LA_E_BUFFER_SIZE
*/
func GetProductVersionFeatureFlag(name string, enabled *bool, data *string) int {
   cName := goToCString(name)
   var cEnabled C.uint
   var cData = getCArray()
   status := C.GetProductVersionFeatureFlag(cName, &cEnabled, &cData[0], maxCArrayLength)
   freeCString(cName)
   *enabled = cEnabled > 0
   *data = ctoGoString(&cData[0])
   return int(status)
}

/*
   FUNCTION: GetLicenseMetadata()

   PURPOSE: Gets the license metadata as set in the dashboard.

   PARAMETERS:
   * key - metadata key to retrieve the value
   * value - pointer to a string that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METADATA_KEY_NOT_FOUND, LA_E_BUFFER_SIZE
*/
func GetLicenseMetadata(key string, value *string) int {
	cKey := goToCString(key)
	var cValue = getCArray()
	status := C.GetLicenseMetadata(cKey, &cValue[0], maxCArrayLength)
	*value = ctoGoString(&cValue[0])
	freeCString(cKey)
	return int(status)
}

/*
   FUNCTION: GetLicenseMeterAttribute()

   PURPOSE: Gets the license meter attribute allowed uses, total and gross uses.

   PARAMETERS:
   * name - name of the meter attribute
   * allowedUses - pointer to the integer that receives the value
   * totalUses - pointer to the integer that receives the value
   * grossUses - pointer to the integer that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METER_ATTRIBUTE_NOT_FOUND
*/
func GetLicenseMeterAttribute(name string, allowedUses *uint, totalUses *uint, grossUses *uint) int {
	cName := goToCString(name)
	var cAllowedUses C.uint
	var cTotalUses C.uint
	var cGrossUses C.uint
	status := C.GetLicenseMeterAttribute(cName, &cAllowedUses, &cTotalUses, &cGrossUses)
	*allowedUses = uint(cAllowedUses)
	*totalUses = uint(cTotalUses)
	*grossUses = uint(cGrossUses)
	freeCString(cName)
	return int(status)
}

/*
   FUNCTION: GetLicenseKey()

   PURPOSE: Gets the license key used for activation.

   PARAMETERS:
   * licenseKey - pointer to a string that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_BUFFER_SIZE
*/
func GetLicenseKey(licenseKey *string) int {
	var cLicenseKey = getCArray()
	status := C.GetLicenseKey(&cLicenseKey[0], maxCArrayLength)
	*licenseKey = ctoGoString(&cLicenseKey[0])
	return int(status)
}

/*
   FUNCTION: GetLicenseAllowedActivations()

   PURPOSE: Gets the allowed activations of the license.

   PARAMETERS:
   * allowedActivations - pointer to the integer that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED
*/
func GetLicenseAllowedActivations(allowedActivations *uint) int {
	var cAllowedActivations C.uint
	status := C.GetLicenseAllowedActivations(&cAllowedActivations)
	*allowedActivations = uint(cAllowedActivations)
	return int(status)
}

/*
   FUNCTION: GetLicenseTotalActivations()

   PURPOSE: Gets the total activations of the license.

   PARAMETERS:
   * totalActivations - pointer to the integer that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED
*/
func GetLicenseTotalActivations(totalActivations *uint) int {
	var cTotalActivations C.uint
	status := C.GetLicenseTotalActivations(&cTotalActivations)
	*totalActivations = uint(cTotalActivations)
	return int(status)
}

/*
   FUNCTION: GetLicenseExpiryDate()

   PURPOSE: Gets the license expiry date timestamp.

   PARAMETERS:
   * expiryDate - pointer to the integer that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED
*/
func GetLicenseExpiryDate(expiryDate *uint) int {
	var cExpiryDate C.uint
	status := C.GetLicenseExpiryDate(&cExpiryDate)
	*expiryDate = uint(cExpiryDate)
	return int(status)
}

/*
    FUNCTION: GetLicenseMaintenanceExpiryDate()

    PURPOSE: Gets the license maintenance expiry date timestamp.

    PARAMETERS:
    * maintenanceExpiryDate - pointer to the integer that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_TIME, LA_E_TIME_MODIFIED
*/
func GetLicenseMaintenanceExpiryDate(maintenanceExpiryDate *uint) int {
   var cMaintenanceExpiryDate C.uint
	status := C.GetLicenseMaintenanceExpiryDate(&cMaintenanceExpiryDate)
	*maintenanceExpiryDate = uint(cMaintenanceExpiryDate)
	return int(status)
}

/*
   FUNCTION: GetLicenseUserEmail()

   PURPOSE: Gets the email associated with license user.

   PARAMETERS:
   * email - pointer to a string that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED,
   LA_E_BUFFER_SIZE
*/
func GetLicenseUserEmail(email *string) int {
	var cEmail = getCArray()
	status := C.GetLicenseUserEmail(&cEmail[0], maxCArrayLength)
	*email = ctoGoString(&cEmail[0])
	return int(status)
}

/*
   FUNCTION: GetLicenseUserName()

   PURPOSE: Gets the name associated with the license user.

   PARAMETERS:
   * name - pointer to a string that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED,
   LA_E_BUFFER_SIZE
*/
func GetLicenseUserName(name *string) int {
	var cName = getCArray()
	status := C.GetLicenseUserName(&cName[0], maxCArrayLength)
	*name = ctoGoString(&cName[0])
	return int(status)
}

/*
   FUNCTION: GetLicenseUserCompany()

   PURPOSE: Gets the company associated with the license user.

   PARAMETERS:
   * company - pointer to a string that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED,
   LA_E_BUFFER_SIZE
*/
func GetLicenseUserCompany(company *string) int {
	var cCompany = getCArray()
	status := C.GetLicenseUserCompany(&cCompany[0], maxCArrayLength)
	*company = ctoGoString(&cCompany[0])
	return int(status)
}

/*
   FUNCTION: GetLicenseUserMetadata()

   PURPOSE: Gets the metadata associated with the license user.

   PARAMETERS:
   * key - metadata key to retrieve the value
   * value - pointer to a string that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METADATA_KEY_NOT_FOUND, LA_E_BUFFER_SIZE
*/
func GetLicenseUserMetadata(key string, value *string) int {
	cKey := goToCString(key)
	var cValue = getCArray()
	status := C.GetLicenseUserMetadata(cKey, &cValue[0], maxCArrayLength)
	*value = ctoGoString(&cValue[0])
	freeCString(cKey)
	return int(status)
}

/*
   FUNCTION: GetLicenseType()

   PURPOSE: Gets the license type (node-locked or hosted-floating).

   PARAMETERS:
   * licenseType - pointer to a string that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED,
   LA_E_BUFFER_SIZE
*/
func GetLicenseType(licenseType *string) int {
	var cLicenseType = getCArray()
	status := C.GetLicenseType(&cLicenseType[0], maxCArrayLength)
	*licenseType = ctoGoString(&cLicenseType[0])
	return int(status)
}

/*
   FUNCTION: GetActivationMetadata()

   PURPOSE: Gets the activation metadata.

   PARAMETERS:
   * key - metadata key to retrieve the value
   * value - pointer to a string that receives the value

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_METADATA_KEY_NOT_FOUND, LA_E_BUFFER_SIZE
*/
func GetActivationMetadata(key string, value *string) int {
	cKey := goToCString(key)
	var cValue = getCArray()
	status := C.GetActivationMetadata(cKey, &cValue[0], maxCArrayLength)
	*value = ctoGoString(&cValue[0])
	freeCString(cKey)
	return int(status)
}

/*
    FUNCTION: GetActivationMode()

    PURPOSE: Gets the mode of activation (online or offline).

    PARAMETERS:
    * initialMode - pointer to a buffer that receives the initial mode of activation
    * initialModeLength - size of the buffer pointed to by the initialMode parameter
    * currentMode - pointer to a buffer that receives the current mode of activation
    * currentModeLength - size of the buffer pointed to by the currentMode parameter

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_TIME_MODIFIED,
    LA_E_BUFFER_SIZE
*/
func GetActivationMode(initialMode *string, currentMode *string) int {
   var cInitialMode = getCArray()
   var cCurrentMode = getCArray()
   status := C.GetActivationMode(&cInitialMode[0],maxCArrayLength, &cCurrentMode[0], maxCArrayLength)
   *initialMode = ctoGoString(&cInitialMode[0])
   *currentMode = ctoGoString(&cCurrentMode[0])
   return int(status)
}

/*
   FUNCTION: GetActivationMeterAttributeUses()

   PURPOSE: Gets the meter attribute uses consumed by the activation.

   PARAMETERS:
   * name - name of the meter attribute
   * allowedUses - pointer to the integer that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METER_ATTRIBUTE_NOT_FOUND
*/
func GetActivationMeterAttributeUses(name string, uses *uint) int {
	cName := goToCString(name)
	var cUses C.uint
	status := C.GetActivationMeterAttributeUses(cName, &cUses)
	*uses = uint(cUses)
	freeCString(cName)
	return int(status)
}

/*
   FUNCTION: GetServerSyncGracePeriodExpiryDate()

   PURPOSE: Gets the server sync grace period expiry date timestamp.

   PARAMETERS:
   * expiryDate - pointer to the integer that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED
*/
func GetServerSyncGracePeriodExpiryDate(expiryDate *uint) int {
	var cExpiryDate C.uint
	status := C.GetServerSyncGracePeriodExpiryDate(&cExpiryDate)
	*expiryDate = uint(cExpiryDate)
	return int(status)
}

/*
   FUNCTION: GetTrialActivationMetadata()

   PURPOSE: Gets the trial activation metadata.

   PARAMETERS:
   * key - metadata key to retrieve the value
   * value - pointer to a string that receives the value

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_METADATA_KEY_NOT_FOUND, LA_E_BUFFER_SIZE
*/
func GetTrialActivationMetadata(key string, value *string) int {
	cKey := goToCString(key)
	var cValue = getCArray()
	status := C.GetTrialActivationMetadata(cKey, &cValue[0], maxCArrayLength)
	*value = ctoGoString(&cValue[0])
	freeCString(cKey)
	return int(status)
}

/*
   FUNCTION: GetTrialExpiryDate()

   PURPOSE: Gets the trial expiry date timestamp.

   PARAMETERS:
   * trialExpiryDate - pointer to the integer that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED
*/
func GetTrialExpiryDate(trialExpiryDate *uint) int {
	var cTrialExpiryDate C.uint
	status := C.GetTrialExpiryDate(&cTrialExpiryDate)
	*trialExpiryDate = uint(cTrialExpiryDate)
	return int(status)
}

/*
   FUNCTION: GetTrialId()

   PURPOSE: Gets the trial activation id. Used in case of trial extension.

   PARAMETERS:
   * trialId - pointer to a string that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED,
   LA_E_BUFFER_SIZE
*/
func GetTrialId(trialId *string) int {
	var cTrialId = getCArray()
	status := C.GetTrialId(&cTrialId[0], maxCArrayLength)
	*trialId = ctoGoString(&cTrialId[0])
	return int(status)
}

/*
   FUNCTION: GetLocalTrialExpiryDate()

   PURPOSE: Gets the trial expiry date timestamp.

   PARAMETERS:
   * trialExpiryDate - pointer to the integer that receives the value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME_MODIFIED
*/
func GetLocalTrialExpiryDate(trialExpiryDate *uint) int {
	var cTrialExpiryDate C.uint
	status := C.GetLocalTrialExpiryDate(&cTrialExpiryDate)
	*trialExpiryDate = uint(cTrialExpiryDate)
	return int(status)
}

/*
   FUNCTION: GetLibraryVersion()

   PURPOSE: Gets the version of this library.

   PARAMETERS:
   * libraryVersion - pointer to a buffer that receives the value of the string
   * length - size of the buffer pointed to by the libraryVersion parameter

   RETURN CODES: LA_OK, LA_E_BUFFER_SIZE
*/
func GetLibraryVersion(libraryVersion *string) int {
	var cLibraryVersion = getCArray()
	status := C.GetLibraryVersion(&cLibraryVersion[0], maxCArrayLength)
	*libraryVersion = ctoGoString(&cLibraryVersion[0])
	return int(status)
}

/*
   FUNCTION: CheckForReleaseUpdate()

   PURPOSE: Checks whether a new release is available for the product.

   This function should only be used if you manage your releases through
   Cryptlex release management API.

   PARAMETERS:
   * platform - release platform e.g. windows, macos, linux
   * version - current release version
   * channel - release channel e.g. stable
   * releaseUpdateCallback - name of the callback function.

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_RELEASE_VERSION_FORMAT
*/
func CheckForReleaseUpdate(platform string, version string, channel string, callbackFunction func(int)) int {
	cPlatform := goToCString(platform)
	cVersion := goToCString(version)
	cChannel := goToCString(channel)
	status := C.CheckForReleaseUpdate(cPlatform, cVersion, cChannel, (C.CallbackType)(unsafe.Pointer(C.releaseUpdateCallbackCgoGateway)))
	releaseUpdateCallbackFuncion = callbackFunction
	freeCString(cPlatform)
	freeCString(cVersion)
	freeCString(cChannel)
	return int(status)
}

/*
   FUNCTION: ActivateLicense()

   PURPOSE: Activates the license by contacting the Cryptlex servers. It
   validates the key and returns with encrypted and digitally signed token
   which it stores and uses to activate your application.

   This function should be executed at the time of registration, ideally on
   a button click.

   RETURN CODES: LA_OK, LA_EXPIRED, LA_SUSPENDED, LA_E_REVOKED, LA_FAIL, LA_E_PRODUCT_ID,
   LA_E_INET, LA_E_VM, LA_E_TIME, LA_E_ACTIVATION_LIMIT, LA_E_SERVER, LA_E_CLIENT,
   LA_E_AUTHENTICATION_FAILED, LA_E_LICENSE_TYPE, LA_E_COUNTRY, LA_E_IP, LA_E_RATE_LIMIT, LA_E_LICENSE_KEY
*/
func ActivateLicense() int {
	status := C.ActivateLicense()
	return int(status)
}

/*
   FUNCTION: ActivateLicenseOffline()

   PURPOSE: Activates your licenses using the offline activation response file.

   PARAMETERS:
   * filePath - path of the offline activation response file.

   RETURN CODES: LA_OK, LA_EXPIRED, LA_FAIL, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_OFFLINE_RESPONSE_FILE
   LA_E_VM, LA_E_TIME, LA_E_FILE_PATH, LA_E_OFFLINE_RESPONSE_FILE_EXPIRED
*/
func ActivateLicenseOffline(filePath string) int {
	cFilePath := goToCString(filePath)
	status := C.ActivateLicenseOffline(cFilePath)
	freeCString(cFilePath)
	return int(status)
}

/*
   FUNCTION: GenerateOfflineActivationRequest()

   PURPOSE: Generates the offline activation request needed for generating
   offline activation response in the dashboard.

   PARAMETERS:
   * filePath - path of the file for the offline request.

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_FILE_PERMISSION
*/
func GenerateOfflineActivationRequest(filePath string) int {
	cFilePath := goToCString(filePath)
	status := C.GenerateOfflineActivationRequest(cFilePath)
	freeCString(cFilePath)
	return int(status)
}

/*
   FUNCTION: DeactivateLicense()

   PURPOSE: Deactivates the license activation and frees up the corresponding activation
   slot by contacting the Cryptlex servers.

   This function should be executed at the time of de-registration, ideally on
   a button click.

   RETURN CODES: LA_OK, LA_E_DEACTIVATION_LIMIT, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME
   LA_E_LICENSE_KEY, LA_E_INET, LA_E_SERVER, LA_E_RATE_LIMIT, LA_E_TIME_MODIFIED
*/
func DeactivateLicense() int {
	status := C.DeactivateLicense()
	return int(status)
}

/*
   FUNCTION: GenerateOfflineDeactivationRequest()

   PURPOSE: Generates the offline deactivation request needed for deactivation of
   the license in the dashboard and deactivates the license locally.

   A valid offline deactivation file confirms that the license has been successfully
   deactivated on the user's machine.

   PARAMETERS:
   * filePath - path of the file for the offline request.

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_FILE_PERMISSION,
   LA_E_TIME, LA_E_TIME_MODIFIED
*/
func GenerateOfflineDeactivationRequest(filePath string) int {
	cFilePath := goToCString(filePath)
	status := C.GenerateOfflineDeactivationRequest(cFilePath)
	freeCString(cFilePath)
	return int(status)
}

/*
   FUNCTION: IsLicenseGenuine()

   PURPOSE: It verifies whether your app is genuinely activated or not. The verification is
   done locally by verifying the cryptographic digital signature fetched at the time of
   activation.

   After verifying locally, it schedules a server check in a separate thread. After the
   first server sync it periodically does further syncs at a frequency set for the license.

   In case server sync fails due to network error, and it continues to fail for fixed
   number of days (grace period), the function returns LA_GRACE_PERIOD_OVER instead of LA_OK.

   This function must be called on every start of your program to verify the activation
   of your app.

   RETURN CODES: LA_OK, LA_EXPIRED, LA_SUSPENDED, LA_GRACE_PERIOD_OVER, LA_FAIL,
   LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_TIME, LA_E_TIME_MODIFIED

   NOTE: If application was activated offline using ActivateLicenseOffline() function, you
   may want to set grace period to 0 to ignore grace period.
*/
func IsLicenseGenuine() int {
	status := C.IsLicenseGenuine()
	return int(status)
}

/*
   FUNCTION: IsLicenseValid()

   PURPOSE: It verifies whether your app is genuinely activated or not. The verification is
   done locally by verifying the cryptographic digital signature fetched at the time of
   activation.

   This is just an auxiliary function which you may use in some specific cases, when you
   want to skip the server sync.

   RETURN CODES: LA_OK, LA_EXPIRED, LA_SUSPENDED, LA_GRACE_PERIOD_OVER, LA_FAIL,
   LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_TIME, LA_E_TIME_MODIFIED

   NOTE: You may want to set grace period to 0 to ignore grace period.
*/
func IsLicenseValid() int {
	status := C.IsLicenseValid()
	return int(status)
}

/*
   FUNCTION: ActivateTrial()

   PURPOSE: Starts the verified trial in your application by contacting the
   Cryptlex servers.

   This function should be executed when your application starts first time on
   the user's computer, ideally on a button click.

   RETURN CODES: LA_OK, LA_TRIAL_EXPIRED, LA_FAIL, LA_E_PRODUCT_ID, LA_E_INET,
   LA_E_VM, LA_E_TIME, LA_E_SERVER, LA_E_CLIENT, LA_E_COUNTRY, LA_E_IP, LA_E_RATE_LIMIT
*/
func ActivateTrial() int {
	status := C.ActivateTrial()
	return int(status)
}

/*
   FUNCTION: ActivateTrialOffline()

   PURPOSE: Activates your trial using the offline activation response file.

   PARAMETERS:
   * filePath - path of the offline activation response file.

   RETURN CODES: LA_OK, LA_TRIAL_EXPIRED, LA_FAIL, LA_E_PRODUCT_ID, LA_E_OFFLINE_RESPONSE_FILE
   LA_E_VM, LA_E_TIME, LA_E_FILE_PATH, LA_E_OFFLINE_RESPONSE_FILE_EXPIRED
*/
func ActivateTrialOffline(filePath string) int {
	cFilePath := goToCString(filePath)
	status := C.ActivateTrialOffline(cFilePath)
	freeCString(cFilePath)
	return int(status)
}

/*
   FUNCTION: GenerateOfflineTrialActivationRequest()

   PURPOSE: Generates the offline trial activation request needed for generating
   offline trial activation response in the dashboard.

   PARAMETERS:
   * filePath - path of the file for the offline request.

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_FILE_PERMISSION
*/
func GenerateOfflineTrialActivationRequest(filePath string) int {
	cFilePath := goToCString(filePath)
	status := C.GenerateOfflineTrialActivationRequest(cFilePath)
	freeCString(cFilePath)
	return int(status)
}

/*
   FUNCTION: IsTrialGenuine()

   PURPOSE: It verifies whether trial has started and is genuine or not. The
   verification is done locally by verifying the cryptographic digital signature
   fetched at the time of trial activation.

   This function must be called on every start of your program during the trial period.

   RETURN CODES: LA_OK, LA_TRIAL_EXPIRED, LA_FAIL, LA_E_TIME, LA_E_PRODUCT_ID, LA_E_TIME_MODIFIED

*/
func IsTrialGenuine() int {
	status := C.IsTrialGenuine()
	return int(status)
}

/*
   FUNCTION: ActivateLocalTrial()

   PURPOSE: Starts the local(unverified) trial.

   This function should be executed when your application starts first time on
   the user's computer.

   PARAMETERS:
   * trialLength - trial length in days

   RETURN CODES: LA_OK, LA_LOCAL_TRIAL_EXPIRED, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME_MODIFIED

   NOTE: The function is only meant for local(unverified) trials.
*/
func ActivateLocalTrial(trialLength uint) int {
	cTrialLength := (C.uint)(trialLength)
	status := C.ActivateLocalTrial(cTrialLength)
	return int(status)
}

/*
   FUNCTION: IsLocalTrialGenuine()

   PURPOSE: It verifies whether trial has started and is genuine or not. The
   verification is done locally.

   This function must be called on every start of your program during the trial period.

   RETURN CODES: LA_OK, LA_LOCAL_TRIAL_EXPIRED, LA_FAIL, LA_E_PRODUCT_ID,
   LA_E_TIME_MODIFIED

   NOTE: The function is only meant for local(unverified) trials.
*/
func IsLocalTrialGenuine() int {
	status := C.IsLocalTrialGenuine()
	return int(status)
}

/*
   FUNCTION: ExtendLocalTrial()

   PURPOSE: Extends the local trial.

   PARAMETERS:
   * trialExtensionLength - number of days to extend the trial

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME_MODIFIED

   NOTE: The function is only meant for local(unverified) trials.
*/
func ExtendLocalTrial(trialExtensionLength uint) int {
	cTrialExtensionLength := (C.uint)(trialExtensionLength)
	status := C.ExtendLocalTrial(cTrialExtensionLength)
	return int(status)
}

/*
   FUNCTION: IncrementActivationMeterAttributeUses()

   PURPOSE: Increments the meter attribute uses of the activation.

   PARAMETERS:
   * name - name of the meter attribute
   * increment - the increment value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METER_ATTRIBUTE_NOT_FOUND,
   LA_E_INET, LA_E_TIME, LA_E_SERVER, LA_E_CLIENT, LA_E_METER_ATTRIBUTE_USES_LIMIT_REACHED,
   LA_E_AUTHENTICATION_FAILED, LA_E_COUNTRY, LA_E_IP, LA_E_RATE_LIMIT, LA_E_LICENSE_KEY

*/
func IncrementActivationMeterAttributeUses(name string, increment uint) int {
	cName := goToCString(name)
	cIncrement := (C.uint)(increment)
	status := C.IncrementActivationMeterAttributeUses(cName, cIncrement)
	freeCString(cName)
	return int(status)
}

/*
   FUNCTION: DecrementActivationMeterAttributeUses()

   PURPOSE: Decrements the meter attribute uses of the activation.

   PARAMETERS:
   * name - name of the meter attribute
   * decrement - the decrement value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METER_ATTRIBUTE_NOT_FOUND,
   LA_E_INET, LA_E_TIME, LA_E_SERVER, LA_E_CLIENT, LA_E_RATE_LIMIT, LA_E_LICENSE_KEY,
   LA_E_AUTHENTICATION_FAILED, LA_E_COUNTRY, LA_E_IP, LA_E_ACTIVATION_NOT_FOUND

   NOTE: If the decrement is more than the current uses, it resets the uses to 0.
*/
func DecrementActivationMeterAttributeUses(name string, decrement uint) int {
	cName := goToCString(name)
	cDecrement := (C.uint)(decrement)
	status := C.DecrementActivationMeterAttributeUses(cName, cDecrement)
	freeCString(cName)
	return int(status)
}

/*
   FUNCTION: ResetActivationMeterAttributeUses()

   PURPOSE: Resets the meter attribute uses consumed by the activation.

   PARAMETERS:
   * name - name of the meter attribute
   * decrement - the decrement value

   RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METER_ATTRIBUTE_NOT_FOUND,
   LA_E_INET, LA_E_TIME, LA_E_SERVER, LA_E_CLIENT, LA_E_RATE_LIMIT, LA_E_LICENSE_KEY,
   LA_E_AUTHENTICATION_FAILED, LA_E_COUNTRY, LA_E_IP, LA_E_ACTIVATION_NOT_FOUND
*/
func ResetActivationMeterAttributeUses(name string) int {
	cName := goToCString(name)
	status := C.ResetActivationMeterAttributeUses(cName)
	freeCString(cName)
	return int(status)
}

/*
   FUNCTION: Reset()

   PURPOSE: Resets the activation and trial data stored in the machine.

   This function is meant for developer testing only.

   RETURN CODES: LA_OK, LA_E_PRODUCT_ID

   NOTE: The function does not reset local(unverified) trial data.
*/
func Reset() int {
	status := C.Reset()
	return int(status)
}
