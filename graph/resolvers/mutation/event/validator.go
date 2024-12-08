package event

import (
	"context"
	"net/url"
	"time"

	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/timeline"
	"timeline/backend/ent"
	entEvent "timeline/backend/ent/event"
	"timeline/backend/graph/model"
	"timeline/backend/lib/utils"

	domainUser "timeline/backend/domain/user"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/exp/maps"
)

type (
	BaseValidator interface {
		GetBaseValidEventInput(GQLInput, context.Context) (*BaseValidEventInput, error)
	}
	GQLInput struct {
		TimelineID  int
		Date        time.Time
		Type        *model.TimelineType
		Title       *string
		Description *string
		ShowTime    *bool
		URL         *string
		Tags        []string
	}

	BaseValidEventInput struct {
		Timeline    *ent.Timeline
		EventType   entEvent.Type
		Date        time.Time
		Title       string
		Description string
		ShowTime    bool
		Url         *url.URL
		Tags        []string
	}

	baseValidatorImpl struct {
		timelineModel timeline.Timeline
		userExtractor domainUser.UserExtractor
	}
)

func NewBaseValidator(timelineModel timeline.Timeline, userExtractor domainUser.UserExtractor) BaseValidator {
	return baseValidatorImpl{timelineModel: timelineModel, userExtractor: userExtractor}
}

func (b baseValidatorImpl) GetBaseValidEventInput(input GQLInput, ctx context.Context) (*BaseValidEventInput, error) {
	p := bluemonday.StrictPolicy()
	timelineEntity, err := b.timelineModel.GetTimeline(input.TimelineID)
	if err != nil {
		return nil, err
	}

	token := appContext.GetToken(ctx)
	user, err := b.userExtractor.ExtractUserFromToken(ctx, token)
	if err != nil {
		return nil, err
	}
	errCheckUser := b.timelineModel.CheckUserTimeline(timelineEntity, user.ID)

	if errCheckUser != nil {
		return nil, errCheckUser
	}

	var eventType entEvent.Type
	if input.Type == nil {
		eventType = entEvent.Type(model.TimelineTypeDefault)
	} else {
		eventType = entEvent.Type(input.Type.String())
	}

	groupedTags := make(map[string]string)
	for _, tagInput := range input.Tags {
		groupedTags[p.Sanitize(tagInput)] = p.Sanitize(tagInput)
	}
	var showTime bool
	if input.ShowTime != nil {
		showTime = *input.ShowTime
	} else {
		showTime = true
	}

	return &BaseValidEventInput{
		Timeline:    timelineEntity,
		EventType:   eventType,
		Date:        input.Date,
		Title:       p.Sanitize(utils.DerefString(input.Title)),
		Description: p.Sanitize(utils.DerefString(input.Description)),
		ShowTime:    showTime,
		Url:         b.getLink(input.URL),
		Tags:        maps.Values(groupedTags),
	}, nil
}

func (b baseValidatorImpl) getLink(link *string) *url.URL {
	linkString := utils.DerefString(link)
	if linkString != "" {
		link, err := url.ParseRequestURI(linkString)
		if err != nil {
			panic(err)
		}
		return link
	} else {
		return nil
	}
}
