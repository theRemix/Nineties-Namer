<template>
  <div class="tab-panel">
    <ul>
      <li v-on:click="activeTab = 'curl'" class="tab-active">Curl</li>
      <li v-on:click="activeTab = 'go'" >Go</li>
      <li v-on:click="activeTab = 'js'" >Js</li>
    </ul>
    <contents v-bind:active-tab="activeTab"></contents>
  </div>
</template>

<script>
import Contents from './Contents'

export default {
  name: 'api-tab-panel',
  components: {
    Contents
  },
  data: () => ({
    activeTab: 'curl'
  })}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>

