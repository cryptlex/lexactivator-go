// Copyright 2025 Cryptlex, LLP. All rights reserved.

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

#ifdef _WIN32
void newReleaseUpdateCallbackCgoGateway(int status, unsigned short* releaseJson, void* unused)
{
	void newReleaseUpdateCallbackWrapper(int, unsigned short*);
	newReleaseUpdateCallbackWrapper(status, releaseJson);
}
#else
void newReleaseUpdateCallbackCgoGateway(int status, const char* releaseJson, void* unused)
{
	void newReleaseUpdateCallbackWrapper(int, const char*);
	newReleaseUpdateCallbackWrapper(status, releaseJson);
}
#endif
*/
import "C"
