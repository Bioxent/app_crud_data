package main

import (
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func Test_createCustomerHandler(t *testing.T) {

	type args struct {
		c  *fiber.Ctx
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "get correct status",
			args: assert.Condition()
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createCustomerHandler(tt.args.c, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("createCustomerHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}



func Test_editCustomerHandler(t *testing.T) {
	type args struct {
		c  *fiber.Ctx
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := editCustomerHandler(tt.args.c, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("editCustomerHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}



