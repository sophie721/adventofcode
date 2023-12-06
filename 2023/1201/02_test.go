package _1201

import "testing"

func Test_run01Two(t *testing.T) {
	type args struct {
		input []string
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "default input",
			args: args{
				input: []string{
					"two1nine",
					"eightwothree",
					"abcone2threexyz",
					"xtwone3four",
					"4nineeightseven2",
					"zoneight234",
					"7pqrstsixteen",
				},
			},
			want:    281,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Run01Two(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run01Two() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Run01Two() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDigitAndString(t *testing.T) {
	type args struct {
		s            string
		lookFromHead bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "match",
			args: args{
				s:            "two9",
				lookFromHead: true,
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findDigitAndString(tt.args.s, tt.args.lookFromHead)
			if (err != nil) != tt.wantErr {
				t.Errorf("findDigitAndString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findDigitAndString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_letter2Int(t *testing.T) {
	type args struct {
		b       []rune
		reverse bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "full letter",
			args: args{
				b:       []rune("eight"),
				reverse: false,
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "letter merge together",
			args: args{
				b:       []rune("abcone"),
				reverse: false,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := letter2Int(tt.args.b, tt.args.reverse)
			if (err != nil) != tt.wantErr {
				t.Errorf("letter2Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("letter2Int() got = %v, want %v", got, tt.want)
			}
		})
	}
}
