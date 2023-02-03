package course

import (
	"reflect"
	"testing"
)

func TestCourse_AddCourse(t *testing.T) {
	type fields struct {
		ID   string
		Name string
	}
	type args struct {
		course Course
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Course
		wantErr bool
	}{
		{
			name: "Add new course",
			args: args{
				course: Course{
					ID:   "3",
					Name: "Science",
				},
			},
			want: append(DummyCourses, Course{
				ID:   "3",
				Name: "Science",
			}),
		},
		{
			name: "Add new course error",
			args: args{
				course: Course{
					ID:   "1",
					Name: "Math",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Course{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			got, err := c.AddCourse(tt.args.course)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddCourse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddCourse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourse_GetCourseByName(t *testing.T) {
	type fields struct {
		ID   string
		Name string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Course
		wantErr bool
	}{
		{
			name: "Get course by name",
			args: args{
				name: "Math",
			},
			want: &Course{
				ID:   "1",
				Name: "Math",
			},
		},
		{
			name: "Get course by name error",
			args: args{
				name: "Sport",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Course{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			got, err := c.GetCourseByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCourseByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCourseByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
