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
body{
  margin:0;
  padding:0;
}
#app {
  font-family: 'Montserrat', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  color: #2c3e50;
  margin-top: 60px;
  /*max-width: 980px;*/
  /*width: 100%;*/
  margin: 0 auto;
  content: "";
  display: table;
  clear: both;
}
@media (min-width: 930px) {
  #app{
    width: auto;
  }
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
  font-size: 15px;
  font-family: 'Montserrat', sans-serif;
}

.button:hover{
  background-color: #000000;
  color: white;
}
input {
  border: none;
  width: 100%;
  border-bottom: 5px solid #000000;
  font-size: 20px;
  color: #000000;
  font-weight: bold;
  padding-bottom: 5px;
  font-family: 'Montserrat', sans-serif;
  background-color: white;
}

input[type="text"]{
  
}

input[type="text"]:focus{
  outline: none;
}

input:nth-child(1) {
  margin-bottom: 30px;
}
input::placeholder {
  color: #C4C4C4;
  font-style: italic;
  font-family: 'Montserrat', sans-serif;
}
a {
  text-decoration: none;
  color: #000000;
}
.random, .hash {
  border: 1px solid #000000;
  margin-left: 20px;
  margin-right: 20px;
  margin-bottom: 20px;
}
@media (min-width: 930px) {
  .random, .hash{
    margin-left: 0;
    margin-right: 0;
  }
}
.inner {
  padding: 20px;
}
h2 {
  color: #FFFFFF;
  background: #000000;
  font-size: 10px;
  margin: 0;
  padding: 5px;
  padding-left: 10px;
  text-align: left;
  text-transform: uppercase;
  letter-spacing: 2px;
  font-family: 'Montserrat', sans-serif;
}
@media (min-width: 930px) {
  .about, .right {
    display: inline-block;
    vertical-align: top;
  }
  .right {
    width: 70%;
    box-sizing: border-box;
    padding-right: 20px;
  }
  .inner, #app {
    padding: 30px;
  }
}

.api{
  margin-left: 20px;
  margin-right: 20px;
}
@media (min-width: 930px) {
  .api{
    margin-left: 0;
    margin-right: 0;
  }
}
</style>
