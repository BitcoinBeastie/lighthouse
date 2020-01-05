package model

import (
	"encoding/json"
	"time"

	"github.com/lbryio/lighthouse/app/es/index"

	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/lbryio/lbry.go/v2/extras/null"
	"github.com/lbryio/lbry.go/v2/extras/util"

	"github.com/sirupsen/logrus"
	"gopkg.in/olivere/elastic.v6"
)

type Claim struct {
	ID                     uint64                 `json:"id,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	ClaimID                string                 `json:"claimId,omitempty"`
	Channel                *null.String           `json:"channel,omitempty"`
	ChannelClaimID         *null.String           `json:"channel_claim_id,omitempty"`
	BidState               string                 `json:"bid_state,omitempty"`
	EffectiveAmount        uint64                 `json:"effective_amount,omitempty"`
	TransactionTimeUnix    null.Uint64            `json:"-"` //Could be null in mempool
	TransactionTime        *null.Time             `json:"transaction_time,omitempty"`
	ChannelEffectiveAmount uint64                 `json:"certificate_amount,omitempty"`
	JSONValue              null.String            `json:"-"`
	Value                  map[string]interface{} `json:"value,omitempty"`
	Title                  *null.String           `json:"title,omitempty"`
	Description            *null.String           `json:"description,omitempty"`
	ReleaseTimeUnix        null.Uint64            `json:"-"`
	ReleaseTime            *null.Time             `json:"release_time,omitempty"`
	ContentType            *null.String           `json:"content_type,omitempty"`
	CertValid              bool                   `json:"cert_valid,omitempty"`
	ClaimType              *null.String           `json:"claim_type,omitempty"`
	FrameWidth             *null.Uint64           `json:"frame_width,omitempty"`
	FrameHeight            *null.Uint64           `json:"frame_height,omitempty"`
	Duration               *null.Uint64           `json:"duration,omitempty"`
	NSFW                   bool                   `json:"nsfw,omitempty"`
	ViewCnt                *null.Uint64           `json:"view_cnt,omitempty"`
	SubCnt                 *null.Uint64           `json:"sub_cnt,omitempty"`
	ThumbnailURL           *null.String           `json:"thumbnail_url,omitempty"`
	Fee                    *null.Float64          `json:"fee,omitempty"`
}

func NewClaim() Claim {
	return Claim{
		Channel:         util.PtrToNullString(""),
		ChannelClaimID:  util.PtrToNullString(""),
		TransactionTime: util.PtrToNullTime(time.Time{}),
		Title:           util.PtrToNullString(""),
		Description:     util.PtrToNullString(""),
		ReleaseTime:     util.PtrToNullTime(time.Time{}),
		ContentType:     util.PtrToNullString(""),
		ClaimType:       util.PtrToNullString(""),
		FrameWidth:      util.PtrToNullUint64(0),
		FrameHeight:     util.PtrToNullUint64(0),
		Duration:        util.PtrToNullUint64(0),
		ViewCnt:         util.PtrToNullUint64(0),
		SubCnt:          util.PtrToNullUint64(0),
		ThumbnailURL:    util.PtrToNullString(""),
		Fee:             util.PtrToNullFloat64(0),
	}
}

func (c Claim) Add(p *elastic.BulkProcessor) {
	r := elastic.NewBulkIndexRequest().Index(index.Claims).Type(index.ClaimType).Id(c.ClaimID).Doc(c)
	p.Add(r)
}

func (c Claim) Delete(p *elastic.BulkProcessor) {
	r := elastic.NewBulkDeleteRequest().Index(index.Claims).Type(index.ClaimType).Id(c.ClaimID)
	p.Add(r)
}

func (c Claim) Update(p *elastic.BulkProcessor) {
	r := elastic.NewBulkUpdateRequest().Index(index.Claims).Type(index.ClaimType).Id(c.ClaimID).Doc(c)
	p.Add(r)
}

func (c Claim) AsJSON() string {
	data, err := json.Marshal(&c)
	if err != nil {
		logrus.Error(errors.Err(err))
		return ""
	}
	return string(data)

}
