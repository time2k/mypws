<template>
  <div class="row">
    <div class="col-12 col-sm-6">
    <line-chart v-if="loaded" :chartdata="chartdata['tempf']" :options="options" />
    </div>
    <div class="col-12 col-sm-6">
    <line-chart v-if="loaded" :chartdata="chartdata['humidity']" :options="options" />
    </div>
    <div class="col-12 col-sm-6">
    <line-chart v-if="loaded" :chartdata="chartdata['rainin']" :options="options" />
    </div>
    <div class="col-12 col-sm-6">
    <line-chart v-if="loaded" :chartdata="chartdata['winddir']" :options="options" />
    </div>
    <div class="col-12 col-sm-6">
    <line-chart v-if="loaded" :chartdata="chartdata['windspeedmph']" :options="options" />
    </div>
    <div class="col-12 col-sm-6">
    <line-chart v-if="loaded" :chartdata="chartdata['windgustdir']" :options="options" />
    </div>
    <div class="col-12 col-sm-6">
    <line-chart v-if="loaded" :chartdata="chartdata['windgustmph']" :options="options" />
    </div>
    <div class="col-12 col-sm-6">
    <line-chart v-if="loaded" :chartdata="chartdata['solarradiation']" :options="options" />
    </div>
</div>
</template>

<style>
</style>

<script>
  import axios from 'axios'
  import LineChart from './LineChart.vue'

  export default {
    name: "HistoryData",
    components: {
      LineChart
    },
    data() {
        return {
            api :'https://mypws.astrofans.net/api/data/get?devicename={devicename}&interval={interval}',
            devicename: this.$route.params.devicename,
            interval: "weekly",
            PWSDataType: {
                dateutc:{name:"数据时间",unit:"datetime-utc"},
                createdatelocal:{name:"上报时间",unit:"datetime"},
                winddir:{name:"风向",unit:"dir"},
                windspeedmph:{name:"风速",unit:"mph"},
                windgustmph:{name:"阵风风速",unit:"mph"},
                windgustdir:{name:"阵风风向",unit:"dir"},
                humidity:{name:"湿度",unit:"%"},
                tempf:{name:"温度",unit:"f"},
                rainin:{name:"1小时降水量",unit:"inchrain"},
                baromin:{name:"气压",unit:"inchpress"},
                UV:{name:"UV指数",unit:"uvindex"},
                solarradiation:{name:"光照",unit:"wm2"},
                indoortempf:{name:"室内温度",unit:"f"},
                indoorhumidity:{name:"室内湿度",unit:"%"},
            },
            area:"china",
            loaded: false,
            chartdata: {},
            options: {}
        }
    },
    async mounted () {
        await this.getHistoryData()
    },
    methods: {
        async getHistoryData () {
            this.api = this.api.replace('{devicename}', this.devicename).replace('{interval}', this.interval)
            let res = await axios.get(this.api)
            const ret = {}
            if (res.status === 200 && res.data.code == 0 && res.data.data.list != null) {
                for(var i in res.data.data.list) {
                    var eachdata = res.data.data.list[i]
                    
                    for(let key in eachdata) {
                        if(this.PWSDataType[key] != undefined) {
                            if(ret[key] == undefined) {
                                ret[key] = {datasets:[],labels:[]}
                                ret[key].datasets.push({data:[],label:"",fill:false,borderColor:"#3399cc"})
                            }
                            let ev = this.changeValue(key, eachdata[key], this.area)
                            ret[key].datasets[0].data.push(ev.value)
                            ret[key].datasets[0].label = this.PWSDataType[key].name + "("+ev.unit+")"
                        
                            ret[key].labels.push(eachdata.dateutc)
                        }

                    }
                }
            }
            this.chartdata = ret
            this.loaded = true
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
                        ret = {value:value,unit:""}
                        break
                    case "mph":
                        ret = {value:(value*0.44704).toFixed(2),unit:"米/秒"}
                        break
                    case "f":
                        if(value > -9999) {
                            ret = {value:((value-32)/1.8).toFixed(2),unit:"摄氏度"}
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
    }
  }
</script>

<style>
  .small {
    max-width: 600px;
    margin:  150px auto;
  }
</style>