package lexactivator

type ReleaseFile struct {
	Name       string `json:"name"`
	URL        string `json:"url"`
	Size       int    `json:"size"`
	Downloads  int    `json:"downloads"`
	Extension  string `json:"extension"`
	Checksum   string `json:"checksum"`
	Secured    bool   `json:"secured"`
	ReleaseID  string `json:"releaseId"`
	ID         string `json:"id"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type Release struct {
    Published   bool      	  `json:"published"`
    Private     bool       	  `json:"private"`
    ProductID   string    	  `json:"productId"`
    Name        string    	  `json:"name"`
    Channel     string    	  `json:"channel"`
    Version     string    	  `json:"version"`
    Platform    string    	  `json:"platform"`
    Platforms   []string  	  `json:"platforms"`
    Notes       string    	  `json:"notes"`
    TotalFiles  int       	  `json:"totalFiles"`
    PublishedAt string    	  `json:"publishedAt"`
    Files       []ReleaseFile `json:"files"`
    ID          string    	  `json:"id"`
    CreatedAt   string 		  `json:"createdAt"`
    UpdatedAt   string 		  `json:"updatedAt"`
}