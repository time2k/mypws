package mytypedef

type PWSData struct {
	Action   string `json:"action" form:"action" query:"action" validate:"required"`
	ID       string `json:"ID" form:"ID" query:"ID" validate:"required"`
	PASSWORD string `json:"PASSWORD" form:"PASSWORD" query:"PASSWORD" validate:"required"`
	PWSOutputData
}

type PWSOutputData struct {
	Dateutc          string  `json:"dateutc" form:"dateutc" query:"dateutc"`
	Winddir          int     `json:"winddir" form:"winddir" query:"winddir"`
	Windspeedmph     float64 `json:"windspeedmph" form:"windspeedmph" query:"windspeedmph"`
	Windgustmph      float64 `json:"windgustmph" form:"windgustmph" query:"windgustmph"`
	Windgustdir      int     `json:"windgustdir" form:"windgustdir" query:"windgustdir"`
	Windspdmph_avg2m float64 `json:"windspdmph_avg2m" form:"windspdmph_avg2m" query:"windspdmph_avg2m"`
	Windgustmph_10m  float64 `json:"windgustmph_10m" form:"windgustmph_10m" query:"windgustmph_10m"`
	Windgustdir_10m  int     `json:"windgustdir_10m" form:"windgustdir_10m" query:"windgustdir_10m"`
	Humidity         int     `json:"humidity" form:"humidity" query:"humidity"`
	Dewptf           float64 `json:"dewptf" form:"dewptf" query:"dewptf"`
	Tempf            float64 `json:"tempf" form:"tempf" query:"tempf"`
	Rainin           float64 `json:"rainin" form:"rainin" query:"rainin"`
	Dailyrainin      float64 `json:"dailyrainin" form:"dailyrainin" query:"dailyrainin"`
	Baromin          float64 `json:"baromin" form:"baromin" query:"baromin"`
	UV               int     `json:"UV" form:"UV" query:"UV"`
	Solarradiation   float64 `json:"solarradiation" form:"solarradiation" query:"solarradiation"`
	Indoortempf      float64 `json:"indoortempf" form:"indoortempf" query:"indoortempf"`
	Indoorhumidity   int     `json:"indoorhumidity" form:"indoorhumidity" query:"indoorhumidity"`
	Softwaretype     string  `json:"softwaretype" form:"softwaretype" query:"softwaretype"`
}

type PWSDeviceInfo struct {
	DeviceName string
	DeviceID   string
	Password   string
}

/*
CREATE TABLE `pws_data` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `deviceid` int(11) unsigned NOT NULL DEFAULT '0',
  `dateutc` datetime NOT NULL,
  `createdatelocal` datetime NOT NULL COMMENT 'data insert local time',
  `winddir` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '0-360 instantaneous wind direction',
  `windspeedmph` decimal(10,3) NOT NULL DEFAULT '0.000' COMMENT 'mph instantaneous wind speed',
  `windgustmph` decimal(10,3) NOT NULL DEFAULT '0.000' COMMENT 'mph current wind gust, using software specific time period',
  `windgustdir` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '0-360 using software specific time period',
  `windspdmph_avg2m` decimal(10,3) NOT NULL DEFAULT '0.000' COMMENT 'mph 2 minute average wind speed mph',
  `windgustmph_10m` decimal(10,3) NOT NULL DEFAULT '0.000' COMMENT 'mph past 10 minutes wind gust mph',
  `windgustdir_10m` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '0-360 past 10 minutes wind gust direction',
  `humidity` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '% outdoor humidity 0-100%',
  `dewptf` decimal(10,3) NOT NULL DEFAULT '0.000' COMMENT 'F outdoor dewpoint F',
  `tempf` decimal(10,3) NOT NULL DEFAULT '0.000' COMMENT 'F outdoor temperature',
  `rainin` decimal(10,3) DEFAULT '0.000' COMMENT 'rain inches over the past hour -- the accumulated rainfall in the past 60 min',
  `dailyrainin` decimal(10,3) NOT NULL DEFAULT '0.000' COMMENT 'rain inches so far today in local time',
  `baromin` decimal(10,3) NOT NULL DEFAULT '0.000' COMMENT 'barometric pressure inches',
  `UV` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'UV index',
  `solarradiation` decimal(10,3) NOT NULL DEFAULT '0.000' COMMENT 'W/m^2',
  `indoortempf` decimal(10,3) DEFAULT '0.000' COMMENT 'F indoor temperature F',
  `indoorhumidity` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '% indoor humidity 0-100',
  `softwaretype` varchar(30) NOT NULL DEFAULT '' COMMENT 'ie: WeatherLink, VWS, WeatherDisplay',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/
