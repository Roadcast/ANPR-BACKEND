// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CamerasColumns holds the columns for the "cameras" table.
	CamerasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "name", Type: field.TypeString},
		{Name: "model", Type: field.TypeString},
		{Name: "imei", Type: field.TypeString, Unique: true},
		{Name: "location", Type: field.TypeString, SchemaType: map[string]string{"postgres": "GEOMETRY(Point, 4326)"}},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "police_station_id", Type: field.TypeUUID, Nullable: true},
	}
	// CamerasTable holds the schema information for the "cameras" table.
	CamerasTable = &schema.Table{
		Name:       "cameras",
		Columns:    CamerasColumns,
		PrimaryKey: []*schema.Column{CamerasColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cameras_police_stations_camera",
				Columns:    []*schema.Column{CamerasColumns[8]},
				RefColumns: []*schema.Column{PoliceStationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "camera_name",
				Unique:  false,
				Columns: []*schema.Column{CamerasColumns[3]},
				Annotation: &entsql.IndexAnnotation{
					Types: map[string]string{
						"mysql":    "FULLTEXT",
						"postgres": "GIN",
					},
				},
			},
		},
	}
	// CarsColumns holds the columns for the "cars" table.
	CarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "make", Type: field.TypeString},
		{Name: "model", Type: field.TypeString},
		{Name: "year", Type: field.TypeInt},
		{Name: "registration", Type: field.TypeString, Unique: true},
		{Name: "color", Type: field.TypeString},
	}
	// CarsTable holds the schema information for the "cars" table.
	CarsTable = &schema.Table{
		Name:       "cars",
		Columns:    CarsColumns,
		PrimaryKey: []*schema.Column{CarsColumns[0]},
	}
	// PermissionsColumns holds the columns for the "permissions" table.
	PermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "name", Type: field.TypeString},
		{Name: "can_read", Type: field.TypeBool, Default: false},
		{Name: "can_create", Type: field.TypeBool, Default: false},
		{Name: "can_update", Type: field.TypeBool, Default: false},
		{Name: "can_delete", Type: field.TypeBool, Default: false},
		{Name: "role_permissions", Type: field.TypeUUID, Nullable: true},
	}
	// PermissionsTable holds the schema information for the "permissions" table.
	PermissionsTable = &schema.Table{
		Name:       "permissions",
		Columns:    PermissionsColumns,
		PrimaryKey: []*schema.Column{PermissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "permissions_roles_permissions",
				Columns:    []*schema.Column{PermissionsColumns[8]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PoliceStationsColumns holds the columns for the "police_stations" table.
	PoliceStationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "name", Type: field.TypeString},
		{Name: "location", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "GEOMETRY(Point, 4326)"}},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "identifier", Type: field.TypeString, Unique: true},
		{Name: "parent_station_id", Type: field.TypeUUID, Nullable: true},
	}
	// PoliceStationsTable holds the schema information for the "police_stations" table.
	PoliceStationsTable = &schema.Table{
		Name:       "police_stations",
		Columns:    PoliceStationsColumns,
		PrimaryKey: []*schema.Column{PoliceStationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "police_stations_police_stations_child_stations",
				Columns:    []*schema.Column{PoliceStationsColumns[7]},
				RefColumns: []*schema.Column{PoliceStationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "policestation_name",
				Unique:  false,
				Columns: []*schema.Column{PoliceStationsColumns[3]},
				Annotation: &entsql.IndexAnnotation{
					Types: map[string]string{
						"mysql":    "FULLTEXT",
						"postgres": "GIN",
					},
				},
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "name", Type: field.TypeString},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "police_station_id", Type: field.TypeUUID, Nullable: true},
		{Name: "role_id", Type: field.TypeUUID},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_police_stations_users",
				Columns:    []*schema.Column{UsersColumns[8]},
				RefColumns: []*schema.Column{PoliceStationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_roles_users",
				Columns:    []*schema.Column{UsersColumns[9]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_name",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[3]},
				Annotation: &entsql.IndexAnnotation{
					Types: map[string]string{
						"mysql":    "FULLTEXT",
						"postgres": "GIN",
					},
				},
			},
		},
	}
	// VehicleDataColumns holds the columns for the "vehicle_data" table.
	VehicleDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp with time zone"}},
		{Name: "plate_bounding_box", Type: field.TypeJSON, Nullable: true},
		{Name: "plate_channel", Type: field.TypeInt, Nullable: true},
		{Name: "plate_is_exist", Type: field.TypeBool, Nullable: true},
		{Name: "plate_color", Type: field.TypeString, Nullable: true},
		{Name: "plate_number", Type: field.TypeString, Nullable: true},
		{Name: "plate_type", Type: field.TypeString, Nullable: true},
		{Name: "plate_region", Type: field.TypeString, Nullable: true},
		{Name: "plate_upload_num", Type: field.TypeInt, Nullable: true},
		{Name: "snap_allow_user", Type: field.TypeBool, Nullable: true},
		{Name: "snap_allow_user_end_time", Type: field.TypeString, Nullable: true},
		{Name: "snap_defence_code", Type: field.TypeString, Nullable: true},
		{Name: "snap_device_id", Type: field.TypeString, Nullable: true},
		{Name: "snap_in_car_people_num", Type: field.TypeInt, Nullable: true},
		{Name: "snap_lan_no", Type: field.TypeInt, Nullable: true},
		{Name: "snap_open_strobe", Type: field.TypeBool, Nullable: true},
		{Name: "vehicle_bounding_box", Type: field.TypeJSON, Nullable: true},
		{Name: "vehicle_series", Type: field.TypeString, Nullable: true},
	}
	// VehicleDataTable holds the schema information for the "vehicle_data" table.
	VehicleDataTable = &schema.Table{
		Name:       "vehicle_data",
		Columns:    VehicleDataColumns,
		PrimaryKey: []*schema.Column{VehicleDataColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CamerasTable,
		CarsTable,
		PermissionsTable,
		PoliceStationsTable,
		RolesTable,
		UsersTable,
		VehicleDataTable,
	}
)

func init() {
	CamerasTable.ForeignKeys[0].RefTable = PoliceStationsTable
	PermissionsTable.ForeignKeys[0].RefTable = RolesTable
	PoliceStationsTable.ForeignKeys[0].RefTable = PoliceStationsTable
	UsersTable.ForeignKeys[0].RefTable = PoliceStationsTable
	UsersTable.ForeignKeys[1].RefTable = RolesTable
}
