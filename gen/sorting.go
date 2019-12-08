package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

func (s CommentSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("comments"), sorts, joins)
}
func (s CommentSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("id")+" "+s.ID.String())
	}

	if s.Reference != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("reference")+" "+s.Reference.String())
	}

	if s.ReferenceID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("referenceID")+" "+s.ReferenceID.String())
	}

	if s.Text != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("text")+" "+s.Text.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdAt")+" "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdBy")+" "+s.CreatedBy.String())
	}

	return nil
}

func (s CompanySortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("companies"), sorts, joins)
}
func (s CompanySortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("id")+" "+s.ID.String())
	}

	if s.Name != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("name")+" "+s.Name.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdAt")+" "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdBy")+" "+s.CreatedBy.String())
	}

	if s.Employees != nil {
		_alias := alias + "_employees"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("company_employees"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("company_id")+" LEFT JOIN "+dialect.Quote(TableName("people"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employee_id")+" = "+dialect.Quote(_alias)+".id")
		err := s.Employees.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s PersonSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("people"), sorts, joins)
}
func (s PersonSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("id")+" "+s.ID.String())
	}

	if s.Name != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("name")+" "+s.Name.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdAt")+" "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdBy")+" "+s.CreatedBy.String())
	}

	if s.Companies != nil {
		_alias := alias + "_companies"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("company_employees"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employee_id")+" LEFT JOIN "+dialect.Quote(TableName("companies"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("company_id")+" = "+dialect.Quote(_alias)+".id")
		err := s.Companies.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
