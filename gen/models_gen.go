// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gen

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CommentFilterType struct {
	And                []*CommentFilterType `json:"AND"`
	Or                 []*CommentFilterType `json:"OR"`
	ID                 *string              `json:"id"`
	IDMin              *string              `json:"idMin"`
	IDMax              *string              `json:"idMax"`
	IDNe               *string              `json:"id_ne"`
	IDMinNe            *string              `json:"idMin_ne"`
	IDMaxNe            *string              `json:"idMax_ne"`
	IDGt               *string              `json:"id_gt"`
	IDMinGt            *string              `json:"idMin_gt"`
	IDMaxGt            *string              `json:"idMax_gt"`
	IDLt               *string              `json:"id_lt"`
	IDMinLt            *string              `json:"idMin_lt"`
	IDMaxLt            *string              `json:"idMax_lt"`
	IDGte              *string              `json:"id_gte"`
	IDMinGte           *string              `json:"idMin_gte"`
	IDMaxGte           *string              `json:"idMax_gte"`
	IDLte              *string              `json:"id_lte"`
	IDMinLte           *string              `json:"idMin_lte"`
	IDMaxLte           *string              `json:"idMax_lte"`
	IDIn               []string             `json:"id_in"`
	IDMinIn            []string             `json:"idMin_in"`
	IDMaxIn            []string             `json:"idMax_in"`
	IDNull             *bool                `json:"id_null"`
	Reference          *string              `json:"reference"`
	ReferenceMin       *string              `json:"referenceMin"`
	ReferenceMax       *string              `json:"referenceMax"`
	ReferenceNe        *string              `json:"reference_ne"`
	ReferenceMinNe     *string              `json:"referenceMin_ne"`
	ReferenceMaxNe     *string              `json:"referenceMax_ne"`
	ReferenceGt        *string              `json:"reference_gt"`
	ReferenceMinGt     *string              `json:"referenceMin_gt"`
	ReferenceMaxGt     *string              `json:"referenceMax_gt"`
	ReferenceLt        *string              `json:"reference_lt"`
	ReferenceMinLt     *string              `json:"referenceMin_lt"`
	ReferenceMaxLt     *string              `json:"referenceMax_lt"`
	ReferenceGte       *string              `json:"reference_gte"`
	ReferenceMinGte    *string              `json:"referenceMin_gte"`
	ReferenceMaxGte    *string              `json:"referenceMax_gte"`
	ReferenceLte       *string              `json:"reference_lte"`
	ReferenceMinLte    *string              `json:"referenceMin_lte"`
	ReferenceMaxLte    *string              `json:"referenceMax_lte"`
	ReferenceIn        []string             `json:"reference_in"`
	ReferenceMinIn     []string             `json:"referenceMin_in"`
	ReferenceMaxIn     []string             `json:"referenceMax_in"`
	ReferenceLike      *string              `json:"reference_like"`
	ReferenceMinLike   *string              `json:"referenceMin_like"`
	ReferenceMaxLike   *string              `json:"referenceMax_like"`
	ReferencePrefix    *string              `json:"reference_prefix"`
	ReferenceMinPrefix *string              `json:"referenceMin_prefix"`
	ReferenceMaxPrefix *string              `json:"referenceMax_prefix"`
	ReferenceSuffix    *string              `json:"reference_suffix"`
	ReferenceMinSuffix *string              `json:"referenceMin_suffix"`
	ReferenceMaxSuffix *string              `json:"referenceMax_suffix"`
	ReferenceNull      *bool                `json:"reference_null"`
	ReferenceID        *string              `json:"referenceID"`
	ReferenceIDMin     *string              `json:"referenceIDMin"`
	ReferenceIDMax     *string              `json:"referenceIDMax"`
	ReferenceIDNe      *string              `json:"referenceID_ne"`
	ReferenceIDMinNe   *string              `json:"referenceIDMin_ne"`
	ReferenceIDMaxNe   *string              `json:"referenceIDMax_ne"`
	ReferenceIDGt      *string              `json:"referenceID_gt"`
	ReferenceIDMinGt   *string              `json:"referenceIDMin_gt"`
	ReferenceIDMaxGt   *string              `json:"referenceIDMax_gt"`
	ReferenceIDLt      *string              `json:"referenceID_lt"`
	ReferenceIDMinLt   *string              `json:"referenceIDMin_lt"`
	ReferenceIDMaxLt   *string              `json:"referenceIDMax_lt"`
	ReferenceIDGte     *string              `json:"referenceID_gte"`
	ReferenceIDMinGte  *string              `json:"referenceIDMin_gte"`
	ReferenceIDMaxGte  *string              `json:"referenceIDMax_gte"`
	ReferenceIDLte     *string              `json:"referenceID_lte"`
	ReferenceIDMinLte  *string              `json:"referenceIDMin_lte"`
	ReferenceIDMaxLte  *string              `json:"referenceIDMax_lte"`
	ReferenceIDIn      []string             `json:"referenceID_in"`
	ReferenceIDMinIn   []string             `json:"referenceIDMin_in"`
	ReferenceIDMaxIn   []string             `json:"referenceIDMax_in"`
	ReferenceIDNull    *bool                `json:"referenceID_null"`
	Text               *string              `json:"text"`
	TextMin            *string              `json:"textMin"`
	TextMax            *string              `json:"textMax"`
	TextNe             *string              `json:"text_ne"`
	TextMinNe          *string              `json:"textMin_ne"`
	TextMaxNe          *string              `json:"textMax_ne"`
	TextGt             *string              `json:"text_gt"`
	TextMinGt          *string              `json:"textMin_gt"`
	TextMaxGt          *string              `json:"textMax_gt"`
	TextLt             *string              `json:"text_lt"`
	TextMinLt          *string              `json:"textMin_lt"`
	TextMaxLt          *string              `json:"textMax_lt"`
	TextGte            *string              `json:"text_gte"`
	TextMinGte         *string              `json:"textMin_gte"`
	TextMaxGte         *string              `json:"textMax_gte"`
	TextLte            *string              `json:"text_lte"`
	TextMinLte         *string              `json:"textMin_lte"`
	TextMaxLte         *string              `json:"textMax_lte"`
	TextIn             []string             `json:"text_in"`
	TextMinIn          []string             `json:"textMin_in"`
	TextMaxIn          []string             `json:"textMax_in"`
	TextLike           *string              `json:"text_like"`
	TextMinLike        *string              `json:"textMin_like"`
	TextMaxLike        *string              `json:"textMax_like"`
	TextPrefix         *string              `json:"text_prefix"`
	TextMinPrefix      *string              `json:"textMin_prefix"`
	TextMaxPrefix      *string              `json:"textMax_prefix"`
	TextSuffix         *string              `json:"text_suffix"`
	TextMinSuffix      *string              `json:"textMin_suffix"`
	TextMaxSuffix      *string              `json:"textMax_suffix"`
	TextNull           *bool                `json:"text_null"`
	UpdatedAt          *time.Time           `json:"updatedAt"`
	UpdatedAtMin       *time.Time           `json:"updatedAtMin"`
	UpdatedAtMax       *time.Time           `json:"updatedAtMax"`
	UpdatedAtNe        *time.Time           `json:"updatedAt_ne"`
	UpdatedAtMinNe     *time.Time           `json:"updatedAtMin_ne"`
	UpdatedAtMaxNe     *time.Time           `json:"updatedAtMax_ne"`
	UpdatedAtGt        *time.Time           `json:"updatedAt_gt"`
	UpdatedAtMinGt     *time.Time           `json:"updatedAtMin_gt"`
	UpdatedAtMaxGt     *time.Time           `json:"updatedAtMax_gt"`
	UpdatedAtLt        *time.Time           `json:"updatedAt_lt"`
	UpdatedAtMinLt     *time.Time           `json:"updatedAtMin_lt"`
	UpdatedAtMaxLt     *time.Time           `json:"updatedAtMax_lt"`
	UpdatedAtGte       *time.Time           `json:"updatedAt_gte"`
	UpdatedAtMinGte    *time.Time           `json:"updatedAtMin_gte"`
	UpdatedAtMaxGte    *time.Time           `json:"updatedAtMax_gte"`
	UpdatedAtLte       *time.Time           `json:"updatedAt_lte"`
	UpdatedAtMinLte    *time.Time           `json:"updatedAtMin_lte"`
	UpdatedAtMaxLte    *time.Time           `json:"updatedAtMax_lte"`
	UpdatedAtIn        []*time.Time         `json:"updatedAt_in"`
	UpdatedAtMinIn     []*time.Time         `json:"updatedAtMin_in"`
	UpdatedAtMaxIn     []*time.Time         `json:"updatedAtMax_in"`
	UpdatedAtNull      *bool                `json:"updatedAt_null"`
	CreatedAt          *time.Time           `json:"createdAt"`
	CreatedAtMin       *time.Time           `json:"createdAtMin"`
	CreatedAtMax       *time.Time           `json:"createdAtMax"`
	CreatedAtNe        *time.Time           `json:"createdAt_ne"`
	CreatedAtMinNe     *time.Time           `json:"createdAtMin_ne"`
	CreatedAtMaxNe     *time.Time           `json:"createdAtMax_ne"`
	CreatedAtGt        *time.Time           `json:"createdAt_gt"`
	CreatedAtMinGt     *time.Time           `json:"createdAtMin_gt"`
	CreatedAtMaxGt     *time.Time           `json:"createdAtMax_gt"`
	CreatedAtLt        *time.Time           `json:"createdAt_lt"`
	CreatedAtMinLt     *time.Time           `json:"createdAtMin_lt"`
	CreatedAtMaxLt     *time.Time           `json:"createdAtMax_lt"`
	CreatedAtGte       *time.Time           `json:"createdAt_gte"`
	CreatedAtMinGte    *time.Time           `json:"createdAtMin_gte"`
	CreatedAtMaxGte    *time.Time           `json:"createdAtMax_gte"`
	CreatedAtLte       *time.Time           `json:"createdAt_lte"`
	CreatedAtMinLte    *time.Time           `json:"createdAtMin_lte"`
	CreatedAtMaxLte    *time.Time           `json:"createdAtMax_lte"`
	CreatedAtIn        []*time.Time         `json:"createdAt_in"`
	CreatedAtMinIn     []*time.Time         `json:"createdAtMin_in"`
	CreatedAtMaxIn     []*time.Time         `json:"createdAtMax_in"`
	CreatedAtNull      *bool                `json:"createdAt_null"`
	UpdatedBy          *string              `json:"updatedBy"`
	UpdatedByMin       *string              `json:"updatedByMin"`
	UpdatedByMax       *string              `json:"updatedByMax"`
	UpdatedByNe        *string              `json:"updatedBy_ne"`
	UpdatedByMinNe     *string              `json:"updatedByMin_ne"`
	UpdatedByMaxNe     *string              `json:"updatedByMax_ne"`
	UpdatedByGt        *string              `json:"updatedBy_gt"`
	UpdatedByMinGt     *string              `json:"updatedByMin_gt"`
	UpdatedByMaxGt     *string              `json:"updatedByMax_gt"`
	UpdatedByLt        *string              `json:"updatedBy_lt"`
	UpdatedByMinLt     *string              `json:"updatedByMin_lt"`
	UpdatedByMaxLt     *string              `json:"updatedByMax_lt"`
	UpdatedByGte       *string              `json:"updatedBy_gte"`
	UpdatedByMinGte    *string              `json:"updatedByMin_gte"`
	UpdatedByMaxGte    *string              `json:"updatedByMax_gte"`
	UpdatedByLte       *string              `json:"updatedBy_lte"`
	UpdatedByMinLte    *string              `json:"updatedByMin_lte"`
	UpdatedByMaxLte    *string              `json:"updatedByMax_lte"`
	UpdatedByIn        []string             `json:"updatedBy_in"`
	UpdatedByMinIn     []string             `json:"updatedByMin_in"`
	UpdatedByMaxIn     []string             `json:"updatedByMax_in"`
	UpdatedByNull      *bool                `json:"updatedBy_null"`
	CreatedBy          *string              `json:"createdBy"`
	CreatedByMin       *string              `json:"createdByMin"`
	CreatedByMax       *string              `json:"createdByMax"`
	CreatedByNe        *string              `json:"createdBy_ne"`
	CreatedByMinNe     *string              `json:"createdByMin_ne"`
	CreatedByMaxNe     *string              `json:"createdByMax_ne"`
	CreatedByGt        *string              `json:"createdBy_gt"`
	CreatedByMinGt     *string              `json:"createdByMin_gt"`
	CreatedByMaxGt     *string              `json:"createdByMax_gt"`
	CreatedByLt        *string              `json:"createdBy_lt"`
	CreatedByMinLt     *string              `json:"createdByMin_lt"`
	CreatedByMaxLt     *string              `json:"createdByMax_lt"`
	CreatedByGte       *string              `json:"createdBy_gte"`
	CreatedByMinGte    *string              `json:"createdByMin_gte"`
	CreatedByMaxGte    *string              `json:"createdByMax_gte"`
	CreatedByLte       *string              `json:"createdBy_lte"`
	CreatedByMinLte    *string              `json:"createdByMin_lte"`
	CreatedByMaxLte    *string              `json:"createdByMax_lte"`
	CreatedByIn        []string             `json:"createdBy_in"`
	CreatedByMinIn     []string             `json:"createdByMin_in"`
	CreatedByMaxIn     []string             `json:"createdByMax_in"`
	CreatedByNull      *bool                `json:"createdBy_null"`
}

