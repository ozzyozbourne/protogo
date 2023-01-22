package enum

import "github/ozzy/grpcgo/protoOut"

func DoEnum() *protoOut.Enumerations {
	return &protoOut.Enumerations{EyeColor: protoOut.EyeColor_EYE_COLOR_BROWN}
}
