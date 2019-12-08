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
  company(id: ID, q: String, filter: CompanyFilterType): Company
  companies(offset: Int, limit: Int = 30, q: String, sort: [CompanySortType!], filter: CompanyFilterType): CompanyResultType
  person(id: ID, q: String, filter: PersonFilterType): Person
  people(offset: Int, limit: Int = 30, q: String, sort: [PersonSortType!], filter: PersonFilterType): PersonResultType
}

type Mutation {
  createComment(input: CommentCreateInput!): Comment!
  updateComment(id: ID!, input: CommentUpdateInput!): Comment!
  deleteComment(id: ID!): Comment!
  deleteAllComments: Boolean!
  createCompany(input: CompanyCreateInput!): Company!
  updateCompany(id: ID!, input: CompanyUpdateInput!): Company!
  deleteCompany(id: ID!): Company!
  deleteAllCompanies: Boolean!
  createPerson(input: PersonCreateInput!): Person!
  updatePerson(id: ID!, input: PersonUpdateInput!): Person!
  deletePerson(id: ID!): Person!
  deleteAllPeople: Boolean!
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

type Company {
  id: ID!
  name: String
  employees: [Person!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  employeesIds: [ID!]!
}

type Person {
  id: ID!
  name: String
  companies: [Company!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  companiesIds: [ID!]!
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
}

input CompanyCreateInput {
  id: ID
  name: String
  employeesIds: [ID!]
}

input CompanyUpdateInput {
  name: String
  employeesIds: [ID!]
}

input CompanySortType {
  id: ObjectSortType
  name: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  employeesIds: ObjectSortType
  employees: PersonSortType
}

input CompanyFilterType {
  AND: [CompanyFilterType!]
  OR: [CompanyFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  name: String
  name_ne: String
  name_gt: String
  name_lt: String
  name_gte: String
  name_lte: String
  name_in: [String!]
  name_like: String
  name_prefix: String
  name_suffix: String
  name_null: Boolean
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
  employees: PersonFilterType
}

type CompanyResultType {
  items: [Company!]!
  count: Int!
}

input PersonCreateInput {
  id: ID
  name: String
  companiesIds: [ID!]
}

input PersonUpdateInput {
  name: String
  companiesIds: [ID!]
}

input PersonSortType {
  id: ObjectSortType
  name: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  companiesIds: ObjectSortType
  companies: CompanySortType
}

input PersonFilterType {
  AND: [PersonFilterType!]
  OR: [PersonFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  name: String
  name_ne: String
  name_gt: String
  name_lt: String
  name_gte: String
  name_lte: String
  name_in: [String!]
  name_like: String
  name_prefix: String
  name_suffix: String
  name_null: Boolean
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
  companies: CompanyFilterType
}

type PersonResultType {
  items: [Person!]!
  count: Int!
}`
)
