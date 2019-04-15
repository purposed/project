package vcs_test

import (
	"testing"

	"github.com/dalloriam/project/vcs"
)

func Test_ParseRepository(t *testing.T) {
	type testCase struct {
		name string

		infoStr string

		expectedName     string
		expectedOwner    string
		expectedProvider string
		expectedURL      string

		wantErr bool
	}

	cases := []testCase{
		{"simple github pattern", "github.com/dalloriam/project", "project", "dalloriam", "github.com", "git@github.com:dalloriam/project.git", false},
		{"bad pattern", "dalloriam/project", "", "", "", "", true},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			outRepo, err := vcs.ParseRepository(tCase.infoStr)

			if (err != nil) != tCase.wantErr {
				t.Errorf("expected error: %v, got err=%v", tCase.wantErr, err)
				return
			}
			if err != nil {
				return
			}

			if outRepo.Name != tCase.expectedName {
				t.Errorf("expected name=%s, got %s", tCase.expectedName, outRepo.Name)
			}

			if outRepo.Owner != tCase.expectedOwner {
				t.Errorf("expected owner=%s, got %s", tCase.expectedOwner, outRepo.Owner)
			}

			if outRepo.Provider != tCase.expectedProvider {
				t.Errorf("expected provider=%s, got %s", tCase.expectedProvider, outRepo.Provider)
			}

			if outRepo.URL != tCase.expectedURL {
				t.Errorf("expected url=%s, got %s", tCase.expectedURL, outRepo.URL)
			}
		})
	}
}
