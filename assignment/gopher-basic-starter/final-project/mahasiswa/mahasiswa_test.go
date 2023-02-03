package mahasiswa

import (
	"final-project/course"
	"reflect"
	"testing"
)

func TestMahasiswa_AddMahasiswa(t *testing.T) {
	type fields struct {
		NIM     string
		Name    string
		Age     int
		Courses []course.Course
	}
	type args struct {
		mahasiswa Mahasiswa
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Mahasiswa
		wantErr bool
	}{
		{
			name: "Test Add Mahasiswa",
			args: args{
				mahasiswa: Mahasiswa{
					NIM:  "789",
					Name: "test",
					Age:  20,
				},
			},
			want: append(DummyMahasiswa, Mahasiswa{
				NIM:  "789",
				Name: "test",
				Age:  20,
			}),
		},
		{
			name: "Test Add Mahasiswa Error",
			args: args{
				mahasiswa: Mahasiswa{
					NIM:  "123",
					Name: "test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mahasiswa{
				NIM:     tt.fields.NIM,
				Name:    tt.fields.Name,
				Age:     tt.fields.Age,
				Courses: tt.fields.Courses,
			}
			got, err := m.AddMahasiswa(tt.args.mahasiswa)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddMahasiswa() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMahasiswa() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMahasiswa_DeleteMahasiswa(t *testing.T) {
	type fields struct {
		NIM     string
		Name    string
		Age     int
		Courses []course.Course
	}
	type args struct {
		nim string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Mahasiswa
		wantErr bool
	}{
		{
			name: "Test Delete Mahasiswa",
			args: args{
				nim: "123",
			},
			want: DummyMahasiswa[:2],
		},
		{
			name: "Test Delete Mahasiswa Error",
			args: args{
				nim: "1234",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mahasiswa{
				NIM:     tt.fields.NIM,
				Name:    tt.fields.Name,
				Age:     tt.fields.Age,
				Courses: tt.fields.Courses,
			}
			got, err := m.DeleteMahasiswa(tt.args.nim)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteMahasiswa() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteMahasiswa() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMahasiswa_GetMahasiswa(t *testing.T) {
	type fields struct {
		NIM     string
		Name    string
		Age     int
		Courses []course.Course
	}
	type args struct {
		nim string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Mahasiswa
		wantErr bool
	}{
		{
			name: "Test Get Mahasiswa",
			args: args{
				nim: "456",
			},
			want: &Mahasiswa{
				NIM:  "456",
				Name: "Jane Doe",
				Age:  20,
				Courses: []course.Course{
					{
						ID:   "1",
						Name: "Math",
					},
					{
						ID:   "2",
						Name: "English",
					},
				},
			},
		},
		{
			name: "Test Get Mahasiswa Error",
			args: args{
				nim: "1234",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mahasiswa{
				NIM:     tt.fields.NIM,
				Name:    tt.fields.Name,
				Age:     tt.fields.Age,
				Courses: tt.fields.Courses,
			}
			got, err := m.GetMahasiswa(tt.args.nim)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMahasiswa() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMahasiswa() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMahasiswa_GetMahasiswaByCourse(t *testing.T) {
	type fields struct {
		NIM     string
		Name    string
		Age     int
		Courses []course.Course
	}
	type args struct {
		courseName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Mahasiswa
		wantErr bool
	}{
		{
			name: "Test Get Mahasiswa By Course",
			args: args{
				courseName: "Math",
			},
			want: []Mahasiswa{
				{
					NIM:  "456",
					Name: "Jane Doe",
					Age:  20,
					Courses: []course.Course{
						{
							ID:   "1",
							Name: "Math",
						},
						{
							ID:   "2",
							Name: "English",
						},
					},
				},
			},
		},
		{
			name: "Test Get Mahasiswa By Course Error",
			args: args{
				courseName: "Sport",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mahasiswa{
				NIM:     tt.fields.NIM,
				Name:    tt.fields.Name,
				Age:     tt.fields.Age,
				Courses: tt.fields.Courses,
			}
			got, err := m.GetMahasiswaByCourse(tt.args.courseName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMahasiswaByCourse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMahasiswaByCourse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
