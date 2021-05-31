package sciconst

import "math"

type Constant struct {
	Description string
	Value       float64
	Uncertainty float64
	Unit        string
}

// mathematical constants
const (
	GoldenRatio = 2.118033988749895 //(1 + math.Sqrt(5)) / 2
)

// SI multipliers
const (
	Yotta = 1e24
	Zetta = 1e21
	Exa   = 1e18
	Peta  = 1e15
	Tera  = 1e12
	Giga  = 1e9
	Mega  = 1e6
	Kilo  = 1e3
	Hecto = 1e2
	Deka  = 1e1
	Deci  = 1e-1
	Centi = 1e-2
	Milli = 1e-3
	Micro = 1e-6
	Nano  = 1e-9
	Pico  = 1e-12
	Femto = 1e-15
	Atto  = 1e-18
	Zepto = 1e-21
)

// binary prefixes
const (
	Kibi = 1024                      // 2**10
	Mebi = 1048576                   // 2**20
	Gibi = 1073741824                // 2**30
	Tebi = 1099511627776             // 2**40
	Pebi = 1125899906842624          // 2**50
	Exbi = 1152921504606846976       // 2**60
	Zebi = 1180591620717411303424    // 2**70
	Yobi = 1208925819614629174706176 // 2**80
)

// physical constants
const (
	C                     = 299792458 //speed of light in vacuum
	SpeedOfLight          = C
	Mu0                   = 1.25663706212e-6 //vacuum mag. permeability
	Epsilon0              = 8.8541878128e-12 //vacuum electric permittivity
	H                     = 6.62607015e-34   //Planck constant
	Planck                = H
	Hbar                  = H / (2 * math.Pi)
	G                     = 6.67430e-11 //Newtonian constant of gravitation
	GravitationalConstant = G
	Gacc                  = 9.80665         //standard acceleration of gravity
	E                     = 1.602176634e-19 //elementary charge
	ElementaryCharge      = E
	R                     = 8.314462618 //molar gas constant
	GasConstant           = R
	Alpha                 = 7.2973525693e-3 //fine-structure constant
	FineStructure         = Alpha
	NA                    = 6.02214076e23 //Avogadro constant
	Avogadro              = NA
	K                     = 1.380649e-23 //Boltzmann constant
	Boltzmann             = K
	Sigma                 = 5.670374419e-8 //Stefan-Boltzmann constant
	StefanBoltzmann       = Sigma
	Wien                  = 2.897771955e-3  //Wien wavelength displacement law constant
	Rydberg               = 10973731.568160 //Rydberg constant
)

// mass in kg
const (
	Gram      = 1e-3
	MetricTon = 1e3
	Grain     = 64.79891e-6
	Lb        = 7000 * Grain // avoirdupois
	Pound     = Lb
	Blob      = Pound * Gacc / 0.0254 // lbf*s**2/in (added in 1.0.0)
	Slinch    = Blob
	Slug      = Blob / 12 // lbf*s**2/foot (added in 1.0.0)
	Oz        = Pound / 16
	Ounce     = Oz
	Stone     = 14 * Pound
	LongTon   = 2240 * Pound
	ShortTon  = 2000 * Pound

	TroyOunce = 480 * Grain // only for metals / gems
	TroyPound = 12 * TroyOunce
	Carat     = 200e-6

	Me = 9.1093837015e-31
	Mp = 1.67262192369e-27
	Mn = 1.67492749804e-27
	Mu = 1.66053906660e-27
)

// angle in rad
const (
	Degree = math.Pi / 180
	Arcmin = Degree / 60
	Arcsec = Arcmin / 60
)

// time in second
const (
	Minute     = 60.0
	Hour       = 60 * Minute
	Day        = 24 * Hour
	Week       = 7 * Day
	Year       = 365 * Day
	JulianYear = 365.25 * Day
)

// length in meter
const (
	Inch             = 0.0254
	Foot             = 12 * Inch
	Yard             = 3 * Foot
	Mile             = 1760 * Yard
	Mil              = Inch / 1000
	Pt               = Inch / 72 // typography
	Point            = Pt
	SurveyFoot       = 1200.0 / 3937
	SurveyMile       = 5280 * SurveyFoot
	NauticalMile     = 1852.0
	Fermi            = 1e-15
	Angstrom         = 1e-10
	Micron           = 1e-6
	AU               = 149597870700.0
	AstronomicalUnit = AU
	LightYear        = JulianYear * C
	Parsec           = AU / Arcsec
)

// pressure in pascal
const (
	Atm        = 101325 //standard atmosphere
	Atmosphere = Atm
	Bar        = 1e5
	Torr       = Atm / 760
	MmHg       = Torr
	Psi        = Pound * Gacc / (Inch * Inch)
)

// area in meter**2
const (
	Hectare = 1e4
	Acre    = 43560 * Foot * Foot
)

// volume in meter**3
const (
	Litre         = 1e-3
	Liter         = Litre
	Gallon        = 231 * Inch * Inch * Inch // US
	GallonUS      = Gallon
	Pint          = GallonUS / 8
	FluidOunce    = GallonUS / 128
	FluidOunceUS  = FluidOunce
	Bbl           = 42 * GallonUS // for oil
	Barrel        = Bbl
	GallonImp     = 4.54609e-3 // UK
	FluidOunceImp = GallonImp / 160
)

// speed in meter per second
const (
	Kmh          = 1e3 / Hour
	Mph          = Mile / Hour
	Mach         = 340.5 // approx value at 15 degrees in 1 atm. is this a common value?
	SpeedOfSound = Mach
	Knot         = NauticalMile / Hour
)

// temperature in kelvin
const (
	ZeroCelsius      = 273.15
	DegreeFahrenheit = 1 / 1.8 // only for differences
)

// energy in joule
const (
	EV           = ElementaryCharge // * 1 Volt
	ElectronVolt = EV
	Calorie      = 4.184
	CalorieTh    = Calorie
	CalorieIT    = 4.1868
	Erg          = 1e-7
	BtuTh        = Pound * DegreeFahrenheit * CalorieTh / Gram
	Btu          = Pound * DegreeFahrenheit * CalorieIT / Gram
	BtuIT        = Btu
	TonTNT       = 1e9 * CalorieTh
)

// power in watt
const (
	Hp         = 550 * Foot * Pound * Gacc
	Horsepower = Hp
)

// force in newton
const (
	Dyn           = 1e-5
	Dyne          = Dyn
	Lbf           = Pound * Gacc
	PoundForce    = Lbf
	Kgf           = Gacc // * 1 kg
	KilogramForce = Kgf
)
