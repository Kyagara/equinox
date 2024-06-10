package val

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 9fef246d3ece1da9515c8941f7a3c7cd57e330fc

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
