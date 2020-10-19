// Time : 2020/10/19 16:50
// Author : Kieran

// logic
package logic

import (
	"begonia2/dispatch"
	"begonia2/dispatch/frame"
)

// service.go something
type WriteFunc = func(result *CallResult, toConnID ...string)

type Service interface {
	caller
	RecvMsg() (msg *Call, wf WriteFunc)
}

func NewService(dp dispatch.Dispatcher) Service {
	c := &service{
		baseLogic: baseLogic{
			dp:       dp,
			waitChan: NewWaitChans(),
		},
	}
	// TODO: add ctx
	return c
}

type service struct {
	baseLogic
}

func (c *service) RecvMsg() (msg *Call, wf WriteFunc) {
	_, f := c.dp.Recv()
	req, ok := f.(*frame.Request)
	if !ok {
		panic("request type error")
	}

	msg = &Call{
		Service: req.Service,
		Fun:     req.Fun,
		Param:   req.Params,
	}

	wf = func(result *CallResult, toConnID ...string) {
		resp := frame.NewResponse(req.ReqId, result.Result, result.Err)
		if toConnID != nil {
			for _, connID := range toConnID {
				c.dp.SendTo(connID, resp)
			}
		} else {
			c.dp.Send(resp)
		}
	}

	return
}