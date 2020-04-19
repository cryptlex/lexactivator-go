package lexactivator

/*
#cgo LDFLAGS: -L./libs -lLexActivator -Wl,-rpath,./libs -Wl,-rpath,./
#include "lexactivator/LexActivator.h"
#include <stdlib.h>
void licenseCallbackCgoGateway(int in); // Forward declaration.
*/
import "C"
import (
	"unsafe"
)

type CallbackType func(int)

var licenseCallbackFuncion CallbackType

const (
	LA_USER      uint = 0
	LA_SYSTEM    uint = 1
	LA_IN_MEMORY uint = 2
)

//export licenseCallbackWrapper
func licenseCallbackWrapper(status int) {
	if licenseCallbackFuncion != nil{
		licenseCallbackFuncion(status)
	}
}

// SetProductFile function as declared in lexactivator/LexActivator.h:82
func SetProductFile(filePath string) int {
	cFilePath := ToCString(filePath)
	status := C.SetProductFile(cFilePath)
	defer C.free(unsafe.Pointer(cFilePath))
	return int(status)
}

// SetProductData function as declared in lexactivator/LexActivator.h:103
func SetProductData(productData string) int {
	cProductData := ToCString(productData)
	status := C.SetProductData(cProductData)
	defer C.free(unsafe.Pointer(cProductData))
	return int(status)
}

// SetProductId function as declared in lexactivator/LexActivator.h:128
func SetProductId(productId string, flags uint) int {
	cProductId := ToCString(productId)
	cFlags := (C.uint)(flags)
	status := C.SetProductId(cProductId, cFlags)
	defer C.free(unsafe.Pointer(cProductId))
	return int(status)
}

// SetLicenseKey function as declared in lexactivator/LexActivator.h:140
func SetLicenseKey(licenseKey string) int {
	cLicenseKey := ToCString(licenseKey)
	status := C.SetLicenseKey(cLicenseKey)
	defer C.free(unsafe.Pointer(cLicenseKey))
	return int(status)
}

// SetLicenseUserCredential function as declared in lexactivator/LexActivator.h:156
func SetLicenseUserCredential(email string, password string) int {
	cEmail := ToCString(email)
	cPassword := ToCString(password)
	status := C.SetLicenseUserCredential(cEmail, cPassword)
	defer C.free(unsafe.Pointer(cEmail))
	defer C.free(unsafe.Pointer(cPassword))
	return int(status)
}

// SetLicenseCallback function as declared in lexactivator/LexActivator.h:175
func SetLicenseCallback(callbackFunction CallbackType) int {
	C.SetLicenseCallback((C.CallbackType)(unsafe.Pointer(C.licenseCallbackCgoGateway)))
	licenseCallbackFuncion = callbackFunction
	return 0
}

// SetActivationMetadata function as declared in lexactivator/LexActivator.h:191
func SetActivationMetadata(key string, value string) int {
	cKey := ToCString(key)
	cValue := ToCString(value)
	status := C.SetActivationMetadata(cKey, cValue)
	defer C.free(unsafe.Pointer(cKey))
	defer C.free(unsafe.Pointer(cValue))
	return int(status)
}

// SetTrialActivationMetadata function as declared in lexactivator/LexActivator.h:208
func SetTrialActivationMetadata(key string, value string) int {
	cKey := ToCString(key)
	cValue := ToCString(value)
	status := C.SetTrialActivationMetadata(cKey, cValue)
	defer C.free(unsafe.Pointer(cKey))
	defer C.free(unsafe.Pointer(cValue))
	return int(status)
}

// SetAppVersion function as declared in lexactivator/LexActivator.h:223
func SetAppVersion(appVersion string) int {
	cAppVersion := ToCString(appVersion)
	status := C.SetAppVersion(cAppVersion)
	defer C.free(unsafe.Pointer(cAppVersion))
	return int(status)
}

// SetOfflineActivationRequestMeterAttributeUses function as declared in lexactivator/LexActivator.h:240
func SetOfflineActivationRequestMeterAttributeUses(name string, uses uint) int {
	cName := ToCString(name)
	cUses := (C.uint)(uses)
	status := C.SetOfflineActivationRequestMeterAttributeUses(cName, cUses)
	defer C.free(unsafe.Pointer(cName))
	return int(status)
}

