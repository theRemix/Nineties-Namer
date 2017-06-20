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
  max-width: 980px;
  margin: 0 auto;
  content: "";
  display: table;
  clear: both;
}
button, .button {
  border: 5px solid #000000;
  padding: 10px;
  text-align: center;
  text-transform: uppercase;
  font-weight: bold;
  letter-spacing: 2px;
  cursor: pointer;
  transition: 0.2s ease all;
}
input {
  border: none;
  width: 100%;
  border-bottom: 5px solid #000000;
  font-size: 20px;
  font-style: italic;
  color: #C4C4C4;
  font-weight: bold;
  padding-bottom: 5px;
}
input:nth-child(1) {
  margin-bottom: 30px;
}
input::placeholder {
  color: #C4C4C4;
}
a {
  text-decoration: none;
  color: #000000;
}
.about, .right {
  display: inline-block;
  vertical-align: top;
}
.right {
  width: 60%;
}
.random, .hash, .api {
  border: 1px solid #000000;
  margin-bottom: 30px;
}
.inner {
  padding: 30px;
}
h2 {
  color: #FFFFFF;
  background: #000000;
  font-size: 12px;
  margin: 0;
  padding: 10px;
  text-align: left;
  text-transform: uppercase;
  letter-spacing: 2px;
}
</style>
