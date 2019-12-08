package gen

import (
	"context"

	"github.com/novacloudcz/graphql-orm/events"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *events.Event) error

	CreateComment     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Comment, err error)
	UpdateComment     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Comment, err error)
	DeleteComment     func(ctx context.Context, r *GeneratedResolver, id string) (item *Comment, err error)
	DeleteAllComments func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryComment      func(ctx context.Context, r *GeneratedResolver, opts QueryCommentHandlerOptions) (*Comment, error)
	QueryComments     func(ctx context.Context, r *GeneratedResolver, opts QueryCommentsHandlerOptions) (*CommentResultType, error)

	CommentCreatedByUser func(ctx context.Context, r *GeneratedResolver, obj *Comment) (res *User, err error)

	CreateCompany      func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Company, err error)
	UpdateCompany      func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Company, err error)
	DeleteCompany      func(ctx context.Context, r *GeneratedResolver, id string) (item *Company, err error)
	DeleteAllCompanies func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryCompany       func(ctx context.Context, r *GeneratedResolver, opts QueryCompanyHandlerOptions) (*Company, error)
	QueryCompanies     func(ctx context.Context, r *GeneratedResolver, opts QueryCompaniesHandlerOptions) (*CompanyResultType, error)

	CompanyEmployees func(ctx context.Context, r *GeneratedResolver, obj *Company) (res []*Person, err error)

	CreatePerson    func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Person, err error)
	UpdatePerson    func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Person, err error)
	DeletePerson    func(ctx context.Context, r *GeneratedResolver, id string) (item *Person, err error)
	DeleteAllPeople func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryPerson     func(ctx context.Context, r *GeneratedResolver, opts QueryPersonHandlerOptions) (*Person, error)
	QueryPeople     func(ctx context.Context, r *GeneratedResolver, opts QueryPeopleHandlerOptions) (*PersonResultType, error)

	PersonCompanies func(ctx context.Context, r *GeneratedResolver, obj *Person) (res []*Company, err error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },

		CreateComment:     CreateCommentHandler,
		UpdateComment:     UpdateCommentHandler,
		DeleteComment:     DeleteCommentHandler,
		DeleteAllComments: DeleteAllCommentsHandler,
		QueryComment:      QueryCommentHandler,
		QueryComments:     QueryCommentsHandler,

		CommentCreatedByUser: CommentCreatedByUserHandler,

		CreateCompany:      CreateCompanyHandler,
		UpdateCompany:      UpdateCompanyHandler,
		DeleteCompany:      DeleteCompanyHandler,
		DeleteAllCompanies: DeleteAllCompaniesHandler,
		QueryCompany:       QueryCompanyHandler,
		QueryCompanies:     QueryCompaniesHandler,

		CompanyEmployees: CompanyEmployeesHandler,

		CreatePerson:    CreatePersonHandler,
		UpdatePerson:    UpdatePersonHandler,
		DeletePerson:    DeletePersonHandler,
		DeleteAllPeople: DeleteAllPeopleHandler,
		QueryPerson:     QueryPersonHandler,
		QueryPeople:     QueryPeopleHandler,

		PersonCompanies: PersonCompaniesHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *events.EventController
}
