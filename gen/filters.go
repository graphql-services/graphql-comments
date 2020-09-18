package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func (f *CommentFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}
func (f *CommentFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("comments"), wheres, whereValues, havings, havingValues, joins)
}
func (f *CommentFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	return nil
}

func (f *CommentFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Reference != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" = ?")
		values = append(values, f.Reference)
	}

	if f.ReferenceNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" != ?")
		values = append(values, f.ReferenceNe)
	}

	if f.ReferenceGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" > ?")
		values = append(values, f.ReferenceGt)
	}

	if f.ReferenceLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" < ?")
		values = append(values, f.ReferenceLt)
	}

	if f.ReferenceGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" >= ?")
		values = append(values, f.ReferenceGte)
	}

	if f.ReferenceLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" <= ?")
		values = append(values, f.ReferenceLte)
	}

	if f.ReferenceIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" IN (?)")
		values = append(values, f.ReferenceIn)
	}

	if f.ReferenceLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ReferenceLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ReferencePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ReferencePrefix))
	}

	if f.ReferenceSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ReferenceSuffix))
	}

	if f.ReferenceNull != nil {
		if *f.ReferenceNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" IS NOT NULL")
		}
	}

	if f.ReferenceID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" = ?")
		values = append(values, f.ReferenceID)
	}

	if f.ReferenceIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" != ?")
		values = append(values, f.ReferenceIDNe)
	}

	if f.ReferenceIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" > ?")
		values = append(values, f.ReferenceIDGt)
	}

	if f.ReferenceIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" < ?")
		values = append(values, f.ReferenceIDLt)
	}

	if f.ReferenceIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" >= ?")
		values = append(values, f.ReferenceIDGte)
	}

	if f.ReferenceIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" <= ?")
		values = append(values, f.ReferenceIDLte)
	}

	if f.ReferenceIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" IN (?)")
		values = append(values, f.ReferenceIDIn)
	}

	if f.ReferenceIDNull != nil {
		if *f.ReferenceIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" IS NOT NULL")
		}
	}

	if f.Text != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" = ?")
		values = append(values, f.Text)
	}

	if f.TextNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" != ?")
		values = append(values, f.TextNe)
	}

	if f.TextGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" > ?")
		values = append(values, f.TextGt)
	}

	if f.TextLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" < ?")
		values = append(values, f.TextLt)
	}

	if f.TextGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" >= ?")
		values = append(values, f.TextGte)
	}

	if f.TextLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" <= ?")
		values = append(values, f.TextLte)
	}

	if f.TextIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" IN (?)")
		values = append(values, f.TextIn)
	}

	if f.TextLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TextLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TextPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TextPrefix))
	}

	if f.TextSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TextSuffix))
	}

	if f.TextNull != nil {
		if *f.TextNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}
