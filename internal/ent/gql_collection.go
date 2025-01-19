// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"go-ent-project/internal/ent/camera"
	"go-ent-project/internal/ent/car"
	"go-ent-project/internal/ent/permission"
	"go-ent-project/internal/ent/policestation"
	"go-ent-project/internal/ent/role"
	"go-ent-project/internal/ent/user"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CameraQuery) CollectFields(ctx context.Context, satisfies ...string) (*CameraQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CameraQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(camera.Columns))
		selectedFields = []string{camera.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "createdAt":
			if _, ok := fieldSeen[camera.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, camera.FieldCreatedAt)
				fieldSeen[camera.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[camera.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, camera.FieldUpdatedAt)
				fieldSeen[camera.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[camera.FieldName]; !ok {
				selectedFields = append(selectedFields, camera.FieldName)
				fieldSeen[camera.FieldName] = struct{}{}
			}
		case "model":
			if _, ok := fieldSeen[camera.FieldModel]; !ok {
				selectedFields = append(selectedFields, camera.FieldModel)
				fieldSeen[camera.FieldModel] = struct{}{}
			}
		case "imei":
			if _, ok := fieldSeen[camera.FieldImei]; !ok {
				selectedFields = append(selectedFields, camera.FieldImei)
				fieldSeen[camera.FieldImei] = struct{}{}
			}
		case "location":
			if _, ok := fieldSeen[camera.FieldLocation]; !ok {
				selectedFields = append(selectedFields, camera.FieldLocation)
				fieldSeen[camera.FieldLocation] = struct{}{}
			}
		case "active":
			if _, ok := fieldSeen[camera.FieldActive]; !ok {
				selectedFields = append(selectedFields, camera.FieldActive)
				fieldSeen[camera.FieldActive] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		c.Select(selectedFields...)
	}
	return nil
}

type cameraPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CameraPaginateOption
}

func newCameraPaginateArgs(rv map[string]any) *cameraPaginateArgs {
	args := &cameraPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &CameraOrder{Field: &CameraOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithCameraOrder(order))
			}
		case *CameraOrder:
			if v != nil {
				args.opts = append(args.opts, WithCameraOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*CameraWhereInput); ok {
		args.opts = append(args.opts, WithCameraFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CarQuery) CollectFields(ctx context.Context, satisfies ...string) (*CarQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CarQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(car.Columns))
		selectedFields = []string{car.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "createdAt":
			if _, ok := fieldSeen[car.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, car.FieldCreatedAt)
				fieldSeen[car.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[car.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, car.FieldUpdatedAt)
				fieldSeen[car.FieldUpdatedAt] = struct{}{}
			}
		case "make":
			if _, ok := fieldSeen[car.FieldMake]; !ok {
				selectedFields = append(selectedFields, car.FieldMake)
				fieldSeen[car.FieldMake] = struct{}{}
			}
		case "model":
			if _, ok := fieldSeen[car.FieldModel]; !ok {
				selectedFields = append(selectedFields, car.FieldModel)
				fieldSeen[car.FieldModel] = struct{}{}
			}
		case "year":
			if _, ok := fieldSeen[car.FieldYear]; !ok {
				selectedFields = append(selectedFields, car.FieldYear)
				fieldSeen[car.FieldYear] = struct{}{}
			}
		case "registration":
			if _, ok := fieldSeen[car.FieldRegistration]; !ok {
				selectedFields = append(selectedFields, car.FieldRegistration)
				fieldSeen[car.FieldRegistration] = struct{}{}
			}
		case "color":
			if _, ok := fieldSeen[car.FieldColor]; !ok {
				selectedFields = append(selectedFields, car.FieldColor)
				fieldSeen[car.FieldColor] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		c.Select(selectedFields...)
	}
	return nil
}

type carPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CarPaginateOption
}

func newCarPaginateArgs(rv map[string]any) *carPaginateArgs {
	args := &carPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &CarOrder{Field: &CarOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithCarOrder(order))
			}
		case *CarOrder:
			if v != nil {
				args.opts = append(args.opts, WithCarOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*CarWhereInput); ok {
		args.opts = append(args.opts, WithCarFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pe *PermissionQuery) CollectFields(ctx context.Context, satisfies ...string) (*PermissionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pe, nil
	}
	if err := pe.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pe, nil
}

func (pe *PermissionQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(permission.Columns))
		selectedFields = []string{permission.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "createdAt":
			if _, ok := fieldSeen[permission.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, permission.FieldCreatedAt)
				fieldSeen[permission.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[permission.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, permission.FieldUpdatedAt)
				fieldSeen[permission.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[permission.FieldName]; !ok {
				selectedFields = append(selectedFields, permission.FieldName)
				fieldSeen[permission.FieldName] = struct{}{}
			}
		case "canRead":
			if _, ok := fieldSeen[permission.FieldCanRead]; !ok {
				selectedFields = append(selectedFields, permission.FieldCanRead)
				fieldSeen[permission.FieldCanRead] = struct{}{}
			}
		case "canCreate":
			if _, ok := fieldSeen[permission.FieldCanCreate]; !ok {
				selectedFields = append(selectedFields, permission.FieldCanCreate)
				fieldSeen[permission.FieldCanCreate] = struct{}{}
			}
		case "canUpdate":
			if _, ok := fieldSeen[permission.FieldCanUpdate]; !ok {
				selectedFields = append(selectedFields, permission.FieldCanUpdate)
				fieldSeen[permission.FieldCanUpdate] = struct{}{}
			}
		case "canDelete":
			if _, ok := fieldSeen[permission.FieldCanDelete]; !ok {
				selectedFields = append(selectedFields, permission.FieldCanDelete)
				fieldSeen[permission.FieldCanDelete] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		pe.Select(selectedFields...)
	}
	return nil
}

type permissionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PermissionPaginateOption
}

func newPermissionPaginateArgs(rv map[string]any) *permissionPaginateArgs {
	args := &permissionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &PermissionOrder{Field: &PermissionOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithPermissionOrder(order))
			}
		case *PermissionOrder:
			if v != nil {
				args.opts = append(args.opts, WithPermissionOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*PermissionWhereInput); ok {
		args.opts = append(args.opts, WithPermissionFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ps *PoliceStationQuery) CollectFields(ctx context.Context, satisfies ...string) (*PoliceStationQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ps, nil
	}
	if err := ps.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ps, nil
}

func (ps *PoliceStationQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(policestation.Columns))
		selectedFields = []string{policestation.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "users":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: ps.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			ps.WithNamedUsers(alias, func(wq *UserQuery) {
				*wq = *query
			})

		case "parentStation":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PoliceStationClient{config: ps.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, policestationImplementors)...); err != nil {
				return err
			}
			ps.withParentStation = query

		case "childStations":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PoliceStationClient{config: ps.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, policestationImplementors)...); err != nil {
				return err
			}
			ps.WithNamedChildStations(alias, func(wq *PoliceStationQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[policestation.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, policestation.FieldCreatedAt)
				fieldSeen[policestation.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[policestation.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, policestation.FieldUpdatedAt)
				fieldSeen[policestation.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[policestation.FieldName]; !ok {
				selectedFields = append(selectedFields, policestation.FieldName)
				fieldSeen[policestation.FieldName] = struct{}{}
			}
		case "location":
			if _, ok := fieldSeen[policestation.FieldLocation]; !ok {
				selectedFields = append(selectedFields, policestation.FieldLocation)
				fieldSeen[policestation.FieldLocation] = struct{}{}
			}
		case "code":
			if _, ok := fieldSeen[policestation.FieldCode]; !ok {
				selectedFields = append(selectedFields, policestation.FieldCode)
				fieldSeen[policestation.FieldCode] = struct{}{}
			}
		case "identifier":
			if _, ok := fieldSeen[policestation.FieldIdentifier]; !ok {
				selectedFields = append(selectedFields, policestation.FieldIdentifier)
				fieldSeen[policestation.FieldIdentifier] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ps.Select(selectedFields...)
	}
	return nil
}

type policestationPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PoliceStationPaginateOption
}

func newPoliceStationPaginateArgs(rv map[string]any) *policestationPaginateArgs {
	args := &policestationPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &PoliceStationOrder{Field: &PoliceStationOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithPoliceStationOrder(order))
			}
		case *PoliceStationOrder:
			if v != nil {
				args.opts = append(args.opts, WithPoliceStationOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*PoliceStationWhereInput); ok {
		args.opts = append(args.opts, WithPoliceStationFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *RoleQuery) CollectFields(ctx context.Context, satisfies ...string) (*RoleQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return r, nil
	}
	if err := r.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RoleQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(role.Columns))
		selectedFields = []string{role.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "permissions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PermissionClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, permissionImplementors)...); err != nil {
				return err
			}
			r.WithNamedPermissions(alias, func(wq *PermissionQuery) {
				*wq = *query
			})

		case "users":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			r.WithNamedUsers(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[role.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, role.FieldCreatedAt)
				fieldSeen[role.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[role.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, role.FieldUpdatedAt)
				fieldSeen[role.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[role.FieldName]; !ok {
				selectedFields = append(selectedFields, role.FieldName)
				fieldSeen[role.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		r.Select(selectedFields...)
	}
	return nil
}

type rolePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []RolePaginateOption
}

func newRolePaginateArgs(rv map[string]any) *rolePaginateArgs {
	args := &rolePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &RoleOrder{Field: &RoleOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithRoleOrder(order))
			}
		case *RoleOrder:
			if v != nil {
				args.opts = append(args.opts, WithRoleOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*RoleWhereInput); ok {
		args.opts = append(args.opts, WithRoleFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "role":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RoleClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, roleImplementors)...); err != nil {
				return err
			}
			u.withRole = query
			if _, ok := fieldSeen[user.FieldRoleID]; !ok {
				selectedFields = append(selectedFields, user.FieldRoleID)
				fieldSeen[user.FieldRoleID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[user.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldCreatedAt)
				fieldSeen[user.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[user.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldUpdatedAt)
				fieldSeen[user.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[user.FieldName]; !ok {
				selectedFields = append(selectedFields, user.FieldName)
				fieldSeen[user.FieldName] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[user.FieldEmail]; !ok {
				selectedFields = append(selectedFields, user.FieldEmail)
				fieldSeen[user.FieldEmail] = struct{}{}
			}
		case "phone":
			if _, ok := fieldSeen[user.FieldPhone]; !ok {
				selectedFields = append(selectedFields, user.FieldPhone)
				fieldSeen[user.FieldPhone] = struct{}{}
			}
		case "active":
			if _, ok := fieldSeen[user.FieldActive]; !ok {
				selectedFields = append(selectedFields, user.FieldActive)
				fieldSeen[user.FieldActive] = struct{}{}
			}
		case "roleID":
			if _, ok := fieldSeen[user.FieldRoleID]; !ok {
				selectedFields = append(selectedFields, user.FieldRoleID)
				fieldSeen[user.FieldRoleID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &UserOrder{Field: &UserOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithUserOrder(order))
			}
		case *UserOrder:
			if v != nil {
				args.opts = append(args.opts, WithUserOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
