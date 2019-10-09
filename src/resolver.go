package src

import (
	"context"

	"github.com/graphql-services/graphql-comments/gen"
	"github.com/novacloudcz/graphql-orm/events"
)

func New(db *gen.DB, ec *events.EventController) *Resolver {
	resolver := NewResolver(db, ec)

	resolver.Handlers.CommentCreatedByUser = func(ctx context.Context, r *gen.GeneratedCommentResolver, obj *gen.Comment) (res *gen.User, err error) {
		if obj.CreatedBy != nil {
			res = &gen.User{ID: *obj.CreatedBy}
		}
		return
	}

	return resolver
}
