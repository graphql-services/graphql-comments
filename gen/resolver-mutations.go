package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/novacloudcz/graphql-orm/events"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

type MutationEvents struct {
	Events []events.Event
}

func EnrichContextWithMutations(ctx context.Context, r *GeneratedResolver) context.Context {
	_ctx := context.WithValue(ctx, KeyMutationTransaction, r.DB.db.Begin())
	_ctx = context.WithValue(_ctx, KeyMutationEvents, &MutationEvents{})
	return _ctx
}
func FinishMutationContext(ctx context.Context, r *GeneratedResolver) (err error) {
	s := GetMutationEventStore(ctx)

	for _, event := range s.Events {
		err = r.Handlers.OnEvent(ctx, r, &event)
		if err != nil {
			return
		}
	}

	tx := r.GetDB(ctx)
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, event := range s.Events {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func RollbackMutationContext(ctx context.Context, r *GeneratedResolver) error {
	tx := r.GetDB(ctx)
	return tx.Rollback().Error
}
func GetMutationEventStore(ctx context.Context) *MutationEvents {
	return ctx.Value(KeyMutationEvents).(*MutationEvents)
}
func AddMutationEvent(ctx context.Context, e events.Event) {
	s := GetMutationEventStore(ctx)
	s.Events = append(s.Events, e)
}

func (r *GeneratedMutationResolver) CreateComment(ctx context.Context, input map[string]interface{}) (item *Comment, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateComment(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateCommentHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Comment, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Comment{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Comment",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes CommentChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["reference"]; ok && (item.Reference != changes.Reference) {
		item.Reference = changes.Reference

		event.AddNewValue("reference", changes.Reference)
	}

	if _, ok := input["referenceID"]; ok && (item.ReferenceID != changes.ReferenceID) {
		item.ReferenceID = changes.ReferenceID

		event.AddNewValue("referenceID", changes.ReferenceID)
	}

	if _, ok := input["text"]; ok && (item.Text != changes.Text) && (item.Text == nil || changes.Text == nil || *item.Text != *changes.Text) {
		item.Text = changes.Text

		event.AddNewValue("text", changes.Text)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}
func (r *GeneratedMutationResolver) UpdateComment(ctx context.Context, id string, input map[string]interface{}) (item *Comment, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateComment(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateCommentHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Comment, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Comment{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Comment",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes CommentChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["reference"]; ok && (item.Reference != changes.Reference) {
		event.AddOldValue("reference", item.Reference)
		event.AddNewValue("reference", changes.Reference)
		item.Reference = changes.Reference
	}

	if _, ok := input["referenceID"]; ok && (item.ReferenceID != changes.ReferenceID) {
		event.AddOldValue("referenceID", item.ReferenceID)
		event.AddNewValue("referenceID", changes.ReferenceID)
		item.ReferenceID = changes.ReferenceID
	}

	if _, ok := input["text"]; ok && (item.Text != changes.Text) && (item.Text == nil || changes.Text == nil || *item.Text != *changes.Text) {
		event.AddOldValue("text", item.Text)
		event.AddNewValue("text", changes.Text)
		item.Text = changes.Text
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteComment(ctx context.Context, id string) (item *Comment, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteComment(ctx, r.GeneratedResolver, id)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteCommentHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Comment, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Comment{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Comment",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("comments")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}
func (r *GeneratedMutationResolver) DeleteAllComments(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllComments(ctx, r.GeneratedResolver)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllCommentsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&Comment{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}
