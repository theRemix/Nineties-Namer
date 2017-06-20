<template>
  <div class="tab-panel">
    <ul>
      <li v-on:click="activeTab = 'curl'" v-bind:class="{ 'tab-active': activeTab==='curl' }">Curl</li>
      <li v-on:click="activeTab = 'go'" v-bind:class="{ 'tab-active': activeTab==='go' }">Go</li>
      <li v-on:click="activeTab = 'js'" v-bind:class="{ 'tab-active': activeTab==='js' }">Js</li>
    </ul>
    <contents v-bind:active-tab="activeTab" class="inner"></contents>
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
ul {
  background: #000000;
  font-size: 12px;
  margin: 0;
  padding: 10px;
  text-align: left;
  text-transform: uppercase;
  letter-spacing: 2px;
  list-style: none;
}
li {
  display: inline-block;
  color: #8C8C8C;
  margin-right: 20px;
  cursor: pointer;
}
.tab-active {
  color: #1AD6C7;
}
</style>