// SetNetworkProxy function as declared in lexactivator/LexActivator.h:262
func SetNetworkProxy(proxy string) int {
	cProxy := ToCString(proxy)
	status := C.SetNetworkProxy(cProxy)
	defer C.free(unsafe.Pointer(cProxy))
	return int(status)
}

// SetCryptlexHost function as declared in lexactivator/LexActivator.h:275
func SetCryptlexHost(host string) int {
	cHost := ToCString(host)
	status := C.SetCryptlexHost(cHost)
	defer C.free(unsafe.Pointer(cHost))
	return int(status)
}

// GetProductMetadata function as declared in lexactivator/LexActivator.h:291
func GetProductMetadata(key string, value *string) int {
	cKey := ToCString(key)
	var cValue = GetCArray()
	status := C.GetProductMetadata(cKey, &cValue[0], MaxCArrayLength)
	*value = C.GoStringN(&cValue[0], MaxGoArrayLength)
	defer C.free(unsafe.Pointer(cKey))
	return int(status)
}

// GetLicenseMetadata function as declared in lexactivator/LexActivator.h:305
func GetLicenseMetadata(key string, value *string) int {
	cKey := ToCString(key)
	var cValue = GetCArray()
	status := C.GetLicenseMetadata(cKey, &cValue[0], MaxCArrayLength)
	*value = C.GoStringN(&cValue[0], MaxGoArrayLength)
	defer C.free(unsafe.Pointer(cKey))
	return int(status)
}

// GetLicenseMeterAttribute function as declared in lexactivator/LexActivator.h:319
func GetLicenseMeterAttribute(name string, allowedUses *uint, totalUses *uint) int {
	cName := ToCString(name)
	var cAllowedUses C.uint
	var cTotalUses C.uint
	status := C.GetLicenseMeterAttribute(cName, &cAllowedUses, &cTotalUses)
	*allowedUses = uint(cAllowedUses)
	*totalUses = uint(cTotalUses)
	defer C.free(unsafe.Pointer(cName))
	return int(status)
}

// GetLicenseKey function as declared in lexactivator/LexActivator.h:332
func GetLicenseKey(licenseKey *string) int {
	var cLicenseKey = GetCArray()
	status := C.GetLicenseKey(&cLicenseKey[0], MaxCArrayLength)
	*licenseKey = C.GoStringN(&cLicenseKey[0], MaxGoArrayLength)
	return int(status)
}

// GetLicenseExpiryDate function as declared in lexactivator/LexActivator.h:344
func GetLicenseExpiryDate(expiryDate *uint) int {
	var cExpiryDate C.uint
	status := C.GetLicenseExpiryDate(&cExpiryDate)
	*expiryDate = uint(cExpiryDate)
	return int(status)
}

// GetLicenseUserEmail function as declared in lexactivator/LexActivator.h:358
func GetLicenseUserEmail(email *string) int {
	var cEmail = GetCArray()
	status := C.GetLicenseUserEmail(&cEmail[0], MaxCArrayLength)
	*email = C.GoStringN(&cEmail[0], MaxGoArrayLength)
	return int(status)
}

// GetLicenseUserName function as declared in lexactivator/LexActivator.h:372
func GetLicenseUserName(name *string) int {
	var cName = GetCArray()
	status := C.GetLicenseUserName(&cName[0], MaxCArrayLength)
	*name = C.GoStringN(&cName[0], MaxGoArrayLength)
	return int(status)
}

// GetLicenseUserCompany function as declared in lexactivator/LexActivator.h:386
func GetLicenseUserCompany(company *string) int {
	var cCompany = GetCArray()
	status := C.GetLicenseUserCompany(&cCompany[0], MaxCArrayLength)
	*company = C.GoStringN(&cCompany[0], MaxGoArrayLength)
	return int(status)
}

// GetLicenseUserMetadata function as declared in lexactivator/LexActivator.h:400
func GetLicenseUserMetadata(key string, value *string) int {
	cKey := ToCString(key)
	var cValue = GetCArray()
	status := C.GetLicenseUserMetadata(cKey, &cValue[0], MaxCArrayLength)
	*value = C.GoStringN(&cValue[0], MaxGoArrayLength)
	defer C.free(unsafe.Pointer(cKey))
	return int(status)
}

