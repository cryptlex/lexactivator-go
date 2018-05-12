package main

/*
#cgo LDFLAGS: -L./libs -lLexActivator
#include <LexActivator.h>
*/
import "C"
import (
	"fmt"
	"os"
)

func initData() {
	var status C.int
	//  status = C.SetProductFile(C.CString("ABSOLUTE_PATH_OF_PRODUCT.DAT_FILE"));
	status = C.SetProductData(C.CString("PASTE_CONTENT_OF_PRODUCT.DAT_FILE"))
	if C.LA_OK != status {
		fmt.Println("Error Code:", status)
		os.Exit(1)
	}

	status = C.SetProductId(C.CString("PASTE_PRODUCT_ID"), C.LA_USER)
	if C.LA_OK != status {
		fmt.Println("Error Code:", status)
		os.Exit(1)
	}

	status = C.SetAppVersion(C.CString("PASTE_YOUR_APP_VERION"))
	if C.LA_OK != status {
		fmt.Println("Error Code:", status)
		os.Exit(1)
	}
}

func activate() {
	var status C.int
	status = C.SetLicenseKey(C.CString("PASTE_LICENCE_KEY"))
	if C.LA_OK != status {
		fmt.Println("Error Code:", status)
		os.Exit(1)
	}

	status = C.SetActivationMetadata(C.CString("key1"), C.CString("value1"))
	if C.LA_OK != status {
		fmt.Println("Error Code:", status)
		os.Exit(1)
	}

	status = C.ActivateLicense()
	if C.LA_OK == status || C.LA_EXPIRED == status || C.LA_SUSPENDED == status {
		fmt.Println("License activated successfully:", status)
	} else {
		fmt.Println("License activation failed:", status)
	}
}

func activateTrial() {
	var status C.int
	status = C.SetTrialActivationMetadata(C.CString("key1"), C.CString("value1"))
	if C.LA_OK != status {
		fmt.Println("Error Code:", status)
		os.Exit(1)
	}

	status = C.ActivateTrial()
	if C.LA_OK == status {
		fmt.Println("Product trial activated successfully!")
	} else if C.LA_TRIAL_EXPIRED == status {
		fmt.Println("Product trial has expired!")
	} else {
		fmt.Println("Product trial activation failed:", status)
	}
}

func main() {
	initData()
	var status C.int
	status = C.IsLicenseGenuine()
	if C.LA_OK == status {
		var expiryDate C.uint32_t
		C.GetLicenseExpiryDate(&expiryDate)
		fmt.Println("License expiry timestamp:", expiryDate)
		fmt.Println("License is genuinely activated!")
	} else if C.LA_EXPIRED == status {
		fmt.Println("License is genuinely activated but has expired!")
	} else if C.LA_SUSPENDED == status {
		fmt.Println("License is genuinely activated but has been suspended!")
	} else if C.LA_GRACE_PERIOD_OVER == status {
		fmt.Println("License is genuinely activated but grace period is over!")
	} else {
		var trialStatus C.int
		trialStatus = C.IsTrialGenuine()
		if C.LA_OK == trialStatus {
			var trialExpiryDate C.uint32_t
			C.GetTrialExpiryDate(&trialExpiryDate)
			fmt.Println("Trial expiry timestamp:", trialExpiryDate)
		} else if C.LA_TRIAL_EXPIRED == trialStatus {
			fmt.Println("Trial has expired!")

			// Time to buy the license and activate the app
			activate()
		} else {
			fmt.Println("Either trial has not started or has been tampered!")
			// Activating the trial
			activateTrial()
		}
	}
}
