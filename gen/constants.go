package gen

type key int

const (
	KeyPrincipalID         key    = iota
	KeyLoaders             key    = iota
	KeyExecutableSchema    key    = iota
	KeyJWTClaims           key    = iota
	KeyMutationTransaction key    = iota
	KeyMutationEvents      key    = iota
	SchemaSDL              string = `scalar Time

type Query {
  comment(id: ID, q: String, filter: CommentFilterType): Comment
  comments(offset: Int, limit: Int = 30, q: String, sort: [CommentSortType!], filter: CommentFilterType): CommentResultType
}

type Mutation {
  createComment(input: CommentCreateInput!): Comment!
  updateComment(id: ID!, input: CommentUpdateInput!): Comment!
  deleteComment(id: ID!): Comment!
  deleteAllComments: Boolean!
}

enum ObjectSortType {
  ASC
  DESC
}

type Comment {
  id: ID!
  reference: String!
  referenceID: ID!
  text: String
  createdByUser: User
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
}

extend type User @key(fields: "id") {
  id: ID! @external
}

input CommentCreateInput {
  id: ID
  reference: String!
  referenceID: ID!
  text: String
}

input CommentUpdateInput {
  reference: String
  referenceID: ID
  text: String
}

input CommentSortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  reference: ObjectSortType
  referenceMin: ObjectSortType
  referenceMax: ObjectSortType
  referenceID: ObjectSortType
  referenceIDMin: ObjectSortType
  referenceIDMax: ObjectSortType
  text: ObjectSortType
  textMin: ObjectSortType
  textMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
}

input CommentFilterType {
  AND: [CommentFilterType!]
  OR: [CommentFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  reference: String
  referenceMin: String
  referenceMax: String
  reference_ne: String
  referenceMin_ne: String
  referenceMax_ne: String
  reference_gt: String
  referenceMin_gt: String
  referenceMax_gt: String
  reference_lt: String
  referenceMin_lt: String
  referenceMax_lt: String
  reference_gte: String
  referenceMin_gte: String
  referenceMax_gte: String
  reference_lte: String
  referenceMin_lte: String
  referenceMax_lte: String
  reference_in: [String!]
  referenceMin_in: [String!]
  referenceMax_in: [String!]
  reference_like: String
  referenceMin_like: String
  referenceMax_like: String
  reference_prefix: String
  referenceMin_prefix: String
  referenceMax_prefix: String
  reference_suffix: String
  referenceMin_suffix: String
  referenceMax_suffix: String
  reference_null: Boolean
  referenceID: ID
  referenceIDMin: ID
  referenceIDMax: ID
  referenceID_ne: ID
  referenceIDMin_ne: ID
  referenceIDMax_ne: ID
  referenceID_gt: ID
  referenceIDMin_gt: ID
  referenceIDMax_gt: ID
  referenceID_lt: ID
  referenceIDMin_lt: ID
  referenceIDMax_lt: ID
  referenceID_gte: ID
  referenceIDMin_gte: ID
  referenceIDMax_gte: ID
  referenceID_lte: ID
  referenceIDMin_lte: ID
  referenceIDMax_lte: ID
  referenceID_in: [ID!]
  referenceIDMin_in: [ID!]
  referenceIDMax_in: [ID!]
  referenceID_null: Boolean
  text: String
  textMin: String
  textMax: String
  text_ne: String
  textMin_ne: String
  textMax_ne: String
  text_gt: String
  textMin_gt: String
  textMax_gt: String
  text_lt: String
  textMin_lt: String
  textMax_lt: String
  text_gte: String
  textMin_gte: String
  textMax_gte: String
  text_lte: String
  textMin_lte: String
  textMax_lte: String
  text_in: [String!]
  textMin_in: [String!]
  textMax_in: [String!]
  text_like: String
  textMin_like: String
  textMax_like: String
  text_prefix: String
  textMin_prefix: String
  textMax_prefix: String
  text_suffix: String
  textMin_suffix: String
  textMax_suffix: String
  text_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
  createdBy_null: Boolean
}

type CommentResultType {
  items: [Comment!]!
  count: Int!
}`
)