// GetLicenseType function as declared in lexactivator/LexActivator.h:414
func GetLicenseType(licenseType *string) int {
	var cLicenseType = GetCArray()
	status := C.GetLicenseType(&cLicenseType[0], MaxCArrayLength)
	*licenseType = C.GoStringN(&cLicenseType[0], MaxGoArrayLength)
	return int(status)
}

// GetActivationMetadata function as declared in lexactivator/LexActivator.h:428
func GetActivationMetadata(key string, value *string) int {
	cKey := ToCString(key)
	var cValue = GetCArray()
	status := C.GetActivationMetadata(cKey, &cValue[0], MaxCArrayLength)
	*value = C.GoStringN(&cValue[0], MaxGoArrayLength)
	defer C.free(unsafe.Pointer(cKey))
	return int(status)
}

// GetActivationMeterAttributeUses function as declared in lexactivator/LexActivator.h:441
func GetActivationMeterAttributeUses(name string, uses *uint) int {
	cName := ToCString(name)
	var cUses C.uint
	status := C.GetActivationMeterAttributeUses(cName, &cUses)
	*uses = uint(cUses)
	defer C.free(unsafe.Pointer(cName))
	return int(status)
}

// GetServerSyncGracePeriodExpiryDate function as declared in lexactivator/LexActivator.h:453
func GetServerSyncGracePeriodExpiryDate(expiryDate *uint) int {
	var cExpiryDate C.uint
	status := C.GetServerSyncGracePeriodExpiryDate(&cExpiryDate)
	*expiryDate = uint(cExpiryDate)
	return int(status)
}

// GetTrialActivationMetadata function as declared in lexactivator/LexActivator.h:468
func GetTrialActivationMetadata(key string, value *string) int {
	cKey := ToCString(key)
	var cValue = GetCArray()
	status := C.GetTrialActivationMetadata(cKey, &cValue[0], MaxCArrayLength)
	*value = C.GoStringN(&cValue[0], MaxGoArrayLength)
	defer C.free(unsafe.Pointer(cKey))
	return int(status)
}

// GetTrialExpiryDate function as declared in lexactivator/LexActivator.h:480
func GetTrialExpiryDate(trialExpiryDate *uint) int {
	var cTrialExpiryDate C.uint
	status := C.GetTrialExpiryDate(&cTrialExpiryDate)
	*trialExpiryDate = uint(cTrialExpiryDate)
	return int(status)
}

// GetTrialId function as declared in lexactivator/LexActivator.h:494
func GetTrialId(trialId *string) int {
	var cTrialId = GetCArray()
	status := C.GetTrialId(&cTrialId[0], MaxCArrayLength)
	*trialId = C.GoStringN(&cTrialId[0], MaxGoArrayLength)
	return int(status)
}

// GetLocalTrialExpiryDate function as declared in lexactivator/LexActivator.h:506
func GetLocalTrialExpiryDate(trialExpiryDate *uint) int {
	var cTrialExpiryDate C.uint
	status := C.GetLocalTrialExpiryDate(&cTrialExpiryDate)
	*trialExpiryDate = uint(cTrialExpiryDate)
	return int(status)
}

// CheckForReleaseUpdate function as declared in lexactivator/LexActivator.h:524
// func CheckForReleaseUpdate(Platform string, Version string, Channel string, ReleaseUpdateCallback CallbackType) int {
// 	cPlatform, _ := unpackArgSCSTRTYPE(Platform)
// 	cVersion, _ := unpackArgSCSTRTYPE(Version)
// 	cChannel, _ := unpackArgSCSTRTYPE(Channel)
// 	cReleaseUpdateCallback, _ := ReleaseUpdateCallback.PassValue()
// 	__ret := C.CheckForReleaseUpdate(cPlatform, cVersion, cChannel, cReleaseUpdateCallback)
// 	packSCSTRTYPE(Channel, cChannel)
// 	packSCSTRTYPE(Version, cVersion)
// 	packSCSTRTYPE(Platform, cPlatform)
// 	__v := (int)(__ret)
// 	return __v
// }

// ActivateLicense function as declared in lexactivator/LexActivator.h:540
func ActivateLicense() int {
	status := C.ActivateLicense()
	return int(status)
}

// ActivateLicenseOffline function as declared in lexactivator/LexActivator.h:553
func ActivateLicenseOffline(filePath string) int {
	cFilePath := ToCString(filePath)
	status := C.ActivateLicenseOffline(cFilePath)
	defer C.free(unsafe.Pointer(cFilePath))
	return int(status)
}

