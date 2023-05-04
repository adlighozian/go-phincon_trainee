package usecase

import (
	"contact-go/mocks"
	"contact-go/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_contactUsecase_List(t *testing.T) {
	tests := []struct {
		name       string
		repoResult []model.Contact
		repoErr    error
		want       []model.Contact
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			repoResult: []model.Contact{
				{ID: 1, Name: "jaguar", NoTelp: "999-888-7777"},
				{ID: 2, Name: "Jane_Smith", NoTelp: "555-555-5678"},
				{ID: 3, Name: "jangkrik", NoTelp: "000-000-0000"},
			},
			repoErr: nil,
			want: []model.Contact{
				{ID: 1, Name: "jaguar", NoTelp: "999-888-7777"},
				{ID: 2, Name: "Jane_Smith", NoTelp: "555-555-5678"},
				{ID: 3, Name: "jangkrik", NoTelp: "000-000-0000"},
			},
			wantErr: false,
		},
		{
			name:       "failed",
			repoResult: []model.Contact{},
			repoErr:    assert.AnError,
			want:       []model.Contact{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContactRepo := mocks.NewContactRepository(t)

			mockContactRepo.On("List").Return(tt.repoResult, tt.repoErr)

			uc := NewContactUsecase(mockContactRepo)

			got, err := uc.List()

			if assert.Equal(t, tt.wantErr, err != nil, "contactUsecase.List() error = %v, wantErr %v", err, tt.wantErr) {
				assert.Equal(t, tt.want, got, "contactUsecase.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contactUsecase_Add(t *testing.T) {
	type args struct {
		req *model.ContactRequest
	}
	tests := []struct {
		name       string
		args       args
		repoResult *model.Contact
		repoErr    error
		want       *model.Contact
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				req: &model.ContactRequest{
					Name:   "test",
					NoTelp: "222-222-3232",
				},
			},
			repoResult: &model.Contact{
				ID:     1,
				Name:   "test",
				NoTelp: "222-222-3232",
			},
			repoErr: nil,
			want: &model.Contact{
				ID:     1,
				Name:   "test",
				NoTelp: "222-222-3232",
			},
			wantErr: false,
		},
		{
			name: "failed",
			args: args{
				req: &model.ContactRequest{
					NoTelp: "222-222-3232",
				},
			},
			repoResult: nil,
			repoErr:    assert.AnError,
			want:       nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContact := new(model.Contact)
			mockContact.Name = tt.args.req.Name
			mockContact.NoTelp = tt.args.req.NoTelp

			mockContactRepo := mocks.NewContactRepository(t)

			mockContactRepo.On("Add", mockContact).Return(tt.repoResult, tt.repoErr)

			uc := NewContactUsecase(mockContactRepo)

			got, err := uc.Add(tt.args.req)

			if assert.Equal(t, tt.wantErr, err != nil, "contactUsecase.Add() error = %v, wantErr %v", err, tt.wantErr) {
				assert.Equal(t, tt.want, got, "contactUsecase.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contactUsecase_Detail(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name       string
		args       args
		repoResult *model.Contact
		repoErr    error
		want       *model.Contact
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				id: 1,
			},
			repoResult: &model.Contact{
				ID:     1,
				Name:   "test",
				NoTelp: "222-222-3232",
			},
			repoErr: nil,
			want: &model.Contact{
				ID:     1,
				Name:   "test",
				NoTelp: "222-222-3232",
			},
			wantErr: false,
		},
		{
			name: "failed",
			args: args{
				id: 0,
			},
			repoResult: nil,
			repoErr:    assert.AnError,
			want:       nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContactRepo := mocks.NewContactRepository(t)

			mockContactRepo.On("Detail", tt.args.id).Return(tt.repoResult, tt.repoErr)

			uc := NewContactUsecase(mockContactRepo)

			got, err := uc.Detail(tt.args.id)

			if assert.Equal(t, tt.wantErr, err != nil, "contactUsecase.Detail() error = %v, wantErr %v", err, tt.wantErr) {
				assert.Equal(t, tt.want, got, "contactUsecase.Detail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contactUsecase_Update(t *testing.T) {
	type args struct {
		id  int64
		req *model.ContactRequest
	}
	tests := []struct {
		name       string
		args       args
		repoResult *model.Contact
		repoErr    error
		want       *model.Contact
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				id: 1,
				req: &model.ContactRequest{
					Name:   "test",
					NoTelp: "222-222-4444",
				},
			},
			repoResult: &model.Contact{
				ID:     1,
				Name:   "test",
				NoTelp: "222-222-4444",
			},
			repoErr: nil,
			want: &model.Contact{
				ID:     1,
				Name:   "test",
				NoTelp: "222-222-4444",
			},
			wantErr: false,
		},
		{
			name: "failed",
			args: args{
				id: 1,
				req: &model.ContactRequest{
					NoTelp: "222-222-3232",
				},
			},
			repoResult: nil,
			repoErr:    assert.AnError,
			want:       nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContact := new(model.Contact)
			mockContact.Name = tt.args.req.Name
			mockContact.NoTelp = tt.args.req.NoTelp

			mockContactRepo := mocks.NewContactRepository(t)

			mockContactRepo.On("Update", tt.args.id, mockContact).Return(tt.repoResult, tt.repoErr)

			uc := NewContactUsecase(mockContactRepo)

			got, err := uc.Update(tt.args.id, tt.args.req)

			if assert.Equal(t, tt.wantErr, err != nil, "contactUsecase.Update() error = %v, wantErr %v", err, tt.wantErr) {
				assert.Equal(t, tt.want, got, "contactUsecase.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contactUsecase_Delete(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		repoErr error
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				id: 1,
			},
			repoErr: nil,
			wantErr: false,
		},
		{
			name: "failed",
			args: args{
				id: 0,
			},
			repoErr: assert.AnError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContactRepo := mocks.NewContactRepository(t)

			mockContactRepo.On("Delete", tt.args.id).Return(tt.repoErr)

			uc := NewContactUsecase(mockContactRepo)

			err := uc.Delete(tt.args.id)

			assert.Equal(t, tt.wantErr, err != nil, "contactUsecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}
