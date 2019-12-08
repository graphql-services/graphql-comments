package src

import (
	"context"

	"github.com/graphql-services/graphql-comments/gen"
	"github.com/novacloudcz/graphql-orm/events"
)

func NewResolver(db *gen.DB, ec *events.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{&gen.GeneratedResolver{Handlers: handlers, DB: db, EventController: ec}}
}

type Resolver struct {
	*gen.GeneratedResolver
}

type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

func (r *MutationResolver) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
	ctx = gen.EnrichContextWithMutations(ctx, r.GeneratedResolver)
	err := fn(ctx)
	if err != nil {
		tx := gen.GetTransaction(ctx)
		tx.Rollback()
		return err
	}
	return gen.FinishMutationContext(ctx, r.GeneratedResolver)
}

type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{&gen.GeneratedMutationResolver{r.GeneratedResolver}}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{r.GeneratedResolver}}
}

type CommentResultTypeResolver struct {
	*gen.GeneratedCommentResultTypeResolver
}

func (r *Resolver) CommentResultType() gen.CommentResultTypeResolver {
	return &CommentResultTypeResolver{&gen.GeneratedCommentResultTypeResolver{r.GeneratedResolver}}
}

type CommentResolver struct {
	*gen.GeneratedCommentResolver
}

func (r *Resolver) Comment() gen.CommentResolver {
	return &CommentResolver{&gen.GeneratedCommentResolver{r.GeneratedResolver}}
}

type CompanyResultTypeResolver struct {
	*gen.GeneratedCompanyResultTypeResolver
}

func (r *Resolver) CompanyResultType() gen.CompanyResultTypeResolver {
	return &CompanyResultTypeResolver{&gen.GeneratedCompanyResultTypeResolver{r.GeneratedResolver}}
}

type CompanyResolver struct {
	*gen.GeneratedCompanyResolver
}

func (r *Resolver) Company() gen.CompanyResolver {
	return &CompanyResolver{&gen.GeneratedCompanyResolver{r.GeneratedResolver}}
}

type PersonResultTypeResolver struct {
	*gen.GeneratedPersonResultTypeResolver
}

func (r *Resolver) PersonResultType() gen.PersonResultTypeResolver {
	return &PersonResultTypeResolver{&gen.GeneratedPersonResultTypeResolver{r.GeneratedResolver}}
}

type PersonResolver struct {
	*gen.GeneratedPersonResolver
}

func (r *Resolver) Person() gen.PersonResolver {
	return &PersonResolver{&gen.GeneratedPersonResolver{r.GeneratedResolver}}
}

type UserResolver struct {
	*gen.GeneratedUserResolver
}
