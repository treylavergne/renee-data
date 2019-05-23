package shared

// this is a data representation of how each individual member is stored
type Member struct {
	Account             string              `dynamo:"Account" json:"Account,omitempty"`
	Organization        string              `dynamo:"Organization" json:"Organization,omitempty"`
	Campus              string              `dynamo:"Campus" json:"Campus,omitempty"`
	AgeGroup            string              `dynamo:"AgeGroup,omitempty" json:"AgeGroup,omitempty"`
	ShirtSize           string              `dynamo:"ShirtSize,omitempty" json:"ShirtSize,omitempty"`
	Baptism             Baptism             `dynamo:"Baptism,omitempty" json:"Baptism,omitempty"`
	Birthday            string              `dynamo:"Birthday,omitempty" json:"Birthday,omitempty"`
	Email               string              `dynamo:"Email,omitempty" json:"Email,omitemptyv"`
	FirstName           string              `dynamo:"FirstName,omitempty" json:"FirstName,omitempty"`
	FullName            string              `dynamo:"FullName,omitempty" json:"FullName,omitempty"`
	Gender              string              `dynamo:"Gender,omitempty" json:"Gender,omitempty"`
	LastName            string              `dynamo:"LastName,omitempty" json:"LastName,omitempty"`
	CreatedAt           int64               `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	LastUpdated         int64               `dynamo:"LastUpdated,omitempty" json:"LastUpdated,omitempty"`
	MembershipType      string              `dynamo:"MembershipType,omitempty" json:"MembershipType,omitempty"`
	PhoneNumber         string              `dynamo:"PhoneNumber,omitempty" json:"PhoneNumber,omitempty"`
	ReceiveEmail        bool                `dynamo:"ReceiveEmail,omitempty" json:"ReceiveEmail,omitempty"`
	SlackID             string              `dynamo:"SlackID,omitempty" json:"SlackID,omitempty"`
	UUID                string              `dynamo:"UUID,omitempty" json:"UUID,omitempty"`
	NewCreation         NewCreation         `dynamo:"NewCreation,omitempty" json:"NewCreation,omitempty"`
	FirstDecision       bool                `dynamo:"FirstDecision,omitempty" json:"FirstDecision,omitempty"`
	Rededication        bool                `dynamo:"Rededication,omitempty" json:"Rededication,omitempty"`
	Volunteer           Volunteer           `dynamo:"Volunteer,omitempty" json:"Volunteer,omitempty"`
	InterestGroup       bool                `dynamo:"InterestGroup,omitempty" json:"InterestGroup,omitempty"`
	InterestGroups      []string            `dynamo:"InterestGroups,omitempty" json:"InterestGroups,omitempty"`
	GrowGroup           bool                `dynamo:"GrowGroup,omitempty" json:"GrowGroup,omitempty"`
	GrowGroups          []string            `dynamo:"GrowGroups,omitempty" json:"GrowGroups,omitempty"`
	ThisIsHome          ThisIsHome          `dynamo:"ThisIsHome,omitempty" json:"ThisIsHome,omitempty"`
	DiscoverYourPurpose DiscoverYourPurpose `dynamo:"DiscoverYourPurpose,omitempty" json:"DiscoverYourPurpose,omitempty"`
	SpiritualGifts      SpiritualGiftsData  `dynamo:"SpiritualGifts,omitempty" json:"SpiritualGifts,omitempty"`
}

// this is how an tenant account is stored to determine what org/campus renee belongs to
type Account struct {
	Campus          string `json:"Campus"`
	ClerkNumber     string `json:"ClerkNumber"`
	GiveTrigger     string `json:"GiveTrigger"`
	HomeTrigger     string `json:"HomeTrigger"`
	NCTrigger       string `json:"NCTrigger"`
	Organization    string `json:"Organization"`
	SlackBotToken   string `json:"SlackBotToken"`
	SlackOAuthToken string `json:"SlackOAuthToken"`
	UUID            string `json:"UUID"`
}

type SpiritualGiftsData struct {
	Assessment []map[string]string `dynamo:"Assessment,omitempty" json:"Assessment,omitempty"`
	Gifts      map[string]int      `dynamo:"Gifts,omitempty" json:"Gifts,omitempty"`
}

type Baptism struct {
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}

type Group struct {
	Status bool     `dynamo:"Status,omitempty" json:"Status,omitempty"`
	Groups []Groups `dynamo:"Groups,omitempty" json:"Groups,omitempty"`
}

type Volunteer struct {
	Status bool       `dynamo:"Status,omitempty" json:"Status,omitempty"`
	Teams  []TeamData `dynamo:"Teams,omitempty" json:"Teams,omitempty"`
}

type TeamData struct {
	UUID      string `dynamo:"UUID" json:"UUID,omitempty"`
	GroupName string `dynamo:"GroupName,omitempty" json:"GroupName,omitempty"`
	CreatedAt int64  `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}

type NewCreation struct {
	Status        bool           `dynamo:"Status,omitempty" json:"Status,omitempty"`
	FirstDecision NCFirstTime    `dynamo:"FirstDecision,omitempty" json:"FirstDecision,omitempty"`
	Rededication  NCRededication `dynamo:"Rededication,omitempty" json:"Rededication,omitempty"`
}

type NCFirstTime struct {
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}

type NCRededication struct {
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}

type ThisIsHome struct {
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}

type DiscoverYourPurpose struct {
	Status    bool  `dynamo:"Status,omitempty" json:"Status,omitempty"`
	CreatedAt int64 `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
}

type Groups struct {
	Account      string        `dynamo:"Account,omitempty" json:"Account,omitempty"`
	Organization string        `dynamo:"Organization,omitempty" json:"Organization,omitempty"`
	Campus       string        `dynamo:"Campus,omitempty" json:"Campus,omitempty"`
	UUID         string        `dynamo:"UUID" json:"UUID"`
	Leaders      []Member      `dynamo:"Leaders" json:"Leaders,omitempty"`
	Attendees    []Member      `dynamo:"Attendees" json:"Attendees,omitempty"`
	GroupName    string        `dynamo:"GroupName" json:"GroupName,omitempty"`
	CreatedAt    int64         `dynamo:"CreatedAt" json:"CreatedAt,omitempty"`
	LastUpdated  int64         `dynamo:"LastUpdated" json:"LastUpdated,omitempty"`
	Details      GroupDetails  `dynamo:"Details" json:"Details,omitempty"`
	Capacity     GroupCapacity `dynamo:"Capacity" json:"Capacity,omitempty"`
}

type GroupDetails struct {
	Date        int64  `dynamo:"Date" json:"Date,omitempty"`
	Location    string `dynamo:"Location" json:"Location,omitempty"`
	Time        int64  `dynamo:"Time" json:"Time,omitempty"`
	Description string `dynamo:"Description" json:"Description,omitempty"`
}

type GroupCapacity struct {
	AtCapacity bool `dynamo:"AtCapacity" json:"AtCapacity,omitempty"`
	Limit      int  `dynamo:"Limit" json:"Limit,omitempty"`
	Openings   int  `dynamo:"Openings" json:"Openings,omitempty"`
	Unlimited  bool `dynamo:"Unlimited" json:"Unlimited,omitempty"`
}
