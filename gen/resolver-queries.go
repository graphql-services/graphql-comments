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
	query := CommentQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &CommentResultType{
		EntityResultType: EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
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
		return nil, &NotFoundError{Entity: "Comment"}
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
	giOpts := GetItemsOptions{
		Alias:      TableName("comments"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, giOpts, &items)

	return
}

func (r *GeneratedCommentResultTypeResolver) Count(ctx context.Context, obj *CommentResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Comment{})
}

type GeneratedCommentResolver struct{ *GeneratedResolver }

func (r *GeneratedCommentResolver) CreatedByUser(ctx context.Context, obj *Comment) (res *User, err error) {
	return r.Handlers.CommentCreatedByUser(ctx, r.GeneratedResolver, obj)
}
func CommentCreatedByUserHandler(ctx context.Context, r *GeneratedResolver, obj *Comment) (res *User, err error) {

	err = fmt.Errorf("Resolver handler for CommentCreatedByUser not implemented")

	return
}

type QueryCompanyHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *CompanyFilterType
}

func (r *GeneratedQueryResolver) Company(ctx context.Context, id *string, q *string, filter *CompanyFilterType) (*Company, error) {
	opts := QueryCompanyHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryCompany(ctx, r.GeneratedResolver, opts)
}
func QueryCompanyHandler(ctx context.Context, r *GeneratedResolver, opts QueryCompanyHandlerOptions) (*Company, error) {
	query := CompanyQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &CompanyResultType{
		EntityResultType: EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("companies")+".id = ?", *opts.ID)
	}

	var items []*Company
	giOpts := GetItemsOptions{
		Alias:      TableName("companies"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Company"}
	}
	return items[0], err
}

type QueryCompaniesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*CompanySortType
	Filter *CompanyFilterType
}

func (r *GeneratedQueryResolver) Companies(ctx context.Context, offset *int, limit *int, q *string, sort []*CompanySortType, filter *CompanyFilterType) (*CompanyResultType, error) {
	opts := QueryCompaniesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryCompanies(ctx, r.GeneratedResolver, opts)
}
func QueryCompaniesHandler(ctx context.Context, r *GeneratedResolver, opts QueryCompaniesHandlerOptions) (*CompanyResultType, error) {
	query := CompanyQueryFilter{opts.Q}

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

	return &CompanyResultType{
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

type GeneratedCompanyResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedCompanyResultTypeResolver) Items(ctx context.Context, obj *CompanyResultType) (items []*Company, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("companies"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, giOpts, &items)

	return
}

func (r *GeneratedCompanyResultTypeResolver) Count(ctx context.Context, obj *CompanyResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Company{})
}

type GeneratedCompanyResolver struct{ *GeneratedResolver }

func (r *GeneratedCompanyResolver) Employees(ctx context.Context, obj *Company) (res []*Person, err error) {
	return r.Handlers.CompanyEmployees(ctx, r.GeneratedResolver, obj)
}
func CompanyEmployeesHandler(ctx context.Context, r *GeneratedResolver, obj *Company) (res []*Person, err error) {

	items := []*Person{}
	err = r.DB.Query().Model(obj).Related(&items, "Employees").Error
	res = items

	return
}

func (r *GeneratedCompanyResolver) EmployeesIds(ctx context.Context, obj *Company) (ids []string, err error) {
	ids = []string{}

	items := []*Person{}
	err = r.DB.Query().Model(obj).Select(TableName("people")+".id").Related(&items, "Employees").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

type QueryPersonHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *PersonFilterType
}

func (r *GeneratedQueryResolver) Person(ctx context.Context, id *string, q *string, filter *PersonFilterType) (*Person, error) {
	opts := QueryPersonHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryPerson(ctx, r.GeneratedResolver, opts)
}
func QueryPersonHandler(ctx context.Context, r *GeneratedResolver, opts QueryPersonHandlerOptions) (*Person, error) {
	query := PersonQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &PersonResultType{
		EntityResultType: EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("people")+".id = ?", *opts.ID)
	}

	var items []*Person
	giOpts := GetItemsOptions{
		Alias:      TableName("people"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Person"}
	}
	return items[0], err
}

type QueryPeopleHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*PersonSortType
	Filter *PersonFilterType
}

func (r *GeneratedQueryResolver) People(ctx context.Context, offset *int, limit *int, q *string, sort []*PersonSortType, filter *PersonFilterType) (*PersonResultType, error) {
	opts := QueryPeopleHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPeople(ctx, r.GeneratedResolver, opts)
}
func QueryPeopleHandler(ctx context.Context, r *GeneratedResolver, opts QueryPeopleHandlerOptions) (*PersonResultType, error) {
	query := PersonQueryFilter{opts.Q}

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

	return &PersonResultType{
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

type GeneratedPersonResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedPersonResultTypeResolver) Items(ctx context.Context, obj *PersonResultType) (items []*Person, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("people"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, giOpts, &items)

	return
}

func (r *GeneratedPersonResultTypeResolver) Count(ctx context.Context, obj *PersonResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Person{})
}

type GeneratedPersonResolver struct{ *GeneratedResolver }

func (r *GeneratedPersonResolver) Companies(ctx context.Context, obj *Person) (res []*Company, err error) {
	return r.Handlers.PersonCompanies(ctx, r.GeneratedResolver, obj)
}
func PersonCompaniesHandler(ctx context.Context, r *GeneratedResolver, obj *Person) (res []*Company, err error) {

	items := []*Company{}
	err = r.DB.Query().Model(obj).Related(&items, "Companies").Error
	res = items

	return
}

func (r *GeneratedPersonResolver) CompaniesIds(ctx context.Context, obj *Person) (ids []string, err error) {
	ids = []string{}

	items := []*Company{}
	err = r.DB.Query().Model(obj).Select(TableName("companies")+".id").Related(&items, "Companies").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}
