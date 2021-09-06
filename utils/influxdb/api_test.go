package influxdb

import "testing"

func TestStore(t *testing.T) {
	Init()

	type args struct {
		name  string
		tags  map[string]string
		value float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				name: "something",
				tags: map[string]string{
					"pod": "1",
				},
				value: 123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Store(tt.args.name, tt.args.tags, tt.args.value)
		})
	}
}
