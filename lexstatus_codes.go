// Copyright 2020 Cryptlex, LLC. All rights reserved.

package lexactivator

// int enumeration from lexactivator/int.h int =4
const (

	// Success code.
	LA_OK int = 0

	// Failure code.
	LA_FAIL int = 1

	// The license has expired or system time has been tampered with. Ensure your date and time settings are correct.
	LA_EXPIRED int = 20

	// The license has been suspended.
	LA_SUSPENDED int = 21

	// The grace period for server sync is over.
	LA_GRACE_PERIOD_OVER int = 22

	// The trial has expired or system time has been tampered with. Ensure your date and time settings are correct.
	LA_TRIAL_EXPIRED int = 25

	// The local trial has expired or system time has been tampered
	// with. Ensure your date and time settings are correct.
	LA_LOCAL_TRIAL_EXPIRED int = 26

	// A new update is available for the product. This means a new release has been published for the product.
	LA_RELEASE_UPDATE_AVAILABLE int = 30

	// No new update is available for the product. The current version is latest.
	LA_RELEASE_NO_UPDATE_AVAILABLE int = 31

	// Invalid file path.
	LA_E_FILE_PATH int = 40

	// Invalid or corrupted product file.
	LA_E_PRODUCT_FILE int = 41

	// Invalid product data.
	LA_E_PRODUCT_DATA int = 42

	// The product id is incorrect.
	LA_E_PRODUCT_ID int = 43

	// Insufficient system permissions. Occurs when LA_SYSTEM flag is used
	// but application is not run with admin privileges.
	LA_E_SYSTEM_PERMISSION int = 44

	// No permission to write to file.
	LA_E_FILE_PERMISSION int = 45

	// Fingerprint couldn't be generated because Windows Management Instrumentation (WMI)
	// service has been disabled. This error is specific to Windows only.
	LA_E_WMIC int = 46

	// The difference between the network time and the system time is
	// more than allowed clock offset.
	LA_E_TIME int = 47

	// Failed to connect to the server due to network error.
	LA_E_INET int = 48

	// Invalid network proxy.
	LA_E_NET_PROXY int = 49

	// Invalid Cryptlex host url.
	LA_E_HOST_URL int = 50

	// The buffer size was smaller than required.
	LA_E_BUFFER_SIZE int = 51

	// App version length is more than 256 characters.
	LA_E_APP_VERSION_LENGTH int = 52

	// The license has been revoked.
	LA_E_REVOKED int = 53

	// Invalid license key.
	LA_E_LICENSE_KEY int = 54

	// Invalid license type. Make sure floating license is not being used.
	LA_E_LICENSE_TYPE int = 55

	// Invalid offline activation response file.
	LA_E_OFFLINE_RESPONSE_FILE int = 56

	// The offline activation response has expired.
	LA_E_OFFLINE_RESPONSE_FILE_EXPIRED int = 57

	// The license has reached it's allowed activations limit.
	LA_E_ACTIVATION_LIMIT int = 58

	// The license activation was deleted on the server.
	LA_E_ACTIVATION_NOT_FOUND int = 59

	// The license has reached it's allowed deactivations limit.
	LA_E_DEACTIVATION_LIMIT int = 60

	// Trial not allowed for the product.
	LA_E_TRIAL_NOT_ALLOWED int = 61

	// Your account has reached it's trial activations limit.
	LA_E_TRIAL_ACTIVATION_LIMIT int = 62

	// Machine fingerprint has changed since activation.
	LA_E_MACHINE_FINGERPRINT int = 63

	// Metadata key length is more than 256 characters.
	LA_E_METADATA_KEY_LENGTH int = 64

	// Metadata value length is more than 256 characters.
	LA_E_METADATA_VALUE_LENGTH int = 65

	// The license has reached it's metadata fields limit.
	LA_E_ACTIVATION_METADATA_LIMIT int = 66

	// The trial has reached it's metadata fields limit.
	LA_E_TRIAL_ACTIVATION_METADATA_LIMIT int = 67

	// The metadata key does not exist.
	LA_E_METADATA_KEY_NOT_FOUND int = 68

	// The system time has been tampered (backdated).
	LA_E_TIME_MODIFIED int = 69

	// Invalid version format.
	LA_E_RELEASE_VERSION_FORMAT int = 70

	// Incorrect email or password.
	LA_E_AUTHENTICATION_FAILED int = 71

	// The meter attribute does not exist.
	LA_E_METER_ATTRIBUTE_NOT_FOUND int = 72

	// The meter attribute has reached it's usage limit.
	LA_E_METER_ATTRIBUTE_USES_LIMIT_REACHED int = 73

	// Application is being run inside a virtual machine / hypervisor,
	// and activation has been disallowed in the VM.
	LA_E_VM int = 80

	// Country is not allowed.
	LA_E_COUNTRY int = 81

	// IP address is not allowed.
	LA_E_IP int = 82

	// Rate limit for API has reached, try again later.
	LA_E_RATE_LIMIT int = 90

	// Server error.
	LA_E_SERVER int = 91

	// Client error.
	LA_E_CLIENT int = 92
)
