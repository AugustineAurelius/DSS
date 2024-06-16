package retry

import (
	"context"
	"errors"
	"reflect"
	"time"
)

func WrapForRetry(f any, params ...any) (func() error, error) {
	if reflect.TypeOf(f).Kind() != reflect.Func {
		return nil, errors.New("type is not a func")
	}

	return func() error {
		return func(args ...any) error {
			vargs := make([]reflect.Value, len(args))
			for n, v := range args {
				vargs[n] = reflect.ValueOf(v)
			}
			out := reflect.ValueOf(f).Call(vargs)
			err, _ := out[0].Interface().(error)

			return err
		}(params...)
	}, nil
}

func Do(attempts int, t time.Duration, f func() error) error {

	for i := 0; i < attempts; i++ {
		err := f()
		if err == nil {
			return nil
		}
		<-time.After(t)
	}

	return errors.New("all attempts wasted")
}

func WrapWithCtx(f any, params ...any) (func(context.Context) error, error) {

	return func(ctx context.Context) error {
		return func(ctx context.Context, args ...any) error {
			vargs := make([]reflect.Value, len(args)+1)
			vargs[0] = reflect.ValueOf(ctx)

			for n, v := range args {
				vargs[n+1] = reflect.ValueOf(v)
			}
			out := reflect.ValueOf(f).Call(vargs)
			err, _ := out[0].Interface().(error)

			return err
		}(ctx, params...)

	}, nil
}

func DoCtx(t time.Duration, f func(context.Context) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	done := make(chan error)
	go func() {
		defer close(done)
		done <- f(ctx)
	}()

	select {
	case err := <-done:
		return err

	case <-ctx.Done():
		return ctx.Err()

	}
}

func Loop(f func() error, t time.Duration) {
	for range time.Tick(t) {
		err := f()
		if err != nil {
			continue
		}
	}
}

func InfLoop(f func() error) {
	for {
		err := f()
		if err != nil {
			continue
		}
	}

}
