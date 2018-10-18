package compass_study_search

const url = "https://connpass.com/api/v1/event/?keyword=python"

type KeywordOd struct {
	keywords []string
}

type KeywordAnd struct {
	keywords []string
}

type Searcher interface {
	Get()
}

type Response struct {
	start  int   `json:"result_start"`
	perNum int   `json:"results_available"`
	count  int   `json:"results_returned"`
	events Event `json:"events"`
}

type Responser interface {
	Start() int
	PerNum() int
	Count() int
	Eventer
}

type Event struct {
	url              string                 `json:"event_url"`
	eventType        string                 `json:"event_type"`
	ownerNickName    string                 `json:"owner_nickname"`
	ownerId          int                    `json:"owner_id"`
	ownerDisplayName string                 `json:"owner_display_name"`
	series           map[string]interface{} `json:"series"`
	updatedAt        string                 `json:"updated_at"`
	lat              string                 `json:"lat"`
	lon              string                 `json:"lon"`
	startedAt        string                 `json:"started_at"`
	hashTag          string                 `json:"hash_tag"`
	title            string                 `json:"title"`
	id               int                    `json:"event_id"`
	waiting          int                    `json:"waiting"`
	limit            int                    `json:"limit"`
}

type Eventer interface {
	Url() string
	EventType() string
	OwnerNickName() string
	OwnerId() int
	OwnerDisplayName() string
	Series() map[string]interface{}
	UpdatedAt() string
	Lat() string
	Lon() string
	StartedAt() string
	HashTag() string
	Title() string
	Id() int
	Waiting() int
	Limit() int
}

func keywordOr(keywords []string) {

}

func keywordAnd() {

}
