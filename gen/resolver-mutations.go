package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
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

	tx := GetTransaction(ctx)
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
func GetTransaction(ctx context.Context) *gorm.DB {
	return ctx.Value(KeyMutationTransaction).(*gorm.DB)
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
	tx := GetTransaction(ctx)

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

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateComment(ctx context.Context, id string, input map[string]interface{}) (item *Comment, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateComment(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateCommentHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Comment, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Comment{}
	now := time.Now()
	tx := GetTransaction(ctx)

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
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteCommentHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Comment, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Comment{}
	now := time.Now()
	tx := GetTransaction(ctx)

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

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteAllComments(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllComments(ctx, r.GeneratedResolver)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllCommentsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := GetTransaction(ctx)
	err := tx.Delete(&Comment{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) CreateCompany(ctx context.Context, input map[string]interface{}) (item *Company, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateCompany(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateCompanyHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Company, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Company{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Company",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes CompanyChanges
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

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["employeesIds"]; exists {
		items := []Person{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employees")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateCompany(ctx context.Context, id string, input map[string]interface{}) (item *Company, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateCompany(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateCompanyHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Company, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Company{}
	now := time.Now()
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Company",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes CompanyChanges
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

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["employeesIds"]; exists {
		items := []Person{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employees")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteCompany(ctx context.Context, id string) (item *Company, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteCompany(ctx, r.GeneratedResolver, id)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteCompanyHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Company, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Company{}
	now := time.Now()
	tx := GetTransaction(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Company",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("companies")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteAllCompanies(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllCompanies(ctx, r.GeneratedResolver)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllCompaniesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := GetTransaction(ctx)
	err := tx.Delete(&Company{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) CreatePerson(ctx context.Context, input map[string]interface{}) (item *Person, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreatePerson(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreatePersonHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Person, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Person{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Person",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PersonChanges
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

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["companiesIds"]; exists {
		items := []Company{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Companies")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdatePerson(ctx context.Context, id string, input map[string]interface{}) (item *Person, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdatePerson(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdatePersonHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Person, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Person{}
	now := time.Now()
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Person",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PersonChanges
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

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["companiesIds"]; exists {
		items := []Company{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Companies")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeletePerson(ctx context.Context, id string) (item *Person, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeletePerson(ctx, r.GeneratedResolver, id)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeletePersonHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Person, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Person{}
	now := time.Now()
	tx := GetTransaction(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Person",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("people")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteAllPeople(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllPeople(ctx, r.GeneratedResolver)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllPeopleHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := GetTransaction(ctx)
	err := tx.Delete(&Person{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}
