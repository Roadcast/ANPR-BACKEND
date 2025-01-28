// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-ent-project/internal/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// RoleID holds the value of the "role_id" field.
	RoleID uuid.UUID `json:"role_id,omitempty"`
	// PoliceStationID holds the value of the "police_station_id" field.
	PoliceStationID uuid.UUID `json:"police_station_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Role holds the value of the role edge.
	Role []*Role `json:"role,omitempty"`
	// PoliceStation holds the value of the police_station edge.
	PoliceStation []*PoliceStation `json:"police_station,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedRole          map[string][]*Role
	namedPoliceStation map[string][]*PoliceStation
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RoleOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// PoliceStationOrErr returns the PoliceStation value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PoliceStationOrErr() ([]*PoliceStation, error) {
	if e.loadedTypes[1] {
		return e.PoliceStation, nil
	}
	return nil, &NotLoadedError{edge: "police_station"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldActive:
			values[i] = new(sql.NullBool)
		case user.FieldName, user.FieldEmail, user.FieldPassword, user.FieldPhone:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID, user.FieldRoleID, user.FieldPoliceStationID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				u.Active = value.Bool
			}
		case user.FieldRoleID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value != nil {
				u.RoleID = *value
			}
		case user.FieldPoliceStationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field police_station_id", values[i])
			} else if value != nil {
				u.PoliceStationID = *value
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryRole queries the "role" edge of the User entity.
func (u *User) QueryRole() *RoleQuery {
	return NewUserClient(u.config).QueryRole(u)
}

// QueryPoliceStation queries the "police_station" edge of the User entity.
func (u *User) QueryPoliceStation() *PoliceStationQuery {
	return NewUserClient(u.config).QueryPoliceStation(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", u.Active))
	builder.WriteString(", ")
	builder.WriteString("role_id=")
	builder.WriteString(fmt.Sprintf("%v", u.RoleID))
	builder.WriteString(", ")
	builder.WriteString("police_station_id=")
	builder.WriteString(fmt.Sprintf("%v", u.PoliceStationID))
	builder.WriteByte(')')
	return builder.String()
}

// NamedRole returns the Role named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedRole(name string) ([]*Role, error) {
	if u.Edges.namedRole == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedRole[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedRole(name string, edges ...*Role) {
	if u.Edges.namedRole == nil {
		u.Edges.namedRole = make(map[string][]*Role)
	}
	if len(edges) == 0 {
		u.Edges.namedRole[name] = []*Role{}
	} else {
		u.Edges.namedRole[name] = append(u.Edges.namedRole[name], edges...)
	}
}

// NamedPoliceStation returns the PoliceStation named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedPoliceStation(name string) ([]*PoliceStation, error) {
	if u.Edges.namedPoliceStation == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedPoliceStation[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedPoliceStation(name string, edges ...*PoliceStation) {
	if u.Edges.namedPoliceStation == nil {
		u.Edges.namedPoliceStation = make(map[string][]*PoliceStation)
	}
	if len(edges) == 0 {
		u.Edges.namedPoliceStation[name] = []*PoliceStation{}
	} else {
		u.Edges.namedPoliceStation[name] = append(u.Edges.namedPoliceStation[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User
