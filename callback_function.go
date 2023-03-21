// Copyright 2023 Cryptlex, LLC. All rights reserved.

package lexactivator

/*
#include <stdio.h>

// The gateway functions
void licenseCallbackCgoGateway(int status)
{
	void licenseCallbackWrapper(int);
	licenseCallbackWrapper(status);
}

void releaseUpdateCallbackCgoGateway(int status)
{
	void releaseUpdateCallbackWrapper(int);
	releaseUpdateCallbackWrapper(status);
}

void newReleaseUpdateCallbackCgoGateway(int status, char *releaseJson)
{
	void newReleaseUpdateCallbackWrapper(int, char *);
	newReleaseUpdateCallbackWrapper(status, releaseJson);
}
*/
import "C"
