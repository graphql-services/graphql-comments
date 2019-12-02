package gen

type key int

const (
	KeyPrincipalID      key    = iota
	KeyLoaders          key    = iota
	KeyExecutableSchema key    = iota
	KeyJWTClaims        key    = iota
	SchemaSDL           string = `scalar Time

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
  reference: ObjectSortType
  referenceID: ObjectSortType
  text: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
}

input CommentFilterType {
  AND: [CommentFilterType!]
  OR: [CommentFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  reference: String
  reference_ne: String
  reference_gt: String
  reference_lt: String
  reference_gte: String
  reference_lte: String
  reference_in: [String!]
  reference_like: String
  reference_prefix: String
  reference_suffix: String
  reference_null: Boolean
  referenceID: ID
  referenceID_ne: ID
  referenceID_gt: ID
  referenceID_lt: ID
  referenceID_gte: ID
  referenceID_lte: ID
  referenceID_in: [ID!]
  referenceID_null: Boolean
  text: String
  text_ne: String
  text_gt: String
  text_lt: String
  text_gte: String
  text_lte: String
  text_in: [String!]
  text_like: String
  text_prefix: String
  text_suffix: String
  text_null: Boolean
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
}

type CommentResultType {
  items: [Comment!]!
  count: Int!
}`
)
