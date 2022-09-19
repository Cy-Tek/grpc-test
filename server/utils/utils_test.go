package utils

import (
	"reflect"
	"testing"
)

type testAddress struct {
	Street string
	City   string
	Zip    string
}
type testUser struct {
	Name    string
	Address testAddress
}

func TestGetValueFromPath(t *testing.T) {
	type args struct {
		path   string
		object any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "Handles 1 level deep",
			args: args{
				path: "Description",
				object: struct {
					Description string
					Completed   bool
				}{Description: "hello"},
			},
			want:    "hello",
			wantErr: false,
		},
		{
			name: "Handles 2 levels deep",
			args: args{
				path: "user.isValid",
				object: struct {
					User struct{ IsValid bool }
				}{
					User: struct {
						IsValid bool
					}{IsValid: true},
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Handles nested pointers to structs",
			args: args{
				path: "user.location.street",
				object: struct {
					User struct {
						Location *testAddress
					}
				}{
					User: struct {
						Location *testAddress
					}{
						Location: &testAddress{
							Street: "2008 Galaxy Drive",
						},
					},
				},
			},
			want:    "2008 Galaxy Drive",
			wantErr: false,
		},
		{
			name: "Handles returning a struct",
			args: args{
				path: "User.Address",
				object: struct {
					User testUser
				}{
					User: testUser{
						Name: "Josh",
						Address: testAddress{
							Street: "212 Baker Street",
							City:   "London",
							Zip:    "23193",
						},
					},
				},
			},
			want: testAddress{
				Street: "212 Baker Street",
				City:   "London",
				Zip:    "23193",
			},
			wantErr: false,
		},
		{
			"Handles an invalid path",
			args{
				path: "User",
				object: struct {
					Test string
				}{Test: "testing"},
			},
			nil,
			true,
		},
		{
			"Handles an empty path",
			args{
				path: "",
				object: struct {
					Test string
				}{Test: "testing"},
			},
			nil,
			true,
		},
		{
			"Handles lowercase path names",
			args{
				path: "test",
				object: struct {
					Test string
				}{Test: "testing"},
			},
			"testing",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetValueFromPath(tt.args.path, tt.args.object)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetValueFromPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValueFromPath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
