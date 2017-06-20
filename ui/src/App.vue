<template>
  <div id="app">
    <about></about>
    <div class="right">
      <random></random>
      <hash></hash>
      <api></api>
    </div>
  </div>
</template>

<script>
import About from './components/About'
import Random from './components/Random'
import Hash from './components/Hash'
import Api from './components/Api'

export default {
  name: 'app',
  components: {
    About,
    Random,
    Hash,
    Api
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  color: #2c3e50;
  margin-top: 60px;
  background: red;
  max-width: 980px;
  margin: 0 auto;
  content: "";
  display: table;
  clear: both;
}
.about, .right {
  display: inline-block;
  vertical-align: top;
}
.about {
  background: blue;
}
.right {
  background: green;
  width: 60%;
}
.random, .hash, .api {
  border: 1px solid #000000;
}
h2 {
  color: #FFFFFF;
  background: #000000;
  font-size: 12px;
  margin: 0;
  padding: 5px;
  text-align: left;
}
</style>
