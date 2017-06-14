<template>
  <div class="api">
    <h2>Use the service!</h2>
    <h3>it's pretty useful wherever you need randomly generated names</h3>
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
