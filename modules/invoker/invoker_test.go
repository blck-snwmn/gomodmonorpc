package invoker

import "testing"

func Test_extractKey(t *testing.T) {

	type args struct {
		fullmethod string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{fullmethod: "/proto.order.v1.OrderService/Order"},
			want:    "proto.order.v1.OrderService",
			wantErr: false,
		},
		{
			name:    "failed",
			args:    args{fullmethod: "proto.order.v1.OrderService/Order"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractKey(tt.args.fullmethod)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
