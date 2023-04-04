package main

import (
	"reflect"
	"testing"
)

func TestBuildHierarchy(t *testing.T) {
	tests := []struct {
		name     string
		content  []*WorkerInPut
		expected WorkerOutPut
	}{
		{
			name: "shouldProcessWithSuccess1",
			content: []*WorkerInPut{
				{
					Id:       1,
					Name:     "CEO",
					ParentId: 0,
				},
				{
					Id:       2,
					Name:     "CTO",
					ParentId: 1,
				},
				{
					Id:       3,
					Name:     "CFO",
					ParentId: 1,
				},
				{
					Id:       4,
					Name:     "Manager 1",
					ParentId: 2,
				},
				{
					Id:       5,
					Name:     "Manager 2",
					ParentId: 2,
				},
			},
			expected: WorkerOutPut{
				Id:   1,
				Name: "CEO",
				Children: []WorkerOutPut{
					{
						Id:   2,
						Name: "CTO",
						Children: []WorkerOutPut{
							{
								Id:       4,
								Name:     "Manager 1",
								Children: nil,
							},
							{
								Id:       5,
								Name:     "Manager 2",
								Children: nil,
							},
						},
					},
					{
						Id:       3,
						Name:     "CFO",
						Children: nil,
					},
				},
			},
		},
	}

	for _, test := range tests {
		out := BuildHierarchy(1, test.content)
		if !reflect.DeepEqual(out, test.expected) {
			t.Errorf("%s => [result: %v, expected: %v]", test.name, out, test.expected)
		}
	}
}
