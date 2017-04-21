// Auto-generated by avdl-compiler v1.3.13 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/signup.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type SignupRes struct {
	PassphraseOk bool `codec:"passphraseOk" json:"passphraseOk"`
	PostOk       bool `codec:"postOk" json:"postOk"`
	WriteOk      bool `codec:"writeOk" json:"writeOk"`
}

type CheckUsernameAvailableArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Username  string `codec:"username" json:"username"`
}

type SignupArg struct {
	SessionID   int        `codec:"sessionID" json:"sessionID"`
	Email       string     `codec:"email" json:"email"`
	InviteCode  string     `codec:"inviteCode" json:"inviteCode"`
	Passphrase  string     `codec:"passphrase" json:"passphrase"`
	Username    string     `codec:"username" json:"username"`
	DeviceName  string     `codec:"deviceName" json:"deviceName"`
	DeviceType  DeviceType `codec:"deviceType" json:"deviceType"`
	StoreSecret bool       `codec:"storeSecret" json:"storeSecret"`
	SkipMail    bool       `codec:"skipMail" json:"skipMail"`
	GenPGPBatch bool       `codec:"genPGPBatch" json:"genPGPBatch"`
}

type InviteRequestArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Email     string `codec:"email" json:"email"`
	Fullname  string `codec:"fullname" json:"fullname"`
	Notes     string `codec:"notes" json:"notes"`
}

type CheckInvitationCodeArg struct {
	SessionID      int    `codec:"sessionID" json:"sessionID"`
	InvitationCode string `codec:"invitationCode" json:"invitationCode"`
}

type GetInvitationCodeArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type SignupInterface interface {
	CheckUsernameAvailable(context.Context, CheckUsernameAvailableArg) error
	Signup(context.Context, SignupArg) (SignupRes, error)
	InviteRequest(context.Context, InviteRequestArg) error
	CheckInvitationCode(context.Context, CheckInvitationCodeArg) error
	GetInvitationCode(context.Context, int) (string, error)
}

func SignupProtocol(i SignupInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.signup",
		Methods: map[string]rpc.ServeHandlerDescription{
			"checkUsernameAvailable": {
				MakeArg: func() interface{} {
					ret := make([]CheckUsernameAvailableArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]CheckUsernameAvailableArg)
					if !ok {
						err = rpc.NewTypeError((*[]CheckUsernameAvailableArg)(nil), args)
						return
					}
					err = i.CheckUsernameAvailable(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"signup": {
				MakeArg: func() interface{} {
					ret := make([]SignupArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SignupArg)
					if !ok {
						err = rpc.NewTypeError((*[]SignupArg)(nil), args)
						return
					}
					ret, err = i.Signup(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"inviteRequest": {
				MakeArg: func() interface{} {
					ret := make([]InviteRequestArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]InviteRequestArg)
					if !ok {
						err = rpc.NewTypeError((*[]InviteRequestArg)(nil), args)
						return
					}
					err = i.InviteRequest(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"checkInvitationCode": {
				MakeArg: func() interface{} {
					ret := make([]CheckInvitationCodeArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]CheckInvitationCodeArg)
					if !ok {
						err = rpc.NewTypeError((*[]CheckInvitationCodeArg)(nil), args)
						return
					}
					err = i.CheckInvitationCode(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getInvitationCode": {
				MakeArg: func() interface{} {
					ret := make([]GetInvitationCodeArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetInvitationCodeArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetInvitationCodeArg)(nil), args)
						return
					}
					ret, err = i.GetInvitationCode(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type SignupClient struct {
	Cli rpc.GenericClient
}

func (c SignupClient) CheckUsernameAvailable(ctx context.Context, __arg CheckUsernameAvailableArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.signup.checkUsernameAvailable", []interface{}{__arg}, nil)
	return
}

func (c SignupClient) Signup(ctx context.Context, __arg SignupArg) (res SignupRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.signup.signup", []interface{}{__arg}, &res)
	return
}

func (c SignupClient) InviteRequest(ctx context.Context, __arg InviteRequestArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.signup.inviteRequest", []interface{}{__arg}, nil)
	return
}

func (c SignupClient) CheckInvitationCode(ctx context.Context, __arg CheckInvitationCodeArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.signup.checkInvitationCode", []interface{}{__arg}, nil)
	return
}

func (c SignupClient) GetInvitationCode(ctx context.Context, sessionID int) (res string, err error) {
	__arg := GetInvitationCodeArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.signup.getInvitationCode", []interface{}{__arg}, &res)
	return
}
