package gen

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

type QueryCommentHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *CommentFilterType
}

func (r *GeneratedQueryResolver) Comment(ctx context.Context, id *string, q *string, filter *CommentFilterType) (*Comment, error) {
	opts := QueryCommentHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryComment(ctx, r.GeneratedResolver, opts)
}
func QueryCommentHandler(ctx context.Context, r *GeneratedResolver, opts QueryCommentHandlerOptions) (*Comment, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := CommentQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &CommentResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("comments")+".id = ?", *opts.ID)
	}

	var items []*Comment
	giOpts := GetItemsOptions{
		Alias:      TableName("comments"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

type QueryCommentsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*CommentSortType
	Filter *CommentFilterType
}

func (r *GeneratedQueryResolver) Comments(ctx context.Context, offset *int, limit *int, q *string, sort []*CommentSortType, filter *CommentFilterType) (*CommentResultType, error) {
	opts := QueryCommentsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryComments(ctx, r.GeneratedResolver, opts)
}
func QueryCommentsHandler(ctx context.Context, r *GeneratedResolver, opts QueryCommentsHandlerOptions) (*CommentResultType, error) {
	query := CommentQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &CommentResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedCommentResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedCommentResultTypeResolver) Items(ctx context.Context, obj *CommentResultType) (items []*Comment, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("comments"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*Comment{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

func (r *GeneratedCommentResultTypeResolver) Count(ctx context.Context, obj *CommentResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("comments"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &Comment{})
}

type GeneratedCommentResolver struct{ *GeneratedResolver }

func (r *GeneratedCommentResolver) CreatedByUser(ctx context.Context, obj *Comment) (res *User, err error) {
	return r.Handlers.CommentCreatedByUser(ctx, r.GeneratedResolver, obj)
}
func CommentCreatedByUserHandler(ctx context.Context, r *GeneratedResolver, obj *Comment) (res *User, err error) {

	err = fmt.Errorf("Resolver handler for CommentCreatedByUser not implemented")

	return
}