type CommentSortType struct {
	ID             *ObjectSortType `json:"id"`
	IDMin          *ObjectSortType `json:"idMin"`
	IDMax          *ObjectSortType `json:"idMax"`
	Reference      *ObjectSortType `json:"reference"`
	ReferenceMin   *ObjectSortType `json:"referenceMin"`
	ReferenceMax   *ObjectSortType `json:"referenceMax"`
	ReferenceID    *ObjectSortType `json:"referenceID"`
	ReferenceIDMin *ObjectSortType `json:"referenceIDMin"`
	ReferenceIDMax *ObjectSortType `json:"referenceIDMax"`
	Text           *ObjectSortType `json:"text"`
	TextMin        *ObjectSortType `json:"textMin"`
	TextMax        *ObjectSortType `json:"textMax"`
	UpdatedAt      *ObjectSortType `json:"updatedAt"`
	UpdatedAtMin   *ObjectSortType `json:"updatedAtMin"`
	UpdatedAtMax   *ObjectSortType `json:"updatedAtMax"`
	CreatedAt      *ObjectSortType `json:"createdAt"`
	CreatedAtMin   *ObjectSortType `json:"createdAtMin"`
	CreatedAtMax   *ObjectSortType `json:"createdAtMax"`
	UpdatedBy      *ObjectSortType `json:"updatedBy"`
	UpdatedByMin   *ObjectSortType `json:"updatedByMin"`
	UpdatedByMax   *ObjectSortType `json:"updatedByMax"`
	CreatedBy      *ObjectSortType `json:"createdBy"`
	CreatedByMin   *ObjectSortType `json:"createdByMin"`
	CreatedByMax   *ObjectSortType `json:"createdByMax"`
}

type User struct {
	ID string `json:"id"`
}

type _Service struct {
	Sdl *string `json:"sdl"`
}

type ObjectSortType string

const (
	ObjectSortTypeAsc  ObjectSortType = "ASC"
	ObjectSortTypeDesc ObjectSortType = "DESC"
)

var AllObjectSortType = []ObjectSortType{
	ObjectSortTypeAsc,
	ObjectSortTypeDesc,
}

func (e ObjectSortType) IsValid() bool {
	switch e {
	case ObjectSortTypeAsc, ObjectSortTypeDesc:
		return true
	}
	return false
}

func (e ObjectSortType) String() string {
	return string(e)
}

func (e *ObjectSortType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ObjectSortType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ObjectSortType", str)
	}
	return nil
}

func (e ObjectSortType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
