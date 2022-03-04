<template>
<div class="container" style="">
    <div class="page-header" id="banner">
        <div class="row">
          <div class="col-lg-8 col-md-7 col-sm-6">
            <h1>pws实时数据</h1>
            <p class="lead">{{devicename}}</p>
          </div>
          <div class="col-lg-4 col-md-5 col-sm-6">
          </div>
        </div>
      </div>
  <div class="row">
    <div class="col-6 col-sm-3" v-for="(item,key) in pwsinfo" :key="key">
        <div class="card border-primary mb-3">
        <div class="card-header" style="text-align:center"><img :src="PWSDataType[key].icon" class="img-responsive"><p>{{PWSDataType[key].name}}</p></div>
        <div class="card-body">
            <p class="card-text" style="text-align:center" :id="key">{{item.value}}&nbsp;{{item.unit}}</p>
        </div>
        </div>
    </div>
  </div>
</div>
</template>

<style>

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


export default {
    name: 'RealtimeData',
    data() {
        return {
            api :'https://mypws.astrofans.net/data/get?devicename={devicename}&interval={interval}',
            devicename: this.$route.query.devicename,
            interval: "realtime",
            PWSDataType: {
                dateutc:{name:"数据时间",icon:pic0,unit:"datetime-utc"},
                createdatelocal:{name:"上报时间",icon:pic0,unit:"datetime"},
                winddir:{name:"风向",icon:pic1,unit:"dir"},
                windspeedmph:{name:"风速",icon:pic2,unit:"mph"},
                windgustmph:{name:"阵风风速",icon:pic2,unit:"mph"},
                windgustdir:{name:"阵风风向",icon:pic1,unit:"dir"},
                humidity:{name:"湿度",icon:pic3,unit:"%"},
                tempf:{name:"温度",icon:pic4,unit:"f"},
                rainin:{name:"1小时降水量",icon:pic5,unit:"inchrain"},
                baromin:{name:"气压",icon:pic6,unit:"inchpress"},
                UV:{name:"UV指数",icon:pic7,unit:"uvindex"},
                solarradiation:{name:"光照",icon:pic8,unit:"wm2"},
                indoortempf:{name:"室内温度",icon:pic4,unit:"f"},
                indoorhumidity:{name:"室内湿度",icon:pic3,unit:"%"},
            },
            pwsinfo:{},
            area:"china"
        }
    },
    async created () {
        await this.getRealtimeData()
    },
    methods: {
        async getRealtimeData () {
            this.api = this.api.replace('{devicename}', this.devicename).replace('{interval}', this.interval)
            let res = await axios.get(this.api)
            const ret = {}
            if (res.status === 200 && res.data.data != null) {
                for(let key in res.data.data) {
                    if(this.PWSDataType[key] != undefined) {
                        ret[key] = this.changeValue(key, res.data.data[key], this.area)
                    }
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
                            ret = {value:((value-32)/1.8).toFixed(1),unit:"摄氏度"}
                        } else {
                            ret = {value:"--",unit:"摄氏度"}
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
