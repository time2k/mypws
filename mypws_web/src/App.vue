<template>
  <div id="app">
    <div class="container" style="">
    <nav class="navbar navbar-expand-lg navbar-dark bg-primary">
        <div class="container-fluid">
            <a class="navbar-brand" href="#">MyPWS</a>
            <button class="navbar-toggler collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#navbarColor01" aria-controls="navbarColor01" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
            </button>

            <div class="collapse navbar-collapse show" id="navbarColor01">
            <ul class="navbar-nav me-auto">
                <!--<li class="nav-item">
                  <router-link to="/home" class="nav-link" active-class="active">首页</router-link>
                </li>-->
                <li class="nav-item">
                <router-link :to="'/realtime/'+this.$route.params.devicename" class="nav-link" active-class="active">实时数据</router-link>
                </li>
                <li class="nav-item">
                <router-link :to="'/history/'+this.$route.params.devicename" class="nav-link" active-class="active">历史数据</router-link>
                </li>
                <!--<li class="nav-item">
                <router-link to="/about" class="nav-link" active-class="active">关于</router-link>
                </li>-->
            </ul>
            </div>
        </div>
    </nav>
    <router-view />
    </div>
  </div>
</template>

<script>
  export default {
    name: "app",
  }
</script>