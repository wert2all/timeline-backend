package event

import (
	"context"
	"time"

	"timeline/backend/domain/db/cursor"
	"timeline/backend/ent"
	"timeline/backend/ent/event"
	"timeline/backend/ent/timeline"

	"entgo.io/ent/dialect/sql"
)

type Repository interface {
	FindByID(ID int) (*ent.Event, error)
	CreateEvent(date time.Time, eventType event.Type) (*ent.Event, error)
	UpdateEvent(*ent.EventUpdateOne) (*ent.Event, error)
	Delete(context.Context, *ent.Event) error

	FindTimelineEventsByCursor(timelineID int, withPrivate bool, cursor *cursor.Cursor, limit int) ([]*ent.Event, error)
}

type repositoryImpl struct {
	client  *ent.Client
	context context.Context
}

func (r repositoryImpl) FindTimelineEventsByCursor(timelineID int, withPrivate bool, cursor *cursor.Cursor, limit int) ([]*ent.Event, error) {
	query := r.client.Event.Query().
		Where(event.HasTimelineWith(timeline.ID(timelineID))).
		Order(
			event.ByDate(sql.OrderDesc()),
			event.ByID(sql.OrderDesc()),
		).
		Limit(limit)

	if !withPrivate {
		query = query.Where(event.Private(false))
	}

	if cursor != nil {
		query = query.Where(func(s *sql.Selector) {
			s.Where(sql.P(func(b *sql.Builder) {
				b.WriteString("(").Ident(event.FieldDate).WriteByte(',').Ident(event.FieldID).WriteString(")").
					WriteOp(sql.OpLTE).
					WriteString("(").Arg(cursor.Timestamp).WriteString(", ").Arg(cursor.ID).WriteString(")")
			}))
		})
	}
	return query.All(r.context)
}

func (r repositoryImpl) Delete(ctx context.Context, event *ent.Event) error {
	return r.client.Event.DeleteOne(event).Exec(ctx)
}

func (r repositoryImpl) UpdateEvent(event *ent.EventUpdateOne) (*ent.Event, error) {
	return event.Save(r.context)
}

func (r repositoryImpl) CreateEvent(date time.Time, eventType event.Type) (*ent.Event, error) {
	return r.client.Event.Create().SetDate(date).SetType(eventType).Save(r.context)
}

func (r repositoryImpl) FindByID(ID int) (*ent.Event, error) {
	return r.client.Event.Query().Where(event.ID(ID)).Only(r.context)
}

func NewRepository(ctx context.Context, client *ent.Client) Repository {
	return repositoryImpl{context: ctx, client: client}
}
