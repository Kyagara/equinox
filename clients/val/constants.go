package val

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 8096d0e7127558ddf4df50a0227b4100b5d54a2f

// Platform routes for Valorant.
type PlatformRoute string

const (
	// Asia-Pacific.
	AP PlatformRoute = "ap"
	// Brazil.
	BR PlatformRoute = "br"
	// Special esports platform.
	ESPORTS PlatformRoute = "esports"
	// Europe.
	EU PlatformRoute = "eu"
	// Korea.
	KR PlatformRoute = "kr"
	// Latin America.
	LATAM PlatformRoute = "latam"
	// North America.
	NA PlatformRoute = "na"
)

func (route PlatformRoute) String() string {
	switch route {
	case AP:
		return "ap"
	case BR:
		return "br"
	case ESPORTS:
		return "esports"
	case EU:
		return "eu"
	case KR:
		return "kr"
	case LATAM:
		return "latam"
	case NA:
		return "na"
	default:
		return string(route)
	}
}
