<template>
  <div class="tab-panel">
    <ul>
      <li v-on:click="activeTab = 'curl'" v-bind:class="{ 'tab-active': activeTab==='curl' }">Curl</li>
      <li v-on:click="activeTab = 'go'" v-bind:class="{ 'tab-active': activeTab==='go' }">Go</li>
      <li v-on:click="activeTab = 'js'" v-bind:class="{ 'tab-active': activeTab==='js' }">Js</li>
    </ul>
    <contents v-bind:active-tab="activeTab" class="inner code-content"></contents>

    <i v-if='activeTab === "go"'>See: <a href="https://github.com/theRemix/Nineties-Namer/blob/master/routes/router.go">Example on Github</a></i>

    <i v-if='activeTab === "js"'>See: <a href="https://github.com/theRemix/Nineties-Namer/blob/master/ui/src/components/Random.vue">Random Example on Github</a></i>
    <i v-if='activeTab === "js"'>See: <a href="https://github.com/theRemix/Nineties-Namer/blob/master/ui/src/components/Hash.vue">Names Example on Github</a></i>
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
  font-size: 10px;
  margin: 0;
  padding: 5px;
  padding-left: 10px;
  text-align: left;
  text-transform: uppercase;
  letter-spacing: 2px;
  list-style: none;
}
li {
  display: inline-block;
  color: #8C8C8C;
  margin-right: 10px;
  cursor: pointer;
}
.code-content {
  background: #2D2D2D;
  color: #BFBFBF;
  margin-bottom: 10px;
}
.tab-active {
  color: #1AD6C7;
}
i {
  display: block;
  margin-bottom: 10px;
  font-family: 'Raleway', sans-serif;
  font-weight: bold;
}
</style>

