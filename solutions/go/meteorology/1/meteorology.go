package meteorology
import "fmt"
type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (tu TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	if int(tu) < 0 || int(tu) >= len(units) {
		return ""
	}
	return units[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (te Temperature) String() string{
    return fmt.Sprintf("%d %s", te.degree, te.unit.String())
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (su SpeedUnit) String() string{
    units := []string{"km/h","mph",}
    if int(su) < 0 || int(su) >= len(units){
        return ""
    }
    return units[su]
}
type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (sd Speed) String() string{
    return fmt.Sprintf("%d %s", sd.magnitude, sd.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData

func (me MeteorologyData) String() string{
    return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", me.location, me.temperature.String(), me.windDirection, me.windSpeed.String(), me.humidity)
}