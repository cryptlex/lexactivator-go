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
*/
import "C"