// GenerateOfflineActivationRequest function as declared in lexactivator/LexActivator.h:566
func GenerateOfflineActivationRequest(filePath string) int {
	cFilePath := ToCString(filePath)
	status := C.GenerateOfflineActivationRequest(cFilePath)
	defer C.free(unsafe.Pointer(cFilePath))
	return int(status)
}

// DeactivateLicense function as declared in lexactivator/LexActivator.h:580
func DeactivateLicense() int {
	status := C.DeactivateLicense()
	return int(status)
}

// GenerateOfflineDeactivationRequest function as declared in lexactivator/LexActivator.h:597
func GenerateOfflineDeactivationRequest(filePath string) int {
	cFilePath := ToCString(filePath)
	status := C.GenerateOfflineDeactivationRequest(cFilePath)
	defer C.free(unsafe.Pointer(cFilePath))
	return int(status)
}

// IsLicenseGenuine function as declared in lexactivator/LexActivator.h:621
func IsLicenseGenuine() int {
	status := C.IsLicenseGenuine()
	return int(status)
}

// IsLicenseValid function as declared in lexactivator/LexActivator.h:638
func IsLicenseValid() int {
	status := C.IsLicenseValid()
	return int(status)
}

// ActivateTrial function as declared in lexactivator/LexActivator.h:652
func ActivateTrial() int {
	status := C.ActivateTrial()
	return int(status)
}

// ActivateTrialOffline function as declared in lexactivator/LexActivator.h:665
func ActivateTrialOffline(filePath string) int {
	cFilePath := ToCString(filePath)
	status := C.ActivateTrialOffline(cFilePath)
	defer C.free(unsafe.Pointer(cFilePath))
	return int(status)
}

// GenerateOfflineTrialActivationRequest function as declared in lexactivator/LexActivator.h:678
func GenerateOfflineTrialActivationRequest(filePath string) int {
	cFilePath := ToCString(filePath)
	status := C.GenerateOfflineTrialActivationRequest(cFilePath)
	defer C.free(unsafe.Pointer(cFilePath))
	return int(status)
}

// IsTrialGenuine function as declared in lexactivator/LexActivator.h:692
func IsTrialGenuine() int {
	status := C.IsTrialGenuine()
	return int(status)
}

// ActivateLocalTrial function as declared in lexactivator/LexActivator.h:709
func ActivateLocalTrial(trialLength uint) int {
	cTrialLength := (C.uint)(trialLength)
	status := C.ActivateLocalTrial(cTrialLength)
	return int(status)
}

// IsLocalTrialGenuine function as declared in lexactivator/LexActivator.h:724
func IsLocalTrialGenuine() int {
	status := C.IsLocalTrialGenuine()
	return int(status)
}

// ExtendLocalTrial function as declared in lexactivator/LexActivator.h:738
func ExtendLocalTrial(trialExtensionLength uint) int {
	cTrialExtensionLength := (C.uint)(trialExtensionLength)
	status := C.ExtendLocalTrial(cTrialExtensionLength)
	return int(status)
}

// IncrementActivationMeterAttributeUses function as declared in lexactivator/LexActivator.h:754
func IncrementActivationMeterAttributeUses(name string, increment uint) int {
	cName := ToCString(name)
	cIncrement := (C.uint)(increment)
	status := C.IncrementActivationMeterAttributeUses(cName, cIncrement)
	defer C.free(unsafe.Pointer(cName))
	return int(status)
}

// DecrementActivationMeterAttributeUses function as declared in lexactivator/LexActivator.h:771
func DecrementActivationMeterAttributeUses(name string, decrement uint) int {
	cName := ToCString(name)
	cDecrement := (C.uint)(decrement)
	status := C.DecrementActivationMeterAttributeUses(cName, cDecrement)
	defer C.free(unsafe.Pointer(cName))
	return int(status)
}

// ResetActivationMeterAttributeUses function as declared in lexactivator/LexActivator.h:786
func ResetActivationMeterAttributeUses(name string) int {
	cName := ToCString(name)
	status := C.ResetActivationMeterAttributeUses(cName)
	defer C.free(unsafe.Pointer(cName))
	return int(status)
}

// Reset function as declared in lexactivator/LexActivator.h:799
func Reset() int {
	status := C.Reset()
	return int(status)
}
