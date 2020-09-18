package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

func (s CommentSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("comments"), sorts, joins)
}
func (s CommentSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Reference != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("reference"), Direction: s.Reference.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("reference") + ")", Direction: s.ReferenceMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("reference") + ")", Direction: s.ReferenceMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("referenceID"), Direction: s.ReferenceID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("referenceID") + ")", Direction: s.ReferenceIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("referenceID") + ")", Direction: s.ReferenceIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Text != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("text"), Direction: s.Text.String()}
		*sorts = append(*sorts, sort)
	}

	if s.TextMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("text") + ")", Direction: s.TextMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.TextMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("text") + ")", Direction: s.TextMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	return nil
}
