package arithmetic_sequence

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArithmeticSequence_SumOfFiveTerms(t *testing.T) {
	type args struct {
		firstTerm float64
		lastTerm  float64
	}
	tests := []struct {
		name     string
		args     args
		expected float64
		wantErr  bool
		errMsg   string
	}{
		{
			name: "firstTerm: 1, lastTerm: 5",
			args: args{
				firstTerm: 1,
				lastTerm:  5,
			},
			expected: 15,
		},
		{
			name: "firstTerm: 0, lastTerm: 0",
			args: args{
				firstTerm: 0,
				lastTerm:  0,
			},
			wantErr: true,
			errMsg:  "firstTerm and lastTerm cannot be zero",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// arrange
			a := NewArithmeticSequence()
			a.SetFirstTerm(tt.args.firstTerm)
			a.SetLastTerm(tt.args.lastTerm)

			// act
			got, err := a.SumOfFiveTerms()
			if tt.wantErr {
				assert.EqualError(t, err, tt.errMsg)
				return
			}

			// assert
			assert.Equal(t, tt.expected, got)

		})
	}
}

func TestArithmeticSequence_SumOfTenTerms(t *testing.T) {
	type args struct {
		firstTerm float64
		lastTerm  float64
	}
	tests := []struct {
		name     string
		args     args
		expected float64
		wantErr  bool
		errMsg   string
	}{
		{
			name: "firstTerm: 1, lastTerm: 5",
			args: args{
				firstTerm: 1,
				lastTerm:  5,
			},
			expected: 30,
		},
		{
			name: "firstTerm: 0, lastTerm: 0",
			args: args{
				firstTerm: 0,
				lastTerm:  0,
			},
			wantErr: true,
			errMsg:  "firstTerm and lastTerm cannot be zero",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// arrange
			a := NewArithmeticSequence()
			a.SetFirstTerm(tt.args.firstTerm)
			a.SetLastTerm(tt.args.lastTerm)

			// act
			got, err := a.SumOfTenTerms()
			if tt.wantErr {
				assert.EqualError(t, err, tt.errMsg)
				return
			}

			// assert
			assert.Equal(t, tt.expected, got)

		})
	}
}

type ArithmeticSequence struct {
	firstTerm float64
	lastTerm  float64
}

func NewArithmeticSequence() *ArithmeticSequence {
	return &ArithmeticSequence{}
}

func (a *ArithmeticSequence) SetFirstTerm(firstTerm float64) {
	a.firstTerm = firstTerm
}

func (a *ArithmeticSequence) SetLastTerm(lastTerm float64) {
	a.lastTerm = lastTerm
}

func (a *ArithmeticSequence) SumOfFiveTerms() (float64, error) {
	return a.sumOfNTerms(5)
}

func (a *ArithmeticSequence) SumOfTenTerms() (float64, error) {
	return a.sumOfNTerms(10)
}

func (a *ArithmeticSequence) sumOfNTerms(n float64) (float64, error) {

	if a.firstTerm == 0 && a.lastTerm == 0 {
		return 0, errors.New("firstTerm and lastTerm cannot be zero")
	}

	return (a.firstTerm + a.lastTerm) * n / 2, nil
}
