<template>
    <div class="row">
        <div class="page-header" id="banner">
            <div class="row">
            <div class="col-lg-8 col-md-7 col-sm-6">
            </div>
            <div class="col-lg-4 col-md-5 col-sm-6">
            </div>
            </div>
        </div>
        <div class="col-12 col-sm-6">
            <div class="card border-light mb-3 main-data">
            <div class="card-header" style="background-color:#f8f8f8">设备{{devicename}}实时数据
            <p class="datetxt">数据时间：{{dateutc.value}} {{dateutc.unit}}&nbsp;&nbsp;上报时间：{{createdatelocal.value}} {{createdatelocal.unit}}</p>
            </div>
            <div class="card-body" v-if="pwsinfo.tempf">
                <h1 class="card-title">
                <!--<b-img :src="PWSDataType.tempf.icon" fluid></b-img>-->
                {{PWSDataType.tempf.name}}
                {{pwsinfo.tempf.value}}&nbsp;{{pwsinfo.tempf.unit}}
                </h1>
            </div>
            <div class="card-body" v-if="pwsinfo.windspeedmph">
                <h2 class="card-title">
                <!--<b-img :src="PWSDataType.windspeedmph.icon" fluid></b-img>-->
                {{PWSDataType.windspeedmph.name}}&nbsp;{{pwsinfo.windspeedmph.value}}&nbsp;{{pwsinfo.windspeedmph.unit}}
                {{PWSDataType.winddir.name}}&nbsp;{{pwsinfo.winddir.value}}&nbsp;{{pwsinfo.winddir.unit}}
                </h2>
                <h6 class="card-subtitle" style="float:left;padding-left:2px;">
                {{PWSDataType.windgustmph.name}}&nbsp;{{pwsinfo.windgustmph.value}}&nbsp;{{pwsinfo.windgustmph.unit}}
                {{PWSDataType.windgustdir.name}}&nbsp;{{pwsinfo.windgustdir.value}}&nbsp;{{pwsinfo.windgustdir.unit}}
                </h6>
            </div>
            <div class="card-body">
                <span class="badge rounded-pill bg-success" style="font-size:14px;margin:5px;" v-if="pwsinfo.humidity">{{PWSDataType.humidity.name}}&nbsp;{{pwsinfo.humidity.value}}&nbsp;{{pwsinfo.humidity.unit}}</span>
                <span class="badge rounded-pill bg-primary" style="font-size:14px;margin:5px;" v-if="pwsinfo.dewpoint">{{PWSDataType.dewpoint.name}}&nbsp;{{pwsinfo.dewpoint.value}}&nbsp;{{pwsinfo.dewpoint.unit}}</span>
                <span class="badge rounded-pill bg-primary" style="font-size:14px;margin:5px;" v-if="pwsinfo.baromin">{{PWSDataType.baromin.name}}&nbsp;{{pwsinfo.baromin.value}}&nbsp;{{pwsinfo.baromin.unit}}</span>
                <span class="badge rounded-pill bg-primary" style="font-size:14px;margin:5px;" v-if="pwsinfo.rainin">{{PWSDataType.rainin.name}}&nbsp;{{pwsinfo.rainin.value}}&nbsp;{{pwsinfo.rainin.unit}}</span>
                <span class="badge rounded-pill bg-light" style="font-size:14px;margin:5px;" v-if="pwsinfo.UV">{{PWSDataType.UV.name}}&nbsp;{{pwsinfo.UV.value}}&nbsp;{{pwsinfo.UV.unit}}</span>
                <span class="badge rounded-pill bg-light" style="font-size:14px;margin:5px;" v-if="pwsinfo.solarradiation">{{PWSDataType.solarradiation.name}}&nbsp;{{pwsinfo.solarradiation.value}}&nbsp;{{pwsinfo.solarradiation.unit}}</span>
                <span class="badge rounded-pill bg-primary" style="font-size:14px;margin:5px;" v-if="pwsinfo.indoortempf">{{PWSDataType.indoortempf.name}}&nbsp;{{pwsinfo.indoortempf.value}}&nbsp;{{pwsinfo.indoortempf.unit}}</span>
                <span class="badge rounded-pill bg-primary" style="font-size:14px;margin:5px;" v-if="pwsinfo.indoorhumidity">{{PWSDataType.indoorhumidity.name}}&nbsp;{{pwsinfo.indoorhumidity.value}}&nbsp;{{pwsinfo.indoorhumidity.unit}}</span>
                <span class="badge rounded-pill bg-primary" style="font-size:14px;margin:5px;" v-if="pwsinfo.indoordewpoint">{{PWSDataType.indoordewpoint.name}}&nbsp;{{pwsinfo.indoordewpoint.value}}&nbsp;{{pwsinfo.indoordewpoint.unit}}</span>
            </div>
            </div>
        </div>
        <div class="col-12 col-sm-6">
            <div class="card border-light mb-3">
            <div class="card-header">设备{{devicename}}监控截图</div>
            <div class="card-body" v-if="cctvimg">
                <b-img :src="cctvimg" fluid alt=""></b-img>
            </div>
            </div>
        </div>
    </div>
    
</template>

<style>
.datetxt {
    display:inline;font-size:12px;padding-left:20px;
}
.main-data {
    background: rgba(0, 0, 0, .4) url('../assets/bg.png') no-repeat center center;
}
</style>

<script>
import axios from 'axios'
import pic0 from "../assets/clock.png"
import pic1 from "../assets/winddirection.png"
import pic2 from "../assets/windspeed.png"
import pic3 from "../assets/humidity.png"
import pic4 from "../assets/temperature.png"
import pic5 from "../assets/rain.png"
import pic6 from "../assets/airpressure.png"
import pic7 from "../assets/uv.png"
import pic8 from "../assets/light.png"

