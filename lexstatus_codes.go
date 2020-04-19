package lexactivator

// int enumeration from lexactivator/int.h:4
const (
	LA_OK                                   int = 0
	LA_FAIL                                 int = 1
	LA_EXPIRED                              int = 20
	LA_SUSPENDED                            int = 21
	LA_GRACE_PERIOD_OVER                    int = 22
	LA_TRIAL_EXPIRED                        int = 25
	LA_LOCAL_TRIAL_EXPIRED                  int = 26
	LA_RELEASE_UPDATE_AVAILABLE             int = 30
	LA_RELEASE_NO_UPDATE_AVAILABLE          int = 31
	LA_E_FILE_PATH                          int = 40
	LA_E_PRODUCT_FILE                       int = 41
	LA_E_PRODUCT_DATA                       int = 42
	LA_E_PRODUCT_ID                         int = 43
	LA_E_SYSTEM_PERMISSION                  int = 44
	LA_E_FILE_PERMISSION                    int = 45
	LA_E_WMIC                               int = 46
	LA_E_TIME                               int = 47
	LA_E_INET                               int = 48
	LA_E_NET_PROXY                          int = 49
	LA_E_HOST_URL                           int = 50
	LA_E_BUFFER_SIZE                        int = 51
	LA_E_APP_VERSION_LENGTH                 int = 52
	LA_E_REVOKED                            int = 53
	LA_E_LICENSE_KEY                        int = 54
	LA_E_LICENSE_TYPE                       int = 55
	LA_E_OFFLINE_RESPONSE_FILE              int = 56
	LA_E_OFFLINE_RESPONSE_FILE_EXPIRED      int = 57
	LA_E_ACTIVATION_LIMIT                   int = 58
	LA_E_ACTIVATION_NOT_FOUND               int = 59
	LA_E_DEACTIVATION_LIMIT                 int = 60
	LA_E_TRIAL_NOT_ALLOWED                  int = 61
	LA_E_TRIAL_ACTIVATION_LIMIT             int = 62
	LA_E_MACHINE_FINGERPRINT                int = 63
	LA_E_METADATA_KEY_LENGTH                int = 64
	LA_E_METADATA_VALUE_LENGTH              int = 65
	LA_E_ACTIVATION_METADATA_LIMIT          int = 66
	LA_E_TRIAL_ACTIVATION_METADATA_LIMIT    int = 67
	LA_E_METADATA_KEY_NOT_FOUND             int = 68
	LA_E_TIME_MODIFIED                      int = 69
	LA_E_RELEASE_VERSION_FORMAT             int = 70
	LA_E_AUTHENTICATION_FAILED              int = 71
	LA_E_METER_ATTRIBUTE_NOT_FOUND          int = 72
	LA_E_METER_ATTRIBUTE_USES_LIMIT_REACHED int = 73
	LA_E_VM                                 int = 80
	LA_E_COUNTRY                            int = 81
	LA_E_IP                                 int = 82
	LA_E_RATE_LIMIT                         int = 90
	LA_E_SERVER                             int = 91
	LA_E_CLIENT                             int = 92
)
