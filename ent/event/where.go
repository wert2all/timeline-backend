// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"
	"timeline/backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDate, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldTime, v))
}

// ShowTime applies equality check predicate on the "showTime" field. It's identical to ShowTimeEQ.
func ShowTime(v bool) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldShowTime, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDescription, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldURL, v))
}

// PreviewlyImageID applies equality check predicate on the "previewly_image_id" field. It's identical to PreviewlyImageIDEQ.
func PreviewlyImageID(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldPreviewlyImageID, v))
}

// Private applies equality check predicate on the "private" field. It's identical to PrivateEQ.
func Private(v bool) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldPrivate, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldCreatedAt, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldDate, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldType, vs...))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldTime, v))
}

// TimeContains applies the Contains predicate on the "time" field.
func TimeContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldTime, v))
}

// TimeHasPrefix applies the HasPrefix predicate on the "time" field.
func TimeHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldTime, v))
}

// TimeHasSuffix applies the HasSuffix predicate on the "time" field.
func TimeHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldTime, v))
}

// TimeIsNil applies the IsNil predicate on the "time" field.
func TimeIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldTime))
}

// TimeNotNil applies the NotNil predicate on the "time" field.
func TimeNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldTime))
}

// TimeEqualFold applies the EqualFold predicate on the "time" field.
func TimeEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldTime, v))
}

// TimeContainsFold applies the ContainsFold predicate on the "time" field.
func TimeContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldTime, v))
}

// ShowTimeEQ applies the EQ predicate on the "showTime" field.
func ShowTimeEQ(v bool) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldShowTime, v))
}

// ShowTimeNEQ applies the NEQ predicate on the "showTime" field.
func ShowTimeNEQ(v bool) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldShowTime, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldTitle))
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldTitle))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldDescription, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldURL, v))
}

// URLIsNil applies the IsNil predicate on the "url" field.
func URLIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldURL))
}

// URLNotNil applies the NotNil predicate on the "url" field.
func URLNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldURL))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldURL, v))
}

// PreviewlyImageIDEQ applies the EQ predicate on the "previewly_image_id" field.
func PreviewlyImageIDEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldPreviewlyImageID, v))
}

// PreviewlyImageIDNEQ applies the NEQ predicate on the "previewly_image_id" field.
func PreviewlyImageIDNEQ(v int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldPreviewlyImageID, v))
}

// PreviewlyImageIDIn applies the In predicate on the "previewly_image_id" field.
func PreviewlyImageIDIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldPreviewlyImageID, vs...))
}

// PreviewlyImageIDNotIn applies the NotIn predicate on the "previewly_image_id" field.
func PreviewlyImageIDNotIn(vs ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldPreviewlyImageID, vs...))
}

// PreviewlyImageIDGT applies the GT predicate on the "previewly_image_id" field.
func PreviewlyImageIDGT(v int) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldPreviewlyImageID, v))
}

// PreviewlyImageIDGTE applies the GTE predicate on the "previewly_image_id" field.
func PreviewlyImageIDGTE(v int) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldPreviewlyImageID, v))
}

// PreviewlyImageIDLT applies the LT predicate on the "previewly_image_id" field.
func PreviewlyImageIDLT(v int) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldPreviewlyImageID, v))
}

// PreviewlyImageIDLTE applies the LTE predicate on the "previewly_image_id" field.
func PreviewlyImageIDLTE(v int) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldPreviewlyImageID, v))
}

// PreviewlyImageIDIsNil applies the IsNil predicate on the "previewly_image_id" field.
func PreviewlyImageIDIsNil() predicate.Event {
	return predicate.Event(sql.FieldIsNull(FieldPreviewlyImageID))
}

// PreviewlyImageIDNotNil applies the NotNil predicate on the "previewly_image_id" field.
func PreviewlyImageIDNotNil() predicate.Event {
	return predicate.Event(sql.FieldNotNull(FieldPreviewlyImageID))
}

// PrivateEQ applies the EQ predicate on the "private" field.
func PrivateEQ(v bool) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldPrivate, v))
}

// PrivateNEQ applies the NEQ predicate on the "private" field.
func PrivateNEQ(v bool) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldPrivate, v))
}

// HasTimeline applies the HasEdge predicate on the "timeline" edge.
func HasTimeline() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TimelineTable, TimelineColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTimelineWith applies the HasEdge predicate on the "timeline" edge with a given conditions (other predicates).
func HasTimelineWith(preds ...predicate.Timeline) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newTimelineStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newTagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(sql.NotPredicates(p))
}