import dewcacl from "../components/dew.js"

export default {
    name: 'RealtimeData',
    data() {
        return {
            api :'https://mypws.astrofans.net/api/data/get?devicename={devicename}&interval={interval}',
            devicename: this.$route.params.devicename,
            interval: "realtime",
            PWSDataType: {
                dateutc:{name:"数据时间",icon:pic0,unit:"datetime-utc"},
                createdatelocal:{name:"上报时间",icon:pic0,unit:"datetime"},
                tempf:{name:"温度",icon:pic4,unit:"f"},
                humidity:{name:"湿度",icon:pic3,unit:"%"},
                windspeedmph:{name:"风速",icon:pic2,unit:"mph"},
                baromin:{name:"气压",icon:pic6,unit:"inchpress"},
                winddir:{name:"风向",icon:pic1,unit:"dir"},
                windgustmph:{name:"阵风风速",icon:pic2,unit:"mph"},
                windgustdir:{name:"阵风风向",icon:pic1,unit:"dir"},
                rainin:{name:"1小时降水量",icon:pic5,unit:"inchrain"},
                UV:{name:"UV指数",icon:pic7,unit:"uvindex"},
                solarradiation:{name:"光照",icon:pic8,unit:"wm2"},
                indoortempf:{name:"室内温度",icon:pic4,unit:"f"},
                indoorhumidity:{name:"室内湿度",icon:pic3,unit:"%"},
                dewpoint:{name:"露点",icon:pic4,unit:"f"},
                indoordewpoint:{name:"室内露点",icon:pic4,unit:"f"},
            },
            pwsinfo:{},
            dateutc:{},
            createdatelocal:{},
            area:"china",
            cctvimg: "https://mypws.astrofans.net/files/"+this.$route.params.devicename+".jpg"
        }
    },
    async mounted () {
        await this.getRealtimeData()
    },
    methods: {
        async getRealtimeData () {
            this.api = this.api.replace('{devicename}', this.devicename).replace('{interval}', this.interval)
            let res = await axios.get(this.api)
            const ret = {}
            if (res.status === 200 && res.data.data != null) {
                var data = res.data.data
                for(let key in this.PWSDataType) {
                    if(data[key] != undefined) {
                        if(key == "dateutc") {
                            this.dateutc = this.changeValue(key, res.data.data[key], this.area)
                        } else if(key == "createdatelocal") {
                            this.createdatelocal = this.changeValue(key, res.data.data[key], this.area)
                        } else {
                            ret[key] = this.changeValue(key, res.data.data[key], this.area)
                        }
                    }
                }
                if(ret.humidity.value != "--" && ret.tempf.value != "--") {
                    var dp1 = dewcacl.dewPoint(ret.humidity.value, dewcacl.stringToFloat(ret.tempf.value)+dewcacl.C_OFFSET)-dewcacl.C_OFFSET
                    dp1 = dewcacl.truncate(dp1,2,2);
                    ret.dewpoint = {value:dp1, unit:"℃"}
                }
                 else {
                    ret.dewpoint = {value:"--", unit:"℃"}
                }
                if(ret.indoorhumidity.value != "--" && ret.indoortempf.value != "--") {
                    var dp2 = dewcacl.dewPoint(ret.indoorhumidity.value, dewcacl.stringToFloat(ret.indoortempf.value)+dewcacl.C_OFFSET)-dewcacl.C_OFFSET
                    dp2 = dewcacl.truncate(dp2,2,2);
                   ret.indoordewpoint = {value:dp2, unit:"℃"}
                } else {
                    ret.indoordewpoint = {value:"--", unit:"℃"}
                }
                this.pwsinfo = ret
            }
        },
        changeValue(key,value,area) {
            let ret  = {}
            if(area == "china") {
                switch(this.PWSDataType[key].unit) {
                    case "datetime-utc":
                        ret = {value:value,unit:"UTC"}
                        break
                    case "datetime":
                        ret = {value:value,unit:""}
                        break
                    case "dir":
                        if(value==0) {
                            ret = {value:"北",unit:""}
                        } else if(value>0 && value<90) {
                            ret = {value:"东北",unit:""}
                        } else if(value==90) {
                            ret = {value:"东",unit:""}
                        } else if(value>90 && value<180) {
                            ret = {value:"东南",unit:""}
                        } else if(value==180) {
                            ret = {value:"南",unit:""}
                        } else if(value>180 && value<270) {
                            ret = {value:"西南",unit:""}
                        }  else if(value == 270) {
                            ret = {value:"西",unit:""}
                        } else if(value>270 && value<360) {
                            ret = {value:"西北",unit:""}
                        }
                        break
                    case "mph":
                        ret = {value:(value*0.44704).toFixed(2),unit:"米/秒"}
                        break
                    case "f":
                        if(value > -9999) {
                            ret = {value:((value-32)/1.8).toFixed(1),unit:"℃"}
                        } else {
                            ret = {value:"--",unit:"℃"}
                        }
                        break
                    case "%":
                        if(value > -9999) {
                            ret = {value:value,unit:"%"}
                        } else {
                            ret = {value:"--",unit:"%"}
                        }
                        break
                    case "inchrain":
                        ret = {value:(value*25.4).toFixed(1),unit:"毫米"}
                        break
                    case "inchpress":
                        if(value > -9999) {
                            ret = {value:(value*33.8638816).toFixed(1),unit:"百帕"}
                        } else {
                            ret = {value:"--",unit:"百帕"}
                        }
                        break
                    case "uvindex":
                        ret = {value:value,unit:"(1-5)"}
                        break
                    case "wm2":
                        ret = {value:(value/0.0079).toFixed(1),unit:"lux"}
                        break
                }
            }
            return ret
        }
    },
}
</script>