func (f *CommentFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.ReferenceMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") = ?")
		values = append(values, f.ReferenceMin)
	}

	if f.ReferenceMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") = ?")
		values = append(values, f.ReferenceMax)
	}

	if f.ReferenceMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") != ?")
		values = append(values, f.ReferenceMinNe)
	}

	if f.ReferenceMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") != ?")
		values = append(values, f.ReferenceMaxNe)
	}

	if f.ReferenceMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") > ?")
		values = append(values, f.ReferenceMinGt)
	}

	if f.ReferenceMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") > ?")
		values = append(values, f.ReferenceMaxGt)
	}

	if f.ReferenceMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") < ?")
		values = append(values, f.ReferenceMinLt)
	}

	if f.ReferenceMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") < ?")
		values = append(values, f.ReferenceMaxLt)
	}

	if f.ReferenceMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") >= ?")
		values = append(values, f.ReferenceMinGte)
	}

	if f.ReferenceMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") >= ?")
		values = append(values, f.ReferenceMaxGte)
	}

	if f.ReferenceMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") <= ?")
		values = append(values, f.ReferenceMinLte)
	}

	if f.ReferenceMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") <= ?")
		values = append(values, f.ReferenceMaxLte)
	}

	if f.ReferenceMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") IN (?)")
		values = append(values, f.ReferenceMinIn)
	}

	if f.ReferenceMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") IN (?)")
		values = append(values, f.ReferenceMaxIn)
	}

	if f.ReferenceMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ReferenceMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ReferenceMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ReferenceMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ReferenceMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ReferenceMinPrefix))
	}

	if f.ReferenceMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ReferenceMaxPrefix))
	}

	if f.ReferenceMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ReferenceMinSuffix))
	}

	if f.ReferenceMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ReferenceMaxSuffix))
	}

	if f.ReferenceIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") = ?")
		values = append(values, f.ReferenceIDMin)
	}

	if f.ReferenceIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") = ?")
		values = append(values, f.ReferenceIDMax)
	}

	if f.ReferenceIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") != ?")
		values = append(values, f.ReferenceIDMinNe)
	}

	if f.ReferenceIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") != ?")
		values = append(values, f.ReferenceIDMaxNe)
	}

	if f.ReferenceIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") > ?")
		values = append(values, f.ReferenceIDMinGt)
	}

	if f.ReferenceIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") > ?")
		values = append(values, f.ReferenceIDMaxGt)
	}

	if f.ReferenceIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") < ?")
		values = append(values, f.ReferenceIDMinLt)
	}

	if f.ReferenceIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") < ?")
		values = append(values, f.ReferenceIDMaxLt)
	}

	if f.ReferenceIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") >= ?")
		values = append(values, f.ReferenceIDMinGte)
	}

	if f.ReferenceIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") >= ?")
		values = append(values, f.ReferenceIDMaxGte)
	}

	if f.ReferenceIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") <= ?")
		values = append(values, f.ReferenceIDMinLte)
	}

	if f.ReferenceIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") <= ?")
		values = append(values, f.ReferenceIDMaxLte)
	}

	if f.ReferenceIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") IN (?)")
		values = append(values, f.ReferenceIDMinIn)
	}

	if f.ReferenceIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") IN (?)")
		values = append(values, f.ReferenceIDMaxIn)
	}

	if f.TextMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") = ?")
		values = append(values, f.TextMin)
	}

	if f.TextMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") = ?")
		values = append(values, f.TextMax)
	}

	if f.TextMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") != ?")
		values = append(values, f.TextMinNe)
	}

	if f.TextMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") != ?")
		values = append(values, f.TextMaxNe)
	}

	if f.TextMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") > ?")
		values = append(values, f.TextMinGt)
	}

	if f.TextMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") > ?")
		values = append(values, f.TextMaxGt)
	}

	if f.TextMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") < ?")
		values = append(values, f.TextMinLt)
	}

	if f.TextMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") < ?")
		values = append(values, f.TextMaxLt)
	}

	if f.TextMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") >= ?")
		values = append(values, f.TextMinGte)
	}

	if f.TextMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") >= ?")
		values = append(values, f.TextMaxGte)
	}

	if f.TextMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") <= ?")
		values = append(values, f.TextMinLte)
	}

	if f.TextMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") <= ?")
		values = append(values, f.TextMaxLte)
	}

	if f.TextMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") IN (?)")
		values = append(values, f.TextMinIn)
	}

	if f.TextMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") IN (?)")
		values = append(values, f.TextMaxIn)
	}

	if f.TextMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TextMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TextMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TextMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TextMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TextMinPrefix))
	}

	if f.TextMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TextMaxPrefix))
	}

	if f.TextMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TextMinSuffix))
	}

	if f.TextMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TextMaxSuffix))
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *CommentFilterType) AndWith(f2 ...*CommentFilterType) *CommentFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &CommentFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *CommentFilterType) OrWith(f2 ...*CommentFilterType) *CommentFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &CommentFilterType{
		Or: append(_f2, f),
	}
}
