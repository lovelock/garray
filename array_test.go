package array

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {

	complexJson := `
	{
  "name": "John Doe",
  "age": 30,
  "email": "john.doe@example.com",
  "isActive": true,
  "address": {
    "street": "123 Main St",
    "city": "Anytown",
    "state": "CA",
    "postalCode": "12345"
  },
  "phoneNumbers": [
    {
      "type": "home",
      "number": "555-555-5555"
    },
    {
      "type": "work",
      "number": "555-555-5556"
    }
  ],
  "projects": [
    {
      "name": "Project Alpha",
      "status": "completed",
      "tasks": [
        {
          "name": "Task 1",
          "dueDate": "2023-10-01",
          "completed": true
        },
        {
          "name": "Task 2",
          "dueDate": "2023-10-15",
          "completed": false
        }
      ]
    },
    {
      "name": "Project Beta",
      "status": "in progress",
      "tasks": [
        {
          "name": "Task 3",
          "dueDate": "2023-11-01",
          "completed": false
        },
        {
          "name": "Task 4",
          "dueDate": "2023-12-01",
          "completed": false
        }
      ]
    }
  ],
  "preferences": {
    "contactMethod": "email",
    "newsletterSubscribed": true,
    "languages": ["English", "Spanish", "German"]
  }
}
`

	var jsonMap map[string]any
	json.Unmarshal([]byte(complexJson), &jsonMap)

	type args struct {
		structure any
		keys      []string
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "Get the status of the second task of the first project OK",
			args: args{
				structure: jsonMap,
				keys:      []string{"projects", "0", "tasks", "1", "completed"},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Get work phone number OK",
			args: args{
				structure: jsonMap,
				keys:      []string{"phoneNumbers", "1", "number"},
			},
			want:    "555-555-5556",
			wantErr: false,
		},
		{
			name: "Get from map OK",
			args: args{
				structure: map[string]any{
					"a": "b",
					"c": "d",
				},
				keys: []string{"a"},
			},
			want:    "b",
			wantErr: false,
		},
		{
			name: "Get from map key not found",
			args: args{
				structure: map[string]any{
					"a": "b",
					"c": "d",
				},
				keys: []string{"e"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Get from nested map OK",
			args: args{
				structure: map[string]any{
					"a": "b",
					"c": map[string]any{
						"d": "e",
						"f": "g",
					},
				},
				keys: []string{"c", "d"},
			},
			want:    "e",
			wantErr: false,
		},
		{
			name: "Get from nested map key not found",
			args: args{
				structure: map[string]any{
					"a": "b",
					"c": map[string]any{
						"d": "e",
						"f": "g",
					},
				},
				keys: []string{"c", "g"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Get from nested slice OK",
			args: args{
				structure: map[string]any{
					"a": "b",
					"c": []any{
						"e",
						"g",
					},
				},
				keys: []string{"c", "1"},
			},
			want:    "g",
			wantErr: false,
		},
		{
			name: "Get from nested slice of map OK",
			args: args{
				structure: map[string]any{
					"a": "b",
					"c": []any{
						map[string]any{
							"foo": "bar",
						},
						"g",
					},
				},
				keys: []string{"c", "0", "foo"},
			},
			want:    "bar",
			wantErr: false,
		},
		{
			name: "Get from nested slice index out of range",
			args: args{
				structure: map[string]any{
					"a": "b",
					"c": []any{
						"e",
						"g",
					},
				},
				keys: []string{"c", "2"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Get from slice OK",
			args: args{
				structure: []any{
					"a",
					"b",
				},
				keys: []string{"0"},
			},
			want:    "a",
			wantErr: false,
		},
		{
			name: "Get from slice out of range",
			args: args{
				structure: []any{
					"a",
					"b",
				},
				keys: []string{"2"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.structure, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
