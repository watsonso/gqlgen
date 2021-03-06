// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT

package dataloader

import (
	context "context"
	strconv "strconv"
	sync "sync"

	graphql "github.com/vektah/gqlgen/graphql"
	errors "github.com/vektah/gqlgen/neelance/errors"
	introspection "github.com/vektah/gqlgen/neelance/introspection"
	query "github.com/vektah/gqlgen/neelance/query"
	schema "github.com/vektah/gqlgen/neelance/schema"
)

type Resolvers interface {
	Customer_address(ctx context.Context, it *Customer) (*Address, error)
	Customer_orders(ctx context.Context, it *Customer) ([]Order, error)

	Order_items(ctx context.Context, it *Order) ([]Item, error)
	Query_customers(ctx context.Context) ([]Customer, error)
}

func MakeExecutableSchema(resolvers Resolvers) graphql.ExecutableSchema {
	return &executableSchema{resolvers}
}

type executableSchema struct {
	resolvers Resolvers
}

func (e *executableSchema) Schema() *schema.Schema {
	return parsedSchema
}

func (e *executableSchema) Query(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation) *graphql.Response {
	ec := executionContext{resolvers: e.resolvers, variables: variables, doc: doc, ctx: ctx}

	data := ec._query(op.Selections, nil)
	ec.wg.Wait()

	return &graphql.Response{
		Data:   data,
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Mutation(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation) *graphql.Response {
	return &graphql.Response{Errors: []*errors.QueryError{{Message: "mutations are not supported"}}}
}

func (e *executableSchema) Subscription(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation) <-chan *graphql.Response {
	events := make(chan *graphql.Response, 1)
	events <- &graphql.Response{Errors: []*errors.QueryError{{Message: "subscriptions are not supported"}}}
	return events
}

type executionContext struct {
	errors.Builder
	resolvers Resolvers
	variables map[string]interface{}
	doc       *query.Document
	ctx       context.Context
	wg        sync.WaitGroup
}

var addressImplementors = []string{"Address"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _address(sel []query.Selection, it *Address) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, addressImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Address")
		case "id":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.ID

			out.Values[i] = graphql.MarshalInt(res)
		case "street":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Street

			out.Values[i] = graphql.MarshalString(res)
		case "country":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Country

			out.Values[i] = graphql.MarshalString(res)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var customerImplementors = []string{"Customer"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _customer(sel []query.Selection, it *Customer) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, customerImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Customer")
		case "id":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.ID

			out.Values[i] = graphql.MarshalInt(res)
		case "name":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Name

			out.Values[i] = graphql.MarshalString(res)
		case "address":
			badArgs := false
			if badArgs {
				continue
			}
			ec.wg.Add(1)
			go func(i int, field graphql.CollectedField) {
				defer ec.wg.Done()
				res, err := ec.resolvers.Customer_address(ec.ctx, it)
				if err != nil {
					ec.Error(err)
					return
				}

				if res == nil {
					out.Values[i] = graphql.Null
				} else {
					out.Values[i] = ec._address(field.Selections, res)
				}
			}(i, field)
		case "orders":
			badArgs := false
			if badArgs {
				continue
			}
			ec.wg.Add(1)
			go func(i int, field graphql.CollectedField) {
				defer ec.wg.Done()
				res, err := ec.resolvers.Customer_orders(ec.ctx, it)
				if err != nil {
					ec.Error(err)
					return
				}

				arr1 := graphql.Array{}
				for idx1 := range res {
					var tmp1 graphql.Marshaler
					tmp1 = ec._order(field.Selections, &res[idx1])
					arr1 = append(arr1, tmp1)
				}
				out.Values[i] = arr1
			}(i, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var itemImplementors = []string{"Item"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _item(sel []query.Selection, it *Item) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, itemImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Item")
		case "name":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Name

			out.Values[i] = graphql.MarshalString(res)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var orderImplementors = []string{"Order"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _order(sel []query.Selection, it *Order) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, orderImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Order")
		case "id":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.ID

			out.Values[i] = graphql.MarshalInt(res)
		case "date":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Date

			out.Values[i] = graphql.MarshalTime(res)
		case "amount":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Amount

			out.Values[i] = graphql.MarshalFloat(res)
		case "items":
			badArgs := false
			if badArgs {
				continue
			}
			ec.wg.Add(1)
			go func(i int, field graphql.CollectedField) {
				defer ec.wg.Done()
				res, err := ec.resolvers.Order_items(ec.ctx, it)
				if err != nil {
					ec.Error(err)
					return
				}

				arr1 := graphql.Array{}
				for idx1 := range res {
					var tmp1 graphql.Marshaler
					tmp1 = ec._item(field.Selections, &res[idx1])
					arr1 = append(arr1, tmp1)
				}
				out.Values[i] = arr1
			}(i, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var queryImplementors = []string{"Query"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _query(sel []query.Selection, it *interface{}) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, queryImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "customers":
			badArgs := false
			if badArgs {
				continue
			}
			ec.wg.Add(1)
			go func(i int, field graphql.CollectedField) {
				defer ec.wg.Done()
				res, err := ec.resolvers.Query_customers(ec.ctx)
				if err != nil {
					ec.Error(err)
					return
				}

				arr1 := graphql.Array{}
				for idx1 := range res {
					var tmp1 graphql.Marshaler
					tmp1 = ec._customer(field.Selections, &res[idx1])
					arr1 = append(arr1, tmp1)
				}
				out.Values[i] = arr1
			}(i, field)
		case "__schema":
			badArgs := false
			if badArgs {
				continue
			}
			res := ec.introspectSchema()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Schema(field.Selections, res)
			}
		case "__type":
			badArgs := false
			var arg0 string
			if tmp, ok := field.Args["name"]; ok {
				tmp2, err := graphql.UnmarshalString(tmp)
				if err != nil {
					badArgs = true
				}
				arg0 = tmp2
			}
			if badArgs {
				continue
			}
			res := ec.introspectType(arg0)

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(sel []query.Selection, it *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __DirectiveImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Name()

			out.Values[i] = graphql.MarshalString(res)
		case "description":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "locations":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Locations()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler
				tmp1 = graphql.MarshalString(res[idx1])
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "args":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Args()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___InputValue(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(sel []query.Selection, it *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __EnumValueImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Name()

			out.Values[i] = graphql.MarshalString(res)
		case "description":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "isDeprecated":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.IsDeprecated()

			out.Values[i] = graphql.MarshalBoolean(res)
		case "deprecationReason":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.DeprecationReason()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(sel []query.Selection, it *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __FieldImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Name()

			out.Values[i] = graphql.MarshalString(res)
		case "description":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "args":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Args()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___InputValue(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "type":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Type()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "isDeprecated":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.IsDeprecated()

			out.Values[i] = graphql.MarshalBoolean(res)
		case "deprecationReason":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.DeprecationReason()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(sel []query.Selection, it *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __InputValueImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Name()

			out.Values[i] = graphql.MarshalString(res)
		case "description":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "type":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Type()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "defaultValue":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.DefaultValue()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(sel []query.Selection, it *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __SchemaImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Types()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Type(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "queryType":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.QueryType()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "mutationType":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.MutationType()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "subscriptionType":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.SubscriptionType()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "directives":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Directives()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Directive(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(sel []query.Selection, it *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __TypeImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Kind()

			out.Values[i] = graphql.MarshalString(res)
		case "name":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Name()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "description":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "fields":
			badArgs := false
			var arg0 bool
			if tmp, ok := field.Args["includeDeprecated"]; ok {
				tmp2, err := graphql.UnmarshalBoolean(tmp)
				if err != nil {
					badArgs = true
				}
				arg0 = tmp2
			}
			if badArgs {
				continue
			}
			res := it.Fields(arg0)

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Field(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "interfaces":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Interfaces()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Type(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "possibleTypes":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.PossibleTypes()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Type(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "enumValues":
			badArgs := false
			var arg0 bool
			if tmp, ok := field.Args["includeDeprecated"]; ok {
				tmp2, err := graphql.UnmarshalBoolean(tmp)
				if err != nil {
					badArgs = true
				}
				arg0 = tmp2
			}
			if badArgs {
				continue
			}
			res := it.EnumValues(arg0)

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___EnumValue(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "inputFields":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.InputFields()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___InputValue(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "ofType":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.OfType()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var parsedSchema = schema.MustParse("schema {\n    query: Query\n}\n\ntype Query {\n    customers: [Customer!]\n}\n\ntype Customer {\n    id: Int!\n    name: String!\n    address: Address\n    orders: [Order!]\n}\n\ntype Address {\n    id: Int!\n    street: String\n    country: String\n}\n\ntype Order {\n    id: Int!\n    date: Time\n    amount: Float!\n    items: [Item!]\n}\n\ntype Item {\n    name: String\n}\nscalar Time\n")

func (ec *executionContext) introspectSchema() *introspection.Schema {
	return introspection.WrapSchema(parsedSchema)
}

func (ec *executionContext) introspectType(name string) *introspection.Type {
	t := parsedSchema.Resolve(name)
	if t == nil {
		return nil
	}
	return introspection.WrapType(t)
}
