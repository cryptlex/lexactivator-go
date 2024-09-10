package lexactivator

type ReleaseFile struct {
	Size       int    `json:"size"`
	Downloads  int    `json:"downloads"`
    Secured    bool   `json:"secured"`
    Id         string `json:"id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Extension  string `json:"extension"`
	Checksum   string `json:"checksum"`
	ReleaseId  string `json:"releaseId"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type Metadata struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

type Release struct {
    TotalFiles  int       	  `json:"totalFiles"`
    IsPrivate   bool          `json:"isPrivate"`
    Published   bool      	  `json:"published"`
    Id          string    	  `json:"id"`
    CreatedAt   string 		  `json:"createdAt"`
    UpdatedAt   string 		  `json:"updatedAt"`
    Name        string    	  `json:"name"`
    Channel     string    	  `json:"channel"`
    Version     string    	  `json:"version"`
    Notes       string    	  `json:"notes"`
    PublishedAt string    	  `json:"publishedAt"`
    ProductId   string    	  `json:"productId"`
    Platforms   []string  	  `json:"platforms"`
    Files       []ReleaseFile `json:"files"`
}

type OrganizationAddress struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City		 string `json:"city"`
	State		 string `json:"state"`
	Country 	 string `json:"country"`
	PostalCode 	 string `json:"postalCode"`
}

type UserLicense struct {
     // The allowed activations of the license. A value of -1 indicates unlimited number of activations.
    AllowedActivations      int64       `json:"allowedActivations"`
    
     // The allowed activations of the license. A value of -1 indicates unlimited number of deactivations.
    AllowedDeactivations    int64       `json:"allowedDeactivations"`
    
    // The license key.
    Key                     string      `json:"key"`
    
    //The license type (node-locked or hosted-floating).
    Type                    string      `json:"type"`
    
    //License metadata with view_permission set to "user".
    Metadata                []Metadata  `json:"metadata"`
}