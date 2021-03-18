package main

import "testing"

func TestAssignStructFields(t *testing.T) {
	type args struct {
		in     interface{}
		values map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Check when struct with no pointers",
			args: args{
				in: struct {
					FieldString string `json:"field_string"`
					FieldInt    int
				}{
					FieldString: "stroka",
					FieldInt:    107,
				},
				values: map[string]interface{}{"first": 5, "FieldString": "dacota"},
			},
			wantErr: true,
		},
		{
			name: "Check when struct with pointers",
			args: args{
				in: &struct {
					FieldString string `json:"field_string"`
					FieldInt    int
				}{
					FieldString: "stroka",
					FieldInt:    107,
				},
				values: map[string]interface{}{"first": 5, "FieldString": "dacota"},
			},
			wantErr: false,
		},
		{
			name: "Check simple case with 2 changed fields",
			args: args{
				in: &struct {
					FieldString string `json:"field_string"`
					FieldInt    int
				}{
					FieldString: "stroka",
					FieldInt:    107,
				},
				values: map[string]interface{}{"FieldInt": 5, "FieldString": "dacota"},
			},
			wantErr: false,
		},
		{
			name: "Check nested struct",
			args: args{
				in: &struct {
					FieldString string `json:"field_string"`
					FieldInt    int
					Object      struct {
						NestedFloatField float64
					}
				}{
					FieldString: "stroka",
					FieldInt:    107,
					Object:      struct{ NestedFloatField float64 }{NestedFloatField: 2.18},
				},
				values: map[string]interface{}{"first": 5, "FieldString": "dacota", "NestedFloatField": 3.14},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AssignStructFields(tt.args.in, tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("AssignStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
