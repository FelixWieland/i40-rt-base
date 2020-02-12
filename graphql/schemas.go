package graphql

//Machines GraphQL-Schema
type Machines []Machine

//Machine GraphQL-Schema
type Machine struct {
	ID      int     `description:"Unique Machine" kosmo:"required"`
	Plant   Plant   `description:"Plant of the Machine"`
	Type    string  `description:"Type of the Machine"`
	Name    string  `description:"Name of the Machine"`
	Sensors Sensors `description:"All Sensors of the machine"`
	Orders  Orders  `description:"All Orders that ran on the machine"`
}

//Plant GraphQL-Schema
type Plant struct {
	ID   int    `description:"Plant ID" kosmo:"required"`
	Name string `description:"Plant name" kosmo:"required"`
}

//Sensors GraphQL-Schema
type Sensors []Sensor

//Sensor GraphQL-Schema
type Sensor struct {
	ID          int          `description:"Unique Sensor" kosmo:"required"`
	Name        string       `description:"Name of the Sensor" kosmo:"required"`
	Description string       `description:"Description of the Sensor"`
	Values      SensorValues `description:"Recorded values of the Sensor"`
}

//SensorValues GraphQL-Schema
type SensorValues SensorValue

//SensorValue GraphQL-Schema
type SensorValue struct {
	Value     int    `description:"Value"`
	Type      string `description:"Type of the recorded Value"`
	Timestamp uint   `description:"UNIX Timestamp"`
}

//Orders GraphQL-Schema
type Orders []Order

//Order GraphQL-Schema
type Order struct {
	MachineID  int `description:"ID of the Machine the Order ran" kosmo:"required"`
	OrderStart int `description:"UNIX Timestamp of the Order start" kosmo:"required"`
	OrderEnd   int `description:"UNIX Timestamp of the Order end"`
}

//Dashboards GraphQL-Schema
type Dashboards []Dashboard

//Dashboard GraphQL-Schema
type Dashboard struct {
	ID      int     `description:"Unique Dashboard" kosmo:"required"`
	Name    int     `description:"Name of the Dashboard" kosmo:"required"`
	Widgets Widgets `description:"Widgets in the dashboard"`
}

//Widgets GraphQL-Schema
type Widgets []Widget

//Widget GraphQL-Schema
type Widget struct {
	Type        string         `description:"Type of the Widget"`
	Title       string         `description:"Title of the Widget"`
	Description string         `description:"Description of the Widget"`
	Data        WidgetData     `description:"Data used by the Widget"`
	Position    WidgetPosition `description:"Position of the Widget in the Grid"`
}

//WidgetData GraphQL-Schema
type WidgetData struct {
	Machines Machines `description:"Machines used by the Widget"`
}

//WidgetPosition GraphQL-Schema
type WidgetPosition struct {
	I int `description:"ID"`
	X int `description:"X-Position"`
	Y int `description:"Y-Position"`
	W int `description:"Width"`
	H int `description:"Height"`
}
