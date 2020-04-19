package lexactivator

/*

#include <stdio.h>

// The gateway function
void licenseCallbackCgoGateway(int status)
{
	void licenseCallbackWrapper(int);
	licenseCallbackWrapper(status);
}
*/
import "C"
