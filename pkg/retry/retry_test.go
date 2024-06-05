package retry

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type tS struct {
	Name string
}

func TestRetry(t *testing.T) {
	testStruct := &tS{"VAsiliy"}
	stringParam := "123"
	floatParam := 123.0

	f := func(ts *tS, key string, amount float64) error {
		if key != stringParam {
			t.Error("string param is not equal", key, stringParam)
		}

		if amount != floatParam {
			t.Error("float param is not equal", amount, floatParam)
		}
		fmt.Println(ts)

		if rand.Intn(3) > 1 {
			return nil
		}

		return errors.New("some err")
	}

	wraped, err := WrapForRetry(f, testStruct, stringParam, floatParam)
	if err != nil {
		return
	}

	err = Do(3, 1, wraped)
	if err != nil {
		t.Error(err)
		return
	}

}

func TestWithContext(t *testing.T) {

	f := func(ctx context.Context, key string) error {

		fmt.Println(ctx.Deadline())
		fmt.Println(key)
		return nil
	}

	cf, err := WrapWithCtx(f, "123")
	if err != nil {
		return
	}

	err = DoCtx(time.Second, cf)
	if err != nil {
		t.Error(err)
	}
}
