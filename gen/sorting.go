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
