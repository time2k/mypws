<template>
<div class="container" style="">
    <Navbar></Navbar>
</div>
</template>

<script>
export default {
  name: 'Home',
  components: {
    'Navbar': () => import('./navbar'),
  },
}
</script>
