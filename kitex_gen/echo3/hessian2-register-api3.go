package echo3

import (
	"fmt"

	"github.com/kitex-contrib/codec-dubbo/pkg/hessian2"
	codec "github.com/kitex-contrib/codec-dubbo/pkg/iface"
	"github.com/pkg/errors"
)

var objectsApi3 = []interface{}{
	&EchoRequest{},
	&EchoResponse{},
}

func init() {
	hessian2.Register(objectsApi3)
}

func GetTestService3IDLAnnotations() map[string][]string {
	return map[string][]string{
		"JavaClassName": {"kitex.echo3.TestService3"},
	}
}

func (p *EchoRequest) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Message)
	if err != nil {
		return err
	}

	err = e.Encode(p.Obj)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoRequest) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Message)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Obj)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoRequest) JavaClassName() string {
	return "kitex.echo3.EchoRequest"
}

func (p *EchoResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Message)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Message)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoResponse) JavaClassName() string {
	return "kitex.echo3.EchoResponse"
}

func (p *TestService3EchoArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestService3EchoArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestService3EchoResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestService3EchoResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}
