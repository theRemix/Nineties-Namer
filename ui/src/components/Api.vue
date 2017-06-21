<template>
  <div class="api">
    <api-tab-panel></api-tab-panel>
  </div>
</template>

<script>
import ApiTabPanel from './ApiTabPanel/Panel'

export default {
  name: 'api',
  components: {
    ApiTabPanel
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h2, h3 {
  font-weight: normal;
}
</style>
