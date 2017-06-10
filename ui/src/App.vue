<template>
  <div id="app">
    <h1>90s Namer</h1>
    <img src="./assets/logo.png">
    <about></about>
    <random></random>
    <hash></hash>
    <api></api>
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
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
